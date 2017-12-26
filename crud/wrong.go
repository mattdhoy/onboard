package crud

import (
	"database/sql"
	"fmt"
)

// ReadWrongs returns a wrong answer
func ReadWrongs(questionID int, db *sql.DB) []string {
	var wrongs []string
	query := fmt.Sprintf("SELECT * FROM wrong WHERE question_id = %v", questionID)
	rows, err := db.Query(query)
	CheckErr(err)
	for rows.Next() {
		var wrong string
		err := rows.Scan(&wrong)
		CheckErr(err)
		wrongs = append(wrongs, wrong)
	}
	return wrongs
}
