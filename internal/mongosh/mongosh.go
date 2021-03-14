package mongosh

import (
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

const (
	mongoShLinux   = "mongosh"
	mongoShWindows = "mongosh.exe"
	windows        = "windows"
)

func FindBinaryInPath() string {
	binary := mongoShLinux
	if runtime.GOOS == windows {
		binary = mongoShWindows
	}

	if path, err := exec.LookPath(binary); err == nil {
		return path
	}

	return ""
}

func Run(binary, mongoURI, username, password string) error {
	args := []string{"mongosh", mongoURI, "-u", username, "-p", password}
	env := os.Environ()
	err := syscall.Exec(binary, args, env)

	if err != nil {
		return err
	}

	return nil
}
