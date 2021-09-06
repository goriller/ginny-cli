package handle

import (
	"errors"
	"fmt"

	"github.com/fatih/structs"
	"github.com/gorillazer/ginny-cli/ginny/options"
	"github.com/gorillazer/ginny-cli/ginny/util"
	"github.com/iancoleman/strcase"
)

// CreateProto
func CreateProto(serviceName string, validate bool) error {
	conf, err := GetProjectInfo()
	if err != nil {
		util.Error("Failed to get project info, ", err.Error())
		return err
	}

	tmpPath := fmt.Sprintf("%s/%s", conf.ProjectPath, options.TempPath)
	if err := PullTemplate(tmpPath, options.ComponentTemplateRepo); err != nil {
		return err
	}
	protoFile := fmt.Sprintf("%s/proto/demo.proto", tmpPath)
	dstFile := fmt.Sprintf("%s/api/proto/%s.proto", conf.ProjectPath, serviceName)
	if util.Exists(dstFile) {
		return errors.New("File already exists and overwriting is not allowed")
	}
	if validate {
		protoFile = fmt.Sprintf("%s/proto/demo_v.proto", tmpPath)
	}
	if err := util.CopyFile(protoFile, dstFile); err != nil {
		return err
	}
	//
	caseName := strcase.ToCamel(serviceName)
	r := &options.ReplaceKeywords{
		APP_NAME:     conf.ProjectName,
		MODULE_NAME:  conf.ProjectModule,
		SERVICE_NAME: caseName,
	}
	if err := ReplaceFileKeyword([]string{dstFile}, structs.Map(r)); err != nil {
		return err
	}

	if validate {
		protoFile = fmt.Sprintf("%s/proto/validate.proto", tmpPath)
		dstFile = fmt.Sprintf("%s/api/proto/validate.proto", conf.ProjectPath)
		if err := util.CopyFile(protoFile, dstFile); err != nil {
			return err
		}
	}

	if err := ExecCommand(conf.ProjectPath, "make", "proto"); err != nil {
		util.Error(err)
	}

	util.Info("You can modify the proto file, and then execute `make proto` to generate pb code.")
	return nil
}
