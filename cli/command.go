package cli

import (
	"fmt"
	"github.com/saba1122333/todo-cli/task"
	"os"
	"strconv"
	"strings"
)

func parseID(argIndex int) (int, error) {
	if len(os.Args) <= argIndex {
		return 0, fmt.Errorf(ErrMissingID)
	}
	idStr := strings.TrimSpace(os.Args[argIndex])
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return -1, fmt.Errorf(ErrInvalidNumber)
	}
	return id, nil
}
func validArgumentCount(argCount int) error {
	if len(os.Args) != argCount {
		return fmt.Errorf(ErrWrongNumberArguments)
	}
	return nil
}
func getDescription(argIndex int) (string, error) {

	if len(os.Args) <= argIndex {
		return "", fmt.Errorf(ErrDescription)
	}
	description := strings.TrimSpace(os.Args[argIndex])
	if len(description) == 0 {
		return "", fmt.Errorf(ErrDescription)
	}
	return description, nil
}
func parseCommand(argIndex int) (string, error) {
	if len(os.Args) <= argIndex {
		return "", fmt.Errorf(ErrMissingCommand)
	}
	command := strings.ToLower(os.Args[argIndex])
	command = strings.TrimSpace(command)
	return command, nil
}

func HandleDeleteCommand() error {
	if err := validArgumentCount(3); err != nil {
		return err
	}
	id, err := parseID(2)
	if err != nil {
		return err
	}
	return task.DeleteTask(id)
}
func HandleAddCommand() error {
	if err := validArgumentCount(3); err != nil {
		return err
	}
	description, err := getDescription(2)
	if err != nil {
		return err
	}
	return task.AppendTask(description)
}

func HandleUpdateCommand() error {

	if err := validArgumentCount(4); err != nil {
		return err
	}
	id, err := parseID(2)

	if err != nil {
		return err
	}
	description, err := getDescription(3)

	if err != nil {
		return err
	}
	return task.UpdateTask(id, description)
}
func HandleListCommand() error {

	if err := validArgumentCount(2); err == nil {
		_, err := task.ListTasks("all")
		return err
	}
	if err := validArgumentCount(3); err == nil {
		status := os.Args[2]
		_, err := task.ListTasks(status)
		return err
	}
	return nil
}

func HandleMarkCommand(command string) error {

	if err := validArgumentCount(3); err != nil {
		return err
	}
	id, err := parseID(2)
	if err != nil {
		return err
	}
	switch command {
	case "mark-in-progress":
		return task.MarkTask(id, task.Inprogress)
	case "mark-done":
		return task.MarkTask(id, task.Done)
	}
	return fmt.Errorf(ErrInvalidMarkCommand, command)
}

func Run() error {

	command, err := parseCommand(1)
	if err != nil {
		return err
	}
	switch command {
	case "add":
		err = HandleAddCommand()
	case "delete":
		err = HandleDeleteCommand()
	case "update":
		err = HandleUpdateCommand()
	case "list":
		err = HandleListCommand()
	case "mark-in-progress", "mark-done":
		err = HandleMarkCommand(command)
	default:
		err = fmt.Errorf("unknown command: %s", command)
	}
	return err
}
