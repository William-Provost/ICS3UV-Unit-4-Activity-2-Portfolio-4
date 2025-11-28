// Author: William
// Version: 1.0.0
// Date: 2025-11-27
// Fileoverview: This program calculates the sum of even numbers up to a user-entered positive integer.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var n int
	var sum int = 0

	reader := bufio.NewReader(os.Stdin)

	// ask for positive integer
	for {
		fmt.Print("Enter a positive integer: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		n, _ = strconv.Atoi(input)

		if n > 0 {
			break
		} else {
			fmt.Println("Invalid input. Please enter a positive integer:")
		}
	}

	// calculate sum of even numbers
	for counter := 1; counter <= n; counter++ {
		if counter%2 == 0 {
			sum += counter
		}
	}

	fmt.Printf("Sum of even numbers from 1 to %d is: %d\n", n, sum)
	fmt.Println("Done.")
}
