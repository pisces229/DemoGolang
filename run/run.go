package run

import (
	"context"
	"demo.golang/logic"
	"fmt"
)

func Run() {
	ctx := context.TODO()
	//if err := logic.NewLogic().Run(ctx); err != nil {
	//	fmt.Println(err)
	//}
	if err := logic.NewLogic().DefaultQueries(ctx); err != nil {
		fmt.Println(err)
	}
}
