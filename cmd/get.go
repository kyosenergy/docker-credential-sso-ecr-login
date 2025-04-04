package cmd

import (
	"github.com/kyosenergy/docker-credential-sso-ecr-login/internal/ecr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCommand)
}

var getCommand = &cobra.Command{
	Use:   "get",
	Short: "Get ECR authentication token",
	Long:  `Get docker authentication token for AWS Elastic container registry via AWS SSO`,
	Run: func(cmd *cobra.Command, args []string) {
		ecr.GetAuthCredentials()
	},
}
