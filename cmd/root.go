package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "docker-credential-sso-ecr-login",
	Short: "docker-credential-sso-ecr-login is a helper for Docker to use AWS SSO credentials",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is a credential helper for Docker to use AWS SSO credentials.")
		fmt.Println("Please use the `docker-credential-sso-ecr-login` command with the appropriate subcommands.")
	},
}

func SetVersionInfo(version, commit, date string) {
	rootCmd.Version = fmt.Sprintf("%s (Built on %s from Git SHA %s)", version, date, commit)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
