package colorscheme

import "github.com/spf13/cobra"

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a colorscheme",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	ColorschemeCmd.AddCommand(setCmd)
}
