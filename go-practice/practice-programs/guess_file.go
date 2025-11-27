package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter file path: ")

	scanner.Scan()
	file_name := scanner.Text()

	fmt.Println("This is your file path: ", file_name)

	file_suffix := filepath.Ext(file_name)

	if file_suffix == ".csv" || file_suffix == ".json" { 
		fmt.Print("Your file path is valid")
	} else { 
		fmt.Print("Your file path is invalid")
	}
}
