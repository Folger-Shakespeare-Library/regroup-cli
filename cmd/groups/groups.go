package groups

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "groups",
	Short: "Manage groups",
}

func init() {
	Cmd.AddCommand(listCmd)
	Cmd.AddCommand(contactsCmd)
}
