package command

import (
	"fmt"
	"log"
	"os"

	"github.com/gorillazer/ginny-cli/ginny/options"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ginny",
	Short: "Ginny project command line tool",
	Long:  "Command line tool for Ginny project bestpractice",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\x1b[35;1m%s\x1b[0m\n", options.Logo)
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

// Exec
func Exec() {
	// cliBox = packr.NewBox("../template")
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("Command line error %v", err)
		os.Exit(1)
	}
}
