package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main()  {
	if len(os.Args) != 2 {
		log.Fatalf("gitc: unexpected arguments")
	}
	gitCmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := gitCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("git rev-parse: %s (%v)", output, err)
	}
	branchName := string(output[:len(output)-1])
	var commitPrefix string
	if strings.HasPrefix(branchName, "wallet-") {
		commitPrefix += "[" + strings.ToUpper(branchName) + "] "
	}

	gitCmd = exec.Command("git", "commit", "-m", commitPrefix+os.Args[1])
	output, err = gitCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("git commit: %s (%v)", output, err)
	}
	fmt.Fprintf(os.Stdout, string(output))
}
