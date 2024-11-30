package main

import (
	"encoding/csv"
	"flag"
	"fmt"
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

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	var count int
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d.\n", count, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.a {
				count++
			}
		}

		fmt.Println("total correct answer is ", count)
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
