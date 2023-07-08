package main

import (
	"fmt"

	"github.com/ryutah/learn-golang/di/generics-sample/di/fx"
	"github.com/ryutah/learn-golang/di/generics-sample/foobar"
)

func main() {
	fx.New(func(fb *foobar.FooBar) {
		fmt.Println("Start !")
		defer fmt.Println("End !")

		fb.Run()
	}).Run()
}
