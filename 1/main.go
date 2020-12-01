package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	target int64 = 2020
	input  []string
	result []int64
)

func main() {
	// Read file
	b, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	// Convert bytes into proper collection
	input = strings.Split(string(b), "\n")

	for _, val := range input {
		result = result[:0]
		val := transformStringIntoInt(val)
		result = append(result, val)

		if is2020Possible(val, false) {
			fmt.Println("Result:")
			fmt.Println(result)

			break
		}
	}
}

func transformStringIntoInt(val string) int64 {
	integer, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		panic(err)
	}

	return integer
}

func is2020Possible(val int64, stop bool) bool {
	for _, value := range input {
		valueAsInt := transformStringIntoInt(value)

		if len(result) == 2 && !stop {
			result[len(result)-1] = valueAsInt
		} else if len(result) == 3 {
			result[len(result)-1] = valueAsInt
		} else {
			result = append(result, valueAsInt)
		}

		cal := val + valueAsInt

		if (cal < target && !stop && is2020Possible(cal, true)) || cal == target {
			return true
		}
	}

	return false
}
