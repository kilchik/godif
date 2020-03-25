package main

import (
	"os/exec"

	"log"
)

func main()  {
	gits := exec.Command("git", "status")
	gitsOut, err := gits.CombinedOutput()
	if err != nil {
		log.Fatalf("git status: %s (%v)", gitsOut, err)
	}
	log.Println(gitsOut)
}
