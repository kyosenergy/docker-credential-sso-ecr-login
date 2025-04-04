package cmd

import (
	"fmt"

	"github.com/kyosenergy/docker-credential-sso-ecr-login/internal/ecr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(showCachePathCmd())
}

func showCachePathCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "logs",
		Short: "Show logs path",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("[%s] Logs path: %s\n", ecr.HelperName, ecr.GetLogFilePath())
		},
	}
}
