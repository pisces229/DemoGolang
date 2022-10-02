package main

import (
	"demo.golang/run"
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
	run.Run()
}
