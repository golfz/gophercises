package main

import (
	"fmt"
	"gophercises/01_quiz_game/problems"
	"os"
	"strings"
)

const (
	defaultFileName = "problems.csv"
	defaultTimeout  = 30
)

var (
	fileName             string
	useTimer, useShuffle bool
	timeout              int
	score                int
)

func main() {
	for i, v := range os.Args {
		if v == "-h" || v == "--help" {
			fmt.Println("Usage: quiz-game [--file=FILENAME] [--timer=TIMEOUT] [--shuffle]")
			return
		}
		if v == "--file" {
			fileName = os.Args[i+1]
		}
		if v == "--timer" {
			useTimer = true
		}
		if v == "--shuffle" {
			useShuffle = true
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

	problemList, err := problems.GetProblemList(string(dat))
	if err != nil {
		panic(err)
	}

	for _, problem := range problemList {

		fmt.Printf("%d + %d = ", problem.A, problem.B)
		var ans int
		_, err := fmt.Scanf("%d\n", &ans)
		if err != nil {
			panic(err)
		}

		if ans == problem.Ans {
			score++
		}
	}

	fmt.Printf("Your score is %d/%d", score, len(problemList))
}
