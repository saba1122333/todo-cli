package task

import "fmt"

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

func (t Task) String() string {
	return fmt.Sprintf(
		"ID: %d\nDescription: %s\nStatus: %s\nCreated At: %s\nUpdated At: %s",
		t.ID, t.Description, t.Status, t.CreatedAt, t.UpdatedAt,
	)
}
