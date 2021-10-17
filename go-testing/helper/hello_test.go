package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Run test with go test -v
// Run test with go test -v -run=FunctionName (for specific function)
// Run from parent directory go test -v ./... (run test from all package)

// TestMain()
// executed with go test -v
// Package level
func TestMain(m *testing.M) {
	// Before all tests
	fmt.Println("executed before")

	m.Run()

	// After all unit tests executed
	fmt.Println("executed after")
}

func TestGreet1(t *testing.T) {
	res := Greet("Falih")
	if res != "Hello Falih" {
		// error
		t.Fail()
	}

	// Code bellow still executed
	fmt.Println("TestGreet1 done")
}

func TestGreet2(t *testing.T) {
	res := Greet("Fauzan")
	if res != "Hello Fauzan" {
		// error
		t.FailNow()
	}

	// Code bellow is not executed
	fmt.Println("TestGreet2 done")
}

// Give some clue about the error with t.Error()
func TestGreet3(t *testing.T) {
	res := Greet("Falih")
	if res != "Hello Falih" {
		// error
		t.Error("Result must be Hello Falih")
	}

	// Code bellow still executed
	fmt.Println("TestGreet1 done")
}

func TestGreet4(t *testing.T) {
	res := Greet("Fauzan")
	if res != "Hello Fauzan" {
		// error
		t.Fatal("Result must be Hello Fauzan")
	}

	// Code bellow is not executed
	fmt.Println("TestGreet2 done")
}

// With testify/assert
func TestGreet5(t *testing.T) {
	res := Greet("Falih")

	// assert automatically call Fail(), require call FailNow()
	assert.Equal(t, "Hello Falih", res, "Must be Hello Falih")

	fmt.Println("TestGreet5 done")
}

// Skip test
func TestSkip(t *testing.T) {
	// Skip test for some condition
	if runtime.GOOS == "windows" {
		t.Skip("Skip test for windows OS")
	}

	res := Greet("Falih")

	// assert automatically call Fail(), require call FailNow()
	assert.Equal(t, "Hello Falih", res, "Must be Hello Falih")

	fmt.Println("TestGreet5 done")
}

// Sub test
// Sub test can be executed individually with -run=TestName/SubTestName
func TestSubTest(t *testing.T) {
	// Sub test 1
	t.Run("Greet 9", func(t *testing.T) {
		res := Greet("Falih")

		// assert automatically call Fail(), require call FailNow()
		assert.Equal(t, "Hello Falih", res, "Must be Hello Falih")

		fmt.Println("Greet 9 done")
	})

	// Sub test 2
	t.Run("Greet 10", func(t *testing.T) {
		res := Greet("Fauzan")

		// assert automatically call Fail(), require call FailNow()
		assert.Equal(t, "Hello Fauzan", res, "Must be Hello Fauzan")

		fmt.Println("Greet 10 done")
	})
}

// Table Test
func TestGreetTable(t *testing.T) {

	// Slice to iterate
	tests := []struct {
		name, request, expected string
	}{
		{
			name:     "Greet Falih",
			request:  "Falih",
			expected: "Hello Falih",
		},
		{
			name:     "Greet Fauzan",
			request:  "Fauzan",
			expected: "Hello Fauzan",
		},
		{
			name:     "Greet Falfal",
			request:  "Falfal",
			expected: "Hello Falfal",
		},
	}

	// Iterate
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := Greet(test.request)

			// assert automatically call Fail(), require call FailNow()
			assert.Equal(t, test.expected, res)

			fmt.Println("Subtest", test.name, "done")
		})
	}
}

// Mock object with testify/mock
