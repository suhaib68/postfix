// calculator project calculator.go
package postfix

import (
	"fmt"
	"strconv"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (value string, notEmpty bool) {
	if s.IsEmpty() {
		notEmpty = false
	} else {
		notEmpty = true
		index := len(*s) - 1
		value = (*s)[index]
		*s = (*s)[0:index]
	}
	return
}
func Postfix(Text string) (P []string) {
	var stack Stack
	var num string

	Text += " "
	for _, l := range Text {
		letter := string(l)
		if _, err := strconv.Atoi(letter); err != nil {
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
			for !stack.IsEmpty() && stack[len(stack)-1] != "(" {
				value, _ := stack.Pop()
				P = append(P, value)
			}
			stack.Pop()
		}
		if operator(letter) {
			if stack.IsEmpty() || (stack)[len(stack)-1] == "(" {
				stack.Push(letter)
			} else {
				for !stack.IsEmpty() && (stack)[len(stack)-1] != "(" && operatorPrecedence(letter) <= operatorPrecedence((stack)[len(stack)-1]) {
					value, _ := stack.Pop()
					P = append(P, value)
				}
				stack.Push(letter)
			}
		}
	}

	for !stack.IsEmpty() {
		value, _ := stack.Pop()
		P = append(P, value)
	}
	stack = Stack{}
	return
}

func CalcPostfix(P []string) (ans float64) {
	var stack Stack
	for _, value := range P {
		if _, err := strconv.Atoi(value); err == nil {
			stack.Push(value)
		}
		if operator(value) {
			Astr, _ := stack.Pop()
			Bstr, _ := stack.Pop()
			A, _ := strconv.ParseFloat(Astr, 64)
			B, _ := strconv.ParseFloat(Bstr, 64)
			res := op(value, A, B)
			stack.Push(fmt.Sprint(res))
		}
	}
	val, _ := stack.Pop()
	ans, _ = strconv.ParseFloat(val, 64)
	return
}

func operator(value string) (isoperator bool) {
	array := [5]string{"+", "-", "*", "/", "^"}
	for i := 0; i < len(array); i++ {
		if array[i] == value {
			isoperator = true
		}
	}
	return
}

func operatorPrecedence(value string) (pre int) {
	if value == "(" || value == ")" {
		pre = 3
	} else if value == "*" || value == "/" {
		pre = 2
	} else if value == "+" || value == "-" {
		pre = 1
	}
	return
}

func op(value string, A, B float64) (res float64) {
	if value == "+" {
		res = B + A
	} else if value == "-" {
		res = B - A
	} else if value == "*" {
		res = B * A
	} else if value == "/" {
		res = B / A
	}
	return
}
