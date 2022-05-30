package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x, y uint64
}

type Matrix struct {
	data [10][10] bool
}

// Converts the number in string s to uint64
// Trimming it in case there are whitespaces
// In case of error exits with rc 1
func readNumber(s string) uint64 {
	trimmed := strings.Trim(s, " ")
	num, err := strconv.ParseInt(trimmed, 10, 0)

	if err != nil {
		fmt.Fprint(os.Stderr, "Expected numbers, got: \"", trimmed, "\"\n")
		fmt.Fprint(os.Stderr, err, "\n")
		os.Exit(1)
	}

	if num < 0 {
		fmt.Fprint(os.Stderr, "Expected positive integers!\n")
		os.Exit(1)
	}

	return uint64(num)
}

// Reads coordinates in the format "%d, %d\n" from standard input
// Returns slice of read coordinates
// In case of error exits with rc 1
func readCoords() []Coord {
	// Read the starting positions
	var ret []Coord
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		sep := strings.Split(scanner.Text(), ",")
		if len(sep) != 2 {
			fmt.Fprint(os.Stderr, "Invalid input format\n")
			os.Exit(1)
		}

		x := readNumber(sep[0])
		y := readNumber(sep[1])
		ret = append(ret, Coord{x, y})
	}
	return ret
}

func (m Matrix) String() string {
	ret := ""
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			if m.data[x][y] {
				ret += "X"
			} else {
				ret += "_"
			}
		}
		ret += "\n"
	}
	return ret
}

func main() {
	m := Matrix{}
	startingCoords := readCoords()

	for _, c := range(startingCoords) {
		m.data[c.x][c.y] = true
	}

	fmt.Println(m)
}
