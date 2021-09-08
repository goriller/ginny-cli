package handle

import (
	"errors"
	"fmt"

	"github.com/fatih/structs"
	"github.com/gorillazer/ginny-cli/ginny/options"
	"github.com/gorillazer/ginny-cli/ginny/util"
	"github.com/iancoleman/strcase"
)

func CreateRepo(repoName string, database []string) error {
	conf, err := GetProjectInfo()
	if err != nil {
		util.Error("Failed to get project info, ", err.Error())
		return err
	}

	tmpPath := fmt.Sprintf("%s/%s", conf.ProjectPath, options.TempPath)
	if err := PullTemplate(tmpPath, options.ComponentTemplateRepo); err != nil {
		return err
	}

	srcFile := fmt.Sprintf("%s/repositories/tpl.go", tmpPath)
	dstFile := fmt.Sprintf("%s/internal/repositories/%s.go", conf.ProjectPath, repoName)
	if util.Exists(dstFile) {
		return errors.New("File already exists and overwriting is not allowed")
	}
	if err := util.CopyFile(srcFile, dstFile); err != nil {
		return err
	}
	// replace
	caseName := strcase.ToCamel(repoName)
	r := &options.ReplaceKeywords{
		APP_NAME:    conf.ProjectName,
		MODULE_NAME: conf.ProjectModule,
		REPO_NAME:   caseName,
	}

	// replace provider
	providerFile := fmt.Sprintf("%s/internal/repositories/provider.go", conf.ProjectPath)
	if !util.Exists(providerFile) {
		if err := util.CopyFile(fmt.Sprintf("%s/repositories/provider.go", tmpPath), providerFile); err != nil {
			return err
		}
	}
	m := structs.Map(r)

	m[options.RepoReplaceAnchor[1]] = ""
	m[options.RepoReplaceAnchor[2]] = ""
	m[options.RepoReplaceAnchor[3]] = options.RepoReplaceAnchorValue[3]([]string{caseName})
	m[options.RepoReplaceAnchor[4]] = ""
	m[options.RepoReplaceAnchor[5]] = ""
	m[options.RepoReplaceAnchor[6]] = ""

	for _, db := range database {
		switch db {
		case "mysql":
			m[options.RepoReplaceAnchor[1]] = fmt.Sprintf("%v\n%s \"%s\"", m[options.RepoReplaceAnchor[1]], "mysql", options.DatabaseRepo["mysql"])
			m[options.RepoReplaceAnchor[2]] = fmt.Sprintf("%v\n%s", m[options.RepoReplaceAnchor[2]], "mysql.Provider,")
			m[options.RepoReplaceAnchor[4]] = fmt.Sprintf("%v\n%s %s", m[options.RepoReplaceAnchor[4]], "mysql", options.DatabaseMap["mysql"])
			m[options.RepoReplaceAnchor[5]] = fmt.Sprintf("%v\n%s %s,", m[options.RepoReplaceAnchor[5]], "mysql", options.DatabaseMap["mysql"])
			m[options.RepoReplaceAnchor[6]] = fmt.Sprintf("%v\n%s:%s,", m[options.RepoReplaceAnchor[6]], "mysql", "mysql")
		case "mongo":
			m[options.RepoReplaceAnchor[1]] = fmt.Sprintf("%v\n%s \"%s\"", m[options.RepoReplaceAnchor[1]], "mongo", options.DatabaseRepo["mongo"])
			m[options.RepoReplaceAnchor[2]] = fmt.Sprintf("%v\n%s", m[options.RepoReplaceAnchor[2]], "mongo.Provider,")
			m[options.RepoReplaceAnchor[4]] = fmt.Sprintf("%v\n%s %s", m[options.RepoReplaceAnchor[4]], "mongo", options.DatabaseMap["mongo"])
			m[options.RepoReplaceAnchor[5]] = fmt.Sprintf("%v\n%s %s,", m[options.RepoReplaceAnchor[5]], "mongo", options.DatabaseMap["mongo"])
			m[options.RepoReplaceAnchor[6]] = fmt.Sprintf("%v\n%s:%s,", m[options.RepoReplaceAnchor[6]], "mongo", "mongo")
		case "redis":
			m[options.RepoReplaceAnchor[1]] = fmt.Sprintf("%v\n%s \"%s\"", m[options.RepoReplaceAnchor[1]], "redis", options.DatabaseRepo["redis"])
			m[options.RepoReplaceAnchor[2]] = fmt.Sprintf("%v\n%s", m[options.RepoReplaceAnchor[2]], "redis.Provider,")
			m[options.RepoReplaceAnchor[4]] = fmt.Sprintf("%v\n%s %s", m[options.RepoReplaceAnchor[4]], "redis", options.DatabaseMap["redis"])
			m[options.RepoReplaceAnchor[5]] = fmt.Sprintf("%v\n%s %s,", m[options.RepoReplaceAnchor[5]], "redis", options.DatabaseMap["redis"])
			m[options.RepoReplaceAnchor[6]] = fmt.Sprintf("%v\n%s:%s,", m[options.RepoReplaceAnchor[6]], "redis", "redis")
		}
	}
	m[options.RepoReplaceAnchor[1]] = fmt.Sprintf("%v \n%s", m[options.RepoReplaceAnchor[1]], options.RepoReplaceAnchor[1])
	m[options.RepoReplaceAnchor[2]] = fmt.Sprintf("%v \n%s", m[options.RepoReplaceAnchor[2]], options.RepoReplaceAnchor[2])
	m[options.RepoReplaceAnchor[4]] = fmt.Sprintf("%v \n%s", m[options.RepoReplaceAnchor[4]], options.RepoReplaceAnchor[4])
	m[options.RepoReplaceAnchor[5]] = fmt.Sprintf("%v \n%s", m[options.RepoReplaceAnchor[5]], options.RepoReplaceAnchor[5])
	m[options.RepoReplaceAnchor[6]] = fmt.Sprintf("%v \n%s", m[options.RepoReplaceAnchor[6]], options.RepoReplaceAnchor[6])

	// replace /cmd/provider.go
	appProviderFile := conf.ProjectPath + "/cmd/provider.go"
	m[options.AppReplaceAnchor[1]] = options.AppReplaceAnchorValue[1]([]string{"repo", conf.ProjectModule, appProviderFile})
	m[options.AppReplaceAnchor[2]] = options.AppReplaceAnchorValue[2]([]string{"repo", appProviderFile})

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
