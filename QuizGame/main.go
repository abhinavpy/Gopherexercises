package main

import (
	"time"
	"flag"
	"fmt"
	"os"
	"encoding/csv"
)

func main() {
	correct := 0
	csvFilename := flag.String("csv","problems.csv", "a csv file of the format 'question,answer'")
	timeLimit := flag.Int("limit", 30, "time limit to answer questions in teh quiz")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("ERROR:\tFailed to open %s\n",*csvFilename))
	}
	
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("ERROR: Failed to parse the csv file!")
	}

	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, p := range problems {
		fmt.Printf("Problem #%d: What is %s, sir? = \n", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n",&answer)
			answerCh <- answer	
		}()
		select {
		case <-timer.C:	
			fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
		
	}

	
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
