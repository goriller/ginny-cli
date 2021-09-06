package handle

import (
	"errors"
	"fmt"

	"github.com/fatih/structs"
	"github.com/gorillazer/ginny-cli/ginny/options"
	"github.com/gorillazer/ginny-cli/ginny/util"
	"github.com/iancoleman/strcase"
)

func CreateGrpc(serverName, types string) error {
	conf, err := GetProjectInfo()
	if err != nil {
		util.Error("Failed to get project info, ", err.Error())
		return err
	}

	tmpPath := fmt.Sprintf("%s/%s", conf.ProjectPath, options.TempPath)
	if err := PullTemplate(tmpPath, options.ComponentTemplateRepo); err != nil {
		return err
	}

	if types == "server" {
		if err := createServer(serverName, tmpPath, conf); err != nil {
			return err
		}
	} else {
		if err := createClient(serverName, tmpPath, conf); err != nil {
			return err
		}
	}

	if err := ExecCommand(conf.ProjectPath, "go", "mod", "tidy"); err != nil {
		return err
	}

	return nil
}

func createServer(serverName, tmpPath string, conf *options.ProjectInfo) error {
	srcFile := fmt.Sprintf("%s/rpc/server/tpl.go", tmpPath)
	dstFile := fmt.Sprintf("%s/internal/rpc/server/%s.go", conf.ProjectPath, serverName)
	if util.Exists(dstFile) {
		return errors.New("File already exists and overwriting is not allowed")
	}
	if err := util.CopyFile(srcFile, dstFile); err != nil {
		return err
	}
	// replace provider
	providerFile := fmt.Sprintf("%s/internal/rpc/server/provider.go", conf.ProjectPath)
	if !util.Exists(providerFile) {
		if err := util.CopyFile(fmt.Sprintf("%s/rpc/server/provider.go", tmpPath), providerFile); err != nil {
			return err
		}
	}
	// rpcProviderFile
	rpcProviderFile := fmt.Sprintf("%s/internal/rpc/provider.go", conf.ProjectPath)
	if !util.Exists(rpcProviderFile) {
		if err := util.CopyFile(fmt.Sprintf("%s/rpc/provider.go", tmpPath), rpcProviderFile); err != nil {
			return err
		}
	}
	// proto file
	protoFile := fmt.Sprintf("%s/api/proto/%s.proto", conf.ProjectPath, serverName)
	if !util.Exists(protoFile) {
		if err := CreateProto(serverName, true); err != nil {
			return err
		}
	}

	// replace service
	caseName := strcase.ToCamel(serverName)
	r := &options.ReplaceKeywords{
		APP_NAME:    conf.ProjectName,
		MODULE_NAME: conf.ProjectModule,
		SERVER_NAME: caseName,
	}
	m := structs.Map(r)
	m[options.ServerReplaceAnchor[1]] = options.ServerReplaceAnchorValue[1]([]string{conf.ProjectModule})
	m[options.ServerReplaceAnchor[2]] = options.ServerReplaceAnchorValue[2]([]string{serverName, caseName})
	m[options.ServerReplaceAnchor[3]] = options.ServerReplaceAnchorValue[3]([]string{caseName, serverName})
	m[options.ServerReplaceAnchor[4]] = options.ServerReplaceAnchorValue[4]([]string{caseName})

	// replace /cmd/provider.go
	appProviderFile := conf.ProjectPath + "/cmd/provider.go"
	m[options.AppReplaceAnchor[1]] = options.AppReplaceAnchorValue[1]([]string{"grpc_server", conf.ProjectModule, appProviderFile})
	m[options.AppReplaceAnchor[2]] = options.AppReplaceAnchorValue[2]([]string{"grpc_server", appProviderFile})
	m[options.AppReplaceAnchor[3]] = options.AppReplaceAnchorValue[3]([]string{"grpc_server", appProviderFile})
	m[options.AppReplaceAnchor[4]] = options.AppReplaceAnchorValue[4]([]string{"grpc_server", appProviderFile})

	if err := ReplaceFileKeyword([]string{dstFile, providerFile, appProviderFile}, m); err != nil {
		return err
	}

	return nil
}

func createClient(clientName, tmpPath string, conf *options.ProjectInfo) error {
	srcFile := fmt.Sprintf("%s/rpc/client/tpl.go", tmpPath)
	dstFile := fmt.Sprintf("%s/internal/rpc/client/%s.go", conf.ProjectPath, clientName)
	if util.Exists(dstFile) {
		return errors.New("File already exists and overwriting is not allowed")
	}
	if err := util.CopyFile(srcFile, dstFile); err != nil {
		return err
	}
	// replace provider
	providerFile := fmt.Sprintf("%s/internal/rpc/client/provider.go", conf.ProjectPath)
	if !util.Exists(providerFile) {
		if err := util.CopyFile(fmt.Sprintf("%s/rpc/client/provider.go", tmpPath), providerFile); err != nil {
			return err
		}
	}
	// rpcProviderFile
	rpcProviderFile := fmt.Sprintf("%s/internal/rpc/provider.go", conf.ProjectPath)
	if !util.Exists(rpcProviderFile) {
		if err := util.CopyFile(fmt.Sprintf("%s/rpc/provider.go", tmpPath), rpcProviderFile); err != nil {
			return err
		}
	}
	// replace service
	caseName := strcase.ToCamel(clientName)
	r := &options.ReplaceKeywords{
		APP_NAME:    conf.ProjectName,
		MODULE_NAME: conf.ProjectModule,
		SERVER_NAME: caseName,
	}
	m := structs.Map(r)
	m[options.ClientReplaceAnchor[1]] = options.ClientReplaceAnchorValue[1]([]string{caseName})

	// replace /cmd/provider.go
	appProviderFile := conf.ProjectPath + "/cmd/provider.go"
	m[options.AppReplaceAnchor[1]] = options.AppReplaceAnchorValue[1]([]string{"grpc_client", conf.ProjectModule, appProviderFile})
	m[options.AppReplaceAnchor[2]] = options.AppReplaceAnchorValue[2]([]string{"grpc_client", appProviderFile})
	m[options.AppReplaceAnchor[3]] = options.AppReplaceAnchorValue[3]([]string{"grpc_client", appProviderFile})
	m[options.AppReplaceAnchor[4]] = options.AppReplaceAnchorValue[4]([]string{"grpc_client", appProviderFile})

	if err := ReplaceFileKeyword([]string{dstFile, providerFile, appProviderFile}, m); err != nil {
		return err
	}
	return nil
}
