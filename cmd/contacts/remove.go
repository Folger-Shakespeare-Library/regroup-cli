package contacts

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/Folger-Shakespeare-Library/regroup-cli/internal/api"
	"github.com/Folger-Shakespeare-Library/regroup-cli/internal/cfg"
)

var (
	removeFlagEmail  string
	removeFlagGroups []string
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a contact from group(s)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if removeFlagEmail == "" {
			return fmt.Errorf("--email is required")
		}
		if len(removeFlagGroups) == 0 {
			return fmt.Errorf("--group is required")
		}

		if err := cfg.Load(); err != nil {
			return err
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		var groupIDs []string
		for _, slug := range removeFlagGroups {
			id, err := client.ResolveGroupID(slug)
			if err != nil {
				return err
			}
			groupIDs = append(groupIDs, strconv.Itoa(id)+"||Remove")
		}

		resp, err := client.AddContact(api.AddContactParams{
			Email:    removeFlagEmail,
			GroupIDs: strings.Join(groupIDs, ";"),
			UserType: "contact",
		})
		if err != nil {
			return err
		}

		if len(resp.ErrorMessages) > 0 {
			return fmt.Errorf("%s", strings.Join(resp.ErrorMessages, "; "))
		}
		return nil
	},
}

func init() {
	removeCmd.Flags().StringVar(&removeFlagEmail, "email", "", "Email address (required)")
	removeCmd.Flags().StringSliceVar(&removeFlagGroups, "group", nil, "Group slug(s) to remove from (required, repeatable)")
}
