package colorscheme

import (
	"schromp/themer/colorscheme"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a colorscheme",
	Run: func(cmd *cobra.Command, args []string) {
		colorschemes := colorscheme.GenerateColorschemes() // TODO: this should probably be done at root level and passed

		colorschemes[args[0]].Set()

		// Call set function on the 
	},
}

func init() {
	ColorschemeCmd.AddCommand(setCmd)
}
