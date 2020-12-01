package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const some = 2020

func main() {
	// Read file
	b, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	// Convert bytes into proper collection
	var input []string
	input = strings.Split(string(b), "\n")

	for _, val := range input {
		val := transformStringIntoInt(val)
		res, secondVal := is2020Possible(val, input)
		if res {
			fmt.Println(fmt.Sprintf("Result is %d", (val * secondVal)))
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

func is2020Possible(val int64, input []string) (bool, int64) {
	for _, value := range input {
		valueAsInt := transformStringIntoInt(value)
		if val+valueAsInt == some {
			return true, valueAsInt
		}
	}

	return false, 0
}
