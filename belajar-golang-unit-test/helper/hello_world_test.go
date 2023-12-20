package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Iqmal",
			request: "Iqmal",
		},
		{
			name:    "Akur",
			request: "Akur",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Iqmal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Iqmal")
		}
	})
	b.Run("Akur", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Akur")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Iqmal")
	}
}

func BenchmarkHelloWorldAkur(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Akur")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Iqmal",
			request:  "Iqmal",
			expected: "Hello Iqmal",
		},
		{
			name:     "Akur",
			request:  "Akur",
			expected: "Hello Akur",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Iqmal", func(t *testing.T) {
		result := HelloWorld("Iqmal")
		require.Equal(t, "Hello Iqmal", result, "Result must be 'Hello Iqmal'")
	})
	t.Run("Akur", func(t *testing.T) {
		result := HelloWorld("Akur")
		require.Equal(t, "Hello Akur", result, "Result must be 'Hello Akur'")
	})
}

func TestMain(m *testing.M) {
	// Before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// After
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Can not run on Linux")
	}

	result := HelloWorld("Akur")
	require.Equal(t, "Hello Akur", result, "Result must be 'Hello Akur'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Akur")

	require.Equal(t, "Hello Akur", result, "Result must be 'Hello Akur'")
	fmt.Println("TestHelloWorldRequire Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Akur")

	assert.Equal(t, "Hello Akur", result, "Result must be 'Hello Akur'")
	fmt.Println("TestHelloWorldAssert Done")
}

func TestHelloWorldIqmal(t *testing.T) {
	result := HelloWorld("Iqmal")

	if result != "Hello Iqmal" {
		// Error
		// t.Fail()
		t.Error("Result must be 'Hello Iqmal'")
	}

	fmt.Println("TestHelloWorldAkur Done")
}

func TestHelloWorldAkur(t *testing.T) {
	result := HelloWorld("Akur")

	if result != "Hello Akur" {
		// Error
		// t.FailNow()
		t.Fatal("Result must be 'Hello Akur'")
	}

	fmt.Println("TestHelloWorldAkur Done")
}
