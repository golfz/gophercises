package main

import (
	"fmt"
	"os"
	"strings"
)

const defaultFileName = "problems.csv"

type question struct {
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

func getQuestions(problemData string) ([]question, error) {
	//problems := strings.Split(problemData, "\n")
	//for i, v := range problems {
	//
	//}
	return []question{}, nil
}

func getQuestion(problem string) (question, error) {
	//ans := strings.Split(problem, ",")
	return question{a: 1, b: 2, ans: 3}, nil
}
