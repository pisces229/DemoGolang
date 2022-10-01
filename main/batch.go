package main

import (
	"demo.golang/run"
)

func batch() {
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
