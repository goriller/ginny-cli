package options

import "fmt"

var (
	Logo = `

	┌─┐┬┌┐┌┌┐┌┬ ┬
	│ ┬│││││││└┬┘
	└─┘┴┘└┘┘└┘ ┴ 
	
	`
	// 模板仓库地址
	TemplateRepo = "https://github.com/gorillazer/ginny-template.git"
	// 组件模板仓库地址
	ComponentTemplateRepo = "https://github.com/gorillazer/ginny-component-template.git"
	// 数据库组件仓库地址
	DatabaseRepo = map[string]string{
		"mongo": "github.com/gorillazer/ginny-mongo",
		"mysql": "github.com/gorillazer/ginny-mysql",
		"redis": "github.com/gorillazer/ginny-redis",
	}
	DatabaseMap = map[string]string{
		"mysql": "*mysql.SqlBuilder",
		"mongo": "*mongo.Manager",
		"redis": "*redis.Manager",
	}
	// 项目标识
	ProjectFlag = ".ginny.yml"
	// 组件临时缓存目录
	TempPath = ".ginny"
	// 替换锚点
	AnchorFlag          = " 锚点请勿删除! Do not delete this line!"
	HandleReplaceAnchor = map[int]string{
		1: "// HANDLE" + AnchorFlag,
		2: "// HANDLE_PROVIDER" + AnchorFlag,
	}
	HandleReplaceAnchorValue = map[int]func(args []string) string{
		1: func(args []string) string {
			return fmt.Sprintf("%s *%sHandler,\n%s", args[0], args[1], HandleReplaceAnchor[1])
		},
		2: func(args []string) string {
			return fmt.Sprintf("%sHandlerProvider,\n%s", args[0], HandleReplaceAnchor[2])
		},
	}
	ServiceReplaceAnchor = map[int]string{
		1: "// SERVICE_PROVIDER" + AnchorFlag,
	}
	ServiceReplaceAnchorValue = map[int]func(args []string) string{
		1: func(args []string) string {
			return fmt.Sprintf("%sServiceProvider,\n%s", args[0], ServiceReplaceAnchor[1])
		},
	}
	RepoReplaceAnchor = map[int]string{
		1: "// DATABASE_LIB" + AnchorFlag,
		2: "// DATABASE_PROVIDER" + AnchorFlag,
		3: "// REPO_PROVIDER" + AnchorFlag,
		4: "// STRUCT_ATTR" + AnchorFlag,
		5: "// FUNC_PARAM" + AnchorFlag,
		6: "// FUNC_ATTR" + AnchorFlag,
	}
	RepoReplaceAnchorValue = map[int]func(args []string) string{
		3: func(args []string) string {
			return fmt.Sprintf("%sRepositoryProvider,\n%s", args[0], RepoReplaceAnchor[3])
		},
	}
)

// ProjectInfo
type ProjectInfo struct {
	ProjectName   string `yaml:"project_name,omitempty"`
	ProjectPath   string
	ProjectModule string `yaml:"project_module,omitempty"`
}

// ReplaceKeywords 标识字典
type ReplaceKeywords struct {
	APP_NAME     string
	MODULE_NAME  string
	SERVICE_NAME string
	HANDLE_NAME  string
	REPO_NAME    string
}
