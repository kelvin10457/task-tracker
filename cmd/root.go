package cmd

import (
	"context"
	"fmt"
	"log"

	addtask "github.com/kelvin10457/task-tracker/cmd/add_task"
	"github.com/kelvin10457/task-tracker/cmd/delete_task"
	"github.com/kelvin10457/task-tracker/cmd/list_task"
	"github.com/kelvin10457/task-tracker/cmd/markdone_task"
	"github.com/kelvin10457/task-tracker/cmd/markinprogress_task"
	"github.com/kelvin10457/task-tracker/cmd/update_task"
	"github.com/kelvin10457/task-tracker/internal/config"
	contextkey "github.com/kelvin10457/task-tracker/internal/contextKey"
	"github.com/kelvin10457/task-tracker/internal/db"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "task-tracker",
	Short: "This is an app for tracking your tasks with a local db for making the data persistent",
	//with the next field of the cobra.Command struct (PersistentPreRunE) you'll run this function before
	//every function related with a command child of this command (the root command)
	//it's called persistent because it's like inheritance (we use it with the context) to all the sub-commands
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

		//We use config.Load() that returns an instance of struct to cfg that contains the env variables
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		//We use Open from the package db that receives a string DSN (domain source name)
		//We get the dsn from concat elements of the cfg instance of a structure
		store, err := db.Open(cfg.DSN())
		if err != nil {
			return fmt.Errorf("we couldn't connect to mysql because: %v", err)
		}

		//we take the context as it was and we agregate a key and a value (currentContext,key,value)
		//we use the empty struct as the  key only because we need that key to be unique

		//this is useful because cobra copies the context from a command to all his subcommands
		//and we can use the store in anther subcommands
		contextkey.SetStore(cmd, store)

		return nil
	},
}

// this exists just for being a unique key with the valueof the struct store

func Execute() {
	//When we execute the context we run PersistentPreRunE and the following commands
	if err := RootCmd.ExecuteContext(context.Background()); err != nil {
		log.Fatal(err)
	}

	//When we already ran all the functions related to the commands (at the end) We close the db
	store := contextkey.GetStore(RootCmd)
	if store != nil {
		_ = store.Close()
	}
}

func init() {
	//here you add all the subcommands (the ones that preceeds the root command)
	//for the root command that you created before with the next syntax:
	//RootCmd.AddCommand(package.nameOfTheVariable)
	RootCmd.AddCommand(addtask.AddCmd)
	RootCmd.AddCommand(delete_task.DeleteCmd)
	RootCmd.AddCommand(list_task.ListCmd)
	RootCmd.AddCommand(markdone_task.MarkDoneCmd)
	RootCmd.AddCommand(markinprogress_task.MarkInProgressCmd)
	RootCmd.AddCommand(update_task.UpdateCmd)
}
