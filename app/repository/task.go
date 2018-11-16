package repository

import (
	"database/sql"

	"github.com/luizdepra/go-rest-api/app/model"
)

// TaskRepository implements methods to get, create, update and delete Tasks
// from the database.
type TaskRepository struct {
	db *sql.DB
}

// New creates a new TaskRepository.
func New(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

// List retrieves every stored Task.
func (repo *TaskRepository) List() ([]*model.Task, error) {
	query := `
		SELECT id, title, priority, done
		FROM tasks;
	`

	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}

	var id int64
	var title string
	var priority uint8
	var done bool

	taskList := make([]*model.Task, 0)

	for rows.Next() {
		err = rows.Scan(&id, &title, &priority, &done)
		if err != nil {
			return nil, err
		}

		task := &model.Task{
			ID:       id,
			Title:    title,
			Priority: priority,
			Done:     done,
		}

		taskList = append(taskList, task)
	}
	rows.Close()

	return taskList, nil
}

// Get retrieves one stored Task by its id.
func (repo *TaskRepository) Get(id int64) (*model.Task, error) {
	query := `
		SELECT title, priority, done
		FROM tasks
		WHERE id = ?;
	`

	row := repo.db.QueryRow(query, id)

	var title string
	var priority uint8
	var done bool

	err := row.Scan(&title, &priority, &done)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, nil
		default:
			return nil, err
		}
	}

	task := &model.Task{
		ID:       id,
		Title:    title,
		Priority: priority,
		Done:     done,
	}

	return task, nil
}

// Create stores a new Task.
func (repo *TaskRepository) Create(task *model.Task) (*model.Task, error) {
	query := `
		INSERT INTO tasks (title, priority, done)
		VALUES (?, ?, ?);
	`

	statement, err := repo.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	result, err := statement.Exec(task.Title, task.Priority, task.Done)
	if err != nil {
		return nil, err
	}

	lasID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	createdTask, err := repo.Get(lasID)
	if err != nil {
		return nil, err
	}

	return createdTask, nil
}

// Update refreshes a stored Task with new values.
func (repo *TaskRepository) Update(task *model.Task) (*model.Task, error) {
	query := `
		UPDATE tasks
		SET title = ?,
			priority = ?,
			done = ?
		WHERE id = ?;
	`

	statement, err := repo.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	_, err = statement.Exec(task.Title, task.Priority, task.Done, task.ID)
	if err != nil {
		return nil, err
	}

	updatedTask, err := repo.Get(task.ID)
	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

// Delete removes a stored Task.
func (repo *TaskRepository) Delete(id int64) (*model.Task, error) {
	query := `
		DELETE FROM tasks
		WHERE id = ?;
	`

	deletedTask, err := repo.Get(id)
	if err != nil {
		return nil, err
	}

	statement, err := repo.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	_, err = statement.Exec(id)
	if err != nil {
		return nil, err
	}

	return deletedTask, nil
}
