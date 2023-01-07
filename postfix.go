package postfix

import (
	"fmt"
	"math"
	"strconv"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (value string) {
	if !s.IsEmpty() {
		index := len(*s) - 1
		value = (*s)[index]
		*s = (*s)[0:index]
	}
	return
}

func (s *Stack) LastItem() (value string) {
	if !s.IsEmpty() {
		value = (*s)[len(*s)-1]
	}
	return
}

func Calc(Text string) (ans float64) {
	var stack Stack
	var num string
	var P []string

	Text += " "
	for _, l := range Text {
		letter := string(l)
		if _, err := strconv.ParseFloat(letter, 64); err != nil && letter != "." {
			if num != "" {
				P = append(P, num)
			}
			num = ""
		} else {
			num += letter
		}
		if letter == "(" {
			stack.Push(letter)
		}
		if letter == ")" {
			for !stack.IsEmpty() && stack.LastItem() != "(" {
				P = append(P, stack.Pop())
			}
			stack.Pop()
		}
		if operator(letter) {
			if stack.IsEmpty() || stack.LastItem() == "(" {
				stack.Push(letter)
			} else {
				for !stack.IsEmpty() && stack.LastItem() != "(" && operatorPrecedence(letter) <= operatorPrecedence(stack.LastItem()) {
					P = append(P, stack.Pop())
				}
				stack.Push(letter)
			}
		}
	}

	for !stack.IsEmpty() {
		P = append(P, stack.Pop())
	}
	stack = Stack{}

	for _, value := range P {
		if _, err := strconv.ParseFloat(value, 64); err == nil {
			stack.Push(value)
		}
		if operator(value) {
			A, _ := strconv.ParseFloat(stack.Pop(), 64)
			B, _ := strconv.ParseFloat(stack.Pop(), 64)
			res := op(value, A, B)
			stack.Push(fmt.Sprint(res))
		}
	}
	ans, _ = strconv.ParseFloat(stack.Pop(), 64)
	return
}

func operator(value string) (isoperator bool) {
	switch value {
	case "+", "-", "*", "/", "^":
		isoperator = true
	}
	return
}

func operatorPrecedence(value string) (pre int) {
	switch value {
	case "(", ")":
		pre = 4
	case "^":
		pre = 3
	case "*", "/":
		pre = 2
	case "+", "-":
		pre = 1
	}
	return
}

func op(value string, A, B float64) (res float64) {
	switch value {
	case "+":
		res = B + A
	case "-":
		res = B - A
	case "*":
		res = B * A
	case "/":
		res = B / A
	case "^":
		res = math.Pow(B, A)
	}
	return
}
