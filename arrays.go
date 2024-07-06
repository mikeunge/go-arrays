package arrays

import (
	"fmt"
)

type Node[T any] struct {
	data T
	next *Node[T]
}

type Array[T any] struct {
	length int
	head   *Node[T]
}

func New[T any]() Array[T] {
	return Array[T]{
		length: 0,
		head:   nil,
	}
}

func (array *Array[T]) Length() int {
	return array.length
}

func (array *Array[T]) First() T {
	return array.head.data
}

func (array *Array[T]) Last() T {
	current := array.head
	for {
		if current.next == nil {
			return current.data
		}
		current = current.next
	}
}

func (array *Array[T]) Get(position int) (T, error) {
	if position > array.length {
		return Node[T]{}.data, fmt.Errorf("out of bounds")
	}

	current := array.head
	for i := 0; i <= position; i++ {
		if i == position {
			return current.data, nil
		}
		current = current.next
	}
	return Node[T]{}.data, fmt.Errorf("nothing found")
}

func (array *Array[T]) InsertAtPos(position int, data T) {
	newNode := Node[T]{data: data}
	if position <= 0 {
		newNode.next = array.head
		array.head = &newNode
		return
	}

	current := array.head
	for i := 0; i < position-1 && current != nil; i++ {
		current = current.next
	}

	if current == nil {
		array.Push(data)
	} else {
		newNode.next = current.next
		current.next = &newNode
	}
}

func (array *Array[T]) ForEach(fn func(index int, value T)) {
	current := array.head
	index := 0
	for {
		fn(index, current.data)
		if current.next == nil {
			break
		}
		current = current.next
		index++
	}
}

func (array *Array[T]) Pop() T {
	if array.length <= 0 {
		return Node[T]{}.data
	}

	current := array.head
	for {
		nextNode := current.next
		if nextNode.next == nil {
			data := nextNode.data
			current.next = nil
			return data
		}
		current = current.next
	}
}

func (array *Array[T]) Push(data T) {
	newNode := Node[T]{data: data}

	array.length++
	current := array.head
	if current == nil {
		array.head = &newNode
		return
	}

	for {
		if current.next == nil {
			current.next = &newNode
			break
		}
		current = current.next
	}
}

func (array *Array[T]) PrintNodes() {
	for current := array.head; current.next != nil; current = current.next {
		fmt.Println(current.data)
	}
}
