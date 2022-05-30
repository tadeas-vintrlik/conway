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
	data [10][10]bool
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

// Reads coordinates in the format "%d, %d\n" from file on filename
// Returns slice of read coordinates
// In case of error exits with rc 1
func readCoords(filename string) []Coord {
	// Read the starting positions
	var ret []Coord
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprint(os.Stderr, err, "\n")
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
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

func (m Matrix) isAlive(x, y int) bool {
	if x < 0 || y < 0 {
		return false
	}

	if x >= 10 || y >= 10 {
		return false
	}

	return m.data[x][y]
}

// Will return number of alive neighbours for cell at coordinates x and y
func (m Matrix) aliveNeighbours(x, y int) int {
	ret := 0
	if m.isAlive(x-1, y-1) {
		ret++
	}
	if m.isAlive(x, y-1) {
		ret++
	}
	if m.isAlive(x+1, y-1) {
		ret++
	}
	if m.isAlive(x-1, y) {
		ret++
	}
	if m.isAlive(x+1, y) {
		ret++
	}
	if m.isAlive(x-1, y+1) {
		ret++
	}
	if m.isAlive(x, y+1) {
		ret++
	}
	if m.isAlive(x+1, y+1) {
		ret++
	}
	return ret
}

func (m Matrix) NextTurn() Matrix {
	ret := Matrix{}
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			// 1. Any live cell with two or three live neighbours survives.
			if m.isAlive(x, y) {
				neighbours := m.aliveNeighbours(x, y)
				if neighbours == 2 || neighbours == 3 {
					ret.data[x][y] = true
				}
			}
			// 2. Any dead cell with three live neighbours becomes a live cell.
			if !m.isAlive(x, y) {
				if m.aliveNeighbours(x, y) == 3 {
					ret.data[x][y] = true
				}
			}
			// 3. All other live cells die in the next generation. Similarly, all other dead cells stay dead.
			// No need to do anything as by default all booleans are false in go
		}
	}
	return ret
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "usage: ./conway file\n")
		os.Exit(1)
	}
	m := Matrix{}
	seed := readCoords(os.Args[1])

	for _, c := range seed {
		m.data[c.x][c.y] = true
	}

	gen := 0
	for {
		fmt.Println("Generation ", gen, ":\n")
		gen++
		fmt.Println(m)
		fmt.Println("Pres return for next generation, Ctrl-c to stop.")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		m = m.NextTurn()
	}
}
