package main 

import (
	"fmt"
	"sync"
	"time" 
)

func processFile(wg *sync.WaitGroup, file string) { 
	
	fmt.Println("Starting conversion:", file)
	time.Sleep(1 * time.Second)
	fmt.Println("Finished conversion:", file)
	
	wg.Done()
	
} // processFile

func main() { 
	var wg sync.WaitGroup 

	fileNames := []string{
		"file1.csv",
		"file2.json", 
		"file3.csv",
	}
	wg.Add(len(fileNames))

	for _, file := range fileNames { 
		go processFile(&wg, file)		
	} 
	wg.Wait()

	fmt.Println("All conversions complete!")
} // main  

/*

Output should look like: 

Starting conversion: file1.csv
Starting conversion: file2.json
Starting conversion: file3.csv
Finished conversion: file1.csv
Finished conversion: file2.json
Finished conversion: file3.csv
All conversions complete!

*/