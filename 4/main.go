package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

/*
 PART ONE
*/

/*
However, they do remember a few key facts about the password:

    It is a six-digit number.
    The value is within the range given in your puzzle input.
    Two adjacent digits are the same (like 22 in 122345).
    Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).

Other than the range rule, the following are true:

    111111 meets these criteria (double 11, never decreases).
    223450 does not meet these criteria (decreasing pair of digits 50).
    123789 does not meet these criteria (no double).
*/

func partOneHasDoubleDigits(str string) bool {
	var previous uint8 = 0
	for i := 0; i < len(str); i++ {
		var char = str[i]
		if previous == char {
			return true
		}
		previous = char
	}
	return false
}

func isAscending(str string) bool {
	var previous uint8 = 0
	for i := 0; i < len(str); i++ {
		var char = str[i]
		if previous < char {
			previous = char
		}

		if previous > char {
			return false
		}
	}
	return true
}

/*
 PART TWO
*/

/*
Given this additional criterion, but still ignoring the range rule, the following are now true:

    112233 meets these criteria because the digits never decrease and all repeated digits are exactly two digits long.
    123444 no longer meets the criteria (the repeated 44 is part of a larger group of 444).
    111122 meets the criteria (even though 1 is repeated more than twice, it still contains a double 22).
*/

func partTwoHasDoubleDigits(str string) bool {
	var previous uint8 = 0
	// Store repeated characters
	var repeated = make(map[uint8]int)

	for i := 0; i < len(str); i++ {
		var char = str[i]
		if previous == char {
			if val, ok := repeated[char]; ok {
				repeated[char] = val + 1
			} else {
				repeated[char] = 1
			}
		}
		previous = char
	}

	// Extract a list of values
	v := make([]int, 0, len(repeated))

	for  _, value := range repeated {
		v = append(v, value)
	}

	// Find if multiple characters only occur once
	for _, b := range v {
		if b == 1 {
			return true
		}
	}

	return false
}

func main() {
	var input = flag.String("input", "", "Input for challenge")
	flag.Parse()

	var split = strings.Split(*input, "-")

	var start, _ = strconv.Atoi(split[0])
	var end, _ = strconv.Atoi(split[1])

	var partOne = 0
	for i := start; i < end; i++ {
		var stringI = strconv.Itoa(i)
		var ascending = isAscending(stringI)

		if ascending {
			var double = partOneHasDoubleDigits(stringI)
			if double {
				partOne++
			}
		}
	}

	fmt.Printf("Part One: %d\n", partOne)

	var partTwo = 0
	for i := start; i < end; i++ {
		var stringI = strconv.Itoa(i)
		var ascending = isAscending(stringI)

		if ascending {
			var double = partTwoHasDoubleDigits(stringI)
			if double {
				partTwo++
			}
		}
	}

	fmt.Printf("Part Two: %d", partTwo)
}
