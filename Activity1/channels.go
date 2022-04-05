package main

import "fmt"

func producer(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
		fmt.Println(i)
	}

	close(channel)
}

func main() {
	ch := make(chan int, 10)
	go producer(ch) //secondary thread due to function call
	for {           //using as a while loop
		v, ok := <-ch
		if ok == false { //adding if statement with condition
			break //this breaks the while loop
		}
		fmt.Println("Received ", v, ok)
	}
}
