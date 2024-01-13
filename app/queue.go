package app

import "fmt"

func QueueMain() {
	queue := NewQueue()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println(queue.Dequeue()) // 1
	fmt.Println(queue.Dequeue()) // 2
	//queue.Enqueue(4)
	fmt.Println(queue.Dequeue()) // 3
	//fmt.Println(queue.Dequeue()) // 4

	fmt.Println(queue.IsEmpty()) // true
	fmt.Println(queue.Size())    // 0
}

type Queue struct {
	data  []interface{}
	front int
	rear  int
}

func NewQueue() *Queue {
	return &Queue{
		data:  make([]interface{}, 0),
		front: 0,
		rear:  0,
	}
}

func (q *Queue) Enqueue(data interface{}) {
	q.data = append(q.data, data)
	q.rear = q.rear + 1
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	data := q.data[q.front]
	//q.data = q.data[1:]
	q.front = q.front + 1
	return data
}

func (q *Queue) IsEmpty() bool {
	return q.front == q.rear
}

func (q *Queue) Size() int {
	return q.rear - q.front
}
