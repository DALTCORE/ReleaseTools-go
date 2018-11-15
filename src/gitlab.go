package main

import (
	"os"

	gitlab "github.com/daltcore/go-gitlab"
	"github.com/fatih/color"
)

func GitlabCheckOpenMrById(mrid int) (*gitlab.MergeRequest, error) {
	git := gitlab.NewClient(nil, ConfigFile().ApiKey)
	git.SetBaseURL(ConfigFile().ApiUrl)
	mergeRequest, _, err := git.MergeRequests.GetMergeRequest(ConfigFile().Repo, mrid)
	return mergeRequest, err
}

func GitlabMakeIssue(title string, stub string) *gitlab.Issue {
	git := gitlab.NewClient(nil, ConfigFile().ApiKey)
	git.SetBaseURL(ConfigFile().ApiUrl)

	opt := gitlab.CreateIssueOptions{
		Title:       &title,
		Description: &stub,
	}

	Issue, _, err := git.Issues.CreateIssue(ConfigFile().Repo, &opt)

	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}

	return Issue
}

func GitlabMakeMergeRequest(title string, from string, to string) *gitlab.MergeRequest {
	git := gitlab.NewClient(nil, ConfigFile().ApiKey)
	git.SetBaseURL(ConfigFile().ApiUrl)

	opt := gitlab.CreateMergeRequestOptions{
		Title:        gitlab.String(title),
		SourceBranch: gitlab.String(from),
		TargetBranch: gitlab.String(to),
	}

	MergeRequest, _, _ := git.MergeRequests.CreateMergeRequest(ConfigFile().Repo, &opt)

	return MergeRequest
}

func GitlabMakeBranch(from string, to string) *gitlab.Branch {
	git := gitlab.NewClient(nil, ConfigFile().ApiKey)
	git.SetBaseURL(ConfigFile().ApiUrl)

	opt := gitlab.CreateBranchOptions{
		Branch: gitlab.String(to),
		Ref:    gitlab.String(from),
	}

	Branch, _, _ := git.Branches.CreateBranch(ConfigFile().Repo, &opt)

	return Branch
}

func GitlabMakeTag(from string, to string) *gitlab.Tag {
	git := gitlab.NewClient(nil, ConfigFile().ApiKey)
	git.SetBaseURL(ConfigFile().ApiUrl)

	opt := gitlab.CreateTagOptions{
		TagName: gitlab.String(to),
		Ref:     gitlab.String(from),
	}

	Tag, _, _ := git.Tags.CreateTag(ConfigFile().Repo, &opt)

	return Tag
}
