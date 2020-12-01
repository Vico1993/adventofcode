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
	ref    int64
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
		val := transformStringIntoInt(val)
		ref = val
		result := is2020Possible(val, false)

		if result > 0 {
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

func is2020Possible(val int64, stop bool) int64 {
	result := ref
	for _, value := range input {
		valueAsInt := transformStringIntoInt(value)
		result = ref * valueAsInt

		cal := val + valueAsInt

		if cal < target && !stop {
			ref = result

			if res := is2020Possible(cal, true); res > 0 {
				return res
			}

			ref = result / valueAsInt
		}

		if cal == target {
			return result
		}
	}

	return 0
}
