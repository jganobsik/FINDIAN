package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("hello, please enter a word")
	if scanner.Scan() {
		input := scanner.Text()
		input1 := strings.ToLower(input)
		i := strings.Index(input1, "i")
		n := strings.LastIndex(input1, "n")
		a := strings.Index(input1, "a")
		length := len(input1)
		zerolen := length - 1

		if i == 0 && a != -1 && n == zerolen {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}

	}

}
