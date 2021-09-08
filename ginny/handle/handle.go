package handle

import (
	"errors"
	"fmt"

	"github.com/fatih/structs"
	"github.com/gorillazer/ginny-cli/ginny/options"
	"github.com/gorillazer/ginny-cli/ginny/util"
	"github.com/iancoleman/strcase"
)

func CreateHandle(handleName string) error {
	conf, err := GetProjectInfo()
	if err != nil {
		util.Error("Failed to get project info, ", err.Error())
		return err
	}

	tmpPath := fmt.Sprintf("%s/%s", conf.ProjectPath, options.TempPath)
	if err := PullTemplate(tmpPath, options.ComponentTemplateRepo); err != nil {
		return err
	}

	srcFile := fmt.Sprintf("%s/handlers/tpl.go", tmpPath)
	dstFile := fmt.Sprintf("%s/internal/handlers/%s.go", conf.ProjectPath, handleName)
	if util.Exists(dstFile) {
		return errors.New("File already exists and overwriting is not allowed")
	}
	if err := util.CopyFile(srcFile, dstFile); err != nil {
		return err
	}
	// replace handle
	caseName := strcase.ToCamel(handleName)
	r := &options.ReplaceKeywords{
		APP_NAME:    conf.ProjectName,
		MODULE_NAME: conf.ProjectModule,
		HANDLE_NAME: caseName,
	}

	// replace provider
	providerFile := fmt.Sprintf("%s/internal/handlers/provider.go", conf.ProjectPath)
	if !util.Exists(providerFile) {
		if err := util.CopyFile(fmt.Sprintf("%s/handlers/provider.go", tmpPath), providerFile); err != nil {
			return err
		}
	}
	m := structs.Map(r)
	m[options.HandleReplaceAnchor[1]] = options.HandleReplaceAnchorValue[1]([]string{handleName, caseName})
	m[options.HandleReplaceAnchor[2]] = options.HandleReplaceAnchorValue[2]([]string{caseName})

	// replace /cmd/provider.go
	appProviderFile := conf.ProjectPath + "/cmd/provider.go"
	m[options.AppReplaceAnchor[1]] = options.AppReplaceAnchorValue[1]([]string{"handler", conf.ProjectModule, appProviderFile})
	m[options.AppReplaceAnchor[2]] = options.AppReplaceAnchorValue[2]([]string{"handler", appProviderFile})

	if err := ReplaceFileKeyword([]string{dstFile, providerFile, appProviderFile}, m); err != nil {
		return err
	}

	if err := ExecCommand(conf.ProjectPath, "go", "mod", "tidy"); err != nil {
		return err
	}

	if err := ExecCommand(conf.ProjectPath, "wire", "./cmd/."); err != nil {
		return err
	}

	_ = GoFmtDir(conf.ProjectPath)

	return nil
}
