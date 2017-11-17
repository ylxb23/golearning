package main

import (
	"fmt"
)

//func main() {
//	testAdds();
//}

func testAdds() {
	for i:=0; i<10; i++ {
		go Adds(i, i);
	}
	fmt.Println("Exit");
}

func Adds(a, b int) {
	z := a+b;
	fmt.Println(z);
}