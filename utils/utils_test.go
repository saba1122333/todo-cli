package utils

import (
	"github.com/saba1122333/todo-cli/task"
	"os"
	"testing"
)

func destroy() {
	os.Remove(task.FileName)
	task.FileName = "Tasks.json"
}

func setup() {
	task.FileName = "utils_test.json"
	os.Remove(task.FileName)
}

func TestFileExists(t *testing.T) {
	setup()
	defer destroy()

	if FileExists() {
		t.Errorf("Expected fileExists() to be false, got true")
	}
	_, err := os.Create(task.FileName)

	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}
	if !FileExists() {
		t.Errorf("Expected fileExists() to be true, got false")
	}

}

func TestCheckOrCreate(t *testing.T) {
	task.FileName = "utils_tasks.json"
	defer destroy()

	os.Remove(task.FileName)
	CheckOrCreate()
	_, err := os.Stat(task.FileName)
	if err != nil {
		t.Errorf("Expected CheckOrCreate() to create file, it did not")
	}

}
