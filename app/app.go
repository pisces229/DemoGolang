package app

import (
	"fmt"
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

func Run() {
	fmt.Println("main")
	// goroutine()
}
