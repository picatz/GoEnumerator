package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

func portScan(targetToScan string, PortStart int, PortEnd int) {
	resultChan := make(chan int)

	var wg sync.WaitGroup
	for port := PortStart; port <= PortEnd; port++ {
		wg.Add(1)

		go func(p int) {
			open := grabBanner(targetToScan, p)
			if open != false {
				resultChan <- p
			}
			wg.Done()
		}(port)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for port := range resultChan {
		openPorts = append(openPorts, port)
	}

}

func grabBanner(ip string, port int) bool {
	// Your testing code here. Return an error, or not.

	var open bool
	connection, err := net.DialTimeout(
		"tcp",
		ip+":"+strconv.Itoa(port),
		time.Second*20)

	if err != nil {
		open = false
		return open
	}

	fmt.Printf("+ Port %d: Open\n", port)
	// See if server offers anything to read
	buffer := make([]byte, 4096)
	connection.SetReadDeadline(time.Now().Add(time.Second * 5))
	numBytesRead, err := connection.Read(buffer)

	if err != nil {
		//		fmt.Println("No banner")
		open = true
		return open
	}

	log.Printf("+ Banner of port %d\n%s\n", port,
		buffer[0:numBytesRead])
	// here we add to map port and banner
	targetPorts[port] = string(buffer[0:numBytesRead])
	open = true
	return open
}
