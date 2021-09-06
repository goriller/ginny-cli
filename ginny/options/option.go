package options

import (
	"fmt"

	"github.com/gorillazer/ginny-cli/ginny/util"
)

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
	AnchorFlag       = " 锚点请勿删除! Do not delete this line!"
	AppReplaceAnchor = map[int]string{
		1: "// CMD_IMPORT" + AnchorFlag,
		2: "// CMD_PROVIDERSET" + AnchorFlag,
		3: "// CMD_SERVEPARAM" + AnchorFlag,
		4: "// CMD_SERVEFUNC" + AnchorFlag,
	}
	AppReplaceAnchorValue = map[int]func(args []string) string{
		1: func(args []string) string { // args[0]=type args[1]=moduleName args[2]=projectDir/cmd/provider.go
			str := ""
			if args[0] == "handler" {
				str = "\"" + args[1] + "/internal/handlers\""
				if b, _ := util.FileHasContainsStr(args[2], str); b {
					str = ""
				}
			} else if args[0] == "service" {
				str = "\"" + args[1] + "/internal/services\""
				if b, _ := util.FileHasContainsStr(args[2], str); b {
					str = ""
				}
			} else if args[0] == "repo" {
				str = "\"" + args[1] + "/internal/repositories\""
				if b, _ := util.FileHasContainsStr(args[2], str); b {
					str = ""
				}
			} else if args[0] == "grpc_server" {
				str = "rpc \"" + args[1] + "/internal/rpc\""
				if b, _ := util.FileHasContainsStr(args[2], str); b {
					str = ""
				} else {
					str += "\nrpc_server \"" + args[1] + "/internal/rpc/server\""
				}
			} else if args[0] == "grpc_client" {
				str = "rpc \"" + args[1] + "/internal/rpc\""
				if b, _ := util.FileHasContainsStr(args[2], str); b {
					str = ""
				} else {
					str += "\nrpc_client \"" + args[1] + "/internal/rpc/client\""
				}
			}

			if str == "" {
				return ""
			}
			return fmt.Sprintf("%s\n%s", str, AppReplaceAnchor[1])
		},
		2: func(args []string) string { // args[0]=type args[1]=projectDir/cmd/provider.go
			str := ""
			if args[0] == "handler" {
				str = "handlers.ProviderSet,"
				if b, _ := util.FileHasContainsStr(args[1], str); b {
					str = ""
				}
			} else if args[0] == "service" {
				str = "services.ProviderSet,"
				if b, _ := util.FileHasContainsStr(args[1], str); b {
					str = ""
				}
			} else if args[0] == "repo" {
				str = "repositories.ProviderSet,"
				if b, _ := util.FileHasContainsStr(args[1], str); b {
					str = ""
				}
			} else if args[0] == "grpc_server" {
				str = "rpc_server.ProviderSet,"
				if b, _ := util.FileHasContainsStr(args[1], str); b {
					str = ""
				} else {
					str += "\n rpc.ProviderSet,"
				}
			} else if args[0] == "grpc_client" {
				str = "rpc_client.ProviderSet,"
				if b, _ := util.FileHasContainsStr(args[1], str); b {
					str = ""
				} else {
					str += "\n rpc.ProviderSet,"
				}
			}
			if str == "" {
				return ""
			}
			return fmt.Sprintf("%s\n%s", str, AppReplaceAnchor[2])
		},
		3: func(args []string) string { // args[0]=type args[1]=projectDir/cmd/provider.go
			str := ""
			if args[0] == "http" {
				str = "hs *http.Server,"
				if b, _ := util.FileHasContainsStr(args[1], "http.Server"); b {
					str = ""
				}
			} else if args[0] == "grpc_server" {
				str = "gs *grpc.Server,"
				if b, _ := util.FileHasContainsStr(args[1], "grpc.Server"); b {
					str = ""
				} else {
					str += "\ncli *consul.Client,"
				}
			} else if args[0] == "grpc_client" {
				str = "gs *grpc.Server,"
				if b, _ := util.FileHasContainsStr(args[1], "grpc.Server"); b {
					str = ""
				} else {
					str += "\ncli *consul.Client,"
				}
			}

			if str == "" {
				return ""
			}
			return fmt.Sprintf("%s\n%s", str, AppReplaceAnchor[3])
		},
		4: func(args []string) string { // args[0]=type args[1]=projectDir/cmd/provider.go
			str := ""
			if args[0] == "http" {
				str = "ginny.HttpServe(hs),"
				if b, _ := util.FileHasContainsStr(args[1], str); b {
					str = ""
				}
			} else if args[0] == "grpc_server" {
				str = "ginny.GrpcServeWithConsul(gs, cli),"
				if b, _ := util.FileHasContainsStr(args[1], "ginny.GrpcServe"); b {
					str = ""
				}
			} else if args[0] == "grpc_client" {
				str = "ginny.GrpcServeWithConsul(gs, cli),"
				if b, _ := util.FileHasContainsStr(args[1], "ginny.GrpcServe"); b {
					str = ""
				}
			}
			if str == "" {
				return ""
			}
			return fmt.Sprintf("%s\n%s", str, AppReplaceAnchor[4])
		},
	}
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
	ServerReplaceAnchor = map[int]string{
		1: "// SERVER_IMPORT" + AnchorFlag,
		2: "// SERVER_PARAM" + AnchorFlag,
		3: "// SERVER_REGIST" + AnchorFlag,
		4: "// SERVER_PROVIDER" + AnchorFlag,
	}
	ServerReplaceAnchorValue = map[int]func(args []string) string{
		1: func(args []string) string {
			return fmt.Sprintf("proto \"%s/api/proto\"\n%s", args[0], ServerReplaceAnchor[1])
		},
		2: func(args []string) string {
			return fmt.Sprintf("%s *%sServer,\n%s", args[0], args[1], ServerReplaceAnchor[2])
		},
		3: func(args []string) string {
			return fmt.Sprintf("proto.Register%sServer(s, %s)\n%s", args[0], args[1], ServerReplaceAnchor[3])
		},
		4: func(args []string) string {
			return fmt.Sprintf("%sServerProvider,\n%s", args[0], ServerReplaceAnchor[4])
		},
	}
	ClientReplaceAnchor = map[int]string{
		1: "// CLIENT_PROVIDER" + AnchorFlag,
	}
	ClientReplaceAnchorValue = map[int]func(args []string) string{
		1: func(args []string) string {
			return fmt.Sprintf("New%sClient,\n%s", args[0], ServerReplaceAnchor[1])
		},
	}
)

// ProjectInfo
type ProjectInfo struct {
	ProjectName   string `yaml:"project_name"`
	ProjectPath   string `yaml:"project_path"`
	ProjectModule string `yaml:"project_module"`
}

// ReplaceKeywords 标识字典
type ReplaceKeywords struct {
	APP_NAME     string
	MODULE_NAME  string
	SERVICE_NAME string
	HANDLE_NAME  string
	REPO_NAME    string
	SERVER_NAME  string
}
