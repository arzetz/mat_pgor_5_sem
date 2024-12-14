package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type buffer struct {
	data     []float64
	capacity int
	start    int
	size     int
}

func (buffer *buffer) Add(value float64) {
	if buffer.size < buffer.capacity {
		buffer.data[(buffer.start+buffer.size)%buffer.capacity] = value
		buffer.size++
	} else {
		buffer.data[buffer.start] = value
		buffer.start = (buffer.start + 1) % buffer.capacity
	}
}

var weight = map[string]int{
	"+":   1,
	"-":   1,
	"*":   2,
	"/":   2,
	"^":   3,
	"sin": 4,
	"cos": 4,
	"tan": 4,
	"cot": 4,
	"ln":  4,
	"log": 4,
}

func isOperator(element string) bool {
	_, exists := weight[element]
	return exists
}

func postfix(infix string, buffer *buffer) string {
	elements := strings.Fields(infix)
	var postfix []string
	var stack []string

	for _, curr_element := range elements {
		if isOperator(curr_element) {
			for len(stack) > 0 && weight[stack[len(stack)-1]] >= weight[curr_element] {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, curr_element)
		} else {
			switch curr_element {
			case "M1":
				postfix = append(postfix, strconv.FormatFloat(buffer.data[0], 'f', 2, 64))
			case "M2":
				postfix = append(postfix, strconv.FormatFloat(buffer.data[1], 'f', 2, 64))
			case "M3":
				postfix = append(postfix, strconv.FormatFloat(buffer.data[2], 'f', 2, 64))
			default:
				postfix = append(postfix, curr_element)
			}

		}
	}

	for len(stack) > 0 {
		postfix = append(postfix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return strings.Join(postfix, " ")
}

func calc(postfix string, buffer *buffer) float64 {
	elements := strings.Fields(postfix)
	var stack []float64
	for _, curr_element := range elements {
		if isOperator(curr_element) {
			if curr_element == "sin" || curr_element == "cos" || curr_element == "tan" || curr_element == "cot" || curr_element == "ln" || curr_element == "log" {
				b := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				var result float64
				switch curr_element {
				case "sin":
					result = math.Sin(b)
				case "cos":
					result = math.Cos(b)
				case "tan":
					result = math.Tan(b)
				case "cot":
					result = 1 / math.Tan(b)
				case "ln":
					result = math.Log(b)
				case "log":
					result = math.Log10(b)
				}
				stack = append(stack, result)
			} else {
				a := stack[len(stack)-2]
				b := stack[len(stack)-1]
				stack = stack[:len(stack)-2]

				var result float64
				switch curr_element {
				case "+":
					result = a + b
				case "-":
					result = a - b
				case "*":
					result = a * b
				case "/":
					result = a / b
				case "^":
					result = math.Pow(a, b)
				}
				stack = append(stack, result)
			}
		} else {
			curr_float, _ := strconv.ParseFloat(curr_element, 64)
			stack = append(stack, curr_float)
		}
	}
	buffer.Add(stack[0])
	fmt.Printf("M1: %f\n", buffer.data[0])
	fmt.Printf("M2: %f\n", buffer.data[1])
	fmt.Printf("M3: %f\n", buffer.data[2])
	return stack[0]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buffer := &buffer{
		data:     make([]float64, 3),
		capacity: 3,
	}
	for {
		scanner.Scan()
		expression := scanner.Text()
		fmt.Printf("Результат: %f\n", calc(postfix(expression, buffer), buffer))
		if expression == "!" {
			break
		}
	}
}
