package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/saba1122333/todo-cli/cli"
	"github.com/saba1122333/todo-cli/utils"
)

func main() {
	utils.CheckOrCreate()
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Fprint(os.Stderr, "Invalid number of arguments. Usage: todo-cli Add <description>")
	}
	var err error
	command := strings.ToLower(os.Args[1])
	command = strings.TrimSpace(command)
	switch command {
	case "add":
		err = cli.HandleAddCommand()
	case "delete":
		err = cli.HandleDeleteCommand()
	case "update":
		err = cli.HandleUpdateCommand()
	case "list":
		err = cli.HandleListCommand()
	case "mark-in-progress":
		err = cli.HandleMarkInProgressCommand()
	case "mark-Done":
		err = cli.HandleMarkDoneCommand()
	default:
		err = fmt.Errorf("unknown command: %s", command)
	}
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		os.Exit(1)
	}
}
