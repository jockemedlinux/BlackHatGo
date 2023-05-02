package main

import (
	"fmt"
)

func main(){
	for i := 1; i <= 100; i++ {
		address := fmt.Sprintf("10.77.0.35:%d", i)
		fmt.Println(address)
	}
}