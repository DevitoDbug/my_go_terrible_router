package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"my_go_terrible_router/pkg/config"
	"time"
)

var Db *sql.DB

type Task struct {
	ID          int64     `json:"id"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	DateCreated time.Time `json:"dateCreated"`
	DateUpdated time.Time `json:"dateUpdated"`
	DateDeleted time.Time `json:"dateDeleted"`
}

func init() {
	err := config.ConnectToDb()
	if err != nil {
		log.Fatalf("Error in connecting to db: %v", err)
		return
	}
	Db = config.GetDBInstance()
}

func (t *Task) CreateTask() error {
	createTaskQuery := `CREATE TABLE IF NOT EXISTS tasks
    (
    id INT AUTO_INCREMENT PRIMARY KEY,
    description VARCHAR(255) NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT false,
    dateCreated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    dateUpdated TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	dateDeleted TIMESTAMP NULL
    );
`
	_, err := Db.Exec(createTaskQuery)
	if err != nil {
		return fmt.Errorf("error creating tasks table: %v", err)
	}

	//populating the db with values
	insertValuesQuery := `INSERT INTO tasks(description, completed,dateCreated,dateUpdated, dateDeleted)
	VALUES (?,FALSE,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP,NULL)
`
	if _, err := Db.Exec(insertValuesQuery, t.Description); err != nil {
		return fmt.Errorf("error inserting values into tasks table: %v", err)
	}
	return nil
}

func GetTasks() ([]Task, error) {
	if Db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	getAllTaskQuery := `SELECT * FROM tasks WHERE dateDeleted IS NULL`

	rows, err := Db.Query(getAllTaskQuery)
	if err != nil {
		return nil, err
	}
	rows.Close()

	//slice to store the returned rows
	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Description, &task.Completed, &task.DateCreated, &task.DateUpdated, &task.DateDeleted)
		if err != nil {
			log.Printf("Error scanning task row: %v", err)
			return nil, err
		}
		tasks = append(tasks, task)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		return nil, err
	}

	return tasks, nil
}

func GetTask(id int) (Task, error) {
	var task Task
	getTaskQueryString := `
		SELECT * FROM tasks WHERE id = ? AND dateDeleted IS NULL 
		`

	row := Db.QueryRow(getTaskQueryString, id)

	err := row.Scan(&task.ID, &task.Description, &task.Completed, &task.DateCreated, &task.DateUpdated, &task.DateDeleted)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Handle the case when no rows are returned (no matching task found)
			return task, fmt.Errorf("task not found for ID: %d", id)
		}
		log.Printf("Error scanning task row: %v", err)
		return task, err
	}
	return task, nil
}

func DeleteTask(id int) error {
	deleteTaskQueryString := `UPDATE tasks SET dateDeleted = current_timestamp WHERE id = ?`

	result, err := Db.Exec(deleteTaskQueryString, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no task found for ID: %d", id)
	}
	return nil
}
