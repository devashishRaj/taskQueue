package main

import (
	"fmt"
	"sync"
	"time"
)

// a unit of work
type Task struct {
	ID   int
	Data string
}

// the task queue
type TaskQueue struct {
	// to enqueue task
	queue chan Task
	// to receivve finished task
	response chan Task
	wg       sync.WaitGroup
}

// creates a new TaskQueue with initialized channels
func NewTaskQueue() *TaskQueue {
	return &TaskQueue{
		queue:    make(chan Task),
		response: make(chan Task),
	}
}

// adds a task to the queue
func (tq *TaskQueue) AddTask(task Task) {
	tq.queue <- task
}

// processes tasks in order , It runs in a goroutine .
func (tq *TaskQueue) ProcessTasks() {
	// once tasks gets done
	defer tq.wg.Done()

	for task := range tq.queue {
		// Simulate processing time , wait for a second
		time.Sleep(time.Second)
		// sending the processed task
		tq.response <- task
	}
}

// start a pool of workers to process tasks
func (tq *TaskQueue) StartWorkerPool(numWorkers int) {
	for i := 1; i <= numWorkers; i++ {
		tq.wg.Add(1)
		go tq.ProcessTasks()
	}
}

// wait for all tasks to be processed
func (tq *TaskQueue) WaitForCompletion() {
	tq.wg.Wait()
	// close the response channel to signal that no more tasks will be processed.
	close(tq.response)
}

// Client function to simulate sending tasks to the TaskQueue
func client(taskQueue *TaskQueue, numWorkers int) {
	for i := 1; i <= numWorkers+1; i++ {
		task := Task{ID: i, Data: fmt.Sprintf("Task %d", i)}
		taskQueue.AddTask(task)
		time.Sleep(time.Millisecond * 500) // Simulate task creation delay
	}
	close(taskQueue.queue) // Close the queue to signal no more tasks will be added
}

// Server function to process completed tasks
func server(taskQueue *TaskQueue) {
	//Continues processing tasks until the response channel is closed.
	for task := range taskQueue.response {
		fmt.Printf("Processed Task %d: %s\n", task.ID, task.Data)
	}
}

func main() {
	taskQueue := NewTaskQueue()
	numOfWorkers := 10

	// Start a pool of worker goroutines
	taskQueue.StartWorkerPool(numOfWorkers)

	// Client goroutine
	go client(taskQueue, numOfWorkers)

	// Server goroutine
	go server(taskQueue)

	// Wait for all tasks to be processed
	taskQueue.WaitForCompletion()
}
