/*
Package goreadmatrix packet to read matrix from a file or through the terminal
*/
package goreadmatrix

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

/*
Termread read an matrix through the terminal.
Where each row read is a row of the matrix with the elements separated by space.

Example:

# Enter the number of rows: 2
# Enter the number of columns: 3
#
# linha 1: 2.4 6.8 7.2
# linha 2: 6.4 5.0 6.3

return:
n int: number of matrix rows
m int: number of matrix columns
a [][]float64: body of matrix
*/
func Termread() (int, int, [][]float64) {
	var a [][]float64
	var n, m int
	var line []string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the number of rows: ")
	fmt.Scan(&n)
	m = n + 1
	fmt.Println()

	for i := 0; i < n; i++ {
		fmt.Printf("line %d: ", i+1)
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1]
		line = strings.Split(string(text), " ")
		a = append(a, lineToFloat(line))
	}

	return n, m, a
}

/*
Fileread reads an matrix from a file where each line of the file is a row of the array,
with the columns separated by space.

File of matrix 3x4 example:

2.4 6.8 7.2 1.0
6.4 5.0 6.3 2.0
3.0 5.6 4.0 8.7Fileread

param:
fname string: the name of the file that contains the matrix

return:
n int: number of matrix rows
m int: number of matrix columns
a [][]float64: body of matrix
*/
func Fileread(fname string) (int, int, [][]float64) {
	var a [][]float64
	var line []string
	bs, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lines := strings.Split(string(bs), "\n")

	for _, text := range lines {
		// if has a blank lines in file jump it
		if text == "" {
			continue
		}
		line = strings.Split(string(text), " ")
		a = append(a, lineToFloat(line))
	}

	n, m := dim(a)

	return n, m, a
}

/*
lineToFloat transforms an array of strings into an array of float64

param:
line: []string

return:
s: []float64
*/
func lineToFloat(line []string) []float64 {
	var s = []float64{}

	for _, l := range line {
		num, err := strconv.ParseFloat(l, 64)
		if err != nil {
			panic(err)
		}
		s = append(s, num)
	}

	return s
}

/*
dim calculate number of rows and columns of array.

param:
a [][]float64:

return:
m int: number of lines
n int: number of columns
*/
func dim(a [][]float64) (int, int) {
	var n, m int

	n = len(a)
	for _, i := range a {
		m = len(i)
		break
	}

	return n, m
}
