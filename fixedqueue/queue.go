package fixedqueue

import "sync"

// Queue of things to do
type Queue struct {
	channel chan interface{}
	size    int
	lock    sync.Mutex
}

// New Queue
func New(size int) *Queue {
	return &Queue{
		channel: make(chan interface{}, size),
		size:    size,
	}
}

//Push an element
func (q *Queue) Push(elem interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.channel) == q.size {
		// we are full, lets drop stuff
		<-q.channel
	}
	q.channel <- elem
}

//Pop and element
func (q *Queue) Pop() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.channel) == 0 {
		return nil
	}
	return <-q.channel
}

//BlPop blocking pop
func (q *Queue) BlPop() interface{} {
	return <-q.channel
}

//Length of the queue
func (q *Queue) Length() int {
	return len(q.channel)
}
