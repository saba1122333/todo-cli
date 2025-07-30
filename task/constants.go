package task

const FileName string = "Tasks.json"

const (
	ErrEmptyList     = "to-do list is empty"
	ErrTaskNotFound  = "task with provided id does not exist"
	ErrInvalidStatus = "invalid status: %s. Valid statuses are: all, todo, in-progress, done"
	ErrMarshal       = "failed to marshal tasks: %w"
	ErrUnMarshal     = "failed to Unmarshal tasks: %w"
	ErrWriteTasks    = "failed to write tasks: %w"
	ErrReadTasks     = "failed to read tasks: %w"
)
