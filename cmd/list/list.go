package list

import "github.com/spf13/cobra"

//aqui si vamos a tener argumentos porque podemos filtrar las done,todo,in-progress
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all ",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {

}
