package nxenv

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/charmbracelet/log"
)

var Logger = log.NewWithOptions(os.Stderr, log.Options{
	ReportTimestamp: true,
	Prefix:          ":",
})

func DetectSystem() {
	if runtime.GOOS != "linux" || runtime.GOARCH != "amd64" {
		Logger.Error("Unsupported operating system detected")
		Logger.Fatal("Support available for linux-amd64 only")
	}
}

func ClearScreen() {
	_, err := fmt.Print("\033[2J\033[H")
	if err == nil {
		return
	}

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

func GetUserHomePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	info, err := os.Stat(homePath)
	if err != nil {
		return "", err
	}

	if !info.IsDir() {
		return "", fmt.Errorf("not a directory: %s", homePath)
	}

	return homePath, nil
}
