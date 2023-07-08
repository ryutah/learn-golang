package foo

import "fmt"

type Foo[V any] interface {
	Print(V)
}

type FooImpl[V any] struct{}

func NewFoo[V any]() Foo[V] {
	return &FooImpl[V]{}
}

func (f *FooImpl[V]) Print(v V) {
	fmt.Printf("%v\n", v)
}
