package main

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

func main() {
	var (
		s string // string read from stdin
		n int    // number of cases to run
	)
	_, err := fmt.Scan(&s)
	if err != nil {
		log.Fatal(err)
	}
	n, err = strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	stackList := make([][]string, n)

	for i := 0; i < n; i++ {
		// Add a new slice at index i.
		stackList[i] = make([]string, 0)

		_, err := fmt.Scan(&s)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		// Build the stack slice using the stack string.
		stackList[i] = strings.Split(s, "")
	}

	// Solve each stack.
	for i, stack := range stackList {
		count, err := solve(stack)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("Case #%d: %d\n", i+1, count)
	}

}

func solve(stack []string) (int, error) {
	var (
		i      int    = 0
		flips  int    = 0
		target string = stack[0]
		err    error
	)

	if len(stack) == 1 {
		if isSolved(stack) {
			return 0, nil
		}
		return 1, nil
	}

	for ; i < len(stack); i++ {
		if isSolved(stack) {
			return flips, nil
		}

		if stack[i] == target {
			continue
		}

		stack, err = flip(stack, i-1)
		if err != nil {
			return -1, err
		}
		target = stack[i]
		flips++
	}

	if !isSolved(stack) {
		return flips + 1, nil
	}

	return flips, nil
}

func flip(stack []string, index int) ([]string, error) {
	if index < 0 || index > len(stack)-1 {
		return nil, fmt.Errorf("%d is an invalid index and must be between 0 and %d", index, len(stack)-1)
	}

	j := index
	for i := 0; i < j; i++ {
		stack[i], stack[j] = stack[j], stack[i]
		stack[i] = flipCake(stack[i])
		stack[j] = flipCake(stack[j])
		j--
	}

	if index%2 == 0 {
		stack[j] = flipCake(stack[j])
	}
	return stack, nil
}

func flipCake(side string) string {
	if side == "+" {
		return "-"
	}
	return "+"
}

func isSolved(stack []string) bool {
	for _, v := range stack {
		if v == "-" {
			return false
		}
	}
	return true
}
