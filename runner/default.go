package runner

import "context"

func (i *Runner) DefaultRunner(ctx context.Context) error {
	err := i.Logic.Run(ctx)
	//err := i.Logic.DefaultQueries(ctx)
	return err
}
