package day_1

import solutions "advent-of-code-2023/internal/solutions"

type day struct {
	exampleInput  string
	exampleAnswer string
}

func (d *day) Day() int { return 1 }

func (d *day) GetExampleInput() string {
	return d.exampleInput
}

func (d *day) GetExampleAnswer() string {
	return d.exampleAnswer
}

func (d *day) SolveFor(input string) string {
	// TODO: Implement solution
	return d.exampleAnswer
}

func Solution() solutions.Solution {
	return &day{
		exampleInput: `1000
		2000
		3000
		
		4000
		
		5000
		6000
		
		7000
		8000
		9000
		
		10000`,
		exampleAnswer: "24000",
	}
}
