package ecr

import (
	"os"
	"path"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

// getHelperPath returns the path to the cache directory for the helper
func getHelperPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(homeDir, ".sso-ecr-login")
}

func GetLogFilePath() string {
	return filepath.Join(getHelperPath(), "logs", "ecr-login.log")
}

func init() {
	// Create Logs directory
	_ = os.MkdirAll(path.Dir(GetLogFilePath()), 0777)

	// Open log file
	file, err := os.OpenFile(GetLogFilePath(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Errorf("[%s] Failed to log to file, using default stderr, HelperName: %s", HelperName, err)
	} else {
		log.SetOutput(file)
	}
}
