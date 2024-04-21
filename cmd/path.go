/*
Copyright © 2024 Chong Wei Jie jackchong398@gmail.com
*/
package cmd

import (
	"github.com/Cwjiee/tracegit/ui"
	"github.com/spf13/cobra"
)

var pathCmd = &cobra.Command{
	Use:   "path",
	Short: "edit your code directory path",
	Long: `input your code directory where you store all your projects.
	It will be the directory where Tracegit looks for your git repositories`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.InitializeScreen(true)
	},
}

func init() {
	rootCmd.AddCommand(pathCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// traceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// traceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
