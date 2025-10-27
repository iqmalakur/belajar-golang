package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSame[T comparable](value1, value2 T) bool {
	// if value1 == value2 {
	// 	return true
	// } else {
	// 	return false
	// }

	return value1 == value2
}

func TestIsSame(t *testing.T) {
	assert.True(t, IsSame[string]("Ucup", "Ucup"))
	assert.True(t, IsSame(100, 100))
	assert.True(t, IsSame(true, true))
}
