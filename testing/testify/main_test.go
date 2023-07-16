package main_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Foo struct {
	Val string
}

func TestSample(t *testing.T) {
	f := &Foo{Val: "foo"}
	f2 := &Foo{Val: "foo2"}

	assert.Equal(t, f, f2, "foo")
}

func TestError(t *testing.T) {
	err1 := errors.New("error1")
	err2 := fmt.Errorf("error2: %w", err1)
	err3 := errors.New("error3")
	err4 := fmt.Errorf("error4: %w", err2)

	t.Log(errors.Is(err1, err2))
	assert.ErrorIs(t, err2, err1)
	assert.ErrorIs(t, err3, err1)
	assert.ErrorIs(t, err4, err1)
}
