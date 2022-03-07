package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_solve(t *testing.T) {
	assert.Equal(t, "abgfedch", solve(3, 7, "abcdefgh"))
	assert.Equal(t, "reviver", solve(1, 7, "reviver"))
	assert.Equal(t, "meramtsirhcyrs", solve(4, 13, "merrychristmas"))
}
