package models

import "database/sql"

// Task a struct containcs Task data
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TaskCollection collection of tasks
type TaskCollection struct {
	Tasks []Task `json:"items"`
}

// GetTasks retrieves tasks from the DB
func GetTasks(db *sql.DB) TaskCollection {
	sql := "SELECT * FROM tasks"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := TaskCollection{}
	for rows.Next() {
		task := Task{}
		err2 := rows.Scan(&task.ID, &task.Name)

		if err2 != nil {
			panic(err2)
		}
		result.Tasks = append(result.Tasks, task)
	}

	return result
}

// PutTask inserts task in DB table
func PutTask(db *sql.DB, name string) (int64, error) {
	sql := "INSERT INTO tasks(name) VALUES(?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	result, err2 := stmt.Exec(name)
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

// DeleteTasks removes record from task table
func DeleteTasks(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM tasks WHERE id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	result, err2 := stmt.Exec(id)
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}
