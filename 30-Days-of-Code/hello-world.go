package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World.")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	name := scanner.Text()
	fmt.Println(name)
}
