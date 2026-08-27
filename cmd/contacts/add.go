package contacts

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"regroup/internal/api"
	"regroup/internal/cfg"
)

var (
	addFlagFirstName  string
	addFlagLastName   string
	addFlagEmail      string
	addFlagPhone      string
	addFlagUsername   string
	addFlagDatabaseID string
	addFlagGroups     []string
	addFlagUserType   string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a contact",
	RunE: func(cmd *cobra.Command, args []string) error {
		if addFlagEmail == "" {
			return fmt.Errorf("--email is required")
		}
		if len(addFlagGroups) == 0 {
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
		for _, slug := range addFlagGroups {
			id, err := client.ResolveGroupID(slug)
			if err != nil {
				return err
			}
			groupIDs = append(groupIDs, strconv.Itoa(id))
		}

		resp, err := client.AddContact(api.AddContactParams{
			FirstName:  addFlagFirstName,
			LastName:   addFlagLastName,
			Email:      addFlagEmail,
			Phone:      addFlagPhone,
			Username:   addFlagUsername,
			DatabaseID: addFlagDatabaseID,
			GroupIDs:   strings.Join(groupIDs, ";"),
			UserType:   addFlagUserType,
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
	addCmd.Flags().StringVar(&addFlagFirstName, "first-name", "", "First name")
	addCmd.Flags().StringVar(&addFlagLastName, "last-name", "", "Last name")
	addCmd.Flags().StringVar(&addFlagEmail, "email", "", "Email address (required)")
	addCmd.Flags().StringVar(&addFlagPhone, "phone", "", "Phone number")
	addCmd.Flags().StringVar(&addFlagUsername, "username", "", "Username")
	addCmd.Flags().StringVar(&addFlagDatabaseID, "databaseid", "", "Database ID")
	addCmd.Flags().StringSliceVar(&addFlagGroups, "group", nil, "Group slug(s) (required, repeatable)")
	addCmd.Flags().StringVar(&addFlagUserType, "type", "contact", "User type (contact, admin)")
}
