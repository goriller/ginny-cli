package command

import (
	"github.com/gorillazer/ginny-cli/ginny/handle"
	"github.com/gorillazer/ginny-cli/ginny/util"
	"github.com/spf13/cobra"
)

func init() {
	protoCmd.Flags().BoolP("validate", "v", false, "Added support for parameter verification")
	rootCmd.AddCommand(protoCmd)
}

var protoCmd = &cobra.Command{
	Use:   "proto",
	Short: "Create proto file",
	Long:  "Create proto file from template",
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if err := CheckArgs(args); err != nil {
			return err
		}
		// 获取参数
		serviceName := args[0]
		flags := cmd.Flags()
		validate, err := flags.GetBool("validate")
		if err != nil {
			return err
		}
		if err := handle.CreateProto(serviceName, validate); err != nil {
			return err
		}

		util.Info("Create new proto file success!")
		util.Info("You can modify the proto file, and then execute `make proto` to generate pb code.")
		return nil
	},
}
