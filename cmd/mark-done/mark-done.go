package markdone

import "github.com/spf13/cobra"

var MarkDoneCmd = &cobra.Command{
	Use:   "mark-done",
	Short: "Use to change the status of a task",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {

}
