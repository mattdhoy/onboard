package crud

import (
	"fmt"

	"github.com/mattdhoy/onboard/engine"
)

// ReadCourse takes an id and returns a slice of questions
// TODO: For all except wrong, the selection takes two DB
// connections, and can be reduced to one by using db
// passthrough methods.
func ReadCourse(courseID int) engine.Course {
	var lessons []engine.Lesson
	var IDs []int
	db := Connect()
	query := fmt.Sprintf("SELECT lesson_id FROM lesson WHERE course_id = %v", courseID)
	rows, err := db.Query(query)
	CheckErr(err)

	for rows.Next() {
		var lesson_id int
		err := rows.Scan(&lesson_id)
		CheckErr(err)
		IDs = append(IDs, lesson_id)
	}
	db.Close()

	for _, v := range IDs {
		lessons = append(lessons, ReadLesson(v))
	}

	course := engine.Course{Lessons: lessons}
	return course
}
