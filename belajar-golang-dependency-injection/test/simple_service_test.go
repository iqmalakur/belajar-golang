package test

import (
	"fmt"
	"iqmalakur/belajar-golang-dependency-injection/simple"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializedService()
	fmt.Println(err)
	fmt.Println(simpleService.SimpleRepository)
}
