package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage:", os.Args[0], "<filename> <solver>")
		fmt.Println("    solver = 1: naive")
		fmt.Println("    solver = 2: greedy")
		return
	}

	// parse input
	file, err := os.Open(os.Args[1])
	check(err)
	defer file.Close()
	reader := bufio.NewReader(file)
	width, height, painted := parse(reader)

	// solve input
	var operations []string
	solver, err := strconv.Atoi(os.Args[2])
	check(err)
	switch (solver) {
	case 1:
		operations = naive(width, height, painted)
	case 2:
		operations = greedy(width, height, painted)
	}

	// write output
	output, err := os.Create("output.txt")
	check(err)
	defer output.Close()
	writer := bufio.NewWriter(output)
	for i := 0; i < len(operations); i++ {
		writer.WriteString(operations[i])
		writer.Write([]byte{10}) // LF byte
	}
	writer.Flush()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLine(reader *bufio.Reader) string {
	line, err := reader.ReadString('\r')
	check(err)
	_, err = reader.ReadString('\n')
	check(err)
	return line[:len(line)-1] // remove trailing \r
}

func parse(reader *bufio.Reader) (int, int, [][]bool) {
	l1 := readLine(reader)
	l1s := strings.Split(l1, ",")
	width, err := strconv.Atoi(l1s[0])
	check(err)
	height, err := strconv.Atoi(l1s[1])
	check(err)

	painted := make([][]bool, height)
	for i := 0; i < height; i++ {
		painted[i] = make([]bool, width)
		line := readLine(reader)
		for j := 0; j < width; j++ {
			if line[j] == 35 { // byte code for #
				painted[i][j] = true
			}
		}
	}

	return width, height, painted
}
