package main

import (
	"fmt"
	"net"
	"sort"
	"flag" //added flag for user input
	"os" // added for user input
)

func worker(ports, results chan int, host string) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", host, p)
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
	banner() // added a "nice" looking banner, lol

	// added argument-flags to specify host and number of ports.
	hostPtr := flag.String("u", "", "The host to scan")
	portsPtr := flag.Int("p", 65535, "The number of ports to scan. Not ranges, not seperated.")
	showHelpPtr := flag.Bool("h", false, "Display help")
	flag.Parse()

	if flag.Lookup("u").Value.String() == "" || *portsPtr < 0 {
		flag.Usage()
	    fmt.Println("\n[-] You need to provide appropriate arguments")
	    os.Exit(0)
	}
	
	if *showHelpPtr {
		flag.Usage()
		os.Exit(0)
	}

	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := 0; i <= cap(ports); i++ {
		go worker(ports, results, *hostPtr)
	}

	go func() {
		for i := 1; i <= *portsPtr; i++ {
			ports <- i
		}
	}()
	
	for i := 0; i < *portsPtr; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	
	close(ports)
	close(results)
	
	// added a check if it returns not open ports, it says so.
	if len(openports) == 0 {
		fmt.Println("[-] No ports found open")
		os.Exit(0)
	}

	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("[+] Port %d is open.\n", port)
	}
	// added a sum of total ports.
	fmt.Printf("\n[+] A total of %d found ports open", len(openports))
}