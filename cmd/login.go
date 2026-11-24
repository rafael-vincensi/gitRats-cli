package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/rafael-vincensi/gitrats-cli/internal/api"
	"github.com/rafael-vincensi/gitrats-cli/internal/config"
	"github.com/spf13/cobra"
)

var (
	loginEmail     string
	loginPassword  string
	loginGithub    bool
	githubClientID string
)
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate with email and password",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		if loginGithub {

			if githubClientID == "" {
				return fmt.Errorf("GitHub Client ID is required (use --client-id or set GITRATS_GITHUB_CLIENT_ID)")
			}
			authURL := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s", githubClientID)
			fmt.Println("Opening browser to authenticate:", authURL)

			if runtime.GOOS == "windows" {
				exec.Command("rundll32", "url.dll,FileProtocolHandler", authURL).Start()
			} else if runtime.GOOS == "darwin" {
				exec.Command("open", authURL).Start()
			} else {
				exec.Command("xdg-open", authURL).Start()
			}

			fmt.Print("Paste the token from your browser: ")
			var pastedToken string
			fmt.Scanln(&pastedToken)
			pastedToken = strings.TrimSpace(pastedToken)
			if pastedToken == "" {
				return fmt.Errorf("Empty token!")
			}

			err = config.Save(config.Config{
				Token:  pastedToken,
				APIURL: client.BaseURL,
			})
			if err != nil {
				return err
			}

			println("Github login successful!")
			return nil
		}

		if loginEmail == "" || loginPassword == "" {
			return fmt.Errorf("email and password are required (or use --github)")
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
	loginCmd.Flags().BoolVar(&loginGithub, "github", false, "Login using GitHub OAuth")
	loginCmd.Flags().StringVar(&githubClientID, "client-id", os.Getenv("GITRATS_GITHUB_CLIENT_ID"), "GitHub OAuth Client ID")

	rootCmd.AddCommand(loginCmd)
}
