package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_solve(t *testing.T) {
	assert.True(t, solve("abc", "ijk"))
	assert.True(t, solve("z", "a"))
	assert.False(t, solve("ppq", "qqp"))
	assert.True(t, solve("atcoder", "atcoder"))
}

func Test_diff(t *testing.T) {
	assert.Equal(t, 0, diff('a', 'a'))
	assert.Equal(t, 1, diff('a', 'b'))
	assert.Equal(t, 2, diff('a', 'c'))
	assert.Equal(t, 1, diff('b', 'c'))
	assert.Equal(t, 1, diff('z', 'a'))
	assert.Equal(t, 25, diff('a', 'z'))
}

func Test_byteShift(t *testing.T) {
	assert.Equal(t, "a", byteShift('a', 0))
	assert.Equal(t, "b", byteShift('a', 1))
	assert.Equal(t, "a", byteShift('z', 1))
	assert.Equal(t, "z", byteShift('a', -1))
	assert.Equal(t, "a", byteShift('a', 26))
}

func Test_strShift1(t *testing.T) {
	assert.Equal(t, "bcdBCD", strShift("abcABC", 1))
	assert.Equal(t, "abcABC", strShift("abcABC", 0))
	assert.Equal(t, "zabZAB", strShift("abcABC", -1))
	assert.Equal(t, "WKLV LV D VHFUHW PHVVDJH",
		strShift("THIS IS A SECRET MESSAGE", 3))
}
