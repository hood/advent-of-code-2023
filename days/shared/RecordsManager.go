package shared

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/libsql/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

func driver() *sql.DB {
	dbUrl := "file:./records.db"

	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", dbUrl, err)
		os.Exit(1)
	}

	db.Exec("CREATE TABLE IF NOT EXISTS records (id INTEGER PRIMARY KEY AUTOINCREMENT, day INTEGER, part INTEGER, time INTEGER)")

	return db
}

func RegisterRecord(day int, part int, time int64) {
	db := driver()

	_, err := db.Exec("INSERT INTO records (day, part, time) VALUES (?, ?, ?)", day, part, time)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to insert record: %s", err)
		os.Exit(1)
	}
}

func GetRecord(day int, part int) int64 {
	db := driver()

	row := db.QueryRow(
		"SELECT time FROM records WHERE day = ? AND part = ?",
		day,
		part,
	)

	var result int64

	err := row.Scan(&result)
	if err != nil {
		// It's fine if we don't find a record.
		if err.Error() == "sql: no rows in result set" {
			return 0
		}

		fmt.Fprintf(os.Stderr, "Failed to get record: %s", err)

		os.Exit(1)
	}

	return result
}
