package fx_test

import (
	"context"
	"testing"

	. "github.com/ryutah/learn-golang/di/generics-sample/di/fx"
	"github.com/ryutah/learn-golang/di/generics-sample/foobar"
)

func TestNew(t *testing.T) {
	var called bool
	app := New(func(*foobar.FooBar) {
		called = true
	})
	if err := app.Start(context.TODO()); err != nil {
		t.Errorf("failed to dependency injection: %v", err)
	}
	if !called {
		t.Error("invoke function is not called")
	}
}
