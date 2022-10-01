package app

import "fmt"

var initVarFirst = func() string {
	fmt.Println("initVarFirst")
	return ""
}()

func init() {
	fmt.Println("init first")
}
