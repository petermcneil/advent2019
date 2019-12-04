package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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

func main() {
	var input = flag.String("input", "", "Input for challenge")
	flag.Parse()

	var err error

	dat, err := ioutil.ReadFile(*input)
	if err != nil {
		log.Fatal(err)
	}

	var count int64 = 0

	for _, line := range strings.Split(string(dat), "\n") {
		if line == "" {
			continue
		}
		mass, err := strconv.ParseInt(line, 10, 0)
		check(err)
		count += calcFuel(mass)
	}
	fmt.Printf("Part One: %v\n", count)
}
