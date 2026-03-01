package main

import (
	"fmt"
	"net"
	"sort"
)

// going to perform port scanning at the specified port & send the results
// if port is open,then port number will be sent over the channel
// otherwise a zero will be sent for closed port across the channel
func scan(ports, results chan int) {
	for p := range ports {
		//website used to learn about port-scanning
		//scanme.nmap.org -> domain
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		fmt.Printf("Address is: %s \n", address)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p

	}

}

func main() {
	ports := make(chan int, 500) //500 values it can hold at first
	results := make(chan int)
	var openPorts []int
	for i := 0; i < cap(ports); i++ { // cap -> built-in function that is used to determine the capacity
		//of a slice channel or array
		go scan(ports, results)
	}
	//sending the port number
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i

		}
	}()
	//collecting the scan results (is implicitly waiting for all workers to finish)
	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)

		}
	}
	close(ports)
	close(results)
	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("%d open\n", port)
	}

}

//how this function is working?
// 500 scan goroutines start and block at -> for p:=range ports,they are waiting for values
//then func() goroutine runs and it start sending ports
//as soon as the port is sent,one waiting goroutine picks it and starts scanning while others keep waiting

//why are we not using waitGroups? -> because we are using channels : are powerful synchronization primitive in Go that can ensure
// that our program waits for certain actions to complete before proceeding. Therefore,when we send a value on a channel,the sending goroutine blocks
//until another goroutine receives that value from the channel
