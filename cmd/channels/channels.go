package channels

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "channels",
	Short: "Manage channels",
}

func init() {
	Cmd.AddCommand(listCmd)
	Cmd.AddCommand(contactsCmd)
}
