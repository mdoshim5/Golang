package main

import (
	"fmt"
)

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element
	}
}

//converts an integer to a string
func int_to_string(number int) string {
	var string_representation string = fmt.Sprint(number)
	return string_representation
}

//converts a string to an integer
func string_to_int(string_representation string) int {
	var number int = 0
	for i := 0; i < len(string_representation); i++ {
		number *= 10
		var digit int = int(string_representation[i] - '0')
		number += digit
	}
	return number
}

//function that evaluates an equation with correct bracket sequence only for operators +, -, *, /
func evaluate_equation(equation string) int {
	var stack Stack // create a stack variable of type Stack

	for i := 0; i < len(equation); i++ {
		//evaluate operation for previous two operand and a operator
		if equation[i] == ')' {
			operand2 := stack.Pop()
			operator := stack.Pop()
			operand1 := stack.Pop()
			stack.Pop() //for removing the opening bracket
			var result int
			//doing the operation according to operator +, -, *, /
			if operator == "+" {
				result = string_to_int(operand1) + string_to_int(operand2)
			} else if operator == "-" {
				result = string_to_int(operand1) - string_to_int(operand2)
			} else if operator == "*" {
				result = string_to_int(operand1) * string_to_int(operand2)
			} else if operator == "/" {
				result = string_to_int(operand1) / string_to_int(operand2)
			}
			stack.Push(int_to_string(result)) //pushing the resultant to the stack
		} else if equation[i] >= '0' && equation[i] <= '9' { //if it is a starting of an integer number
			//find the entire int number
			var num string = string(equation[i])
			for equation[i+1] >= '0' && equation[i+1] <= '9' {
				num += string(equation[i+1])
				i++
			}
			stack.Push(num)
		} else { //pushing the ( and operators
			stack.Push(string(equation[i]))
		}
	}
	result := stack.Pop()        // extracting the result from the stack
	return string_to_int(result) // returning the result
}

func main() {

	var equation string = "((1+2)*(4/2))"        // the equation we want to solve
	var result int = evaluate_equation(equation) // calling the evaluate_equation function to evaluate the equation
	fmt.Println(result)                          // printing the result to the console
}
