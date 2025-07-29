package task

import "time"

func Add(description string) error {
	iD := generateID()
	task := Task{
		ID:          iD,
		Description: description,
		Status:      Todo,
		CreatedAt:   time.Now().Truncate(time.Minute).Format("2006-01-02 15:04")}
	return appendTask(task)

}
func Delete(id int) error {
	return DeleteTask(id)
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
