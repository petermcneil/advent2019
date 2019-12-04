package main

import (
	"flag"
	"strconv"
	"strings"
)

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

func isAscending(str string) bool {
	var previous uint8 = 0
	for i:= 0; i < len(str); i++ {
		var char = str[i]
		if previous == char {
			return true
		}
		previous = char
	}
	return false
}

func hasDoubleDigits(str string) bool {
	var previous uint8 = 0
	for i:= 0; i < len(str); i++ {
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

func main() {
	var input = flag.String("input", "", "Input for challenge")
	flag.Parse()

	var split = strings.Split(*input, "-")

	var start, _ = strconv.Atoi(split[0])
	var end, _ = strconv.Atoi(split[1])

	var count = 0
	for i := start; i < end; i++ {
		var stringI = strconv.Itoa(i)
		var descending = isAscending(stringI)

		if descending {
			var double = hasDoubleDigits(stringI)
			if double {
				count++
			}
		}
	}

	println(count)
}
