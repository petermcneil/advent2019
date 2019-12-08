package main

import (
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

func calcFuel(mass int64) int64 {
	return (mass / 3) - 2
}

func calcFuel2(m int64) int64 {
	var total int64
	newFuel := calcFuel(m)
	for newFuel >= 0 {
		total += newFuel
		newFuel = calcFuel(newFuel)
	}
	return total
}

func main() {
	var input = flag.String("input", "", "Input for challenge")
	flag.Parse()

	dat, err := ioutil.ReadFile(*input)
	check(err)

	var count1 int64 = 0
	var count2 int64 = 0

	for _, line := range strings.Split(string(dat), "\n") {
		if line == "" {
			continue
		}
		mass, err := strconv.ParseInt(line, 10, 0)
		check(err)
		count1 += calcFuel(mass)
	}
	fmt.Printf("Part One: %v\n", count1)

	for _, line := range strings.Split(string(dat), "\n") {
		if line == "" {
			continue
		}
		mass, err := strconv.ParseInt(line, 10, 0)
		check(err)
		count2 += calcFuel2(mass)
	}
	fmt.Printf("Part Two: %v\n", count2)
}
