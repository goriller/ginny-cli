package command

import (
	"github.com/gorillazer/ginny-cli/ginny/handle"
	"github.com/gorillazer/ginny-cli/ginny/util"
	"github.com/spf13/cobra"
)

func init() {
	newCmd.Flags().StringP("module", "m", "", "Define the project module, ex: github.com/demo")
	newCmd.Flags().Bool("grpc", false, "Create a grpc service project")
	newCmd.Flags().Bool("http", true, "Create a http service project")
	rootCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new Ginny project",
	Long:  "Create a new Ginny project from template",
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if err := CheckArgs(args); err != nil {
			return err
		}
		// 获取参数
		projectName := args[0]
		flags := cmd.Flags()
		module, err := flags.GetString("module")
		if err != nil {
			return err
		}
		enableGrpc, err := flags.GetBool("grpc")
		if err != nil {
			return err
		}
		enableHttp, err := flags.GetBool("http")
		if err != nil {
			return err
		}
		arg := []string{}
		if enableHttp || (!enableGrpc && !enableHttp) {
			arg = append(arg, "http")
		}
		if enableGrpc {
			arg = append(arg, "grpc")
		}
		if err := handle.CreateProject(projectName, module, arg...); err != nil {
			return err
		}

		util.Info("Create new project success!")
		return nil
	},
}
