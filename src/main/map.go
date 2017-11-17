package main

import (
	"fmt"
)

/**
map
Attention: 没有用到的变量和导入的包在编译期检查并会报错
*/
//func main() {
//	displayMap();
//	mapAppend();
//}

func mapAppend() {
	fmt.Println("mapAppend:------------------------")
	var dictionary map[string] string = map[string] string{"China":"中国", "Janpan":"日本"};
	fmt.Println(dictionary);
	dictionary["Asia"] = "亚洲";
	fmt.Println(dictionary);
	delete(dictionary, "Janpan");
	fmt.Println(dictionary);
}

func displayMap() {
	fmt.Println("displayMap:-----------------------")
	var persons map[string]PersonInfo;
	persons = make(map[string] PersonInfo);
	
	persons["10001"] = PersonInfo{"10001", "Little Ming", "SH", 23};
	persons["10002"] = PersonInfo{"10002", "Old Wang", "SH", 53};
	
	xiaoming, ok := persons["10001"];
	if ok {
		fmt.Println(ok, "xiaoming: ", xiaoming)
	} else {
		fmt.Println("Did not find person with ID 10001");
	}
	
	xiaoming.address = "上海";
	fmt.Println(ok, "xiaoming: ", xiaoming)
	
	// 声明map变量并指定初始空间大小
	var dictionary map[string] string = make(map[string] string, 5);
	dictionary["China"] = "中国";
	dictionary["Japan"] = "日本";
	dictionary["Asia"] = "亚洲";
	
	fmt.Println("dictionary[\"Asia\"]=", dictionary["Asia"])
	fmt.Println("dictionary[\"China\"]=", dictionary["China"])
	fmt.Println("dictionary[\"Japan\"]=", dictionary["Japan"])
}

type PersonInfo struct {
	ID string
	Name string
	address string
	age int8
}

