package main

import (
	"log"
	"os/exec"
	"project/parser"
	"strings"
)

func main() {
	output, _ := exec.Command("df", "-h").Output()
	op := string(output)
	log.Println("\n", op)
	lines := strings.Split(op, "\n")
	parser.GetResults(lines)
}
