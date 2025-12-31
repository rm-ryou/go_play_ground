package main

import (
	"errors"
)

var (
	ErrDataIsEmpty = errors.New("data is empty")
	ErrDataIsFull  = errors.New("data is full")
)

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

type Queue struct {
	Max        int
	head, tail int
	data       []int
}

func NewQueue(max int) *Queue {
	data := make([]int, max)
	return &Queue{
		Max:  max,
		data: data,
	}
}

func (q *Queue) Enqueue(val int) error {
	if q.head == (q.tail+1)%q.Max {
		return ErrDataIsFull
	}

	q.data[q.tail] = val
	if q.tail+1 == q.Max {
		q.tail = 0
	} else {
		q.tail++
	}

	return nil
}

func (q *Queue) Dequeue() (int, error) {
	if q.head == q.tail {
		return 0, ErrDataIsEmpty
	}

	res := q.data[q.head]
	if q.head+1 == q.Max {
		q.head = 0
	} else {
		q.head++
	}

	return res, nil
}
