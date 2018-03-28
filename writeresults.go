package main

import (
	"fmt"
	"os"
)

func writeResults(TargetToScan string, content []string, filename string) {

	filename = TargetToScan + "/" + filename
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error opening %s because of: %s", filename, err)
	}

	defer file.Close()
	for _, line := range content {
		file.WriteString(line + "\n")

	}
}
