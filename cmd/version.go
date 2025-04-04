package cmd

import (
	"fmt"

	"github.com/kyosenergy/docker-credential-sso-ecr-login/internal/ecr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: fmt.Sprintf("Print the version number of %s", ecr.HelperName),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s %s -- HEAD\n", ecr.HelperName, "0.1.0")
	},
}
