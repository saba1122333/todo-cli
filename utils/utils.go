package utils

import (
	"github.com/saba1122333/todo-cli/task"
	"os"
)

func FileExists() bool {
	_, err := os.Stat(task.FileName)
	return err == nil || !os.IsNotExist(err)
}

func CheckOrCreate() {
	if !FileExists() {
		file, err := os.Create(task.FileName)
		for err != nil {
			file, err = os.Create(task.FileName)
		}
		defer file.Close()
	}
}
