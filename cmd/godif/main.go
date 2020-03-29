package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"log"
)

var skipSuffixes = []string {
	"_mock.go",
	"_test.go",
	"pb.go",
	"pb.goclay.go",
}

func isSkippable(fname string) bool {
	match, _ := regexp.Match(strings.Join(skipSuffixes, "|"), []byte(fname))
	return match
}

func main()  {
	gits := exec.Command("git", "status")
	gitsOut, err := gits.CombinedOutput()
	if err != nil {
		log.Fatalf("git status: %s (%v)", gitsOut, err)
	}
	var toDiff []string
	for _, line := range strings.Split(string(gitsOut), "\n") {
		if strings.Contains(line, "\tmodified:") {
			tokens := strings.Split(line," ")
			filePath := tokens[len(tokens)-1]
			if isSkippable(filePath) {
				continue
			}
			toDiff = append(toDiff, filePath)
		}
	}
	args := []string{"diff"}
	args = append(args, toDiff...)

	diffOut, _ := exec.Command("git", args...).CombinedOutput()
	fmt.Fprintf(os.Stdout, "%s", diffOut)
}
