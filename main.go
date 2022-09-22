package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Input 1 line.
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(scanner.Text())
}
