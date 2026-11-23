package cmd

import (
	"fmt"
	"strings"

	"github.com/rafael-vincensi/gitrats-cli/internal/api"
	"github.com/spf13/cobra"
)

var (
	checkinTitle       string
	checkinDescription string
	checkinGroupCode   string
	checkinImage       string
)

var checkinCmd = &cobra.Command{
	Use:   "checkin",
	Short: "Create new checkin with title, description and image",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		if client.Token == "" {
			return fmt.Errorf("you are not logged in. Run 'gitrats login' first")
		}

		var groups []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			Code string `json:"code"`
		}

		err = client.DoRequest("GET", "/groups", nil, &groups)
		if err != nil {
			return err
		}

		var resolvedGroupID string
		var groupName string

		for _, g := range groups {
			if strings.EqualFold(g.Code, checkinGroupCode) {
				resolvedGroupID = g.ID
				groupName = g.Name
				break
			}
		}

		if resolvedGroupID == "" {
			return fmt.Errorf("group '%s' not found", checkinGroupCode)
		}

		payload := struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			GroupID     string `json:"groupId"`
		}{
			Title:       checkinTitle,
			Description: checkinDescription,
			GroupID:     resolvedGroupID,
		}

		err = client.DoMultipartJSON("/checkins", payload, "file", checkinImage, nil)
		if err != nil {
			return err
		}

		fmt.Printf("Check-in created successfully to group '%s'\n", groupName)
		return nil

	},
}

func init() {
	checkinCmd.Flags().StringVarP(&checkinTitle, "title", "t", "", "Checkin title")
	checkinCmd.Flags().StringVarP(&checkinDescription, "description", "d", "", "Checkin description")
	checkinCmd.Flags().StringVarP(&checkinGroupCode, "group", "g", "", "Checkin group code")
	checkinCmd.Flags().StringVarP(&checkinImage, "image", "i", "", "Checkin image")
	_ = checkinCmd.MarkFlagRequired("title")
	_ = checkinCmd.MarkFlagRequired("group")
	_ = checkinCmd.MarkFlagRequired("image")

	rootCmd.AddCommand(checkinCmd)
}
