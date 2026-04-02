package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {

		if !scanner.Scan() {
			break
		}

		texts := scanner.Text()
		texts = strings.TrimSpace(texts)

		if texts == "Quit" {
			break
		}
		if texts == "" {
			fmt.Println("Error: Emp")
			continue
		}

		splitText := strings.Fields(texts)

		if len(splitText) != 3 {
			fmt.Println("Error: Invalid input!")
			continue
		}

		if splitText[0] != "convert" {
			fmt.Println("Error: The words must start with convert!")
		}

		secondPart := splitText[1]
		thirdpart := splitText[2]

		switch thirdpart {

		case "hex":
			n, err := strconv.ParseInt(secondPart, 16, 64)
			if err != nil {
				fmt.Println("Error:", secondPart+" is not a valid hex! ")
				continue
			}
			fmt.Println("Decimal:", n)
		case "bin":
			n, err := strconv.ParseInt(secondPart, 2, 64)
			if err != nil {
				fmt.Println("Error:", secondPart+" is not a valid binary! ")
				continue
			}
			fmt.Println("Decimal:", n)

		case "dec":
			n, err := strconv.ParseInt(secondPart, 10, 64)
			if err != nil {
				fmt.Println("Error:", secondPart+" is not a valid decimal! ")
				continue
			}
			fmt.Printf("Binary %b\n", n)
			fmt.Printf("Hex %X\n", n)
		}

	}
}
