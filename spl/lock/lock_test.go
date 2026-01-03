package lock

import (
	"fmt"
	"sync"
	"testing"
)

func TestReentrantLock(t *testing.T) {
	l := &sync.Mutex{}
	foo(l)
}

func foo(l sync.Locker) {
	fmt.Println("in foo")
	l.Lock()
	bar(l)
	l.Unlock()
}

func bar(l sync.Locker) {
	l.Lock()
	fmt.Println("in bar")
	l.Unlock()
}
