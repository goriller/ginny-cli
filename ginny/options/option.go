package options

var (
	Logo = `

	┌─┐┬┌┐┌┌┐┌┬ ┬
	│ ┬│││││││└┬┘
	└─┘┴┘└┘┘└┘ ┴ 
	
	`
	// 模板仓库地址
	TemplateRepo = "https://github.com/gorillazer/ginny-template.git"
	// 项目标识
	ProjectFlag = ".ginny.yml"
)

// ProjectInfo
type ProjectInfo struct {
	ProjectName   string `yaml:"project_name,omitempty"`
	ProjectPath   string
	ProjectModule string `yaml:"project_module,omitempty"`
}

// ReplaceKeywords 标识字典
type ReplaceKeywords struct {
	APP_NAME    string
	MODULE_NAME string
}
