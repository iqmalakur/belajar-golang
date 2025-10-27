package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestTypeParameter(t *testing.T) {
	// var result string = Length[string]("Ucup")
	// var result = Length[string]("Ucup")
	var result = Length("Ucup")
	assert.Equal(t, "Ucup", result)

	// var resultNumber int = Length[int](100)
	var resultNumber = Length(100)
	assert.Equal(t, 100, resultNumber)
}
