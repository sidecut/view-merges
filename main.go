package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("git", "log", "--oneline", "--first-parent")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Wait()
	fmt.Printf("%s", &out)

	// log.Printf("Exit code: %v", exitCode)
}
