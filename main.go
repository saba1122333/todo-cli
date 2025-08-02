package main

import (
	"fmt"
	"github.com/saba1122333/todo-cli/cli"
	"os"
)

func main() {
	err := cli.Run()
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		os.Exit(1)
	}
}
