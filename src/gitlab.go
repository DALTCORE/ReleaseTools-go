package main

import (
	gitlab "github.com/daltcore/go-gitlab"
	"fmt"
	"github.com/metal3d/go-slugify"
	"os"
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

	Issue, _, _ := git.Issues.CreateIssue(ConfigFile().Repo, &opt)

	return Issue
}

func SetupFreshRepo(name string) {
	git := gitlab.NewClient(nil, ConfigFile().ApiKey)
	git.SetBaseURL(ConfigFile().ApiUrl)
	slug := slugify.Marshal(name)

	fmt.Println("Creating repository with default branch 'develop'...")
	// Create repo
	gitlabStruct := gitlab.CreateProjectOptions{
		Name: gitlab.String(name),
		Path: gitlab.String(slug),
		NamespaceID: gitlab.Int(5), // curl --header "PRIVATE-TOKEN: TOKEN_HERE" https://git.intothesource.com/api/v4/namespaces?search=source
		MergeRequestsEnabled: gitlab.Bool(true),
		OnlyAllowMergeIfAllDiscussionsAreResolved: gitlab.Bool(true),
		OnlyAllowMergeIfPipelineSucceeds: gitlab.Bool(true),
		ApprovalsBeforeMerge: gitlab.Int(1),
	}
	project, _, e := git.Projects.CreateProject(&gitlabStruct)

	if e != nil {
		color.Red(e.Error())
		os.Exit(1)
	}

	// Inset default stuff in branch
	fileOptions := gitlab.CreateFileOptions{
		Branch: gitlab.String("develop"),
		AuthorEmail: gitlab.String("systeembeheer@intothesource.com"),
		AuthorName: gitlab.String("Releasetool Autogenerator"),
		Content: gitlab.String(assetTemplateMergeRequestMd()),
		CommitMessage:gitlab.String("Adding merge request template"),
	}
	fmt.Println("Adding merge request template")
	_, _, e = git.RepositoryFiles.CreateFile(project.PathWithNamespace, ".gitlab/merge_request_templates/MergeRequest.md", &fileOptions)

	if e != nil {
		color.Red(e.Error())
		os.Exit(1)
	}

	fileOptions = gitlab.CreateFileOptions{
		Branch: gitlab.String("develop"),
		AuthorEmail: gitlab.String("systeembeheer@intothesource.com"),
		AuthorName: gitlab.String("Releasetool Autogenerator"),
		Content: gitlab.String(assetTemplateBugMd()),
		CommitMessage:gitlab.String("Adding bug template"),
	}
	fmt.Println("Adding bug template")
	_, _, e = git.RepositoryFiles.CreateFile(project.PathWithNamespace, ".gitlab/issue_templates/Bug.md", &fileOptions)

	if e != nil {
		color.Red(e.Error())
		os.Exit(1)
	}

	// Create branches
	branchStruct := gitlab.CreateBranchOptions{
		Branch:gitlab.String("staging"),
		Ref:gitlab.String("develop"),
	}
	fmt.Println("Creating branch 'staging' from 'develop'")
	_, _, e = git.Branches.CreateBranch(project.PathWithNamespace, &branchStruct)

	if e != nil {
		color.Red(e.Error())
		os.Exit(1)
	}

	branchStruct = gitlab.CreateBranchOptions{
		Branch:gitlab.String("master"),
		Ref:gitlab.String("develop"),
	}
	fmt.Println("Creating branch 'master' from 'develop'")
	_, _, e = git.Branches.CreateBranch(project.PathWithNamespace, &branchStruct)

	if e != nil {
		color.Red(e.Error())
		os.Exit(1)
	}

	// Protect branches
	branchSettings := gitlab.ProtectBranchOptions{
		DevelopersCanMerge:gitlab.Bool(false),
		DevelopersCanPush:gitlab.Bool(false),
	}

	fmt.Println("Protecting branch 'develop'")
	git.Branches.ProtectBranch(project.PathWithNamespace, "develop", &branchSettings)
	fmt.Println("Protecting branch 'staging'")
	git.Branches.ProtectBranch(project.PathWithNamespace, "staging", &branchSettings)
	fmt.Println("Protecting branch 'master'")
	git.Branches.ProtectBranch(project.PathWithNamespace, "master", &branchSettings)
	fmt.Println("")
	fmt.Println("Project URL:", project.WebURL)
	fmt.Println("")
	fmt.Println("Fresh beginning:")
	fmt.Println("git clone " + project.SSHURLToRepo)
	fmt.Println("git add .")
	fmt.Println("git commit")
	fmt.Println("git push origin develop")
	fmt.Println("")
	fmt.Println("Existing directory:")
	fmt.Println("git remote add origin " + project.SSHURLToRepo)
	fmt.Println("git add .")
	fmt.Println("git commit")
	fmt.Println("git push origin develop --allow-unrelated-histories")
}