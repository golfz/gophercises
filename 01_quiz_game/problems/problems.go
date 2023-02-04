package problems

import (
	"errors"
	"gophercises/01_quiz_game/csv"
	"strconv"
	"strings"
)

type Problem struct {
	A   int
	B   int
	Ans int
}

func GetProblemList(rawData string) ([]Problem, error) {
	var problemList []Problem

	rows, err := csv.New(rawData)
	if err != nil {
		return nil, err
	}

	for _, columns := range rows {
		if len(columns) != 2 {
			continue
		}
		problem, err := getProblem(columns[0], columns[1])
		if err != nil {
			continue
		}
		problemList = append(problemList, problem)
	}

	return problemList, nil
}

func getProblem(question string, ans string) (Problem, error) {
	var problem Problem

	splitQuestions := strings.Split(question, "+")
	if len(splitQuestions) != 2 {
		return Problem{}, errors.New("question format was wrong")
	}

	sA := strings.TrimSpace(splitQuestions[0])
	iA, err := strconv.Atoi(sA)
	if err != nil {
		return Problem{}, err
	}
	problem.A = iA

	sB := strings.TrimSpace(splitQuestions[1])
	iB, err := strconv.Atoi(sB)
	if err != nil {
		return Problem{}, err
	}
	problem.B = iB

	sAns := strings.TrimSpace(ans)
	iAns, err := strconv.Atoi(sAns)
	if err != nil {
		return Problem{}, err
	}
	problem.Ans = iAns

	if problem.A+problem.B != problem.Ans {
		return Problem{}, errors.New("raw problem format was wrong")
	}

	return problem, nil
}
