package main

import (
	"fmt"
	"github.com/saba1122333/todo-cli/cli"
	"github.com/saba1122333/todo-cli/utils"
	"os"
)

func main() {
	utils.CheckOrCreate()
	err := cli.Run()
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		os.Exit(1)
	}
}
