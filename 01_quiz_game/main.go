package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const defaultFileName = "problems.csv"

type problem struct {
	a   int
	b   int
	ans int
}

func main() {
	var fileName string
	for i, v := range os.Args {
		if v == "-f" {
			fileName = os.Args[i+1]
			break
		}
	}

	fileName = strings.TrimSpace(fileName)
	if fileName == "" {
		fmt.Printf("No specify a file name, use default file: %s\n", defaultFileName)
		fileName = defaultFileName
	}

	dat, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dat))
}

func getProblemList(problemData string) ([]problem, error) {
	var problemList []problem

	rawProblemList := strings.Split(problemData, "\n")

	for _, v := range rawProblemList {
		prob, err := getProblem(v)
		if err != nil {
			continue
		}
		problemList = append(problemList, prob)
	}

	return problemList, nil
}

func getProblem(sProblem string) (problem, error) {
	// example: 1+2,3
	sProblem = strings.TrimSpace(sProblem)
	if sProblem == "" {
		return problem{}, errors.New("problem string is empty")
	}

	splitAns := strings.Split(sProblem, ",")
	if len(splitAns) != 2 {
		return problem{}, errors.New("raw problem format was wrong")
	}
	sAns := strings.TrimSpace(splitAns[1])
	ans, err := strconv.Atoi(sAns)
	if err != nil {
		return problem{}, err
	}

	sProblem = strings.TrimSpace(splitAns[0])
	splitProblem := strings.Split(sProblem, "+")
	if len(splitAns) != 2 {
		return problem{}, errors.New("raw problem format was wrong")
	}
	sA := strings.TrimSpace(splitProblem[0])
	a, err := strconv.Atoi(sA)
	if err != nil {
		return problem{}, err
	}

	sB := strings.TrimSpace(splitProblem[1])
	b, err := strconv.Atoi(sB)
	if err != nil {
		return problem{}, err
	}

	if a+b != ans {
		return problem{}, errors.New("raw problem format was wrong")
	}

	return problem{a: a, b: b, ans: ans}, nil
}
