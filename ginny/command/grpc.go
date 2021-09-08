package command

import (
	"github.com/gorillazer/ginny-cli/ginny/handle"
	"github.com/gorillazer/ginny-cli/ginny/util"
	"github.com/spf13/cobra"
)

func init() {
	grpcCmd.Flags().BoolP("server", "s", true, "GRPC service type created, server or client")
	grpcCmd.Flags().BoolP("client", "c", false, "GRPC service type created, server or client")
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
		types := []string{}
		server, err := flags.GetBool("server")
		if err != nil {
			return err
		}
		if server {
			types = append(types, "server")
		}
		client, err := flags.GetBool("client")
		if err != nil {
			return err
		}
		if client {
			types = append(types, "client")
		}
		if err := handle.CreateGrpc(serverName, types...); err != nil {
			return err
		}

		util.Info("Create new service file success!")
		return nil
	},
}
