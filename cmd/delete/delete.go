package delete

import "github.com/spf13/cobra"

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a task",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {

}
