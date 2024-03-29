package util

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func ExecuteCmdWithPrint(name string, args []string, dir string) error {
	command := exec.Command(name, args...)
	command.Dir = dir
	var out bytes.Buffer
	command.Stdout = &out
	if err := command.Run(); err != nil {
		return err
	}

	fmt.Printf("%s", out.String())
	return nil
}

func CheckGitExistsOnDir(dir string) bool {
	if _, err := os.Stat(fmt.Sprintf("%s/.git", dir)); os.IsNotExist(err) {
		return false
	}
	return true
}

func ExecuteCmd(name string, args []string, dir string) error {
	command := exec.Command(name, args...)
	command.Dir = dir
	var out bytes.Buffer
	command.Stdout = &out
	if err := command.Run(); err != nil {
		return err
	}

	return nil
}

func GetCurrentDir() (string, error) {
	return os.Getwd()
}
