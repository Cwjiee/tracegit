/*
Copyright © 2024 Chong Wei Jie jackchong398@gmail.com
*/
package cmd

import (
	"github.com/Cwjiee/tracegit/ui"
	"github.com/spf13/cobra"
)

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "shows git repositories",
	Long: `Show git repositories or direct to other pages to see commit logs
		open editor on a repo.`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.InitializeScreen(false)
	},
}

func init() {
	rootCmd.AddCommand(gitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// traceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// traceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
