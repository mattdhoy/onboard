package crud

import (
	"fmt"

	"github.com/mattdhoy/onboard/engine"
)

// ReadQuestion returns a Question
func ReadQuestion(questionID int) engine.Question {
	db := Connect()
	query := fmt.Sprintf("SELECT * FROM question WHERE question_id = %v", questionID)
	rows, err := db.Query(query)
	CheckErr(err)

	var q engine.Question
	var question_id int
	var quiz_id int
	var problem string
	var answer string

	err = rows.Scan(&question_id, &quiz_id, &problem, &answer)
	CheckErr(err)

	q.Answer = answer
	q.Problem = problem
	q.Wrongs = ReadWrongs(question_id, db)

	db.Close()
	return q
}
