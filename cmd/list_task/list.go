package list_task

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"slices"
	"strings"
	"text/tabwriter"
	"time"

	contextkey "github.com/kelvin10457/task-tracker/internal/contextKey"
	"github.com/spf13/cobra"
)

type task struct {
	id_task          int
	description_task string
	status_task      string
	created_at       time.Time
	updated_at       sql.NullTime
}

// aqui si vamos a tener argumentos porque podemos filtrar las done,todo,in-progress
var ListCmd = &cobra.Command{
	Use:       "list [status]",
	Short:     "list tasks filtering by state if you want",
	ValidArgs: []string{"todo", "in-progress", "done"},
	RunE: func(cmd *cobra.Command, args []string) error {
		store := contextkey.GetStore(cmd)
		if store == nil {
			return fmt.Errorf("we couldn't connect to mysql")
		}

		db := store.DB

		var statusFilter string
		if len(args) > 0 {
			//if there is an arg you take the first one, if it's not statusFilter = ""
			statusFilter = args[0]
		}
		validStatus := []string{"todo", "in-progress", "done"}
		isValidState := slices.Contains(validStatus, statusFilter) //(slice,valueLookedInSlice)
		if !isValidState && statusFilter != "" {
			return fmt.Errorf("the status (%s) is not a valid status", statusFilter)
		}

		//we created this because we need to cancel the query if it don't answer in 5 seconds
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		//we create the variables like this for avoiding a problem with the scope of the variables
		var query string
		var errQuery error
		var rows *sql.Rows
		if statusFilter == "" {
			query = "SELECT id_task,description_task,status_task,created_at,updated_at FROM task ORDER BY created_at DESC;"
			rows, errQuery = db.QueryContext(ctx, query)
		} else {
			query = "SELECT id_task,description_task,status_task,created_at,updated_at FROM task WHERE status_task = ? ORDER BY created_at DESC;"
			rows, errQuery = db.QueryContext(ctx, query, statusFilter)
		}
		if errQuery != nil {
			return fmt.Errorf("error executing the query %s", query)
		}
		defer rows.Close()

		var tasks []task

		for rows.Next() {
			//Scan converts columns read from the database into the following common Go types
			// and special types provided by the sql package
			var t task
			err := rows.Scan(
				&t.id_task,
				&t.description_task,
				&t.status_task,
				&t.created_at,
				&t.updated_at,
			)
			if err != nil {
				return fmt.Errorf("error while iterating rows")
			}
			tasks = append(tasks, t)
		}

		if len(tasks) == 0 {
			if statusFilter == "" {
				fmt.Printf("There is not tasks\n")
			} else {
				fmt.Printf("There is not tasks with the status %s \n", statusFilter)
			}
		}

		printTasksTable(tasks, statusFilter)

		return nil

	},
}

// printTasksTable prints the list of tasks in a neatly formatted table.
func printTasksTable(tasks []task, filter string) {
	// Create tabwriter for column alignment
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	// Title
	if filter != "" {
		fmt.Printf("\n📋 Tasks with Status: %s\n\n", strings.ToUpper(filter))
	} else {
		fmt.Println("\n📋 All Tasks")
	}

	// Header
	fmt.Fprintln(w, "ID\tDESCRIPTION\tSTATUS\tCREATED\tUPDATED")
	fmt.Fprintln(w, "──\t───────────\t──────\t──────\t───────────")

	// Rows
	for _, t := range tasks {
		// Format update date, handling NULL (sql.NullTime)
		updatedStr := "N/A"
		// NOTE: Changed format string to 2006-01-02 15:04 for standard Go format
		if t.updated_at.Valid {
			updatedStr = t.updated_at.Time.Format("2000-00-00 12:00")
		}

		// Get emoji based on status
		statusEmoji := getStatusEmoji(t.status_task)

		fmt.Fprintf(w, "%d\t%s\t%s %s\t%s\t%s\n",
			t.id_task,
			truncateString(t.description_task, 40),
			statusEmoji,
			t.status_task,
			t.created_at.Format("2006-01-02 15:04"),
			updatedStr,
		)
	}

	w.Flush()

	// Summary
	fmt.Printf("\n📊 Total: %d task(s)\n\n", len(tasks))

	// Show statistics if it's the full list
	if filter == "" {
		printStatistics(tasks)
	}
}

// getStatusEmoji returns an emoji for a given status.
func getStatusEmoji(status string) string {
	switch status {
	case "done":
		return "✅" // Neutral checkmark
	case "todo":
		return "📝" // Neutral empty box
	case "in-progress":
		return "🔄" // Neutral arrow/progression
	default:
		return "❓"
	}
}

// truncateString shortens a string if it exceeds maxLen.
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// printStatistics shows task completion statistics.
func printStatistics(tasks []task) {
	var todo, inProgress, done int

	for _, t := range tasks {
		switch t.status_task {
		case "todo":
			todo++
		case "in-progress":
			inProgress++
		case "done":
			done++
		}
	}

	fmt.Println("📈 Statistics:")
	fmt.Printf("   📝 TODO: %d\n", todo)
	fmt.Printf("   📝 IN-PROGRESS: %d\n", inProgress)
	fmt.Printf("   ✅ DONE: %d\n", done)

	if len(tasks) > 0 {
		percentage := float64(done) / float64(len(tasks)) * 100
		fmt.Printf("   📊 Completion Rate: %.1f%%\n", percentage)
	}
}

func init() {

}
