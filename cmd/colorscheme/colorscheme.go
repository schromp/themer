/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package colorscheme

import (
	"github.com/spf13/cobra"
)

// colorschemeCmd represents the colorscheme command
var ColorschemeCmd = &cobra.Command{
	Use:   "colorscheme",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// colorschemeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// colorschemeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
