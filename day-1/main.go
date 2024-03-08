package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
)

func main() {
	// Read line by line
	f, err := os.Open("input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var sum int64

	for scanner.Scan() {
		line := scanner.Text()

		start := 0
		end := utf8.RuneCountInString(line) - 1

		var first *int = nil
		var last *int = nil

		for range line {
			if first != nil && last != nil {
				break
			}

			firstNumber, err := strconv.Atoi(string(line[start]))
			if err == nil && first == nil {
				first = &firstNumber
			}

			lastNumber, err := strconv.Atoi(string(line[end]))
			if err == nil && last == nil {
				last = &lastNumber
			}

			start++
			end--
		}

		fmt.Printf("First Character: %d, Last Character %d\n", *first, *last)

		if first != nil && last != nil {
			sum += int64((*first)*10 + (*last))
		}

	}
	fmt.Printf("First character: %d \n", sum)

	err = scanner.Err()
	check(err)

	// find first and last number in each
	// sum all of them

	// print the result
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
