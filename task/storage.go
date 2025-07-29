package task

import (
	"encoding/json"
	"fmt"

	"os"
)

func WriteTasks(tasks []Task) error {
	jsonData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return fmt.Errorf("failed to marshal tasks: %w", err)
	}
	if err := os.WriteFile(FileName, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write tasks: %w", err)
	}
	return nil
}

func ReadTasks() ([]Task, error) {
	jsonData, err := os.ReadFile(FileName)
	var tasks []Task
	if err != nil {
		return tasks, fmt.Errorf("failed to read tasks: %w", err)
	}

	if err := json.Unmarshal(jsonData, &tasks); err != nil {

		return tasks, fmt.Errorf("failed to Unmarshal tasks: %w", err)
	}
	return tasks, nil
}

func appendTask(task Task) error {
	tasks, err := ReadTasks()
	if err != nil {
		return err // this is already taken care of
	}
	tasks = append(tasks, task)

	if err := WriteTasks(tasks); err != nil {
		return err // this is already taken care of
	}
	return nil

}
func DeleteTask(id int) error {
	tasks, err := ReadTasks()
	if err != nil {
		return fmt.Errorf("failed to read tasks: %w", err)
	}
	if len(tasks) == 0 {
		return fmt.Errorf("unable to delete: to-do list is empty")
	}

	var newTasks []Task
	found := false
	for _, v := range tasks {
		if v.ID != id {
			newTasks = append(newTasks, v)
		} else {
			found = true
		}
	}
	if !found {
		return fmt.Errorf("task with provided id does not exist")
	}
	if err := WriteTasks(newTasks); err != nil {
		return fmt.Errorf("failed to write tasks: %w", err)
	}
	return nil
}
