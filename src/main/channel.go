package main

import (
	"fmt"
)

/** "通信的方式共享内存" */

func main() {
	testCountChannel();
}

func testCountChannel() {
	chs := make([]chan int, 10);
	for i:=10; i<10; i++ {
		chs[i] = make(chan int);
		go CountChannel(chs[i]);
	}
	
	for _, ch := range(chs) {
		<- ch
	}
}

func CountChannel(ch chan int) {
	ch <- 1;
	fmt.Println("Counting");
}
