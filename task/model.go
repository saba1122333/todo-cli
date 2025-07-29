package task

type Status string

const (
	Todo       Status = "todo"
	Inprogress Status = "in-progress"
	Done       Status = "done"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
