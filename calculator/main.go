package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// calculate performs one of the four basic arithmetic operations.
func calculate(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("cannot divide by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operator %q", operator)
	}
}

func readNumber(reader *bufio.Reader, prompt string) (float64, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		return 0, errors.New("invalid number")
	}
	return value, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Calculator")
	fmt.Println("Supported operations: +  -  *  /")
	fmt.Println("Press Ctrl+C to exit.")

	for {
		a, err := readNumber(reader, "First number: ")
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Print("Operator (+, -, *, /): ")
		operatorInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		operator := strings.TrimSpace(operatorInput)

		b, err := readNumber(reader, "Second number: ")
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		result, err := calculate(a, b, operator)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Printf("Result: %v\n\n", result)
	}
}
