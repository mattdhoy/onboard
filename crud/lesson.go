package crud

import (
	"fmt"

	"github.com/mattdhoy/onboard/engine"
)

// ReadLesson takes an id and returns a slice of questions
func ReadLesson(lessonID int) engine.Lesson {
	var quizzes []engine.Quiz
	var IDs []int
	db := Connect()
	query := fmt.Sprintf("SELECT quiz_id FROM quiz WHERE lesson = %v", lessonID)
	rows, err := db.Query(query)
	CheckErr(err)

	for rows.Next() {
		var quiz_id int
		err := rows.Scan(&quiz_id)
		CheckErr(err)
		IDs = append(IDs, quiz_id)
	}
	db.Close()

	for _, v := range IDs {
		quizzes = append(quizzes, ReadQuiz(v))
	}

	lesson := engine.Lesson{Quizzes: quizzes}
	return lesson
}
