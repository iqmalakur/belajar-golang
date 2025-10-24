package test

import (
	"fmt"
	"iqmalakur/belajar-golang-dependency-injection/simple"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService := simple.InitializedService()
	fmt.Println(simpleService.SimpleRepository)
}
