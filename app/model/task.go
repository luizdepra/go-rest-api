package model

// TaskRepository is a dummy list of Task.
var TaskRepository = []*Task{
	&Task{
		ID:       1,
		Title:    "Just Do It!",
		Priority: 0,
		Done:     false,
	},
	&Task{
		ID:       2,
		Title:    "Why do it today when it can be done tomorrow...",
		Priority: 10,
		Done:     false,
	},
	&Task{
		ID:       3,
		Title:    "Already done!",
		Priority: 1,
		Done:     true,
	},
}

// Task represets a task.
type Task struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Priority uint8  `json:"priority"`
	Done     bool   `json:"done"`
}
