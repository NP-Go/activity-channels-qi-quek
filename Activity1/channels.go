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
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}
}
