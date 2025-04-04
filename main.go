package main

import (
	"github.com/kyosenergy/docker-credential-sso-ecr-login/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cmd.SetVersionInfo(version, commit, date)
	cmd.Execute()
}
