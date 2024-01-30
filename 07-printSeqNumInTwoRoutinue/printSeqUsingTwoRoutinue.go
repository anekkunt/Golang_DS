package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var eveCh = make(chan bool)
var oddCh = make(chan bool)

func main() {
	wg.Add(2)
	go printEvenNum()
	go printOddNum()
	eveCh <- true
	wg.Done()
}

func printEvenNum() {
	defer wg.Done()

	for i := 0; i <= 10; i = i + 2 {
		<-eveCh
		fmt.Println(i)
		oddCh <- true
	}

}

func printOddNum() {
	defer wg.Done()

	for i := 1; i <= 10; i = i + 2 {
		<-oddCh
		fmt.Println(i)
		eveCh <- true
	}

}
