package app

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func demoGoroutine() {
	//demoGoroutineMutex()
	demoGoroutineChan()
	//demoGoroutineWaitGroup()
	//demoGoroutineContext()
}

func demoGoroutineMutex() {
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

func demoGoroutineChan() {
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

func demoGoroutineWaitGroup() {
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

func demoGoroutineContext() {
	run := func(root context.Context, name string, millisecond int) {
		select {
		case <-root.Done():
			fmt.Printf("[%s][Done!]\n", name)
			return
		default:
			time.Sleep(time.Duration(millisecond) * time.Millisecond)
			fmt.Printf("[%s][Doing...]\n", name)
			return
		}

		//for {
		//	select {
		//	case <-root.Done():
		//		fmt.Printf("[%10s][Done!]\n", name)
		//		return
		//	default:
		//		fmt.Printf("[%10s][Doing...]\n", name)
		//		time.Sleep(500 * time.Millisecond)
		//	}
		//}
	}
	// root := context.Background()
	root, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	go run(root, "first", 1)
	go run(root, "second", 2)
	go run(root, "third", 3)
	time.Sleep(2500 * time.Millisecond)
}
