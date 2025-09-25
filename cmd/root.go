package cmd

import (
	"github.com/kelvin10457/task-tracker/cmd/parser"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "tokenparser",
	Short: "Parser and validates jwts",
}

func init() {
	//here you add all the commands that you created before with the next syntax:
	RootCmd.AddCommand(parser.ParserCmd)
}
