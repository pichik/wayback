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

	tool.RegisterTool("wayback", "wayback machine", map[string]string{})

	if len(os.Args) < 2 {
		tool.PrintDefaultHelp()
	}
	iTool, _ := tools.GetTool(os.Args[1])

	if iTool == nil {
		tool.PrintDefaultHelp()
	}

	iTool.SetupFlags()
	tool.SetupFlags()
	tool.UpdateFlagUsageHelp()
	tool.ParseFlags(os.Args[2:])

	output.CompileFilters()

	iTool.SetupInput(inputLines)

	// tool.Setup(inputLines)

	fmt.Println("")
}

func cleanup() {
	output.CloseAllFiles()
}
