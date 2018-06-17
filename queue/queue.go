package queue

import (
	"errors"
)

type Queue struct {
	//Queue represnets the data structure queue. Its implementation follows
	// the algorithm in CLRS, which utilizes a circular array
	a        []int
	tail     int
	head     int
	capacity int
	//Length is the number of elments of in queue
	Length int
}

const (
	autoResize = 0 << iota
	manuResize
)

var (
	QueueCapNonPositveErr = errors.New("cannot initialize a queue with non-positive number of elements")
	QueueOverFlowErr      = errors.New("queue overflows: queue is full. Invoke q.resize() to increase its capacity")
	QueueUnderFlowErr     = errors.New("queue underflow: no more elements to be dequeued")
)

//NewQueue returns a new queue with the size of the underlying array being n
func NewQueue(n int) *Queue {
	if n <= 0 {
		panic(QueueCapNonPositveErr)
	}
	return &Queue{a: make([]int, n), head: 0, capacity: n}
}

//EnQueue adds one element x at the end of the queue. When the queue the full,
//the function will return a QueueOverFlowErr
func (q *Queue) EnQueue(x int) error {
	if q.Length >= q.capacity {
		return QueueOverFlowErr
	}
	q.a[q.tail] = x
	q.Length++
	if q.tail == q.capacity-1 {
		q.tail = 0
	} else {
		q.tail++
	}
	if q.head == q.tail {
		return QueueOverFlowErr
	}

	return nil
}

//DeQueue removes one element x from the begining of the queue. When the queue the empty,
//the function will return a QueueUnderFlowErr
func (q *Queue) DeQueue() (int, error) {
	if q.Length <= 0 {
		return 0, QueueUnderFlowErr
	}

	x := q.a[q.head]
	q.Length--
	if q.head == q.capacity-1 {
		q.head = 0
	} else {
		q.head++
	}
	return x, nil
}

//Resize returns a new queue with the size of underlying array doubled
func (q *Queue) Resize() *Queue {
	cap := 2 * q.capacity
	arr := make([]int, cap)
	head := 0
	tail := q.Length
	length := q.Length

	for i := 0; i < q.Length; i++ {
		if i+q.head > q.capacity-1 {
			arr[i] = q.a[i+q.head-q.capacity]
		} else {
			arr[i] = q.a[i+q.head]
		}
	}
	return &Queue{a: arr, head: head, tail: tail, Length: length, capacity: cap}
}
