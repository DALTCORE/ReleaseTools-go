package main

import (
	"context"
	"github.com/fatih/color"
	"github.com/google/go-github/github"
	"os"
	"runtime"
)

func getAssetByOs() string {
	if runtime.GOOS == "windows" {
		return "release-tools-windows-amd64.exe"
	}

	if runtime.GOOS == "linux" {
		return "release-tools-linux-amd64"
	}

	if runtime.GOOS == "darwin" {
		return "release-tools-darwin-amd64"
	}

	return "unknown"
}

func getAssetDownloadUrl() string {
	ctx := context.Background()
	client := github.NewClient(nil)

	opt := github.ListOptions{
		Page:    1,
		PerPage: 1,
	}

	tags, _, _ := client.Repositories.ListTags(ctx, "DALTCORE", "ReleaseTools", &opt)
	for _, element := range tags {
		if RTVERSION != element.GetName() {
			color.Red("%v", "Cannot update release-tool. Current version "+RTVERSION+" equals update version " + element.GetName())
			os.Exit(128)
		}

		return "https://github.com/DALTCORE/ReleaseTools/releases/download/" + element.GetName() + "/" + getAssetByOs()
	}

	return ""
}

func checkIfUpdateAvailible() bool {
	ctx := context.Background()
	client := github.NewClient(nil)

	opt := github.ListOptions{
		Page:    1,
		PerPage: 1,
	}

	tags, _, _ := client.Repositories.ListTags(ctx, "DALTCORE", "ReleaseTools", &opt)
	for _, element := range tags {
		if RTVERSION == "{{VERSION}}" {
			return false
		}
		return RTVERSION != element.GetName()
	}

	return false
}

func getGitVersion() string {
	ctx := context.Background()
	client := github.NewClient(nil)

	opt := github.ListOptions{
		Page:    1,
		PerPage: 1,
	}

	tags, _, _ := client.Repositories.ListTags(ctx, "DALTCORE", "ReleaseTools", &opt)
	for _, element := range tags {
		return element.GetName()
	}

	return "unknown"
}
