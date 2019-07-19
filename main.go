package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/chuckha/appenv/programs"
)

func main() {
	versioners := []Versioner{
		programs.Python,
		programs.Ruby,
		programs.Bash,
		programs.Git,
		programs.SSH,
		programs.Docker,
	}

	printPath := false
	couldNotFind := []string{}
	for _, versioner := range versioners {
		out, err := versioner.Version()
		if err == exec.ErrNotFound {
			couldNotFind = append(couldNotFind, versioner.GetCommand())
			printPath = true
		}
		if strings.HasSuffix(string(out), "\n") {
			fmt.Print(string(out))
			continue
		}
		fmt.Println(string(out))
	}
	if printPath {
		fmt.Println(os.Getenv("PATH"))
	}
}

type Versioner interface {
	Version() ([]byte, error)
	GetCommand() string
}
