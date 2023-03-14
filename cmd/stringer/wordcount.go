package stringer

import (
	"fmt"

	"github.com/ahy69195/CLITutorial2/pkg/stringer"
	"github.com/spf13/cobra"
)

var wordCountCmd = &cobra.Command{
	Use:     "wordcount",
	Aliases: []string{"wrdc"},
	Short:   "Counts number of words",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		i := args[0]
		res := stringer.Wordcount(i)

		fmt.Printf("'%s' has %d words.\n", i, res)
	},
}

func init() {
	rootCmd.AddCommand(wordCountCmd)
}
