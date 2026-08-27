package root

import (
	"github.com/spf13/cobra"
	"regroup/cmd/channels"
	"regroup/cmd/config"
	"regroup/cmd/contacts"
	"regroup/cmd/groups"
)

var Cmd = &cobra.Command{
	Use:   "regroup",
	Short: "CLI for the Regroup Mass Notification API",
}

func init() {
	Cmd.AddCommand(channels.Cmd)
	Cmd.AddCommand(config.Cmd)
	Cmd.AddCommand(contacts.Cmd)
	Cmd.AddCommand(groups.Cmd)
}
