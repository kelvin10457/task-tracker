package update_task

import "github.com/spf13/cobra"

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "update the description of a task",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {

}
