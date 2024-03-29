package main

import (
	"demo.golang/app"
	"fmt"
	"os"
)

func main() {
	fmt.Println("os.Args", os.Args)
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "app":
			app.Run()
		case "batch":
			batch()
		case "backend":
			backend()
		default:
			fmt.Println("unknown", os.Args[0])
		}
	} else {
		fmt.Println("unknown")
	}

}
