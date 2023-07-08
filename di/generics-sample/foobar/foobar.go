package foobar

import "github.com/ryutah/learn-golang/di/generics-sample/bar"

type FooBar struct {
	Bar  bar.Bar[string]
	Bar2 bar.Bar2[int]
}

func NewFooBar(bar bar.Bar[string], bar2 bar.Bar2[int]) *FooBar {
	return &FooBar{
		Bar:  bar,
		Bar2: bar2,
	}
}

func (f *FooBar) Run() {
	f.Bar.Exec("foo")
	f.Bar2.Exec(1)
}
