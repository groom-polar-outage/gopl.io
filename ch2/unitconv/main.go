package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Temperature, length and weight converter.")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString("\n")

}
