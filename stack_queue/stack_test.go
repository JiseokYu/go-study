package main

import (
	"testing"
)

func TestStackPush(t *testing.T) {
	s := NewStack()
	const want = 1
	s.Push(want)
	if s.value[0] != want {
		t.Fatalf(`stack push result = %d, want %d`, s.value[0], want)
	}
}

func TestStackPop(t *testing.T) {
	t.Run("stack에 값이 없을 때", func(t *testing.T) {
		s := NewStack()
		result, err := s.Pop()

		if err == nil {
			t.Errorf("expected err but got '%d'", result)
		}
	})

	t.Run("정상적인 경우", func(t *testing.T) {
		s := NewStack()
		s.Push(10)
		want := 20
		s.Push(want)
		result, err := s.Pop()

		if result != want || err != nil {
			t.Errorf("expected '%d' but got value '%d', error '%d'", want, result, err)
		}
	})
}

func TestStackTop(t *testing.T) {
	t.Run("stack에 값이 없을 때", func(t *testing.T) {
		s := NewStack()
		result, err := s.Top()

		if err == nil || result != 0 {
			t.Errorf("expected err but got value: '%d', error: '%q'", result, err)
		}
	})

	t.Run("정상적인 경우", func(t *testing.T) {
		s := NewStack()
		s.Push(10)
		want := 20
		s.Push(want)
		result, err := s.Top()

		if result != want || err != nil {
			t.Errorf("expected '%d' but got value '%d', error '%d'", want, result, err)
		}
	})
}

func TestStackEmpty(t *testing.T) {
	t.Run("stack에 값이 없을 때", func(t *testing.T) {
		s := NewStack()
		result := s.Empty()

		if result != true {
			t.Errorf("expected false but got value: '%t'", result)
		}
	})

	t.Run("stack에 값이 있을 때", func(t *testing.T) {
		s := NewStack()
		s.Push(20)
		result := s.Empty()

		if result != false {
			t.Errorf("expected false but got value: '%t'", result)
		}
	})
}
