package main 

import ( 
	"fmt"
	"time"
)

func main() { 

	start := time.Now()
	ch := make(chan string, 3)

	// branches off to create subroutine
	go func() { 
		for i := 1; i < 4; i++ {
			fileName := fmt.Sprintf("file_%d.txt", i)
			time.Sleep(1 * time.Second)

			// FIRST EXECUTION
			fmt.Println("Processing", fileName, "at", time.Since(start))
			
			ch <- fileName

			// SECOND EXECUTION
			fmt.Println(fileName + " processed")
			
		} // for 
		close(ch)
	}() // subroutine
	
	// stays on main routine, iteratively waits to receive the messages through the channel
	// once message is sent in subroutine through ch <- fileName
	// IMPORTANT: This hangs up and waits to recieve
	for file := range ch { 
		// THIRD EXECUTION 
		fmt.Println("Received", file, "at", time.Since(start))
	} // for 
}