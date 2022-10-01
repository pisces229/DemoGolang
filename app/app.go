package app

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var initVarMain = func() string {
	fmt.Println("initVarMain")
	return ""
}()

func init() {
	fmt.Println("init main 0")
}
func init() {
	fmt.Println("init main 1")
}

func App() {
	fmt.Println("main")
	// goroutine()
}

func onc() {
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
