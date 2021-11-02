package rpm

import (
	"kript/stack"
	"strconv"
	"strings"
)

func Rpm(expr string) float32 {
	stack := stack.NewStack(len(expr))
	splited := strings.Split(expr, " ")
	for _, v := range splited {
		// fmt.Println(stack)
		if val, err := strconv.ParseFloat(v, 32); err == nil {
			stack.Push(float32(val))
		} else {
			rhs, err := stack.Pop()
			if err != nil {
				panic("")
			}
			lhs, err := stack.Pop()
			if err != nil {
				panic("")
			}
			switch v {
			case "+":
				stack.Push(lhs + rhs)
			case "-":
				stack.Push(lhs - rhs)
			case "*":
				stack.Push(lhs * rhs)
			case "/":
				stack.Push(lhs / rhs)
			}
		}
	}
	return stack.Top()
}
