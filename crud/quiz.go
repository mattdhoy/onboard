package crud

import (
	"fmt"

	"github.com/mattdhoy/onboard/engine"
)

// ReadQuiz takes an id and returns a slice of questions
func ReadQuiz(quizID int) engine.Quiz {
	var questions []engine.Question
	var IDs []int
	db := Connect()
	query := fmt.Sprintf("SELECT question_id FROM question WHERE quiz_id = %v", quizID)
	rows, err := db.Query(query)
	CheckErr(err)

	for rows.Next() {
		var question_id int
		err = rows.Scan(&question_id)
		CheckErr(err)
		IDs = append(IDs, question_id)
	}
	db.Close()

	for _, v := range IDs {
		questions = append(questions, ReadQuestion(v))
	}

	quiz := engine.Quiz{Questions: questions}
	return quiz
}
