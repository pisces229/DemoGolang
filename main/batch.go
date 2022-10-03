package main

import (
	"fmt"

	"demo.golang/runner"
	"demo.golang/singleton"
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
	err := runner.DefaultRunner()
	fmt.Println("batch:", err)
}
