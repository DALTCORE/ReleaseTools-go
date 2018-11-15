package main

import (
	"bufio"
	"fmt"
	ioutil "io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/tcnksm/go-gitconfig"
)

/**
 * Default questions
 */
const ASK_VERSION = "What version?"
const ASK_TITLE = "Whats the summary of this merge request?"
const ASK_MRID = "What merge request id?"
const ASK_REPONAME = "Whats the name of the new repository?"

const ASK_ACCEPT = "Acceptation"
const ASK_PRODUC = "Production"

/**
 * Struct for slice AskedQuestions
 */
type questions struct {
	question string
	awnser   string
}

/**
 * Slice with all asked questions
 */
var AskedQuestions []questions

var active bool = false

/**
 * Ask an question and save it to a slice so if asked again,
 * we can just return the already asked answer
 */
func askQuestion(question string) string {
	if len(AskedQuestions) != 0 {
		for index, _ := range AskedQuestions {
			if AskedQuestions[index].question == question {
				active = false
				return AskedQuestions[index].awnser
			}
		}
	}

	reader := bufio.NewReader(os.Stdin)

	// For some reason when the iteration is high enough
	// the platform will print two times? Weird, work-a-round
	// in use by var active
	if active == false {
		fmt.Print(question + ": ")
		active = true
	}

	text, _ := reader.ReadString('\n')

	text = strings.Replace(text, "\r", "", -1)
	text = strings.Replace(text, "\n", "", -1)

	if len(text) < 1 {
		return askQuestion(question)
	}

	AskedQuestions = append(AskedQuestions, questions{question, text})
	active = false
	return text
}

/**
 * Set answer before asking!
 */
func setAwnser(question string, answer string) {
	AskedQuestions = append(AskedQuestions, questions{question, answer})
}

/**
 * Ask user for confirmation
 */
func askConfirmation() bool {
	var s string

	fmt.Printf("Do you want to continue? (y/N): ")
	_, err := fmt.Scan(&s)
	if err != nil {
		panic(err)
	}

	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if s == "y" || s == "yes" {
		return true
	}

	return false
}

func askMergeType() string {
	mergeRequestTypes := [8]string{
		"Other",
		"New feature",
		"Bug fix",
		"Feature change",
		"New deprecation",
		"Feature removal",
		"Security fix",
		"Style fix",
	}

	for index, element := range mergeRequestTypes {
		fmt.Printf("[%d] %s\n", index, element)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is this merge request about? [0-7]: ")
	index, _ := reader.ReadString('\n')

	index = strings.Replace(index, "\r", "", -1)
	index = strings.Replace(index, "\n", "", -1)

	int64Index, _ := strconv.ParseInt(index, 10, 64)

	if len(mergeRequestTypes) >= int(int64Index) {
		return mergeRequestTypes[int64Index]
	} else {
		fmt.Println("\nYour input of", index, "is not valid for this list.\nTry again, bright light :-) :\n")
		return askMergeType()
	}
}

func askReleaseType() string {
	mergeRequestTypes := [2]string{
		"Acceptation",
		"Production",
	}

	for index, element := range mergeRequestTypes {
		fmt.Printf("[%d] %s\n", index, element)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What kind of release are you doing? [0-1]: ")
	index, _ := reader.ReadString('\n')

	index = strings.Replace(index, "\r", "", -1)
	index = strings.Replace(index, "\n", "", -1)

	int64Index, _ := strconv.ParseInt(index, 10, 64)

	if len(mergeRequestTypes) >= int(int64Index) {
		return mergeRequestTypes[int64Index]
	} else {
		fmt.Println("\nYour input of", index, "is not valid for this list.\nTry again, bright light :-) :\n")
		return askReleaseType()
	}
}

func askUsername() string {

	reader := bufio.NewReader(os.Stdin)

	name, _ := gitconfig.Username()

	fmt.Print("What is your name [", name, "]: ")
	text, _ := reader.ReadString('\n')

	text = strings.Replace(text, "\r", "", -1)
	text = strings.Replace(text, "\n", "", -1)

	if len(text) < 3 {
		text = name
	}

	return text
}

func askChangelogSummary() string {
	reader := bufio.NewReader(os.Stdin)

	name := GetLastCommitMessage()

	fmt.Print(ASK_TITLE+" [", GetLastCommitMessage(), "]: ")
	text, _ := reader.ReadString('\n')

	text = strings.Replace(text, "\r", "", -1)
	text = strings.Replace(text, "\n", "", -1)

	if len(text) < 3 {
		text = name
	}

	return text
}

func askVersion() string {

	if len(AskedQuestions) != 0 {
		for index, _ := range AskedQuestions {
			if AskedQuestions[index].question == ASK_VERSION {
				active = false
				return AskedQuestions[index].awnser
			}
		}
	}

	var re = regexp.MustCompile(`\#\#\s(?P<major>\d+).(?P<minor>\d+).(?P<hotfix>\d+)\s\((?P<year>\d+)-(?P<month>\d+)-(?P<day>\d+)\)`)
	b, _ := ioutil.ReadFile(ChangelogFile())
	text := ""

	if len(b) > 0 {
		str := string(b)
		m := reSubMatchMap(re, str)

		reader := bufio.NewReader(os.Stdin)

		hotfix, _ := strconv.ParseInt(m["hotfix"], 0, 0)
		newHotfixVersion := (1 + hotfix)

		name := m["major"] + "." + m["minor"] + "." + strconv.Itoa(int(newHotfixVersion))

		fmt.Print(ASK_VERSION+" [next in increment is '", name, "']: ")
		text, _ = reader.ReadString('\n')

		if len(text) < 3 {
			text = name
		}

	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(ASK_VERSION + ":")
		text, _ = reader.ReadString('\n')
	}

	text = strings.Replace(text, "\r", "", -1)
	text = strings.Replace(text, "\n", "", -1)

	if len(text) < 3 {
		fmt.Println(fmt.Errorf("No version found in %s, faling back to 0.0.0", ChangelogFile()))
		text = "0.0.0"
	}

	AskedQuestions = append(AskedQuestions, questions{ASK_VERSION, text})

	return text
}

func reSubMatchMap(r *regexp.Regexp, str string) map[string]string {
	match := r.FindStringSubmatch(str)
	subMatchMap := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i != 0 {
			subMatchMap[name] = match[i]
		}
	}
	return subMatchMap
}
