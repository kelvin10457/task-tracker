package update_task

import (
	"errors"
	"fmt"
	"strconv"

	contextkey "github.com/kelvin10457/task-tracker/internal/contextKey"
	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:   "update <id> <newDescription>",
	Short: "update the description of a task",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		store := contextkey.GetStore(cmd)
		if store == nil {
			return errors.New("we couldn't connect to mysql")
		}

		db := store.DB

		idString := args[0]
		description := args[1]

		id, err := strconv.Atoi(idString)
		if err != nil {
			return errors.New("the id might be a number")
		}
		if description == "" {
			return errors.New("you might enter a description")
		}

		query := "UPDATE task SET description_task = ?,created_at = NOW() WHERE id_task = ?;"
		result, err := db.Exec(
			query,
			description,
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
