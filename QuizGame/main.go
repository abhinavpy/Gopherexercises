package main

import (
	"flag"
	"fmt"
	"os"
	"encoding/csv"
)

func main() {
	correct := 0
	csvFilename := flag.String("csv","problems.csv", "a csv file of the format 'question,answer'")
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

	for i, p := range problems {
		fmt.Printf("Problem #%d: What is %s, sir? = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			fmt.Println("Correct")
			correct++
		} else {
			fmt.Println("Incorrect")
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
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
