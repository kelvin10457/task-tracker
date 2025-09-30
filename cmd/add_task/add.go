package addtask

import (
	"errors"
	"fmt"

	contextkey "github.com/kelvin10457/task-tracker/internal/contextKey"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add <description>",
	Short: "create a task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		store := contextkey.GetStore(cmd)
		if store == nil {
			return errors.New("we couldn't connect to mysql")
		}

		db := store.DB

		description := args[0]
		if description == "" {
			return errors.New("you might enter a description")
		}
		query := `
		INSERT INTO task
		(description_task,status_task,created_at,updated_at)
		VALUES (?,?,NOW(),null);`

		result, err := db.Exec(
			query,
			description,
			"todo",
		)
		if err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		fmt.Printf("✅ TASK CREATED SUCCESFULLY WITH ID %d:", id)

		return nil
	},
}
