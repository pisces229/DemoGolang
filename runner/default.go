package runner

import (
	"demo.golang/logic"
)

func (i *Runner) DefaultRunner() error {
	err := logic.NewLogic().Run(i.Context)
	//err := logic.NewLogic().DefaultQueries(i.Context)
	return err
}
