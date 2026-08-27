package channels

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"regroup/internal/api"
	"regroup/internal/cfg"
)

var (
	listFlagName     string
	listFlagPage     int
	listFlagCount    int
	listFlagSortBy   string
	listFlagSortType string
	listFlagAll      bool
	listFlagOutput   string
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List channels",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cfg.Load(); err != nil {
			return err
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		resp, err := client.ListChannels(api.ListChannelsParams{
			Name:     listFlagName,
			Page:     listFlagPage,
			Count:    listFlagCount,
			SortBy:   listFlagSortBy,
			SortType: listFlagSortType,
			All:      listFlagAll,
		})
		if err != nil {
			return err
		}

		switch listFlagOutput {
		case "table":
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "ID\tNAME\tDESCRIPTION\tSUBSCRIBERS\tPRIVACY\tPERMISSION")
			for _, ch := range resp.Results {
				fmt.Fprintf(w, "%d\t%s\t%s\t%d\t%s\t%s\n",
					ch.ID, ch.Name, ch.Description, ch.Subscribers, ch.AdminPrivacy, ch.Permission)
			}
			return w.Flush()
		default:
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")
			return enc.Encode(resp.Results)
		}
	},
}

func init() {
	listCmd.Flags().StringVar(&listFlagName, "name", "", "Filter by name")
	listCmd.Flags().IntVar(&listFlagPage, "page", 0, "Page number")
	listCmd.Flags().IntVar(&listFlagCount, "count", 0, "Results per page")
	listCmd.Flags().StringVar(&listFlagSortBy, "sort-by", "", "Sort field (name, description, subscribers, sharing, admin_privacy)")
	listCmd.Flags().StringVar(&listFlagSortType, "sort-type", "", "Sort direction (ASC, DESC)")
	listCmd.Flags().BoolVar(&listFlagAll, "all", true, "Fetch all channels (use --all=false to paginate)")
	listCmd.Flags().StringVarP(&listFlagOutput, "output", "o", "json", "Output format (json, table)")
}
