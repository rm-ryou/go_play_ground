package main

import "errors"

var ErrDataIsEmpty = errors.New("data is empty")

type Stack struct {
	size int
	data []int
}

func (s *Stack) Push(val int) {
	s.size++
	s.data = append(s.data, val)
}

func (s *Stack) Pop() (int, error) {
	if s.size == 0 {
		return 0, ErrDataIsEmpty
	}

	res := s.data[s.size-1]
	s.data = s.data[:s.size-1]

	return res, nil
}
