package storage

import (
	"database/sql"

	"github.com/DePavelPo/task-manager-cli/models"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func OpenDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./task-manager.db")
	if err != nil {
		logrus.Fatalf("open sqlite3 db error: %v", err)
	}

	return db
}

func CloseDB(db *sql.DB) {
	db.Close()
}

func Migrate(db *sql.DB) {
	query := `
		create table if not exists task
		(
			id integer primary key autoincrement,
			title text not null,
			completed integer not null default 0,
			created_at datetime not null default CURRENT_TIMESTAMP
		);
	`

	if _, err := db.Exec(query); err != nil {
		logrus.Fatalf("migrate sqlite3 db error: %v", err)
	}
}

func InsertTask(title string, db *sql.DB) error {
	query := `
		insert into task(title) values (?);
	`

	if _, err := db.Exec(query, title); err != nil {
		return err
	}

	return nil
}

func SelectTasks(db *sql.DB) ([]models.Task, error) {
	query := `
		select id, title, completed, created_at from task;
	`

	rows, err := db.Query(query)
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

func intToBool(i int) bool {
	return i != 0
}

func MarkTask(id uint64, completed bool, db *sql.DB) error {
	query := `
		update task set completed = ?
			where id = ?;
	`

	if _, err := db.Exec(query, &id, &completed); err != nil {
		return err
	}

	return nil
}
