package handle

import (
	"github.com/gorillazer/ginny-cli/ginny/options"
	"github.com/gorillazer/ginny-cli/ginny/util"
	"gopkg.in/yaml.v3"
)

// CreateProject 创建项目
func CreateProject(projectName, moduleName string, args ...string) error {
	d, err := GetCurrentDir()
	if err != nil {
		util.Error("Failed to get project directory, ", err.Error())
		return err
	}
	projectDir := d + "/" + projectName
	if err := PullTemplate(projectDir); err != nil {
		return err
	}
	// 删除多余文件
	if err := util.RemoveFile(projectDir + "/.git"); err != nil {
		return err
	}

	if moduleName == "" {
		moduleName = projectName
	}
	kb := &options.ReplaceKeywords{
		APP_NAME:    projectName,
		MODULE_NAME: moduleName,
	}
	// 替换关键字
	if err := ReplaceFileKeyword(util.GetFiles(projectDir), kb); err != nil {
		return err
	}

	// 写入项目标识
	p := &options.ProjectInfo{
		ProjectName:   projectName,
		ProjectModule: moduleName,
	}
	if err := writeProjectFlag(projectDir, p); err != nil {
		return err
	}

	return nil
}

// writeProjectFlag
func writeProjectFlag(projectDir string, p *options.ProjectInfo) error {
	bt, err := yaml.Marshal(p)
	if err != nil {
		return err
	}
	return util.WriteToFile(projectDir+"/"+options.ProjectFlag, bt)
}
