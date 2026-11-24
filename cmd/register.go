package cmd

import (
	"fmt"

	"github.com/rafael-vincensi/gitrats-cli/internal/api"
	"github.com/spf13/cobra"
)

var (
	registerUsername string
	registerName     string
	registerEmail    string
	registerPassword string
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Create an account",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		reqBody := map[string]string{
			"username": registerUsername,
			"name":     registerName,
			"email":    registerEmail,
			"password": registerPassword,
		}

		err = client.DoRequest("POST", "/users", reqBody, nil)
		if err != nil {
			return err
		}

		fmt.Println("Account created! Run 'gitrats login'")
		return nil
	},
}

func init() {
	registerCmd.Flags().StringVarP(&registerUsername, "username", "u", "", "User username")
	registerCmd.Flags().StringVarP(&registerName, "name", "n", "", "User name")
	registerCmd.Flags().StringVarP(&registerEmail, "email", "e", "", "User email")
	registerCmd.Flags().StringVarP(&registerPassword, "password", "p", "", "User password")
	_ = registerCmd.MarkFlagRequired("email")
	_ = registerCmd.MarkFlagRequired("password")

	rootCmd.AddCommand(registerCmd)
}
