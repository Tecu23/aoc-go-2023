package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

		numbers := make([]int8, 0, 10)

		for _, c := range line {

			number, err := strconv.Atoi(string(c))
			if err == nil {
				numbers = append(numbers, int8(number))
			}
		}

		fmt.Printf("First Character: %d, Last Character %d\n", numbers[0], numbers[len(numbers)-1])

		sum += int64(numbers[0]*10 + numbers[len(numbers)-1])

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
