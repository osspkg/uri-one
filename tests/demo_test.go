package demo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSomething(t *testing.T) {
	var a string = "Hello"
	var b string = "Hello"
	require.Equal(t, a, b, "The two words should be the same.")
}
