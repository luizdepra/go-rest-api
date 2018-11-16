package model

// Task represets a task.
type Task struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Priority uint8  `json:"priority"`
	Done     bool   `json:"done"`
}
