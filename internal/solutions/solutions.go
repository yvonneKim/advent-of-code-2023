package solutions

type Solution interface {
	Day() int
	SolveFor(string) string
	GetExampleInput() string
	GetExampleAnswer() string
}
