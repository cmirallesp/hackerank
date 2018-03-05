package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReduced(t *testing.T) {
	assert.Equal(t, Reduced("aaabb"), "a", "Reduced('aaabb')!='a'")
	assert.Equal(t, Reduced("abaabb"), "ab", "Reduced('ababb')!='ab'")
	assert.Equal(t, Reduced("aabb"), "Empty String", "Reduced('aabb')!='Empty String'")
}

/*func BenchmarkReduced(t *testing.B) {*/
//for i := 0; i < t.N; i++ {
//Reduced("aaaaaaaaabaa")
//}
/*}*/
