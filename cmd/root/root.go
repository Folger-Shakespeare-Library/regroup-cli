package root

import (
	"github.com/spf13/cobra"
	"github.com/Folger-Shakespeare-Library/regroup-cli/cmd/channels"
	"github.com/Folger-Shakespeare-Library/regroup-cli/cmd/config"
	"github.com/Folger-Shakespeare-Library/regroup-cli/cmd/contacts"
	"github.com/Folger-Shakespeare-Library/regroup-cli/cmd/groups"
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
