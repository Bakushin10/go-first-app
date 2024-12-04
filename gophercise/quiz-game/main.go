package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	// "go/printer"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {

	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of question,answer")

	timeLimit := flag.Int("limit", 2, "the limit for the quiz is seconds")
	flag.Parse()

	_ = csvFileName
	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Printf("Failed to open the CSV file : %s", *csvFileName)
		os.Exit(1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		fmt.Printf("Failed to Parse the provided CSV file")
		os.Exit(1)
	}

	problems := parseLines(lines)
	fmt.Println(problems)

	answerCh := make(chan string)
	count := 0

	go func() {
		for {
			var answer string
			if _, err := fmt.Scanf("%s\n", &answer); err == nil {
				answerCh <- answer
			} else {
				close(answerCh)
				return
			}
		}
	}()

	for i, p := range problems {

		fmt.Printf("problem #%d %s = ", i+1, p.q)
		timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d.\n", count, len(problems))
			os.Exit(0) // finish the program

		case answer := <-answerCh:
			if answer == p.a {
				count++
			}
		}

	}
}

func parseLines(lines [][]string) []problem {

	problems := []problem{}

	for _, val := range lines {
		problems = append(problems, problem{
			q: val[0],
			a: strings.TrimSpace(val[1]),
		})
	}

	return problems
}
