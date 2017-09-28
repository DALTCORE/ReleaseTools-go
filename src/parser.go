package main

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	ioutil "io/ioutil"
	"os"
	"strings"
)

func ParsePlaybook(name string) bool {

	b, err := ioutil.ReadFile(ReleaseToolPlaybookDirectory() + DirSep() + name + ".rtp")
	if err != nil {
		fmt.Print(err)
	}

	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(b))

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

			if strings.Contains(fromBranch, ":version") {
				fromBranch = strings.Replace(fromBranch, ":version", askQuestion(ASK_VERSION), -1)
			}

			if strings.Contains(toBranch, ":version") {
				toBranch = strings.Replace(toBranch, ":version", askQuestion(ASK_VERSION), -1)
			}

			fmt.Println("Merge request from", fromBranch, "to", toBranch)
		}

		if viper.IsSet("playbook.gitlab.make_branch") {
			fromBranch := viper.GetString("playbook.gitlab.make_branch.from")
			toBranch := viper.GetString("playbook.gitlab.make_branch.to")

			if strings.Contains(fromBranch, ":version") {
				fromBranch = strings.Replace(fromBranch, ":version", askQuestion(ASK_VERSION), -1)
			}

			if strings.Contains(toBranch, ":version") {
				toBranch = strings.Replace(toBranch, ":version", askQuestion(ASK_VERSION), -1)
			}

			fmt.Println("Make branch from", fromBranch, "to", toBranch)
		}
	}

	if viper.IsSet("playbook.mattermost") {
		if viper.IsSet("playbook.mattermost.notify") {
			channel := viper.GetString("playbook.mattermost.notify.channel")
			message := viper.GetString("playbook.mattermost.notify.message")

			fmt.Println("Notify channel", channel, "with message", message)
		}
	}
}
