package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MAX_RED   = 12
	MAX_GREEN = 13
	MAX_BLUE  = 14
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var sum int64

	for scanner.Scan() {
		line := scanner.Text()

		game_details := strings.Split(line, ":")

		sets := strings.Split(game_details[1], ";")

		// ok := true
		var max_r, max_g, max_b int

		for _, s := range sets {
			cubes := strings.Split(s, ",")

			for _, c := range cubes {
				values := strings.Split(c, " ")

				switch values[2] {
				case "blue":
					value, _ := strconv.Atoi(values[1])
					if value > max_b {
						max_b = value
					}
					// if value > MAX_BLUE {
					// 	ok = false
					// 	break
					// }
				case "green":
					value, _ := strconv.Atoi(values[1])
					if value > max_g {
						max_g = value
					}
					// if value > MAX_GREEN {
					// 	ok = false
					// 	break
					// }
				case "red":
					value, _ := strconv.Atoi(values[1])
					if value > max_r {
						max_r = value
					}
					// if value > MAX_RED {
					// 	ok = false
					// 	break
					// }

				}

			}

			// if !ok {
			// 	break
			// }
		}
		// if ok {
		// 	id := strings.Split(game_details[0], " ")
		//
		// 	value, _ := strconv.Atoi(id[1])
		//
		// 	sum += int64(value)
		// }

		fmt.Printf("r:%d g:%d b:%d\n", max_r, max_g, max_b)
		prod := int64(max_b) * int64(max_g) * int64(max_r)
		sum += prod

	}

	fmt.Print(sum)
}
