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
		for err != nil {
			file, err = os.Create(task.FileName)
		}
		defer file.Close()
	}
}
