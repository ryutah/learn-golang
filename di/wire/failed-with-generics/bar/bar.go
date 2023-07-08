package bar

import (
	"github.com/ryutah/learn-golang/di/wire/failed-with-generics/foo"
)

type Bar[V any] struct {
	Foo foo.Foo[V]
}

func NewBar[V any](foo foo.Foo[V]) Bar[V] {
	return Bar[V]{Foo: foo}
}

func (b *Bar[V]) Exec(arg V) {
	b.Foo.Print(arg)
}

type Bar2[V any] struct {
	Foo foo.Foo[V]
}

func NewBar2[V any](foo foo.Foo[V]) Bar[V] {
	return Bar[V]{Foo: foo}
}

func (b *Bar2[V]) Exec(arg V) {
	b.Foo.Print(arg)
}
