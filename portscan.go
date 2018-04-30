package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

// startBanner spins up a handful of async workers
func startBannerGrabbers(num int, target string, portsIn <-chan int) <-chan int {
	portsOut := make(chan int)

	var wg sync.WaitGroup

	wg.Add(num)

	for i := 0; i < num; i++ {
		go func() {
			for p := range portsIn {
				if grabBanner(target, p) {
					portsOut <- p
				}
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(portsOut)
	}()

	return portsOut

}

func portScan(targetToScan string, portStart int, portEnd int) []int {
	ports := make(chan int)

	go func() {
		for port := portStart; port <= portEnd; port++ {
			ports <- port
		}
		close(ports)
	}()

	resultChan := startBannerGrabbers(16, targetToScan, ports)

	//var openPorts []int
	for port := range resultChan {
		openPorts = append(openPorts, port)
	}

	return openPorts
}

//var targetPorts = make(map[int]string)

func grabBanner(ip string, port int) bool {
	connection, err := net.DialTimeout(
		"tcp",
		ip+":"+strconv.Itoa(port),
		time.Second*5)

	if err != nil {
		fmt.Printf(".")
		return false
	}
	defer connection.Close() // you should close this!

	fmt.Printf("\n+ Port %d: Open\n", port)
	buffer := make([]byte, 4096)
	connection.SetReadDeadline(time.Now().Add(time.Second * 5))
	numBytesRead, err := connection.Read(buffer)

	if err != nil {
		return true
	}

	fmt.Printf("\n+ Banner of port %d\n+ %s\n", port, buffer[0:numBytesRead])
	// here we add to map port and banner
	// ******* MAPS ARE NOT SAFE FOR CONCURRENT WRITERS ******
	// ******************* CHANGE THIS *******************
	targetPorts[port] = string(buffer[0:numBytesRead])

	return true
}
