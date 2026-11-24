package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/rafael-vincensi/gitrats-cli/internal/api"
	"github.com/spf13/cobra"
)

var meCmd = &cobra.Command{
	Use:   "me",
	Short: "Display your profile, streak, and weekly GitHub commits",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		if client.Token == "" {
			return fmt.Errorf("you are not logged in. Run 'gitrats login' first")
		}

		var profile struct {
			Email          string `json:"email"`
			Username       string `json:"username"`
			Name           string `json:"name"`
			Bio            string `json:"bio"`
			AvatarUrl      string `json:"avatarUrl"`
			GithubId       string `json:"githubId"`
			GithubUsername string `json:"githubUsername"`
			CurrentStreak  *int   `json:"currentStreak"`
			HighestStreak  *int   `json:"highestStreak"`
			TotalCheckIns  *int   `json:"totalCheckIns"`
			WeeklyCommits  *int   `json:"weeklyCommits"`
		}

		err = client.DoRequest("GET", "/users/me", nil, &profile)
		if err != nil {
			return err
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintf(w, "\nUsername:\t%s\n", profile.Username)
		fmt.Fprintf(w, "Name:\t%s\n", profile.Name)
		fmt.Fprintf(w, "Bio:\t%s\n", profile.Bio)
		fmt.Fprintf(w, "Github username:\t%s\n", profile.GithubUsername)

		if profile.CurrentStreak != nil {
			fmt.Fprintf(w, "Current streak:\t%d\n", *profile.CurrentStreak)
		}

		if profile.HighestStreak != nil {
			fmt.Fprintf(w, "Highest Streak:\t%d\n", *profile.HighestStreak)
		}

		if profile.TotalCheckIns != nil {
			fmt.Fprintf(w, "Total checkIns:\t%d\n", *profile.TotalCheckIns)
		}

		if profile.WeeklyCommits != nil {
			fmt.Fprintf(w, "Weekly commits:\t%d\n", *profile.WeeklyCommits)
		}

		return w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(meCmd)
}
