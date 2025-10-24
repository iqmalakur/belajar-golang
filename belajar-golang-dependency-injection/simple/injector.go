//go:build wireinject
// +build wireinject

// use `wire gen [package name]` to generate dependency injection using wire
// example : wire gen iqmalakur/belajar-golang-dependency-injection/simple
// or just run `wire` in simple folder as active folder in terminal

package simple

import "github.com/google/wire"

func InitializedService() *SimpleService {
	wire.Build(
		NewSimpleRepository,
		NewSimpleService,
	)
	return nil
}
