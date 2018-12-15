package fixedqueue

import "sync"

type Queue struct {
	channel chan interface{}
	size    int
	lock    sync.Mutex
}

func New(size int) *Queue {
	return &Queue{
		channel: make(chan interface{}, size),
		size:    size,
	}
}

func (q *Queue) Push(elem interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.channel) == q.size {
		// we are full, lets drop stuff
		<-q.channel
	}
	q.channel <- elem
}

func (q *Queue) Pop() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.channel) == 0 {
		return nil
	}
	return <-q.channel
}

func (q *Queue) BlPop() interface{} {
	return <-q.channel
}

func (q *Queue) Length() int {
	return len(q.channel)
}
