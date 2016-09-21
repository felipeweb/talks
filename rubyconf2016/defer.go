package main

import (
	"fmt"
	"time"
)

func main() {
	defer world()
	fmt.Println("Hello")
}

func world() {
	time.Sleep(5 * time.Second)
	fmt.Println("World")
}