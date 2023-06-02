package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ShiraazMoollatjie/goluhn"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter string for LUHN Validation: ")
	name, _ := reader.ReadString('\n')

	name = strings.TrimSpace(name)

	err := goluhn.Validate(name)

	if err != nil {
		fmt.Printf("❌ FAILED VALIDATION")
	} else {
		fmt.Printf("✅ PASSED VALIDATION")
	}
}
