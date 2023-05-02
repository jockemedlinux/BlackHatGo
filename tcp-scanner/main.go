package main
// go build -ldflags "-w -s"
import (
	"fmt"
	"net"
)


func main(){
	for i := 1; i <= 100; i++ {
		address := fmt.Sprintf("10.77.0.35:%d", i)
		// easily use parallel nc -lvp ::: 22 44 60 80 100
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// port is closed or filtered.
			continue
		}
		conn.Close()
		fmt.Printf("Port %d is open.\n", i)
	}
}