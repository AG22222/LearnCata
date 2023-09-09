package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Operate func(int, int) int

type Expression struct {
	X, Y      int
	Operation Operate
}

// Map of single digits
var singledigits = map[string]int{
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "X": 10, "I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9,
}

// Map of operators "+" "-" and funcs
var operators = map[string]Operate{
	"+": func(x, y int) int { return x + y },
	"-": func(x, y int) int { return x - y },
	"*": func(x, y int) int { return x * y },
	"/": func(x, y int) int { return x / y },
}

// Filling Expression structure
func (exp Expression) FillingExpression(stringarr []string) Expression {
	for _, elem := range stringarr {
		_, ok := singledigits[elem]

		if ok {
			exp.X = singledigits[stringarr[0]]
			exp.Y = singledigits[stringarr[2]]
		} else {
			exp.Operation = operators[elem]
		}

	}
	return Expression{exp.X, exp.Y, exp.Operation}
}

// Preparing input condition with trim spaces
func PreparingInputCondition(condition string) []string {
	stringArr := []string{}
	conditionArr := strings.Split(condition, " ")

	for _, str := range conditionArr {
		if str != " " {
			stringArr = append(stringArr, str)
		}
	}
	return stringArr
}

func main() {
	r := bufio.NewReader(os.Stdin)
	inputCondition, _ := r.ReadString('\n')
	inputCondition = strings.TrimSpace(inputCondition)
	preparedCondition := PreparingInputCondition(inputCondition)
	fmt.Println("Prepared condition: ", PreparingInputCondition(inputCondition))

	expression := Expression{}

	completeExpression := expression.FillingExpression(preparedCondition)

	result := Expression{
		X:         completeExpression.X,
		Y:         completeExpression.Y,
		Operation: completeExpression.Operation,
	}

	fmt.Println("Result of operation: ", result.Operation(result.X, result.Y))

}
