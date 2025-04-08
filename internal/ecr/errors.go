package ecr

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func handleError(err error) {
	if err != nil {
		out := fmt.Sprintf("[%s] %s", HelperName, err)

		log.Error(out)
		fmt.Println(out)

		os.Exit(1)
	}
}
