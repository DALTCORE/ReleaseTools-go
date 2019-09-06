package main

import (
	"bytes"
	"github.com/mitchellh/go-homedir"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func CurrentExecutablePath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func CurrentWorkingDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func ChangelogsDirectory() string {
	var directory bytes.Buffer

	directory.WriteString(CurrentWorkingDirectory())
	directory.WriteString(DirSep())
	directory.WriteString("changelogs")

	var osCompatibleDir string = filepath.FromSlash(directory.String())

	dir, err := filepath.Abs(osCompatibleDir)
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func ChangelogReleasedDirectory() string {
	var directory bytes.Buffer

	directory.WriteString(ChangelogsDirectory())
	directory.WriteString(DirSep())
	directory.WriteString("released")

	var osCompatibleDir string = filepath.FromSlash(directory.String())

	dir, err := filepath.Abs(osCompatibleDir)
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func ChangelogUnreleasedDirectory() string {
	var directory bytes.Buffer

	directory.WriteString(ChangelogsDirectory())
	directory.WriteString(DirSep())
	directory.WriteString("unreleased")

	var osCompatibleDir string = filepath.FromSlash(directory.String())

	dir, err := filepath.Abs(osCompatibleDir)
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func ReleaseToolsConfigFile() string {
	var directory bytes.Buffer

	directory.WriteString(CurrentWorkingDirectory())
	directory.WriteString(DirSep())
	directory.WriteString(".release-tool")

	var osCompatibleDir string = filepath.FromSlash(directory.String())

	dir, err := filepath.Abs(osCompatibleDir)
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func ReleaseToolsHomeConfigFile() string {
	rtFileHome, _ := homedir.Dir()
	rtFileHome = rtFileHome + DirSep() + ".release-tool"

	var osCompatibleDir string = filepath.FromSlash(rtFileHome)

	dir, err := filepath.Abs(osCompatibleDir)
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func ReleaseToolDirectory() string {
	var directory bytes.Buffer

	directory.WriteString(CurrentWorkingDirectory())
	directory.WriteString(DirSep())
	directory.WriteString(".release-tools")

	var osCompatibleDir string = filepath.FromSlash(directory.String())

	dir, err := filepath.Abs(osCompatibleDir)
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func ReleaseToolStubDirectory() string {
	var directory bytes.Buffer

	directory.WriteString(ReleaseToolDirectory())
	directory.WriteString(DirSep())
	directory.WriteString("stubs")

	var osCompatibleDir string = filepath.FromSlash(directory.String())

	dir, err := filepath.Abs(osCompatibleDir)
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func ReleaseToolPlaybookDirectory() string {
	var directory bytes.Buffer

	directory.WriteString(ReleaseToolDirectory())
	directory.WriteString(DirSep())
	directory.WriteString("playbooks")

	var osCompatibleDir string = filepath.FromSlash(directory.String())

	dir, err := filepath.Abs(osCompatibleDir)
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func ChangelogFile() string {
	var directory bytes.Buffer

	directory.WriteString(CurrentWorkingDirectory())
	directory.WriteString(DirSep())
	directory.WriteString("CHANGELOG.md")

	var osCompatibleDir string = filepath.FromSlash(directory.String())

	dir, err := filepath.Abs(osCompatibleDir)
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func FileExists(filename string) bool {
	var file bytes.Buffer
	file.WriteString(filename)

	var osCompatibleDir string = filepath.FromSlash(file.String())

	if _, err := os.Stat(osCompatibleDir); os.IsNotExist(err) {
		return false
	}

	return true
}

func DirSep() string {
	if runtime.GOOS == "windows" {
		return "\\"
	}

	return "/"
}
