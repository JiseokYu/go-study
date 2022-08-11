package main

import (
	"errors"
)

type queue struct {
	value []int
}

func NewQueue() *queue {
	q := queue{}
	return &q
}

func (q *queue) Enqueue(n int) {
	(*q).value = append((*q).value, n)
}

func (q *queue) Dequeue() (int, error) {
	if q.Empty() {
		return 0, errors.New("queue is empty")
	}
	var r = (*q).value[0]
	(*q).value = (*q).value[1:]
	return r, nil
}

func (q *queue) Peek() (int, error) {
	if q.Empty() {
		return 0, errors.New("queue is empty")
	}
	return (*q).value[0], nil
}

func (q *queue) Empty() bool {
	return len((*q).value) == 0
}
