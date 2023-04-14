package main

import "github.com/google/uuid"

func main() {
	for i := 0; i < 5; i++ {
		println(uuid.New().String())
	}
}
