package fixedqueue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	queue := New(3)
	assert.Equal(t, 0, queue.Length())
	for i := 0; i < 10; i++ {
		queue.Push(i)
	}
	assert.Equal(t, 3, queue.Length())
	a1 := queue.Pop()
	assert.NotNil(t, a1)
	assert.Equal(t, 2, queue.Length())
	a2 := queue.Pop()
	assert.NotNil(t, a2)
	assert.Equal(t, 1, queue.Length())
	a3 := queue.Pop()
	assert.NotNil(t, a3)
	assert.Equal(t, 0, queue.Length())
	//a4 := queue.Pop()
	//assert.Nil(t, a4)
}

func TestBlocking(t *testing.T) {
	stack := make([]interface{}, 0)
	queue := New(5)
	done := make(chan bool)
	n := 5
	go func() {
		for i := 0; i < n; i++ {
			e := queue.BlPop()
			stack = append(stack, e)
		}
		done <- true
	}()
	for i := 0; i < n; i++ {
		queue.Push(i)
	}
	<-done
	fmt.Println(stack)
	assert.Len(t, stack, n)
}
