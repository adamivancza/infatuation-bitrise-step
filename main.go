package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var commitMessage = os.Getenv("GIT_CLONE_COMMIT_MESSAGE_BODY")
	fmt.Printf("%s", commitMessage)
	var re = regexp.MustCompile(`Co-authored-by: .*`)
	var replaced = re.ReplaceAllString(s, "")
	replaced = replaced + "Authored by: The Infatuation"
	fmt.Printf("%s", replaced)

	cmdLog, err := exec.Command("bitrise", "envman", "add", "--key", "GIT_CLONE_COMMIT_MESSAGE_BODY", "--value", replaced).CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to expose output with envman, error: %#v | output: %s", err, cmdLog)
		os.Exit(1)
	}
	
	os.Exit(0)
}
