package main

import (
	"log"
	"net"
	"strconv"
	"time"
)

func portScan(TargetToScan string, PortStart int, PortEnd int, openPorts []int) []int {
	activeThreads := 0
	doneChannel := make(chan bool)

	for port := PortStart; port <= PortEnd; port++ {
		go grabBanner(TargetToScan, port, doneChannel)
		activeThreads++
	}

	// Wait for all threads to finish
	for activeThreads > 0 {
		<-doneChannel
		activeThreads--
	}
	return openPorts
}

func grabBanner(ip string, port int, doneChannel chan bool) {
	connection, err := net.DialTimeout(
		"tcp",
		ip+":"+strconv.Itoa(port),
		time.Second*10)
	if err != nil {
		doneChannel <- true
		return
	}
	// append open port to slice
	openPorts = append(openPorts, port)

	log.Printf("Port %d: Open\n", port)
	// See if server offers anything to read
	buffer := make([]byte, 4096)
	connection.SetReadDeadline(time.Now().Add(time.Second * 5))
	// Set timeout
	numBytesRead, err := connection.Read(buffer)
	if err != nil {
		doneChannel <- true
		return
	}
	log.Printf("Banner from port %d\n%s\n", port,
		buffer[0:numBytesRead])

	doneChannel <- true
	return
}
