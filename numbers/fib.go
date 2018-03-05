package numbers

//import "fmt"

func Fib(n int) int {
	if n <= 1 {
		return 0
	} else {
		return Fib(n-1) + Fib(n-2)
	}
}

func fibMem(n int, mem []int) int {
	if n <= 1 {
		return 1
	} else {
		if mem[n-1] == 0 {
			mem[n-1] = fibMem(n-1, mem)
		}
		if mem[n-2] == 0 {
			mem[n-2] = fibMem(n-2, mem)
		}
		return mem[n-1] + mem[n-2]
	}
}

func FibMem(n int) int {
	mem := make([]int, n)
	//fmt.Printf("1] %v\n", mem)
	result := fibMem(n, mem)
	//fmt.Printf("2] %v\n", mem)
	return result
}
