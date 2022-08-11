package main

import (
	"testing"
)

func TestQueuePush(t *testing.T) {
	s := NewQueue()
	const want = 1
	s.Enqueue(want)
	if s.value[0] != want {
		t.Fatalf(`queue push result = %d, want %d`, s.value[0], want)
	}
}

func TestQueuePop(t *testing.T) {
	t.Run("queue에 값이 없을 때", func(t *testing.T) {
		s := NewQueue()
		result, err := s.Dequeue()

		if err == nil {
			t.Errorf("expected err but got '%d'", result)
		}
	})

	t.Run("정상적인 경우", func(t *testing.T) {
		s := NewQueue()
		want := 10
		s.Enqueue(want)
		s.Enqueue(20)
		result, err := s.Dequeue()

		if result != want || err != nil {
			t.Errorf("expected '%d' but got value '%d', error '%d'", want, result, err)
		}
	})
}

func TestQueueTop(t *testing.T) {
	t.Run("queue에 값이 없을 때", func(t *testing.T) {
		s := NewQueue()
		result, err := s.Peek()

		if err == nil || result != 0 {
			t.Errorf("expected err but got value: '%d', error: '%q'", result, err)
		}
	})

	t.Run("정상적인 경우", func(t *testing.T) {
		s := NewQueue()
		want := 10
		s.Enqueue(want)
		s.Enqueue(20)
		result, err := s.Peek()

		if result != want || err != nil {
			t.Errorf("expected '%d' but got value '%d', error '%d'", want, result, err)
		}
	})
}

func TestQueueEmpty(t *testing.T) {
	t.Run("queue에 값이 없을 때", func(t *testing.T) {
		s := NewQueue()
		result := s.Empty()

		if result != true {
			t.Errorf("expected false but got value: '%t'", result)
		}
	})

	t.Run("queue에 값이 있을 때", func(t *testing.T) {
		s := NewQueue()
		s.Enqueue(10)
		result := s.Empty()

		if result != false {
			t.Errorf("expected false but got value: '%t'", result)
		}
	})
}
