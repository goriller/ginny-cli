package command

import (
	"github.com/gorillazer/ginny-cli/ginny/handle"
	"github.com/gorillazer/ginny-cli/ginny/util"
	"github.com/spf13/cobra"
)

func init() {
	newCmd.Flags().StringP("module", "m", "", "Define the project module, ex: git.tencent.com/demo")
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
		if err := handle.CreateProject(projectName, module); err != nil {
			return err
		}

		util.Info("Create new project success!")
		return nil
	},
}
