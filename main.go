package main

import (
	"advent-of-code-2023/internal/app"
	"advent-of-code-2023/internal/solutions"
	"advent-of-code-2023/internal/solutions/day_1"
)

func main() {
	app := app.InitializeApp()
	app.LoadCookie()

	for _, s := range []solutions.Solution{
		day_1.Solution(),
	} {
		if s.SolveFor(s.GetExampleInput()) == s.GetExampleAnswer() {
			input := app.GetInputForDay(s.Day())
			answer := s.SolveFor(input)

			app.SubmitAnswer(s.Day(), answer)
		}
	}
}
