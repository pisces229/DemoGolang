package main

import (
	"context"
	"demo.golang/runner"
	"demo.golang/singleton"
	"fmt"
)

func batch() {
	fmt.Println(singleton.SingletonConfiguration.Mode)
	// flag.Parse()
	// fmt.Println("cpuprofile:", *cpuprofile)
	// if *cpuprofile != "" {
	// 	file, err := os.Create(*cpuprofile)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	pprof.StartCPUProfile(file)
	// 	defer pprof.StopCPUProfile()
	// }
	runner := runner.NewRunner()

	ctx := context.TODO()
	//ctx := context.Background()
	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(3 * time.Second))
	//defer cancel()

	err := runner.DefaultRunner(ctx)
	fmt.Println("batch:", err)

	//go runner.DefaultRunner(ctx)
	//time.Sleep(5 * time.Second)
}
