package main

import (
	"fmt"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Bar struct {
	value string
}

func (b Bar) Validate() error {
	return validation.ValidateStruct(
		&b,
		validation.Field(
			&b.value,
			validation.Required.Error("必須項目です"),
			validation.Length(1, 10).Error("{{.min}} - {{.max}}文字で入力してください"),
		),
	)
}

type Foo struct {
	bar   Bar
	Value string
}

func (f Foo) Validate() error {
	return validation.ValidateStruct(
		&f,
		validation.Field(&f.bar),
		validation.Field(&f.Value, is.Email),
	)
}

func main() {
	data := "example"
	err := validation.Validate(
		data,
		validation.Required,       // not empty
		validation.Length(5, 100), // length between 5 and 100
		is.URL,                    // is a valid URL
	)
	fmt.Println(err)
	if ve, ok := err.(validation.Error); ok {
		fmt.Println(ve.Code(), ve.Message(), ve.Params())
	}

	foo := Foo{
		bar: Bar{
			value: "example",
		},
		Value: "hoge@gmail.com",
	}
	err = validation.Validate(foo)
	fmt.Println(err)

	foo = Foo{
		bar: Bar{
			value: "",
		},
		Value: "hoge@gmail.com",
	}
	err = validation.Validate(foo)
	fmt.Println(err)

	foo = Foo{
		bar: Bar{
			value: strings.Repeat("a", 1000),
		},
		Value: "hoge@gmail.com",
	}
	err = validation.Validate(foo)
	fmt.Println(err)

	foo = Foo{
		bar: Bar{
			value: strings.Repeat("a", 1000),
		},
		Value: "hogegcom",
	}
	err = validation.Validate(foo)
	fmt.Println(err)

	if ve, ok := err.(validation.Errors); ok {
		for field, v := range ve {
			if e, ok := v.(validation.Error); ok {
				fmt.Println("Pamams:", e.Params(), field)
			}
			fmt.Println(field, v)
		}
	}
}
