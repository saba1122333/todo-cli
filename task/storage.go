package task

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
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
	if err != nil {
		return nil, fmt.Errorf("failed to read tasks: %w", err)
	}
	if len(jsonData) == 0 {
		return []Task{}, nil
	}
	var tasks []Task
	if err := json.Unmarshal(jsonData, &tasks); err != nil {
		return tasks, fmt.Errorf("failed to Unmarshal tasks: %w", err)
	}

	return tasks, nil
}

func AppendTask(description string) error {
	tasks, err := ReadTasks()
	id := generateID()
	task := Task{
		ID:          id,
		Description: description,
		Status:      Todo,
		CreatedAt:   time.Now().Truncate(time.Minute).Format("2006-01-02 15:04")}

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
		return fmt.Errorf("to-do list is empty")
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

func UpdateTask(id int, description string) error {
	tasks, err := ReadTasks()
	if err != nil {
		return fmt.Errorf("failed to read tasks: %w", err)
	}
	if len(tasks) == 0 {
		return fmt.Errorf("to-do list is empty")
	}
	var newTasks []Task
	found := false
	for _, v := range tasks {
		if v.ID == id {
			found = true
			v.Description = description
			v.UpdatedAt = time.Now().Truncate(time.Minute).Format("2006-01-02 15:04")
		}
		newTasks = append(newTasks, v)
	}
	if !found {
		return fmt.Errorf("task with provided id does not exist")
	}
	if err := WriteTasks(newTasks); err != nil {
		return fmt.Errorf("failed to write tasks: %w", err)
	}
	return nil
}

func MarkTaskInProgress(id int) error {
	tasks, err := ReadTasks()
	if err != nil {
		return fmt.Errorf("failed to read tasks: %w", err)
	}
	if len(tasks) == 0 {
		return fmt.Errorf("to-do list is empty")
	}
	var newTasks []Task
	found := false
	for _, v := range tasks {
		if v.ID == id {
			found = true
			v.Status = Inprogress
			v.UpdatedAt = time.Now().Truncate(time.Minute).Format("2006-01-02 15:04")
		}
		newTasks = append(newTasks, v)
	}
	if !found {
		return fmt.Errorf("task with provided id does not exist")
	}
	if err := WriteTasks(newTasks); err != nil {
		return fmt.Errorf("failed to write tasks: %w", err)
	}
	return nil
}

func MarkTaskDone(id int) error {
	tasks, err := ReadTasks()
	if err != nil {
		return fmt.Errorf("failed to read tasks: %w", err)
	}
	if len(tasks) == 0 {
		return fmt.Errorf("to-do list is empty")
	}
	var newTasks []Task
	found := false
	for _, v := range tasks {
		if v.ID == id {
			found = true
			v.Status = Done
			v.UpdatedAt = time.Now().Truncate(time.Minute).Format("2006-01-02 15:04")
		}
		newTasks = append(newTasks, v)
	}
	if !found {
		return fmt.Errorf("task with provided id does not exist")
	}
	if err := WriteTasks(newTasks); err != nil {
		return fmt.Errorf("failed to write tasks: %w", err)
	}
	return nil
}
func ListTasks(status string) error {
	tasks, err := ReadTasks()
	if err != nil {
		return err
	}
	if status != "all" && status != string(Todo) && status != string(Inprogress) && status != string(Done) {
		return fmt.Errorf("invalid status: %s. Valid statuses are: all, todo, in-progress, done", status)
	}
	if status == "all" {
		for _, v := range tasks {
			fmt.Println(v.String())
			fmt.Println()
		}
		return nil
	}
	for _, v := range tasks {
		if string(v.Status) == status {
			fmt.Println(v.String())
			fmt.Println()
		}
	}
	return nil

}
func generateID() int {
	tasks, _ := ReadTasks()
	max := 0
	for _, v := range tasks {
		if v.ID > max {
			max = v.ID
		}
	}
	return max + 1

}
