package fx

import (
	"github.com/ryutah/learn-golang/di/generics-sample/bar"
	"github.com/ryutah/learn-golang/di/generics-sample/foo"
	"github.com/ryutah/learn-golang/di/generics-sample/foobar"
	"go.uber.org/fx"
)

func New(invoke func(*foobar.FooBar)) *fx.App {
	return fx.New(
		fx.Provide(
			foo.NewFoo[string],
			bar.NewBar[string],
			foo.NewFoo[int],
			bar.NewBar2[int],
			foobar.NewFooBar,
		),
		fx.Invoke(invoke),
	)
}
