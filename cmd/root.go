package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "task-manager-cli",
	Short: "A simple task manager on Go",
	Long: `
	The app is designed to help with making and managing simple to-do tasks.
	It provides basic command set and using CLI and uses CLI to interact with tasks.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(doneCmd)
	rootCmd.AddCommand(deleteCmd)

	listCmd.Flags().Bool("completed", false, "for getting only completed tasks")
	listCmd.Flags().Bool("pending", false, "for getting only non-completed tasks")
}
