package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	t.Parallel()

	patterns := map[string]struct {
		input    int
		expected int
	}{
		"success": {
			input:    1,
			expected: 2,
		},
		"fail": {
			input:    1,
			expected: 1,
		},
	}
	for msg, p := range patterns {
		t.Run(msg, func(t *testing.T) {
			actual := foo(p.input)
			assert.Equal(t, p.expected, actual)
		})
	}
}
