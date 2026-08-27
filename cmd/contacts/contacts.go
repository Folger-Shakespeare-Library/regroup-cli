package contacts

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "contacts",
	Short: "Manage contacts",
}

func init() {
	Cmd.AddCommand(addCmd)
	Cmd.AddCommand(listCmd)
	Cmd.AddCommand(removeCmd)
}
