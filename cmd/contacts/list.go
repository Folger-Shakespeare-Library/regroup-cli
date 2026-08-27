package contacts

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
	flagEmail      string
	flagUsername   string
	flagDatabaseID string
	flagPhone      string
	flagPage       int
	flagCount      int
	flagSortBy     string
	flagSortType   string
	flagAll        bool
	flagOutput     string
	flagGroup      string
	flagChannel    string
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List contacts",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cfg.Load(); err != nil {
			return err
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		resp, err := client.ListContacts(api.ListContactsParams{
			Email:      flagEmail,
			Username:   flagUsername,
			DatabaseID: flagDatabaseID,
			Phone:      flagPhone,
			Page:       flagPage,
			Count:      flagCount,
			SortBy:     flagSortBy,
			SortType:   flagSortType,
			All:        flagAll,
			Group:      flagGroup,
			Channel:    flagChannel,
		})
		if err != nil {
			return err
		}

		scoped := flagGroup != "" || flagChannel != ""

		switch flagOutput {
		case "table":
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			if scoped {
				fmt.Fprintln(w, "FIRST\tLAST\tEMAIL\tUSERNAME\tDBID\tPHONE")
				for _, c := range resp.Results {
					fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n",
						c.FirstName, c.LastName, c.Email, c.Username, c.DatabaseID,
						strings.Join(c.PhoneNumbers, ", "))
				}
			} else {
				fmt.Fprintln(w, "FIRST\tLAST\tEMAIL\tUSERNAME\tDBID\tPHONE\tGROUPS")
				for _, c := range resp.Results {
					fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
						c.FirstName, c.LastName, c.Email, c.Username, c.DatabaseID,
						strings.Join(c.PhoneNumbers, ", "),
						strings.Join(c.Groups, ", "))
				}
			}
			return w.Flush()
		default:
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")
			if scoped {
				return enc.Encode(resp.Scoped().Results)
			}
			return enc.Encode(resp.Results)
		}
	},
}

func init() {
	listCmd.Flags().StringVar(&flagEmail, "email", "", "Filter by email")
	listCmd.Flags().StringVar(&flagUsername, "username", "", "Filter by username")
	listCmd.Flags().StringVar(&flagDatabaseID, "databaseid", "", "Filter by database ID")
	listCmd.Flags().StringVar(&flagPhone, "phone", "", "Filter by phone number")
	listCmd.Flags().IntVar(&flagPage, "page", 0, "Page number")
	listCmd.Flags().IntVar(&flagCount, "count", 0, "Results per page")
	listCmd.Flags().StringVar(&flagSortBy, "sort-by", "", "Sort field (first_name, last_name, email, username, databaseid)")
	listCmd.Flags().StringVar(&flagSortType, "sort-type", "", "Sort direction (ASC, DESC)")
	listCmd.Flags().BoolVar(&flagAll, "all", true, "Fetch all contacts (use --all=false to paginate)")
	listCmd.Flags().StringVarP(&flagOutput, "output", "o", "json", "Output format (json, table)")
	listCmd.Flags().StringVar(&flagGroup, "group", "", "Filter by group slug")
	listCmd.Flags().StringVar(&flagChannel, "channel", "", "Filter by channel slug")
}
