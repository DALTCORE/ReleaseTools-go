package main

import (
	"bytes"
	"fmt"
	ioutil "io/ioutil"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

func ParsePlaybook(name string) bool {

	b, err := ioutil.ReadFile(ReleaseToolPlaybookDirectory() + DirSep() + name + ".rtp")
	if err != nil {
		fmt.Print(err)
	}

	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer(b))

	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}

	// Handling playbook!
	if viper.IsSet("playbook") == false {
		fmt.Println("We're not handeling a playbook!")
		os.Exit(-1)
	}

	if viper.IsSet("playbook.version") == false {
		fmt.Println("Unknown playbook version!")
		os.Exit(-1)
	}

	if viper.GetString("playbook.version") == "1.1" {
		fmt.Println("Using playbook parser version:", viper.GetString("playbook.version"))
		fmt.Println()
		versionOneDotOne()
	} else {
		fmt.Println("Unknown playbook version:", viper.GetString("playbook.version"))
		fmt.Println()
		os.Exit(-1)
	}

	return true
}

func versionOneDotOne() {
	if viper.IsSet("playbook.gitlab") {
		if viper.IsSet("playbook.gitlab.merge_request") {
			fromBranch := viper.GetString("playbook.gitlab.merge_request.from")
			toBranch := viper.GetString("playbook.gitlab.merge_request.to")
			mergeTitle := viper.GetString("playbook.gitlab.merge_request.title")

			if strings.Contains(fromBranch, ":version") {
				fromBranch = strings.Replace(fromBranch, ":version", askVersion(), -1)
			}

			if strings.Contains(toBranch, ":version") {
				toBranch = strings.Replace(toBranch, ":version", askVersion(), -1)
			}

			if strings.Contains(mergeTitle, ":version") {
				mergeTitle = strings.Replace(mergeTitle, ":version", askVersion(), -1)
			}

			if strings.Contains(mergeTitle, ":from") {
				mergeTitle = strings.Replace(mergeTitle, ":from", fromBranch, -1)
			}

			if strings.Contains(mergeTitle, ":to") {
				mergeTitle = strings.Replace(mergeTitle, ":to", toBranch, -1)
			}

			fmt.Println("Merge request from", fromBranch, "to", toBranch)
			MergeRequest := GitlabMakeMergeRequest(mergeTitle, fromBranch, toBranch)
			fmt.Println("URL:", MergeRequest.WebURL)
		}

		if viper.IsSet("playbook.gitlab.make_branch") {
			fromBranch := viper.GetString("playbook.gitlab.make_branch.from")
			toBranch := viper.GetString("playbook.gitlab.make_branch.to")

			if strings.Contains(fromBranch, ":version") {
				fromBranch = strings.Replace(fromBranch, ":version", askVersion(), -1)
			}

			if strings.Contains(toBranch, ":version") {
				toBranch = strings.Replace(toBranch, ":version", askVersion(), -1)
			}

			fmt.Println("Make branch from", fromBranch, "to", toBranch)
			GitlabMakeBranch(fromBranch, toBranch)
		}

		if viper.IsSet("playbook.gitlab.create_tag") {
			fromBranch := viper.GetString("playbook.gitlab.create_tag.from")
			tagName := viper.GetString("playbook.gitlab.create_tag.name")

			if strings.Contains(fromBranch, ":version") {
				fromBranch = strings.Replace(fromBranch, ":version", askVersion(), -1)
			}

			if strings.Contains(tagName, ":version") {
				tagName = strings.Replace(tagName, ":version", askVersion(), -1)
			}

			fmt.Println("Create tag from", fromBranch, "name:", tagName)
			GitlabMakeTag(fromBranch, tagName)
		}
	}

	if viper.IsSet("playbook.mattermost") {
		if viper.IsSet("playbook.mattermost.notify") {
			channel := viper.GetString("playbook.mattermost.notify.channel")
			message := viper.GetString("playbook.mattermost.notify.message")

			// Very specific for our usage for our usage
			if strings.Contains(message, ":project") {
				url := StringBefore(ConfigFile().Repo, "/")
				message = strings.Replace(message, ":project", url, -1)
			}

			if strings.Contains(message, ":version") {
				message = strings.Replace(message, ":version", askVersion(), -1)
			}

			if strings.Contains(message, ":repo") {
				message = strings.Replace(message, ":repo", ConfigFile().Repo, -1)
			}

			fmt.Println("Notify channel", channel, "with message", message)
			MattermostNotify(channel, message)
		}
	}
}
