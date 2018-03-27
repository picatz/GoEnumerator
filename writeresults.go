package main

import (
	"fmt"
	"os"
)

func writeResults(content string, TargetToScan string) {

	filename := TargetToScan + "/discovered"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Printf("error opening %s: %s", filename, err)
	}

	defer file.Close()

	fmt.Fprintf(file, content+"\n")
}
