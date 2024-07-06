package arrays_test

import (
	"testing"

	"github.com/mikeunge/go-arrays"
)

func TestArrayInit(t *testing.T) {
	want := "Hello World"
	arr := arrays.New[string]()
	arr.Push("Hello World")

	if arr.First() != want {
		t.Errorf("Failed to initialize new array! Want: %s, got: %s\n", want, arr.First())
	}
}

func TestArrayLast(t *testing.T) {
	want := "World"
	arr := arrays.New[string]()
	arr.Push("Hello")
	arr.Push("New")
	arr.Push("World")

	if arr.Last() != want {
		t.Errorf("Failed to get last pushed to array! Want: %s, got: %s\n", want, arr.Last())
	}
}

func TestArrayGet(t *testing.T) {
	want := "Eve"
	arr := arrays.New[string]()
	arr.Push("Marc")
	arr.Push("Peter")
	arr.Push("Eve")
	arr.Push("Farid")

	if got, err := arr.Get(2); err != nil {
		t.Errorf("Failed to get from array! Want: %s, got: %s\n", want, got)
	}
}

func TestArrayLength(t *testing.T) {
	want := 5
	arr := arrays.New[int]()
	arr.Push(1)
	arr.Push(2)
	arr.Push(3)
	arr.Push(4)
	arr.Push(5)

	if arr.Length() != want {
		t.Errorf("Failed to get correct length! Want: %d, got: %d\n", want, arr.Length())
	}
}

func TestArrayInsert(t *testing.T) {
	var (
		want = 3
		got  int
		err  error
	)

	arr := arrays.New[int]()
	arr.Push(1)
	arr.Push(2)
	arr.InsertAtPos(1, want)

	if got, err = arr.Get(1); err != nil {
		t.Errorf("Failed to get from array! Error: %s\n", err.Error())
	} else if got != want {
		t.Errorf("Failed to get from array! Want: %d, got: %d\n", want, got)
	}
}

func TestArrayForEach(t *testing.T) {
	want := 3110
	arr := arrays.New[int]()
	arr.Push(1389)
	arr.Push(291)
	arr.Push(want)
	arr.Push(14123)
	arr.Push(241)

	arr.ForEach(func(index, value int) {
		if index == 2 && value != 3110 {
			t.Errorf("Array foreach failed! Want: %d, got: %d\n", want, value)
		}
	})
}

func TestArrayPop(t *testing.T) {
	want := 200
	arr := arrays.New[int]()
	arr.Push(20)
	arr.Push(100)
	arr.Push(want)

	got := arr.Pop()
	if got != want {
		t.Errorf("Failed to pop from array! Want: %d, got: %d\n", want, got)
	}
}
