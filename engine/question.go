package engine

// Question is a multiple choice question that will be used by quizzes and exams.
type Question struct {
	Problem string
	Answer  string
	Wrongs  []string
}

func newQuestion(prob string, ans string, wrongs []string) *Question {
	q := &Question{
		Problem: prob,
		Answer:  ans,
		Wrongs:  wrongs,
	}
	return q
}
