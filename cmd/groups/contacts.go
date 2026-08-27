package groups

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"regroup/internal/api"
	"regroup/internal/cfg"
)

var (
	contactsFlagOutput string
)

var contactsCmd = &cobra.Command{
	Use:   "contacts <group-slug>",
	Short: "List contacts in a group",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cfg.Load(); err != nil {
			return err
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		resp, err := client.ListContacts(api.ListContactsParams{
			Group: args[0],
			All:   true,
		})
		if err != nil {
			return err
		}

		scoped := resp.Scoped()

		switch contactsFlagOutput {
		case "table":
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "FIRST\tLAST\tEMAIL\tUSERNAME\tDBID\tPHONE")
			for _, c := range scoped.Results {
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n",
					c.FirstName, c.LastName, c.Email, c.Username, c.DatabaseID,
					strings.Join(c.PhoneNumbers, ", "))
			}
			return w.Flush()
		default:
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")
			return enc.Encode(scoped.Results)
		}
	},
}

func init() {
	contactsCmd.Flags().StringVarP(&contactsFlagOutput, "output", "o", "json", "Output format (json, table)")
}
