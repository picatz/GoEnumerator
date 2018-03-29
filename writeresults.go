package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func writeResultsString(TargetToScan string, content []string, filename string) {

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

func writeResultsInt(TargetToScan string, content []int, filename string) {

	filename = TargetToScan + "/" + filename
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error opening %s because of: %s", filename, err)
	}

	defer file.Close()
	for _, line := range content {
		file.WriteString(strconv.Itoa(line) + "\n")
	}
}

func writeResultsMap(TargetToScan string, content map[int]string, filename string) {

	filename = TargetToScan + "/" + filename
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error opening %s because of: %s", filename, err)
	}

	defer file.Close()
	for port, banner := range content {
		line := strconv.Itoa(port) + " " + strings.TrimSpace(banner) + "\n"

		file.WriteString(line)
	}
}
