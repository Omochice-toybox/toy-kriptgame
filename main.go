package main

import (
	"bufio"
	"fmt"
	"kript/rpm"
	"os"
	"strings"
)

func Input(prompt string) (stringInput string) {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	rawInputed := scanner.Text()
	return strings.TrimSpace(rawInputed)
}

func main() {
	expr := Input("> ")
	rst := rpm.Rpm(expr)
	fmt.Printf("%s => %d(%f)", expr, int(rst), rst)
}
