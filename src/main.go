package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/inconshreveable/go-update"
	"github.com/urfave/cli"
	ioutil2 "io/ioutil"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const RTVERSION = "{{VERSION}}"

/**
 * Mainly main things
 */
func main() {

	_, err := net.DialTimeout("tcp", "github.com:443", 10*time.Second)
	if err != nil {
		color.Red("%s", "Cannot connect to network. Some features might be by unstable now!")
	}

	// Auto updater
	if checkIfUpdateAvailible() {
		fmt.Println("There is a update for the release-tool. \nNew version " + getGitVersion() + " \nYou have version: " + RTVERSION + "\n")
		if askConfirmation() {
			downloadUrl := getAssetDownloadUrl()
			resp, err := http.Get(downloadUrl)
			if err != nil {
				color.Red("%v", err)
			}
			defer resp.Body.Close()
			err = update.Apply(resp.Body, update.Options{})
			if err != nil {
				color.Red("%v", err)
				fmt.Println("Going to rollback old version")
				err = update.RollbackError(err)
				if err != nil {
					color.Red("%v", err)
					fmt.Println("Can not recover from failed update!")
					os.Exit(-1);
				}
			}

			fmt.Println("Restart the application to make use of the new version")
			os.Exit(0)
		}
	}

	// Cli app
	app := cli.NewApp()
	app.Name = "ReleaseTools"
	app.Usage = "Releasing made easy"
	app.EnableBashCompletion = true
	app.Version = RTVERSION
	app.Action = func(c *cli.Context) error {
		cli.ShowAppHelp(c)
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "manager:prepare",
			Aliases: []string{"mp"},
			Usage:   "Prepare a new release",
			Action: func(c *cli.Context) error {

				byteArray, err := ioutil2.ReadFile(ReleaseToolStubDirectory() + DirSep() + "prepare.stub")
				if err != nil {
					fmt.Print(err)
				}

				if c.Args().Get(0) != "" {
					setAwnser(ASK_VERSION, c.Args().Get(0))
				}

				text := string(byteArray)

				if strings.Contains(text, ":version") {
					text = strings.Replace(text, ":version", askQuestion(ASK_VERSION), -1)
				}

				if strings.Contains(text, ":repo") {
					text = strings.Replace(text, ":repo", ConfigFile().Repo, -1)
				}

				Issue := GitlabMakeIssue("Release of version v"+askQuestion(ASK_VERSION), text)

				fmt.Println("Issue '" + Issue.Title + "' is posted on:")
				fmt.Println(Issue.WebURL)
				return nil
			},
		},
		{
			Name:    "manager:setup",
			Aliases: []string{"ms"},
			Usage:   "Setup a new Repo and Intothetest/accept environment",
			Action: func(c *cli.Context) error {
				SetupFreshRepo(askQuestion(ASK_REPONAME))
				return nil
			},
		},
		{
			Name:    "manager:changelog",
			Aliases: []string{"mc"},
			Usage:   "Build all changelog entries to CHANGELOG.md",
			Action: func(c *cli.Context) error {

				if c.Args().Get(0) != "" {
					setAwnser(ASK_VERSION, c.Args().Get(0))
				}
				version := askQuestion(ASK_VERSION)


				BuildWholeChangelog(version)

				return nil
			},
		},
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Initialize a ReleaseTools environment",
			Action: func(c *cli.Context) error {

				/**
				  Initialize the .release-tool file
				*/
				if CheckIfReleaseToolInit() == false {
					rtcf := "---\n" +
						"repo: namespace/repo\n"
						//"api_url: https://gitlab.com/api/v4\n" +
						//"api_key: YourRandomApiKey\n" +
						//"mattermost_webhook: https://mattermost.com/hooks/de1pqpn9dprmj\n" +
						//"github_token: YourRandomToken\n" +
						//"# Create a github_token: https://github.com/settings/tokens/new?scopes=repo&description=ReleaseTools-Go\n"

					// write the whole body at once
					err := ioutil2.WriteFile(ReleaseToolsConfigFile(), []byte(rtcf), 0755)
					if err != nil {
						panic(err)
					}
					color.Green("Added " + ReleaseToolsConfigFile() + " with dummy content")
				}

				if CheckIfChangelogDirsAreReady() == false {
					os.Mkdir(ChangelogsDirectory(), 0755)
					color.Green("Added " + ChangelogsDirectory())
				}

				if CheckIfReleaseToolDirIsReady() == false {
					os.Mkdir(ReleaseToolDirectory(), 0755)
					color.Green("Added " + ReleaseToolDirectory())
				}

				if CheckIfReleaseToolStubDirIsReady() == false {
					os.Mkdir(ReleaseToolStubDirectory(), 0755)
					color.Green("Added " + ReleaseToolStubDirectory())
				}

				if CheckIfReleaseToolPlaybookDirIsReady() == false {
					os.Mkdir(ReleaseToolPlaybookDirectory(), 0755)
					color.Green("Added " + ReleaseToolPlaybookDirectory())
				}

				if CheckIfChangelogReleasedDirIsReady() == false {
					os.Mkdir(ChangelogReleasedDirectory(), 0755)
					color.Green("Added " + ChangelogReleasedDirectory())
				}

				if CheckIfChangelogUnreleasedDirIsReady() == false {
					os.Mkdir(ChangelogUnreleasedDirectory(), 0755)
					color.Green("Added " + ChangelogUnreleasedDirectory())
				}

				return nil
			},
		},
		{
			Name:    "status",
			Aliases: []string{"s"},
			Usage:   "Get ReleaseTools environment status",
			Action: func(c *cli.Context) error {

				error := false
				fmt.Println("")
				fmt.Println("Going to test ReleaseTool's filesystem")
				fmt.Println("")
				fmt.Println("OS: " + runtime.GOOS)
				fmt.Println("ARCH: " + runtime.GOARCH)

				fmt.Println("ReleaseTool bin dir:", CurrentExecutablePath())
				fmt.Println("ReleaseTool working dir:", CurrentWorkingDirectory())

				fmt.Println("")

				if CheckIfReleaseToolInit() {
					color.Green("File " + ReleaseToolsConfigFile() + " is found")
				} else {
					error = true
				}

				if CheckIfReleaseToolHomeInit() {
					fmt.Println("")
					color.Magenta("/!\\ This is the global .release-tool file which overrides keys from local config /!\\")
					color.Magenta("/!\\ if the keys are non-existent or if they're empty /!\\")
					color.Cyan("File " + ReleaseToolsHomeConfigFile() + " is found")
					fmt.Println("")
				} else {
					error = true
				}

				if CheckIfReleaseToolDirIsReady() {
					color.Green("Directory " + ReleaseToolDirectory() + " is found")

					if CheckIfReleaseToolStubDirIsReady() {
						color.Green("Directory " + ReleaseToolStubDirectory() + " is found")
					} else {
						error = true
					}

					if CheckIfReleaseToolPlaybookDirIsReady() {
						color.Green("Directory " + ReleaseToolPlaybookDirectory() + " is found")
					} else {
						error = true
					}

				} else {
					error = true
				}

				if CheckIfChangelogDirsAreReady() {
					color.Green("Main changelog dir " + ChangelogsDirectory() + " is found")
				} else {
					error = true
				}

				if CheckIfChangelogReleasedDirIsReady() {
					color.Green("Released changelogs dir " + ChangelogReleasedDirectory() + " is found")
				} else {
					error = true
				}

				if CheckIfChangelogUnreleasedDirIsReady() {
					color.Green("Unreleased changelogs dir " + ChangelogUnreleasedDirectory() + "is found")
				} else {
					error = true
				}

				fmt.Println("")

				if error {
					color.Red("Some errors where found. Do not continue at this point. Call for a release manager!")
				} else {
					color.Green("FS seems to be fine")
				}

				return nil
			},
		},
		{
			Name:    "playbook",
			Aliases: []string{"p"},
			Usage:   "Run a playbook",
			Action: func(c *cli.Context) error {

				if c.Args().Get(0) == "" {
					color.Red("You have to give a playbook name!")
					os.Exit(1)
				}

				if c.Args().Get(1) != "" {
					setAwnser(ASK_VERSION, c.Args().Get(1))
				}

				ParsePlaybook(c.Args().Get(0))
				return nil
			},
		},
		{
			Name:    "auto-update",
			Aliases: []string{"au"},
			Usage:   "Auto update the Release Tools",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "force",
					Usage: "Force update",
				},
			},
			Action: func(c *cli.Context) error {
				if checkIfUpdateAvailible() || c.IsSet("force") {
					fmt.Println("Going to update...")
					downloadUrl := getAssetDownloadUrl()

					fmt.Println("Downloading...")
					resp, err := http.Get(downloadUrl)
					if err != nil {
						color.Red("%v", err)
					}
					defer resp.Body.Close()
					fmt.Println("Going to apply update...")
					err = update.Apply(resp.Body, update.Options{})
					if err != nil {
						color.Red("%v", err)
					}

					fmt.Println("Updating finished...")
				} else {
					fmt.Println("There is no update for the release-tool.\nGit version " + getGitVersion() + " \nYou have version: " + RTVERSION + "\n")
				}
				return nil
			},
		},
		{
			Name:    "changelog",
			Aliases: []string{"c"},
			Usage:   "Make a changelog entry",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "dry-run",
					Usage: "To test commands",
				},
				cli.BoolFlag{
					Name:  "force",
					Usage: "Force write to file",
				},
			},
			Action: func(c *cli.Context) error {

				if RunChecks() == false {
					os.Exit(1)
				}

				if c.IsSet("dry-run") && c.IsSet("force") {
					fmt.Print("You force on a dry run.")
					time.Sleep(1 * time.Second)
					fmt.Print(".")
					time.Sleep(1 * time.Second)
					fmt.Print(".")
					time.Sleep(1 * time.Second)
					fmt.Print(".")
					time.Sleep(1 * time.Second)
					fmt.Print(".")
					time.Sleep(1 * time.Second)
					fmt.Print(".")
					time.Sleep(1 * time.Second)
					fmt.Println("")
					color.Red("No, just no...")
					os.Exit(1)
				}

				var summary string = askQuestion(ASK_TITLE)
				fmt.Println("")
				// var version string = askQuestion(ASK_VERSION)
				var mergeType string = askMergeType()
				fmt.Println("")
				var mrid string = askQuestion(ASK_MRID)
				fmt.Println("")
				var name string = askUsername()
				fmt.Println("")

				clitem := BuildChangelogEntry(MergeRequestSummary{
					summary,
					name,
					mergeType,
					mrid,
				})

				var fileName string = filepath.FromSlash(ChangelogUnreleasedDirectory() + DirSep() + mrid + "-" + GetCurrentBranch() + ".yaml")

				if c.IsSet("dry-run") {
					fmt.Println("")
					fmt.Println("")
					fmt.Println("Filename:", fileName)
					fmt.Println(clitem)
				} else {

					// Nice idea, but MergeRequest always exists...
					mridInt, _ := strconv.ParseInt(mrid, 0, 32)

					mergeRequest, err := GitlabCheckOpenMrById(int(mridInt))

					if err == nil && mergeRequest != nil {
						fmt.Println("Merge request info:")
						fmt.Println()
						fmt.Println("Author: " + mergeRequest.Author.Username)
						fmt.Println("Title:  " + mergeRequest.Title)
						fmt.Println("State:  " + mergeRequest.State)
						fmt.Println("URL:    " + mergeRequest.WebURL)
					}

					if FileExists(fileName) && c.IsSet("force") == false {
						color.Red("File " + mrid + "-" + GetCurrentBranch() + ".yaml already exists.")
						os.Exit(1)
					}

					// write the whole body at once
					err = ioutil2.WriteFile(fileName, []byte(clitem), 0644)
					if err != nil {
						panic(err)
					}
				}

				return nil
			},
		},
	}
	app.Run(os.Args)
}
