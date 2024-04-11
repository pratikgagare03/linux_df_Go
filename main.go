package main

import (
	"fmt"
	"os/exec"
	"project/parser"
	"strings"
)

func main() {
	output, _ := exec.Command("df").Output()
	op := string(output)
	lines := strings.Split(op, "\n")
	parser.GetResults(lines)
	fmt.Println(parser.Readable(479077064))
}
