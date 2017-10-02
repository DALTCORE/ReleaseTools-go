package main

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"strings"
	"fmt"
)

type ErrNotFound struct {
	Key string
}

func ExecGit(args ...string) string {
	gitArgs := args
	var stdout bytes.Buffer
	cmd := exec.Command("git", gitArgs...)

	cmd.Stdout = &stdout
	cmd.Stderr = ioutil.Discard
	cmd.Run()

	str := stdout.String()
	str = strings.TrimLeft(str, "* ")
	str = strings.TrimRight(str, "\000")
	str = strings.TrimRight(str, "\n")
	str = strings.TrimRight(str, "\r")
	str = strings.TrimRight(str, "\\lf")

	str = strings.Replace(str, "/", "-", -1)
	str = strings.Replace(str, "\\", "-", -1)

	return str
}

func GetCurrentBranch() string {
	var stdout bytes.Buffer
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")

	cmd.Stdout = &stdout
	cmd.Stderr = ioutil.Discard
	cmd.Run()

	str := stdout.String()
	str = strings.TrimLeft(str, "* ")
	str = strings.TrimRight(str, "\000")
	str = strings.TrimRight(str, "\n")
	str = strings.TrimRight(str, "\r")
	str = strings.TrimRight(str, "\\lf")

	str = strings.Replace(str, "/", "-", -1)
	str = strings.Replace(str, "\\", "-", -1)

	return str
}