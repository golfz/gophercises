package main

import (
	"fmt"
	"gophercises/01_quiz_game/problems"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	defaultFileName = "problems.csv"
	defaultTimeout  = 30
)

var (
	fileName   string
	useShuffle bool
	timeout    int
	score      int
)

func main() {
	timeout = defaultTimeout

	for i, v := range os.Args {
		if v == "-h" || v == "--help" {
			fmt.Println("Usage: quiz-game [--file=FILENAME] [--timer=TIMEOUT] [--shuffle]")
			return
		}
		if v == "--file" {
			fileName = os.Args[i+1]
		}
		if v == "--timer" {
			sTimeout := os.Args[i+1]
			iTimeout, err := strconv.Atoi(sTimeout)
			if err != nil {
				panic(err)
			}
			timeout = iTimeout
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

	if useShuffle {
		problemList = problems.Shuffle(problemList)
	}

	fmt.Printf("You have %d seconds to answer, press enter for starting: ", timeout)
	fmt.Scanln()

	start := time.Now()

	done := make(chan struct{})
	go AskQuestion(problemList, done)

	select {
	case <-done:
		break
	case <-time.After(time.Duration(timeout) * time.Second):
		fmt.Println("\nTime out!")
	}

	fmt.Printf("Your score is %d/%d\n", score, len(problemList))
	fmt.Printf("Your time is %f seconds", time.Since(start).Seconds())
}

func AskQuestion(problemList []problems.Problem, done chan struct{}) {
	for i, problem := range problemList {
		fmt.Printf("#%d: %d + %d = ", i+1, problem.A, problem.B)
		var ans int
		_, err := fmt.Scanf("%d\n", &ans)
		if err != nil {
			panic(err)
		}

		if ans == problem.Ans {
			score++
		}
	}
	done <- struct{}{}
}
