package numbers

import (
	//"fmt"
	"testing"
)

func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(i)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }

func benchmarkFibMem(i int, b *testing.B) {
	//fmt.Printf("N=> %d\n", b.N)
	for n := 0; n < b.N; n++ {
		FibMem(i)
	}
}

func BenchmarkFibMem1(b *testing.B)  { benchmarkFibMem(1, b) }
func BenchmarkFibMem2(b *testing.B)  { benchmarkFibMem(2, b) }
func BenchmarkFibMem3(b *testing.B)  { benchmarkFibMem(3, b) }
func BenchmarkFibMem10(b *testing.B) { benchmarkFibMem(10, b) }
func BenchmarkFibMem20(b *testing.B) { benchmarkFibMem(20, b) }
func BenchmarkFibMem40(b *testing.B) { benchmarkFibMem(40, b) }
