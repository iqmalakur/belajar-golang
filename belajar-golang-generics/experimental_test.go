package main

import (
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

func ExperimentalMin[T constraints.Ordered](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestExperimentalMin(t *testing.T) {
	assert.Equal(t, 100, ExperimentalMin(100, 200))
	assert.Equal(t, 100.00, ExperimentalMin(100.00, 200.00))
}

func TestExperimentalMaps(t *testing.T) {
	first := map[string]string{
		"Name": "Ucup",
	}
	second := map[string]string{
		"Name": "Ucup",
	}

	assert.True(t, maps.Equal(first, second))
}

func TestExperimentalSlices(t *testing.T) {
	first := []string{"Ucup"}
	second := []string{"Ucup"}

	assert.True(t, slices.Equal(first, second))
}
