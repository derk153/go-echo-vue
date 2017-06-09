package main

import "database/sql"

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS tasks(
      id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
      name VARCHAR NOT NULL
  );
	`

	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}
