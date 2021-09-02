package handle

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/structs"
	"github.com/go-git/go-git/v5"
	"github.com/gorillazer/ginny-cli/ginny/options"
	"github.com/gorillazer/ginny-cli/ginny/util"
)

// GetCurrentDir 获取当前目录
func GetCurrentDir() (string, error) {
	// 获取当前目录
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return dir, nil
}

// PullTemplate 拉取模板
func PullTemplate(dir string) error {
	// Clone the given repository to the given directory
	util.Info("git clone " + options.TemplateRepo)

	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:      options.TemplateRepo,
		Progress: os.Stdout,
	})
	if err != nil {
		util.Error("Clone the given repository error:", err.Error())
		return err
	}
	return nil
}

// GenerateProjectInfo 构造项目标识
func GenerateProjectInfo() {

}

// GetProjectInfo 检查当前目录是否ginny项目，并返回项目信息
func GetProjectInfo() (conf *options.ProjectInfo, err error) {
	return nil, nil
}

// ReplaceFileKeyword
func ReplaceFileKeyword(file []string, r *options.ReplaceKeywords) error {
	m := structs.Map(r)
	for _, f := range file {
		if !util.IsFile(f) {
			continue
		}
		for k, v := range m {
			if m[k] != "" {
				if err := util.ReplaceFile(f, k, fmt.Sprintf("%v", v)); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// ExecCommand
func ExecCommand(dir, name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	util.Info("ExecCommand", cmd)
	err := cmd.Run()
	if err != nil {
		return err
	}
	util.Info("ExecCommand", "End")
	return nil
}
