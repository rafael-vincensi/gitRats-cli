package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/rafael-vincensi/gitrats-cli/internal/api"
	"github.com/spf13/cobra"
)

var groupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "List all groups you belong to and their weekly commits",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		if client.Token == "" {
			return fmt.Errorf("you are not logged in. Run 'gitrats login' first")
		}

		var groups []struct {
			Name          string `json:"name"`
			Code          string `json:"code"`
			MemberCount   int    `json:"memberCount"`
			TotalCheckIns int64  `json:"totalCheckIns"`
			TotalCommits  int64  `json:"totalCommits"`
		}

		err = client.DoRequest("GET", "/groups", nil, &groups)
		if err != nil {
			return err
		}

		if len(groups) == 0 {
			fmt.Println("You don't belong to any groups yet!")
			return nil
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "NAME\tCODE\tCHECK-INS\tCOMMITS")
		for _, g := range groups {
			fmt.Fprintf(w, "%s\t%s\t%d\t%d\n", g.Name, g.Code, g.TotalCheckIns, g.TotalCommits)
		}

		return w.Flush()
	},
}

var groupsJoinCmd = &cobra.Command{
	Use:   "join [invite-code]",
	Short: "Join a group using its invite code",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}
		if client.Token == "" {
			return fmt.Errorf("you are not logged in. Run 'gitrats login' first")
		}

		var group struct {
			Name string `json:"name"`
		}

		code := args[0]

		err = client.DoRequest("POST", fmt.Sprintf("/groups/%s/join", code), nil, &group)
		if err != nil {
			return err
		}

		fmt.Printf("You joined the group '%s'\n", group.Name)
		return nil
	},
}

func init() {
	groupsCmd.AddCommand(groupsJoinCmd)
	rootCmd.AddCommand(groupsCmd)
}
