package main

import (
	"fmt"
	"strings"
	"strconv"
)

// A ParseError indicates an error in converting a word into an integer.
type PParseError struct {
	Index int    // The index into the space-separated list of words.
	Word  string // The word that generated the parse error.
	Error error  // The raw error that precipitated this error, if any.
}

// String returns a human-readable error message.
func (e *PParseError) String() string {
	return fmt.Sprintf("pkg: error parsing %q as int", e.Word)
}

// Parse parses the space-separated words in in put as integers.
func PParse(input string) (numbers []int) {
	fields := strings.Fields(input)
	numbers = pfields2numbers(fields)
	return
}

func pfields2numbers(fields []string) (numbers []int) {
	if len(fields) == 0 {
		panic("no words to parse")
	}
	for idx, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(&PParseError{idx, field, err})
		}
		numbers = append(numbers, num)
	}
	return
}

func main() {
	var examples = []string{
		"1 2 3 4 5",
		"100 50 25 12.5 6.25",
		"2 + 2 = 4",
		"1st class",
		"",
	}

	for _, ex := range examples {
		fmt.Printf("Parsing %q:\n  ", ex)
		nums := PParse(ex)
		fmt.Println(nums)
	}
}
