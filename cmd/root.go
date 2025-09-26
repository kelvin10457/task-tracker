package cmd

import (
	"github.com/kelvin10457/task-tracker/cmd/delete_task"
	"github.com/kelvin10457/task-tracker/cmd/list_task"
	"github.com/kelvin10457/task-tracker/cmd/markdone_task"
	"github.com/kelvin10457/task-tracker/cmd/update_task"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "task-tracker",
	Short: "This is an app for tracking your tasks with a local db for making the data persistent",
	/*PersistentPreRun: func(cmd *cobra.Command) error{
		you need to connect the db with the system and it's finished
	}*/
}

func init() {
	//here you add all the commands that you created before with the next syntax:
	RootCmd.AddCommand(delete_task.DeleteCmd)
	RootCmd.AddCommand(list_task.ListCmd)
	RootCmd.AddCommand(markdone_task.MarkDoneCmd)
	RootCmd.AddCommand(update_task.UpdateCmd)

}
