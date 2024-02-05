package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	Elements []int
}

func (q *Queue) Enqueue(elem int) {
	q.Elements = append(q.Elements, elem)
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		fmt.Println("no more queue is already empty")
		return 0
	}
	lastElem := q.Elements[len(q.Elements)-1]
	q.Elements = q.Elements[:len(q.Elements)-1]
	return lastElem
}

func (q *Queue) IsEmpty() bool {
	return len(q.Elements) == 0
}

func (q Queue) GetLength() int {
	return len(q.Elements)
}

func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty queue")
	}
	return q.Elements[0], nil
}

func main() {
	myQueue := Queue{}

	// Enqueue elements
	myQueue.Enqueue(1)
	myQueue.Enqueue(2)
	myQueue.Enqueue(3)

	// Print length
	fmt.Println("Queue length:", myQueue.GetLength())

	// Peek at the front element
	frontElement, err := myQueue.Peek()
	if err == nil {
		fmt.Println("Front element:", frontElement)
	} else {
		fmt.Println("Error:", err)
	}

	// Dequeue elements
	fmt.Println("Dequeued:", myQueue.Dequeue())
	fmt.Println("Dequeued:", myQueue.Dequeue())

	// Print length after dequeueing
	fmt.Println("Queue length:", myQueue.GetLength())

	// Peek at the front element after dequeueing
	frontElement, err = myQueue.Peek()
	if err == nil {
		fmt.Println("Front element:", frontElement)
	} else {
		fmt.Println("Error:", err)
	}
	fmt.Println("Dequeued:", myQueue.Dequeue())
	// Dequeue from an empty queue
	fmt.Println("Dequeued:", myQueue.Dequeue()) // This will print an error message
}
