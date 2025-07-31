package task

var FileName string = "Tasks.json"

// this should not be const or var but I am new and FileName being const was not cutting it
//  for next time I will write tests along the way which will help me see a better way

const (
	ErrEmptyList     = "to-do list is empty"
	ErrTaskNotFound  = "task with provided id does not exist"
	ErrInvalidStatus = "invalid status: %s. Valid statuses are: all, todo, in-progress, done"
	ErrMarshal       = "failed to marshal tasks: %w"
	ErrUnMarshal     = "failed to Unmarshal tasks: %w"
	ErrWriteTasks    = "failed to write tasks: %w"
	ErrReadTasks     = "failed to read tasks: %w"
)
const (
	SucTaskAdded   = "Task Added successfully (ID:%v)"
	SucTaskDeleted = "Task Deleted successfully (ID:%v)"
	SucTaskUpdated = "Task Updated successfully (ID:%v)"
	SucTaskMarked  = "Task Marked successfully (ID:%v)"
)
