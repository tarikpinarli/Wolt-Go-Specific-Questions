package main

import (
	"fmt"
	"time"
	"sync"
)

type Order struct {
	id int
	item string
}

func (o Order) Process() {
	fmt.Printf("Order %d: Making %s...\n", o.id, o.item)
	time.Sleep(2 * time.Second)
	fmt.Printf("Order %s is done.\n", o.item)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	latte := Order{1, "latte"}
	americano := Order{2, "americano"}

	startTime := time.Now().Second()
	go func() {
		defer wg.Done()
		latte.Process()
	}()
	go func() {
		defer wg.Done()
		americano.Process()
	}()
	fmt.Printf("Orders are being prepared\n")
	wg.Wait()
	endTime := time.Now().Second()
	TotalTime := endTime - startTime
	fmt.Println("All ready")
	fmt.Printf("time: %d\n", TotalTime)
}