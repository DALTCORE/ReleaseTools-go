package main

import (
	"context"
	"github.com/google/go-github/github"
	"runtime"
	"os"
	"github.com/fatih/color"
	"golang.org/x/oauth2"
)

func getAssetByOs() string {
	if runtime.GOOS == "windows" {
		return "release-tool-windows-amd64.exe"
	}

	if runtime.GOOS == "linux" {
		return "release-tool-linux-amd64"
	}

	if runtime.GOOS == "darwin" {
		return "release-tool-darwin-amd64"
	}

	return "unknown"
}

func getAssetDownloadUrl() string {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ConfigFile().GithubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	opt := github.ListOptions{
		Page:    1,
		PerPage: 1,
	}

	tags, _, _ := client.Repositories.ListTags(ctx, "DALTCORE", "ReleaseTools-go", &opt)

	for _, element := range tags {
		if RTVERSION == element.GetName() {
			color.Red("%v", "Cannot update release-tool. Current version " + RTVERSION + " equals update version " + element.GetName())
			os.Exit(128)
		}

		return "https://github.com/DALTCORE/ReleaseTools-go/releases/download/" + element.GetName() + "/" + getAssetByOs()
	}

	return ""
}

func checkIfUpdateAvailible() bool {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ConfigFile().GithubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	opt := github.ListOptions{
		Page:    1,
		PerPage: 1,
	}


	tags, _, _ := client.Repositories.ListTags(ctx, "DALTCORE", "ReleaseTools-go", &opt)

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
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ConfigFile().GithubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	opt := github.ListOptions{
		Page:    1,
		PerPage: 1,
	}

	tags, _, _ := client.Repositories.ListTags(ctx, "DALTCORE", "ReleaseTools-go", &opt)
	for _, element := range tags {
		return element.GetName()
	}

	return "unknown"
}
