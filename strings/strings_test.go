package strings

import (
	//s "github.com/deckarep/golang-set"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountUniqueSubstrings(t *testing.T) {
	result := CountUniqueSubstrings("aabaa", 1, 1)
	fmt.Println(result)
	N := result.Cardinality()
	assert.Equal(t, 1, N, "N!=1")
	assert.True(t, result.Contains("a"))

	result = CountUniqueSubstrings("aabaa", 1, 4)
	N = result.Cardinality()
	assert.Equal(t, 8, N, "N!=8")
	assert.True(t, result.Contains("a", "b", "ab", "ba", "aa", "aba", "baa", "abaa"))
}
