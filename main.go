// This program is a small implementation of the seq command. Seq
// prints a sequence of numbers. This program does not include all
// the features in seq.
// Created by Emile O. E. Antat (eoea)
//
// Globals as variable:
//
// - `padded_length`
//   - Value is changed only when the -w flag is passed to main().
package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

var (
	padded_length int
)

func seq(min, max float64, incr ...float64) {
	// Prints a sequence of numbers, one per line (default),
	// from first (default 1), to near last as possible, in
	// increments of incr (default 1). When first is larger than
	// last, the default incr is -1. All numbers are represented
	// as floating point. Normally integer values are printed as
	// decimal integers if the -w flag is not passed.

	var _incr float64 = 1
	if len(incr) != 0 {
		_incr = incr[0]
	}

	switch min > max {
	case true:
		for min >= max {
			fmt.Printf("%0[2]*[1]v\n", min, padded_length)
			min -= _incr
		}
	default:
		for min <= max {
			fmt.Printf("%0[2]*[1]v\n", min, padded_length)
			min += _incr
		}
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("invalid floating point argument: %v", err)
	}
}

func parseHelper(input []string) {
	// Parses the input and calls the seq for the input.
	switch len(input) {
	case 1:
		min, err := strconv.ParseFloat(input[0], 64)
		checkErr(err)
		seq(min, min)
	case 2:
		min, err := strconv.ParseFloat(input[0], 64)
		checkErr(err)
		max, err := strconv.ParseFloat(input[1], 64)
		checkErr(err)
		seq(min, max)
	case 3:
		min, err := strconv.ParseFloat(input[0], 64)
		checkErr(err)
		_incr, err := strconv.ParseFloat(input[1], 64)
		checkErr(err)
		max, err := strconv.ParseFloat(input[2], 64)
		checkErr(err)
		incr := make([]float64, 1)
		incr[0] = _incr
		seq(min, max, incr...)
	default:
		fmt.Println("usage: goseq [-w] [first [incr] last]")
	}
}

var (
	fixedWidth = flag.Bool("w", false,
		"Left pads the width of numbers with zeros to make them\n"+
			"uniform.")
)

func main() {
	flag.Parse()
	input := flag.Args()

	if *fixedWidth {
		var min string
		var max string
		if len(input) == 2 {
			min = input[0]
			max = input[1]
		}
		if len(input) == 3 {
			min = input[0]
			max = input[2]
		}

		if len(min) > len(max) {
			padded_length = len(min)
		} else {
			padded_length = len(max)
		}
		parseHelper(input)
	} else {
		parseHelper(input)
	}
}
