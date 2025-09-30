package delete_task

import (
	"errors"
	"fmt"
	"strconv"

	contextkey "github.com/kelvin10457/task-tracker/internal/contextKey"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "delete a task with the id",
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
			return errors.New("the id migth be a number")
		}

		query := "DELETE FROM task WHERE id_task = ?;"
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
			return errors.New("TASK WAS NOT DELETED. TRY AGAIN WITH A GOOD ID")
		}

		fmt.Printf("✅ TASK DELETED SUCCESFULLY WITH ID: %d\n", id)

		return nil
	},
}

func init() {

}
