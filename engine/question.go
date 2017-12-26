package engine

// Question is a multiple choice question that will be used by quizzes and exams.
type question struct {
	problem string
	answer  string
	wrongs  []string
}

func newQuestion(prob string, ans string, wrongs []string) *question {
	q := &question{
		problem: prob,
		answer:  ans,
		wrongs:  wrongs,
	}
	return q
}
