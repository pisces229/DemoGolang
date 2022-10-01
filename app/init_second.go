package app

import (
	"fmt"
)

var initVarSecond = func() string {
	fmt.Println("initVarSecond")
	return ""
}()

func init() {
	fmt.Println("init second")
}
