package command

import (
	"github.com/gorillazer/ginny-cli/ginny/handle"
	"github.com/gorillazer/ginny-cli/ginny/util"
	"github.com/spf13/cobra"
)

func init() {
	repoCmd.Flags().StringArrayP("database", "d", []string{}, "Define the database used by the project, support mysql、mongo、redis")
	rootCmd.AddCommand(repoCmd)
}

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Create repository file",
	Long:  "Create repository file from template",
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if err := CheckArgs(args); err != nil {
			return err
		}
		// 获取参数
		repoName := args[0]
		flags := cmd.Flags()
		database, err := flags.GetStringArray("database")
		if err != nil {
			return err
		}
		if err := handle.CreateRepo(repoName, database); err != nil {
			return err
		}

		util.Info("Create new repository file success!")
		return nil
	},
}
