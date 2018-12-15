package fixedqueue

import (
	"sync"
)

// Queue is a fixed size queue
type Queue struct {
	start   int
	end     int
	size    int
	queue   []interface{}
	lock    sync.RWMutex
	channel chan bool
}

// New Queue
func New(size int) *Queue {
	return &Queue{
		start:   0,
		end:     0,
		queue:   make([]interface{}, size),
		channel: make(chan bool, size),
	}
}

// Length of the queue
func (q *Queue) Length() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.size
}

// Push an element
func (q *Queue) Push(elem interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.end++
	if q.end == len(q.queue) {
		q.end = 0
	}
	if q.size < len(q.queue) {
		q.size++
		q.channel <- true
	}
	q.queue[q.end] = elem
}

func (q *Queue) pop() interface{} {
	elem := q.queue[q.start]
	q.start++
	if q.start == len(q.queue) {
		q.start = 0
	}
	q.size--
	return elem
}

// Pop an element
func (q *Queue) Pop() interface{} {
	q.lock.RLock()
	defer q.lock.RUnlock()
	if q.size == 0 {
		return nil
	}
	<-q.channel
	return q.pop()
}

// BlPop blocking pop an element
func (q *Queue) BlPop() interface{} {
	q.lock.RLock()
	defer q.lock.RUnlock()
	<-q.channel
	return q.pop()
}
