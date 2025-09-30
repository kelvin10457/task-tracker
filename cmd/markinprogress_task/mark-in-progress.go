package markinprogress_task

import (
	"errors"
	"fmt"
	"strconv"

	contextkey "github.com/kelvin10457/task-tracker/internal/contextKey"
	"github.com/spf13/cobra"
)

var MarkInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress <id>",
	Short: "update the status of a task with the id",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store := contextkey.GetStore(cmd)
		if store == nil {
			return errors.New("we couldn't connect to mysql")
		}

		db := store.DB

		idString := args[0]

		id, err := strconv.Atoi(idString)
		if err != nil {
			return errors.New("the id might be a number")
		}

		query := "UPDATE task SET status_task = 'in-progress' WHERE id_task = ?;"
		result, err := db.Exec(
			query,
			id,
		)
		if err != nil {
			return err
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return errors.New("TASK WAS NOT UPDATED. TRY AGAIN WITH A GOOD ID")
		}
		fmt.Printf("✅ TASK UPDATED SUCCESFULLY WITH ID: %d\n", id)

		return nil

	},
}

func init() {

}
