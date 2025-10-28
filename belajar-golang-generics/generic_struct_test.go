package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First  T
	Second T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "Ucup",
		Second: "Surucup",
	}

	fmt.Println(data)
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First:  "Ucup",
		Second: "Surucup",
	}

	assert.Equal(t, "Udin", data.ChangeFirst("Udin"))
	assert.Equal(t, "Hello Ucup", data.SayHello("Ucup"))
}
