package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.
	var x uint64
	var y float64
	var z string

	// Read and save an integer, double, and String to your variables.
	fmt.Scan(&x)
	fmt.Scan(&y)
	scanner.Scan()
	z = scanner.Text()

	// Print the sum of both integer variables on a new line.
	fmt.Printf("%d", i+x)

	// Print the sum of the double variables on a new line.
	fmt.Printf("\n%.1f", d+y)

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Printf("\n%s%s", s, z)

}
