package command

import (
	"errors"

	"github.com/gorillazer/ginny-cli/ginny/handle"
	"github.com/gorillazer/ginny-cli/ginny/util"
	"github.com/spf13/cobra"
)

func init() {
	grpcCmd.Flags().StringP("type", "t", "server", "GRPC service type created, server or client")
	rootCmd.AddCommand(grpcCmd)
}

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Create grpc server/client file",
	Long:  "Create grpc server/client file from template",
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if err := CheckArgs(args); err != nil {
			return err
		}
		// 获取参数
		serverName := args[0]
		flags := cmd.Flags()
		types, err := flags.GetString("type")
		if err != nil {
			return err
		}
		if types != "server" && types != "client" {
			return errors.New("The wrong type was entered")
		}
		if err := handle.CreateGrpc(serverName, types); err != nil {
			return err
		}

		util.Info("Create new service file success!")
		return nil
	},
}
