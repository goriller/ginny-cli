package handle

import (
	"strings"

	"github.com/fatih/structs"
	"github.com/gorillazer/ginny-cli/ginny/options"
	"github.com/gorillazer/ginny-cli/ginny/util"
)

// CreateProject 创建项目
func CreateProject(projectName, moduleName string, args ...string) error {
	d, err := GetCurrentDir()
	if err != nil {
		util.Error("Failed to get project directory, ", err.Error())
		return err
	}
	ProjectPath := d + "/" + projectName
	if err := PullTemplate(ProjectPath, options.TemplateRepo); err != nil {
		return err
	}
	// 删除多余文件
	if err := util.RemoveFile(ProjectPath + "/.git"); err != nil {
		return err
	}

	if moduleName == "" {
		moduleName = projectName
	}

	argStr := strings.Join(args, ",")
	if strings.Contains(argStr, "grpc") {
		if err := CreateGrpc(projectName, "server"); err != nil {
			return err
		}
	}

	r := &options.ReplaceKeywords{
		APP_NAME:    projectName,
		MODULE_NAME: moduleName,
	}
	// 替换关键字
	if err := ReplaceFileKeyword(util.GetFiles(ProjectPath), structs.Map(r)); err != nil {
		return err
	}

	// 写入项目标识
	p := &options.ProjectInfo{
		ProjectName:   projectName,
		ProjectModule: moduleName,
	}
	if err := GenerateProjectInfo(ProjectPath, p); err != nil {
		return err
	}

	//
	if err := ExecCommand(ProjectPath, "go", "mod", "tidy"); err != nil {
		return err
	}

	_ = GoFmtDir(ProjectPath)

	return nil
}
