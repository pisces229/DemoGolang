package app

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

func dmeoGoroutine() {
	//goroutineMutex()
	//goroutineChan()
	//goroutineWaitGroup()
	//goroutineContext()
}

func goroutineMutex() {
	var count int32 = 0
	var mutex sync.Mutex
	run := func(name string) {
		for i := 0; i < 5000; i++ {
			mutex.Lock()
			count++
			//fmt.Printf("[%s][%d]\n", name, count)
			mutex.Unlock()
		}
	}
	go run("first")
	go run("second")
	go run("third")
	time.Sleep(5 * time.Second)
	fmt.Printf("[%d]\n", count)
}

func goroutineChan() {
	runChan := make(chan bool)
	doneChan := make(chan bool)
	cancelChan := make(chan bool)
	i := 0
	producer := func() {
		for {
			fmt.Println("producer...")
			select {
			case <-cancelChan:
				fmt.Println("cancel")
				return
			case <-time.After(1 * time.Second):
				i++
				fmt.Println("producer:", i)
				runChan <- true
				<-doneChan
			}
			fmt.Println("...producer")
		}
	}
	consumer := func() {
		for {
			fmt.Println("consumer...")
			select {
			case <-runChan:
				fmt.Println("consumer:", i)
				time.Sleep(1 * time.Second)
				doneChan <- true
			}
			fmt.Println("...consumer")
		}
	}
	go producer()
	go consumer()
	time.Sleep(10 * time.Second)
	cancelChan <- true
	time.Sleep(1 * time.Second)
}

func goroutineWaitGroup() {
	var count int32 = 0
	var waitGroup sync.WaitGroup
	run := func() {
		defer waitGroup.Done()
		for i := 0; i < 5000; i++ {
			count++
		}
	}

	waitGroup.Add(3)
	go run()
	go run()
	go run()
	waitGroup.Wait()

	//waitGroup.Add(1)
	//go run()
	//waitGroup.Wait()
	//waitGroup.Add(1)
	//go run()
	//waitGroup.Wait()
	//waitGroup.Add(1)
	//go run()
	//waitGroup.Wait()

	fmt.Printf("[%d]\n", count)
}

func goroutineContext() {
	var waitGroup sync.WaitGroup
	run := func(root context.Context, name string, timeout bool) {
		ctx, cancel := context.WithTimeout(root, 3*time.Second)
		defer waitGroup.Done()
		defer cancel()
		if timeout {
			time.Sleep(4 * time.Second)
		} else {
			time.Sleep(2 * time.Second)
		}
		select {
		case <-ctx.Done():
			fmt.Printf("[%s][%s]\n", name, strconv.FormatBool(false))
			return
		default:
			fmt.Printf("[%s][%s]\n", name, strconv.FormatBool(true))
			return
		}
	}
	// root := context.Background()
	root, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	waitGroup.Add(3)
	go run(root, "first", false)
	go run(root, "second", true)
	go run(root, "third", false)
	waitGroup.Wait()
}
