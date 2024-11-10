package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pichik/go-modules/output"
	"github.com/pichik/go-modules/tool"
	"github.com/pichik/wayback/tools"
)

var inputLines []string

func SetupPipe() {

	//Read from pipe, if nothing is piped show help
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {

		r := bufio.NewReader(os.Stdin)
		for {
			line, _, err := r.ReadLine()
			if len(line) > 0 {
				inputLines = append(inputLines, string(line))
			}
			if err != nil {
				break
			}
		}
	}
}

func main() {
	defer cleanup()
	SetupPipe()

	if len(os.Args) < 2 {
		tool.PrintDefaultHelp()
	}

	iTool, t := tools.GetTool(os.Args[1])

	if t.Name == "error" {
		tool.PrintDefaultHelp()
	}

	iTool.SetupFlags()
	tool.SetupFlags()
	tool.UpdateFlagUsageHelp()

	output.CompileFilters()

	iTool.Setup(inputLines)

	fmt.Println("")
}

func cleanup() {
	output.CloseAllFiles()
}
