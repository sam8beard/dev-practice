package main 

import ( 
	"bufio"
	"fmt"
	"os"
	// "path/filepath"
)

func main() { 
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter path to file: ")
	scanner.Scan()
	file_path := scanner.Text() 

	data, err := os.Open(file_path)
	
	if err != nil { 
		fmt.Print(err)
	} else { 
		fmt.Println(file_path)
	}

	file_scanner := bufio.NewScanner(data)
	for i:= 0; i < 4 && file_scanner.Scan(); i++ { 
		// skip column row
		if i == 0 { 
			continue
		} else { 
			fmt.Println(file_scanner.Text())
		}
	}
	
}