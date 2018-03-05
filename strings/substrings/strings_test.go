package strings

import (
	//s "github.com/deckarep/golang-set"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountUniqueSubstrings(t *testing.T) {
	N := CountUniqueSubstrings("aabaa", 1, 1)
	assert.Equal(t, 1, N, "N!=1")

	N = CountUniqueSubstrings("aabaa", 1, 4)
	assert.Equal(t, 8, N, "N!=8")
}

func TestArrCountUniqueSubstrings(t *testing.T) {
	result := ArrCountUniqueSubstrings("aabaa", 1, 1)
	assert.Equal(t, 1, result, "N!=1")

	result = ArrCountUniqueSubstrings("aabaa", 1, 4)
	assert.Equal(t, 8, result, "N!=8")

	result = ArrCountUniqueSubstrings("aabaa", 0, 0)
	assert.Equal(t, 1, result, "N!=1")
}

func TestForCountUniqueSubstrings(t *testing.T) {
	result := ForCountUniqueSubstrings("aabaa", 1, 1)
	assert.Equal(t, 1, result, "N!=1")

	result = ForCountUniqueSubstrings("aabaa", 1, 4)
	assert.Equal(t, 8, result, "N!=8")

	result = ForCountUniqueSubstrings("aabaa", 0, 0)
	assert.Equal(t, 1, result, "N!=1")
}
