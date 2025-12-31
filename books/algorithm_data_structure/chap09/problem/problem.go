package problem

import (
	"errors"
	"fmt"
)

var ErrDataIsEmpty = errors.New("data is empty")

type Stack struct {
	tail int
	data []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
	s.tail++
}

func (s *Stack) Pop() (int, error) {
	if s.tail == 0 {
		return 0, ErrDataIsEmpty
	}

	s.tail--
	res := s.data[s.tail]
	s.data = s.data[:s.tail]
	return res, nil
}

func isOperator(v rune) bool {
	return v == '+' || v == '-' || v == '*' || v == '/'
}

func processOperate(x, y int, v rune) int {
	var res int

	switch v {
	case '+':
		res = x + y
	case '-':
		res = x - y
	case '*':
		res = x * y
	case '/':
		res = x / y
	}
	return res
}

func RPN(expre string) (int, error) {
	s := NewStack()

	for _, v := range expre {
		if v >= '0' && v <= '9' {
			s.Push(int(v - '0'))
		}

		if isOperator(v) {
			r, err := s.Pop()
			if err != nil {
				return 0, err
			}
			l, err := s.Pop()
			if err != nil {
				return 0, err
			}
			s.Push(processOperate(l, r, v))
		}
	}

	res, err := s.Pop()
	if err != nil {
		return 0, err
	}
	return res, nil
}

func CheckBrackets(expre string) (bool, error) {
	s := NewStack()

	for i, v := range expre {
		switch v {
		case '(':
			s.Push(i)
		case ')':
			idx, err := s.Pop()
			if err != nil {
				return false, err
			}

			fmt.Printf("%d and %d match\n", idx, i)
		default:
			continue
		}
	}

	var res bool
	if s.tail == 0 {
		res = true
	} else {
		res = false
	}
	return res, nil
}
