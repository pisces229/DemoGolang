package app

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func demoOnce() {
	var counter uint32
	var once sync.Once
	once.Do(func() {
		atomic.AddUint32(&counter, 1)
	})
	fmt.Printf("The counter: %d\n", counter)
	once.Do(func() {
		atomic.AddUint32(&counter, 2)
	})
	fmt.Printf("The counter: %d\n", counter)
	var once2 sync.Once
	once2.Do(func() {
		atomic.AddUint32(&counter, 5)
	})
	fmt.Printf("The counter: %d\n", counter)
}
