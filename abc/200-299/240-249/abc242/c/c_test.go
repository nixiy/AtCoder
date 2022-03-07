package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_add(t *testing.T) {
	t.Run("", func(t *testing.T) {
		a := 10
		add(&a, 5)
		assert.Equal(t, 15, a)
	})

	t.Run("", func(t *testing.T) {
		a := 10
		add(&a, -5)
		assert.Equal(t, 5, a)
	})
}

func Test_solve(t *testing.T) {
	t.Run("", func(t *testing.T) {
		assert.Equal(t, 25, solve(2))
		assert.Equal(t, 203, solve(4))
		assert.Equal(t, 117263, solve(10))
		assert.Equal(t, 248860093, solve(1000000))
	})
}
