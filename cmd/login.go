package cmd

import (
	"fmt"

	"github.com/rafael-vincensi/gitrats-cli/internal/api"
	"github.com/rafael-vincensi/gitrats-cli/internal/config"
	"github.com/spf13/cobra"
)

var (
	loginEmail    string
	loginPassword string
)
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate with email and password",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		reqBody := map[string]string{
			"email":    loginEmail,
			"password": loginPassword,
		}

		var res struct {
			Token string `json:"token"`
		}

		err = client.DoRequest("POST", "/auth/login", reqBody, &res)
		if err != nil {
			return err
		}

		err = config.Save(config.Config{
			Token:  res.Token,
			APIURL: client.BaseURL,
		})
		if err != nil {
			return err
		}

		fmt.Println("Login successful!")

		return nil
	},
}

func init() {
	loginCmd.Flags().StringVarP(&loginEmail, "email", "e", "", "User email")
	loginCmd.Flags().StringVarP(&loginPassword, "password", "p", "", "User password")
	_ = loginCmd.MarkFlagRequired("email")
	_ = loginCmd.MarkFlagRequired("password")

	rootCmd.AddCommand(loginCmd)
}
