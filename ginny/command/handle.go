package command

import (
	"github.com/gorillazer/ginny-cli/ginny/handle"
	"github.com/gorillazer/ginny-cli/ginny/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(handleCmd)
}

var handleCmd = &cobra.Command{
	Use:   "handle",
	Short: "Create handle file",
	Long:  "Create handle file from template",
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if err := CheckArgs(args); err != nil {
			return err
		}
		// 获取参数
		handleName := args[0]

		if err := handle.CreateHandle(handleName); err != nil {
			return err
		}

		util.Info("Create new handle file success!")
		return nil
	},
}
