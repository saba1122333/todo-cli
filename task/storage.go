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
		return fmt.Errorf(ErrMarshal, err)
	}
	if err := os.WriteFile(FileName, jsonData, 0644); err != nil {

		return fmt.Errorf(ErrWriteTasks, err)
	}
	return nil
}

func ReadTasks() ([]Task, error) {
	jsonData, err := os.ReadFile(FileName)
	if err != nil {
		return nil, fmt.Errorf(ErrReadTasks, err)
	}
	if len(jsonData) == 0 {
		return []Task{}, nil
	}
	var tasks []Task
	if err := json.Unmarshal(jsonData, &tasks); err != nil {
		return tasks, fmt.Errorf(ErrUnMarshal, err)
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
		return err
	}
	if len(tasks) == 0 {
		return fmt.Errorf(ErrEmptyList)
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
		return fmt.Errorf(ErrTaskNotFound)
	}
	if err := WriteTasks(newTasks); err != nil {
		return fmt.Errorf(ErrWriteTasks, err)
	}
	return nil
}

func ApplyTaskByID(id int, apply func(task *Task) error) error {
	tasks, err := ReadTasks()
	if err != nil {
		return err //  already taken care of
	}
	if len(tasks) == 0 {
		return fmt.Errorf(ErrEmptyList)
	}
	for i, v := range tasks {
		if v.ID == id {
			if err := apply(&tasks[i]); err != nil {
				return err
			}
			return WriteTasks(tasks)
		}
	}
	return fmt.Errorf(ErrTaskNotFound)
}
func UpdateTask(id int, description string) error {
	return ApplyTaskByID(id, func(task *Task) error {
		task.Description = description
		task.UpdatedAt = time.Now().Truncate(time.Minute).Format("2006-01-02 15:04")
		return nil
	})
}
func MarkTask(id int, status Status) error {
	return ApplyTaskByID(id, func(task *Task) error {
		task.Status = status
		task.UpdatedAt = time.Now().Truncate(time.Minute).Format("2006-01-02 15:04")
		return nil
	})
}

func ListTasks(status string) error {

	if status != "all" && status != string(Todo) && status != string(Inprogress) && status != string(Done) {
		return fmt.Errorf(ErrInvalidStatus, status)
	}
	tasks, err := ReadTasks()
	if err != nil {
		return err
	}
	for _, v := range tasks {
		if status == "all" || string(v.Status) == status {
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
