package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("10.77.0.35:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}		
		conn.Close()
		results <- p
	}
}

// added banner function for aestetics
func banner() {
	fmt.Println(`
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|J|M|L|-|P|O|R|T|S|C|A|N|N|E|R|
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		`)
}

func main(){
	banner() // added a nice looking banner
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := 0; i <= cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()
	
	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	
	close(ports)
	close(results)
	
	// added a check if it returns not open ports, it says so.
	if len(openports) == 0 {
		fmt.Println("[-] No open ports were found")
	}

	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("[+] Port %d is open.\n", port)
	}
}