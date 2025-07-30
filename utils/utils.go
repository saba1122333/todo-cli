package utils

import (
	"github.com/saba1122333/todo-cli/task"
	"os"
)

func fileExists() bool {
	_, err := os.Stat(task.FileName)
	return err == nil || !os.IsNotExist(err)
}

func CheckOrCreate() {
	if !fileExists() {
		file, err := os.Create(task.FileName)
		if err != nil {
			// Don't loop forever. Just fail fast.
			panic("failed to create file: " + err.Error())
		}
		defer file.Close()
	}
}
