package browser

import (
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/google/shlex"
)

// Reference: https://github.com/cli/cli/blob/2ab073d599/pkg/browser/browser.go
func Command(url string) (*exec.Cmd, error) {
	launcher := os.Getenv("BROWSER")
	if launcher != "" {
		return FromLauncher(launcher, url)
	}
	return ForOS(runtime.GOOS, url), nil
}

func ForOS(goos, url string) *exec.Cmd {
	exe := "open"
	var args []string
	switch goos {
	case "darwin":
		args = append(args, url)
	case "windows":
		exe, _ = exec.LookPath("cmd")
		r := strings.NewReplacer("&", "^&")
		args = append(args, "/c", "start", r.Replace(url))
	default:
		exe = linuxExe()
		args = append(args, url)
	}

	cmd := exec.Command(exe, args...)
	cmd.Stderr = os.Stderr
	return cmd
}

func FromLauncher(launcher, url string) (*exec.Cmd, error) {
	args, err := shlex.Split(launcher)
	if err != nil {
		return nil, err
	}

	exe, err := exec.LookPath(args[0])
	if err != nil {
		return nil, err
	}

	args = append(args, url)
	cmd := exec.Command(exe, args[1:]...)
	cmd.Stderr = os.Stderr
	return cmd, nil
}

func linuxExe() string {
	exe := "xdg-open"

	_, err := exec.LookPath(exe)
	if err != nil {
		_, err := exec.LookPath("wslview")
		if err == nil {
			exe = "wslview"
		}
	}

	return exe
}
