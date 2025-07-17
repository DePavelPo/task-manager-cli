package storage

import (
	"database/sql"
	"fmt"

	"github.com/DePavelPo/task-manager-cli/models"
)

type SQLiteStore struct {
	db *sql.DB
}

func NewSQLiteStore(dbPath string) (*SQLiteStore, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	store := &SQLiteStore{db: db}
	if err := store.migrate(); err != nil {
		return nil, err
	}

	return store, nil
}

func (s *SQLiteStore) CloseDB() error {
	return s.db.Close()
}

func (s *SQLiteStore) migrate() error {
	query := `
		CREATE TABLE IF NOT EXISTS task
		(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			completed INTEGER NOT NULL DEFAULT 0,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`

	_, err := s.db.Exec(query)
	return err
}

func (s *SQLiteStore) SaveTask(title string) error {
	query := `
		insert into task(title) values (?);
	`

	_, err := s.db.Exec(query, title)
	return err
}

func (s *SQLiteStore) LoadTasks(completed *bool) ([]models.Task, error) {
	query := `SELECT id, title, completed, created_at FROM task`

	if completed != nil {
		completedInt := boolToInt(*completed)
		query += fmt.Sprint(" where completed =", completedInt)
	}

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]models.Task, 0)
	for rows.Next() {
		var task models.Task
		var completedInt int
		err := rows.Scan(&task.ID, &task.Title, &completedInt, &task.CreatedAt)
		if err != nil {
			return nil, err
		}

		task.Completed = intToBool(completedInt)
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *SQLiteStore) UpdateTask(id uint64, completed bool) error {
	query := `
		update task set completed = ?
			where id = ?;
	`

	completedInt := boolToInt(completed)
	_, err := s.db.Exec(query, &completedInt, &id)
	return err
}

func (s *SQLiteStore) DeleteTask(id uint64) error {
	query := `delete from task where id = ?;`

	_, err := s.db.Exec(query, &id)
	return err
}
