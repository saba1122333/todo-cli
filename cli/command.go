package cli

import (
	"fmt"
	"github.com/saba1122333/todo-cli/task"
	"os"
	"strconv"
	"strings"
)

func HandleDeleteCommand() error {

	if len(os.Args) < 3 {
		return fmt.Errorf("missing task ID")
	}
	idStr := strings.TrimSpace(os.Args[2])
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("Invalid ID, please provide a valid number")
	}

	return task.DeleteTask(id)
}

func HandleAddCommand() error {
	if len(os.Args) != 3 {
		return fmt.Errorf("Wrong number of Arguments")
	}
	description := strings.TrimSpace(os.Args[2])
	if len(description) == 0 {
		return fmt.Errorf("Description cannot be empty.")
	}
	return task.AppendTask(description)
}

func HandleUpdateCommand() error {
	if len(os.Args) != 4 {
		return fmt.Errorf("Wrong number of arguments")
	}
	idStr, description := strings.TrimSpace(os.Args[2]), strings.TrimSpace(os.Args[3])
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("Invalid ID, please provide valid number")
	}
	return task.UpdateTask(id, description)
}
func HandleListCommand() error {
	if len(os.Args) == 2 {
		return task.ListTasks("all")
	}
	if len(os.Args) == 3 {
		status := os.Args[2]
		return task.ListTasks(status)
	}
	return fmt.Errorf("Wrong number of arguments")
}

func HandleMarkInProgressCommand() error {
	if len(os.Args) != 3 {
		return fmt.Errorf("Wrong number of arguments")
	}
	idStr := strings.TrimSpace(os.Args[2])
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("Invalid ID, please provide valid number")
	}
	return task.MarkTaskInProgress(id)
}

func HandleMarkDoneCommand() error {
	if len(os.Args) != 3 {
		return fmt.Errorf("Wrong number of arguments")
	}
	idStr := strings.TrimSpace(os.Args[2])
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("Invalid ID, please provide valid number")
	}
	return task.MarkTaskDone(id)
}
