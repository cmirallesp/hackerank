package main

import (
	"bufio"
	"fmt"
	p "hackerank/strings"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	if p.Pangram(text) {
		fmt.Println("pangram")
	} else {
		fmt.Println("not pangram")
	}
}
