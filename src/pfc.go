package main

import (
	"bytes"
	"github.com/fatih/color"
	"os"
	"path/filepath"
)

func RunChecks() bool {

	if !CheckIfGitExists() {
		return false
	}

	if !CheckIfBranchIsSet() {
		return false
	}

	if !CheckIfReleaseToolInit() {
		return false
	}

	if !CheckIfReleaseToolDirIsReady() {
		return false
	}

	if !CheckIfReleaseToolStubDirIsReady() {
		return false
	}

	if !CheckIfReleaseToolPlaybookDirIsReady() {
		return false
	}

	if !CheckIfChangelogDirsAreReady() {
		return false
	}

	if !CheckIfChangelogReleasedDirIsReady() {
		return false
	}

	if !CheckIfChangelogUnreleasedDirIsReady() {
		return false
	}

	return true
}

func CheckIfGitExists() bool {

	var directory bytes.Buffer

	directory.WriteString(CurrentWorkingDirectory())
	directory.WriteString(DirSep())
	directory.WriteString(".git")

	var osCompatibleDir string = filepath.FromSlash(directory.String())

	if _, err := os.Stat(osCompatibleDir); os.IsNotExist(err) {
		color.Red("Could not find " + osCompatibleDir + " directory. You're in the right place?")
		return false
	}

	return true
}

func CheckIfBranchIsSet() bool {
	if len(ExecGit("branch")) > 3 {
		return true
	}

	color.Red("No... This cannot be done. A branch must be availible!")
	return false
}

func CheckIfReleaseToolInit() bool {
	var directory bytes.Buffer

	directory.WriteString(CurrentWorkingDirectory())
	directory.WriteString(DirSep())
	directory.WriteString(".release-tool")

	var osCompatibleDir string = filepath.FromSlash(directory.String())

	if _, err := os.Stat(osCompatibleDir); os.IsNotExist(err) {
		color.Red("Could not find " + osCompatibleDir + " file. Call for release manager.")
		return false
	}

	return true
}

func CheckIfReleaseToolHomeInit() bool {
	var osCompatibleDir string = filepath.FromSlash(ReleaseToolsHomeConfigFile())

	if _, err := os.Stat(osCompatibleDir); os.IsNotExist(err) {
		return false
	}

	return true
}

func CheckIfReleaseToolDirIsReady() bool {
	var directory bytes.Buffer

	directory.WriteString(CurrentWorkingDirectory())
	directory.WriteString(DirSep())
	directory.WriteString(".release-tools")

	var osCompatibleDir string = filepath.FromSlash(directory.String())

	if _, err := os.Stat(osCompatibleDir); os.IsNotExist(err) {
		color.Red("Could not find " + osCompatibleDir + " directory. Call for release manager.")
		return false
	}

	return true
}

func CheckIfReleaseToolStubDirIsReady() bool {
	var directory bytes.Buffer

	directory.WriteString(CurrentWorkingDirectory())
	directory.WriteString(DirSep())
	directory.WriteString(".release-tools")
	directory.WriteString(DirSep())
	directory.WriteString("stubs")

	var osCompatibleDir string = filepath.FromSlash(directory.String())

	if _, err := os.Stat(osCompatibleDir); os.IsNotExist(err) {
		color.Red("Could not find " + osCompatibleDir + " directory. Call for release manager.")
		return false
	}

	return true
}

func CheckIfReleaseToolPlaybookDirIsReady() bool {
	var directory bytes.Buffer

	directory.WriteString(CurrentWorkingDirectory())
	directory.WriteString(DirSep())
	directory.WriteString(".release-tools")
	directory.WriteString(DirSep())
	directory.WriteString("playbooks")

	var osCompatibleDir string = filepath.FromSlash(directory.String())

	if _, err := os.Stat(osCompatibleDir); os.IsNotExist(err) {
		color.Red("Could not find " + osCompatibleDir + " directory. Call for release manager.")
		return false
	}

	return true
}

func CheckIfChangelogDirsAreReady() bool {
	var directory bytes.Buffer

	directory.WriteString(ChangelogsDirectory())

	var osCompatibleDir string = filepath.FromSlash(directory.String())

	if _, err := os.Stat(osCompatibleDir); os.IsNotExist(err) {
		color.Red("Could not find " + osCompatibleDir + " directory. Call for release manager.")
		return false
	}

	return true
}

func CheckIfChangelogReleasedDirIsReady() bool {
	var directory bytes.Buffer

	directory.WriteString(ChangelogsDirectory())
	directory.WriteString(DirSep())
	directory.WriteString("released")

	var osCompatibleDir string = filepath.FromSlash(directory.String())

	if _, err := os.Stat(osCompatibleDir); os.IsNotExist(err) {
		color.Red("Could not find " + osCompatibleDir + " directory. Call for release manager.")
		return false
	}

	return true
}

func CheckIfChangelogUnreleasedDirIsReady() bool {
	var directory bytes.Buffer

	directory.WriteString(ChangelogsDirectory())
	directory.WriteString(DirSep())
	directory.WriteString("unreleased")

	var osCompatibleDir string = filepath.FromSlash(directory.String())

	if _, err := os.Stat(osCompatibleDir); os.IsNotExist(err) {
		color.Red("Could not find " + osCompatibleDir + " directory. Call for release manager.")
		return false
	}

	return true
}
