package application

import (
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typictx"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typiobj"
	"github.com/typical-go/typical-rest-server/pkg/utility/runkit"
	"github.com/urfave/cli"
	"go.uber.org/dig"
)

type application struct {
	*typictx.Context
}

func (a application) Run(ctx *cli.Context) (err error) {
	di := dig.New()
	defer a.Close(di)
	runkit.GracefulShutdown(func() error {
		return a.Close(di)
	})
	if err = typiobj.Provide(di, a); err != nil {
		return
	}
	if err = typiobj.Prepare(di, a); err != nil {
		return
	}
	runner := a.Application.(typiobj.Runner)
	return di.Invoke(runner.Run())
}

func (a application) Prepare() (preparations []interface{}) {
	if preparer, ok := a.Application.(typiobj.Preparer); ok {
		preparations = append(preparations, preparer.Prepare()...)
	}
	return
}

func (a application) Provide() (constructors []interface{}) {
	constructors = append(constructors, a.Constructors...)
	constructors = append(constructors, a.Modules.Provide()...)
	if provider, ok := a.Application.(typiobj.Provider); ok {
		constructors = append(constructors, provider.Provide()...)
	}
	return
}

func (a application) Destroy() (destructors []interface{}) {
	destructors = append(destructors, a.Modules.Destroy()...)
	if destroyer, ok := a.Application.(typiobj.Destroyer); ok {
		destructors = append(destructors, destroyer.Destroy()...)
	}
	return
}

func (a application) Close(c *dig.Container) (err error) {
	return typiobj.Destroy(c, a)
}
