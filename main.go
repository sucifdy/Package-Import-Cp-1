package main

import (
	"a21hc3NpZ25tZW50/internal" // path go mod
	"fmt"
	"strconv"
	"strings"
)

func AdvanceCalculator(calculate string) float32 {
	if calculate == "" {
		return 0.0
	}

	tokens := strings.Fields(calculate)
	if len(tokens) == 0 {
		return 0.0
	}

	firstNum, _ := strconv.ParseFloat(tokens[0], 32)
	calc := internal.NewCalculator(float32(firstNum))

	for i := 1; i < len(tokens); i += 2 {
		operator := tokens[i]
		nextNum, _ := strconv.ParseFloat(tokens[i+1], 32)

		switch operator {
		case "+":
			calc.Add(float32(nextNum))
		case "-":
			calc.Subtract(float32(nextNum))
		case "*":
			calc.Multiply(float32(nextNum))
		case "/":
			calc.Divide(float32(nextNum))
		}
	}

	return calc.Result()
}

func main() {
	res := AdvanceCalculator("3 * 4 / 2 + 10 - 5")
	fmt.Println(res) // Output: 11.0
}
