package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func generateOutput(input []int)  ([]int, error) {
	for i := 0; i < len(input); i += 4 {
		opcode := input[i]
		if opcode == 99 {
			break
		}
		first := input[i+1]
		second := input[i+2]
		third := input[i+3]

		if third > len(input) || second > len(input) || first > len(input) {
			return nil, errors.New("input length")
		}

		if opcode == 1 {
			input[third] = input[first] + input[second]
		} else if opcode == 2 {
			input[third] = input[first] * input[second]
		}
	}

	return input, nil
}

func part2(saved []int) (int, error) {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			var copied = make([]int, len(saved))
			copy(copied, saved)
			copied[1] = noun
			copied[2] = verb

			output2, err := generateOutput(copied)
			if err != nil {
				continue
			}

			if output2[0]  == 19690720 {
				fmt.Printf("100 * %v + %v\n", noun, verb)
				return (100 * noun) + verb, nil
			}
		}
	}

	return 0, errors.New("")
}

func main() {
	var inputF = flag.String("input", "", "Input for challenge")
	flag.Parse()
	dat, err := ioutil.ReadFile(*inputF)
	check(err)

	var input = []int{}
	for _, line := range strings.Split(string(dat), "\n") {
		ls := strings.Split(line, ",")
		for _, i := range ls {
			if i != "" {
				j, err := strconv.Atoi(i)
				check(err)
				input = append(input, j)
			}
		}
	}

	var output1 = make([]int, len(input))

	var saved = make([]int, len(input))
	copy(saved, input)
	input[1] = 12
	input[2] = 2


	output1, err = generateOutput(input)
	check(err)

	fmt.Printf("Part One: %v\n", output1[0])

	output2, err := part2(saved)

	if err == nil {
		fmt.Printf("Part Two: %v\n", output2)
	}

}
