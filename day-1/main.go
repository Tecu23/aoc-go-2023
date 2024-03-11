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

		// numbers := make([]int8, 0, 10)
		//
		// for _, c := range line {
		//
		// 	number, err := strconv.Atoi(string(c))
		// 	if err == nil {
		// 		numbers = append(numbers, int8(number))
		// 	}
		// }
		//
		// fmt.Printf("First Character: %d, Last Character %d\n", numbers[0], numbers[len(numbers)-1])
		//
		// sum += int64(numbers[0]*10 + numbers[len(numbers)-1])

		num := part2(line)

		sum += int64(num)

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

func part2(line string) int {
	prefixes := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"zero":  0,
	}

	for i := 0; i <= 9; i++ {
		prefixes[ToString(i)] = i
	}

	var first, last int

	for len(line) > 0 {
		for prefix, val := range prefixes {
			if doesStringHavePrefix(line, prefix) {
				if first == 0 {
					first = val
				}

				last = val
				break
			}
		}
		line = line[1:]
	}

	fmt.Printf("%d %d\n", first, last)

	return first*10 + last
}

func doesStringHavePrefix(str string, prefix string) bool {
	if len(str) < len(prefix) {
		return false
	}
	return str[:len(prefix)] == prefix
}

func ToString(arg interface{}) string {
	var str string
	switch arg.(type) {
	case int:
		str = strconv.Itoa(arg.(int))
	case byte:
		b := arg.(byte)
		str = string(rune(b))
	case rune:
		str = string(arg.(rune))
	default:
		panic(fmt.Sprintf("unhandled type for string casting %T", arg))
	}
	return str
}
