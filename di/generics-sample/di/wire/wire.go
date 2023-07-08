package di

import (
	"github.com/google/wire"
	"github.com/ryutah/learn-golang/di/generics-sample/bar"
	"github.com/ryutah/learn-golang/di/generics-sample/foo"
	"github.com/ryutah/learn-golang/di/generics-sample/foobar"
)

var fooSet = wire.NewSet(
	foo.NewFoo[string],
	foo.NewFoo[int],
)

var barSet = wire.NewSet(
	bar.NewBar[string],
	bar.NewBar2[int],
)

func InjectFooBar() *foobar.FooBar {
	panic(wire.Build(
		foobar.NewFooBar,
		fooSet,
		barSet,
	))
}
