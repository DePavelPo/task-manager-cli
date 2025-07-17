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
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	store := &SQLiteStore{db: db}
	if err := store.migrate(); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
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
		INSERT INTO task(title) VALUES (?);
	`

	_, err := s.db.Exec(query, title)
	return err
}

func (s *SQLiteStore) LoadTasks(completed *bool) ([]models.Task, error) {
	baseQuery := `SELECT id, title, completed, created_at FROM task`
	var (
		query string
		args  []any
	)

	if completed != nil {
		query = baseQuery + " WHERE completed = ?"
		args = append(args, boolToInt(*completed))
	} else {
		query = baseQuery
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query tasks: %w", err)
	}
	defer rows.Close()

	tasks := make([]models.Task, 0, 16)
	for rows.Next() {
		var task models.Task
		var completedInt int
		if err := rows.Scan(&task.ID, &task.Title, &completedInt, &task.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan task row: %w", err)
		}
		task.Completed = intToBool(completedInt)
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return tasks, nil
}

func (s *SQLiteStore) UpdateTask(id uint64, completed bool) error {
	query := `
		UPDATE task SET completed = ?
		WHERE id = ?;
	`

	completedInt := boolToInt(completed)
	_, err := s.db.Exec(query, completedInt, id)
	return err
}

func (s *SQLiteStore) DeleteTask(id uint64) error {
	query := `DELETE FROM task WHERE id = ?;`

	_, err := s.db.Exec(query, id)
	return err
}
