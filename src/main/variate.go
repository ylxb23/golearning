package main

import (
	"fmt"
	"os"
)

/*
Go语言内置以下基础类型
布尔类型：bool
整型：int8, uint8(byte), byte, int16, uint16, int, uint, uintptr, int32, uint32, int64, uint64
浮点类型：float32, float64
复数类型：complex64, complex128
字符串：string
字符类型：rune
错误类型：error
指针(pointer)
数组(array)
切片(slice)
字典(map)
通道(chan)
结构体(struct)
接口(interface)
*/

const (
	cat = iota
	dog = iota
	sheep = iota
	horse = iota
	u int64 = 10086
	v float32 = 10000
	w float64 = 10001
	x uintptr = 99999
	cv1 complex64 = 2.3 + 2.5i;
	cv2 complex128 = 2.3 + 2.5i;
)

//func main() {
//	testVariate()
//}

func testVariate() {
	const a, b, c = 1, 3, "5";
	fmt.Printf("a: %v, b: %v, c: %v\n", a, b, c)
	
	var javaHome = os.Getenv("JAVA_HOME")
	fmt.Println(javaHome);
	var goHome = os.Getenv("GOROOT")
	fmt.Println(goHome);
	var path = os.Getenv("PATH")
	fmt.Println(path);
	
	fmt.Println(cat);
	fmt.Println(dog);
	fmt.Println(horse);
	fmt.Println(x);
	fmt.Println(x == uintptr(w));
	fmt.Println(cv1);
	fmt.Println(cv2);
	
	var str string = "hello,中国";
	fmt.Println("---------------------");
	displayStringByte(str);
	fmt.Println("---------------------");
	displayStringUnicode(str);
	fmt.Println("---------------------");
}

func displayStringByte(str string) {
	n := len(str);
	fmt.Println(str, ", size=", n);
	for i:=0; i<n; i++ {
		ch := str[i];
		fmt.Println(i, ch);
	}
}
func displayStringUnicode(str string) {
	for i, ch := range str {
		fmt.Println(i, ch);
	}
}