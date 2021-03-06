package main

import (
	"bufio"
	"fmt"
	ioutil "io/ioutil"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/ghodss/yaml"
	"github.com/imdario/mergo"
	"github.com/olekukonko/tablewriter"
	"github.com/vigneshuvi/GoDateFormat"
)

/**
 * Struct for Merge Request Summaries
 */
type MergeRequestSummary struct {
	Title     string `json:"title"`
	Name      string `json:"author"`
	MergeType string `json:"type"`
	MergeId   string `json:"merge_request"`
}

type ConfigHolder struct {
	GitlabUrl         string `json:"gitlab_url"`
	Group             string `json:"group"`
	Repo              string `json:"repo"`
	ApiUrl            string `json:"api_url"`
	ApiKey            string `json:"api_key"`
	MattermostWebhook string `json:"mattermost_webhook"`
	GithubAccessToken string `json:"github_token"`
}

type Changelogs struct {
	Filename     string
	MergeSummary MergeRequestSummary
}

/**
 * Slice with the parsed changelogs
 */
var ParsedChangelogs []Changelogs

func BuildChangelogEntry(summary MergeRequestSummary) string {
	y, err := yaml.Marshal(summary)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return ""
	}
	return string(y)
}

func ConfigFile() ConfigHolder {
	var c1 ConfigHolder

	b, err := ioutil.ReadFile(ReleaseToolsConfigFile()) // just pass the file name
	if err != nil {
		// fmt.Print(err)
	}

	err = yaml.Unmarshal(b, &c1)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return ConfigHolder{}
	}

	if CheckIfReleaseToolHomeInit() {
		var c2 ConfigHolder
		h, err := ioutil.ReadFile(ReleaseToolsHomeConfigFile()) // just pass the file name
		if err != nil {
			// fmt.Print(err)
		}

		err = yaml.Unmarshal(h, &c2)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return ConfigHolder{}
		}

		mergo.MergeWithOverwrite(&c1, c2)
	}

	return c1
}

func BuildWholeChangelog(version string) {
	files, err := ioutil.ReadDir(ChangelogUnreleasedDirectory())
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}

	for _, f := range files {

		fileName := ChangelogUnreleasedDirectory() + DirSep() + f.Name()
		MrSum := MergeRequestSummary{}

		b, err := ioutil.ReadFile(fileName) // just pass the file name
		if err != nil {
			fmt.Print(err)
		}

		err = yaml.Unmarshal(b, &MrSum)
		if err != nil {
			color.Red("err: %v\n", err)
		}

		ParsedChangelogs = append(ParsedChangelogs, Changelogs{
			Filename:     f.Name(),
			MergeSummary: MrSum,
		})
	}

	if len(ParsedChangelogs) == 0 {
		color.Red("No changelogs to parse!")
		os.Exit(1)
	}

	today := time.Now()
	todayString := today.Format(GoDateFormat.ConvertFormat("yyyy-mm-dd"))

	freshChangelog := "## " + version + " (" + todayString + ")  \n"

	sort.Slice(ParsedChangelogs, func(i, j int) bool {
		switch strings.Compare(ParsedChangelogs[i].MergeSummary.MergeType, ParsedChangelogs[j].MergeSummary.MergeType) {
		case -1:
			return true
		case 1:
			return false
		}

		return false
	})

	lastMergeType := ""
	for _, v := range ParsedChangelogs {
		if v.MergeSummary.MergeType != lastMergeType {
			freshChangelog = freshChangelog + "\n**" + v.MergeSummary.MergeType + "**\n\n"
			lastMergeType = v.MergeSummary.MergeType
		}

		freshChangelog = freshChangelog + "- " + v.MergeSummary.Title + " [!" + v.MergeSummary.MergeId + "] — " + v.MergeSummary.Name + "\n"
	}

	freshChangelog = freshChangelog + "\n"

	for _, v := range ParsedChangelogs {
		freshChangelog = freshChangelog + "[!" + v.MergeSummary.MergeId + "]: <" + ConfigFile().GitlabUrl + "/" + ConfigFile().Repo + "/merge_requests/" + v.MergeSummary.MergeId + "> \"!" + v.MergeSummary.MergeId + "\"\n"
		os.Rename(ChangelogUnreleasedDirectory()+DirSep()+v.Filename, ChangelogReleasedDirectory()+DirSep()+v.Filename)
	}

	b, _ := ioutil.ReadFile(ChangelogFile()) // just pass the file name

	if len(b) > 0 {
		freshChangelog = freshChangelog + "\n\n" + string(b)
	}

	f, err := os.Create(ChangelogFile())
	if err != nil {
		color.Red("Cannot create changelog: " + err.Error())
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	_, err = w.WriteString(freshChangelog)
	if err != nil {
		color.Red("Cannot write changelog: " + err.Error())
	}

}

func ListChangelogs() {
	files, err := ioutil.ReadDir(ChangelogUnreleasedDirectory())
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"MR ID", "Author", "Entry", "Type"})
	table.SetBorder(true)
	table.SetColWidth(1000)

	for _, f := range files {

		fileName := ChangelogUnreleasedDirectory() + DirSep() + f.Name()
		MrSum := MergeRequestSummary{}

		b, err := ioutil.ReadFile(fileName) // just pass the file name
		if err != nil {
			fmt.Print(err)
		}

		err = yaml.Unmarshal(b, &MrSum)
		if err != nil {
			color.Red("err: %v\n", err)
		}

		table.Append([]string{MrSum.MergeId, MrSum.Name, MrSum.Title, MrSum.MergeType})

	}

	if len(files) > 0 {
		table.Render()
	} else {
		color.Red("No unreleased changelogs found!")
	}

}
