package test

import (
	"iqmalakur/belajar-golang-dependency-injection/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("Databases")
	assert.NotNil(t, connection)

	cleanup()
}
