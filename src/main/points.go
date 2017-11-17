package main

/** 指针 */
// * 取指针所指元素
// & 取变量的地址
import (
	logger "log"
)

/**
	
	
*/

//func main() {
//	testPoint();
//	
//	logger.Println("------------------------");
//	
//	var a string;
//	a = "hello";
//	var b = &a;
//	
//	println(a);
//	a = "world"
//	println(*b);
//}

func testPoint() {
	var a *int;	// 声明一个 int 类型指针
	var b = 9;	
	a = &b;		// 将 b的地址赋值给指针 a
	logger.Println("*a = ", *a)
	b = 100;
	logger.Println("*a = ", *a)
	logger.Println("&b = ", &b)	// &b 取 b所在的地址
	logger.Println("a = ", a)	// a 是一个地址
}