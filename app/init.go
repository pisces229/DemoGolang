package app

import (
	"fmt"
)

var initVar = func() string {
	fmt.Println("initVar")
	return ""
}()

func init() {
	fmt.Println("init main 0")
}
func init() {
	fmt.Println("init main 1")
}
