package handle

import (
	"errors"
	"fmt"

	"github.com/fatih/structs"
	"github.com/gorillazer/ginny-cli/ginny/options"
	"github.com/gorillazer/ginny-cli/ginny/util"
	"github.com/iancoleman/strcase"
)

func CreateService(serviceName string) error {
	conf, err := GetProjectInfo()
	if err != nil {
		util.Error("Failed to get project info, ", err.Error())
		return err
	}

	tmpPath := fmt.Sprintf("%s/%s", conf.ProjectPath, options.TempPath)
	if err := PullTemplate(tmpPath, options.ComponentTemplateRepo); err != nil {
		return err
	}

	srcFile := fmt.Sprintf("%s/services/test.go", tmpPath)
	dstFile := fmt.Sprintf("%s/internal/services/%s.go", conf.ProjectPath, serviceName)
	if util.Exists(dstFile) {
		return errors.New("File already exists and overwriting is not allowed")
	}
	if err := util.CopyFile(srcFile, dstFile); err != nil {
		return err
	}
	// replace service
	caseName := strcase.ToCamel(serviceName)
	r := &options.ReplaceKeywords{
		APP_NAME:     conf.ProjectName,
		MODULE_NAME:  conf.ProjectModule,
		SERVICE_NAME: caseName,
	}

	// replace provider
	providerFile := fmt.Sprintf("%s/internal/services/provider.go", conf.ProjectPath)
	if !util.Exists(providerFile) {
		if err := util.CopyFile(fmt.Sprintf("%s/services/provider.go", tmpPath), providerFile); err != nil {
			return err
		}
	}
	m := structs.Map(r)
	m[options.ServiceReplaceAnchor[1]] = options.ServiceReplaceAnchorValue[1]([]string{caseName})
	if err := ReplaceFileKeyword([]string{dstFile, providerFile}, m); err != nil {
		return err
	}

	if err := ExecCommand(conf.ProjectPath, "go", "mod", "tidy"); err != nil {
		return err
	}

	return nil
}
