package main

import (
	"fmt"
)

/*
数组
cap() 函数返回的是数组切片分配的空间大小
len() 函数返回的是数组切片中当前所存储的元素个数
*/

//func main() {
//	displayArraysPro();
//	testCopy();
//}

func testCopy() {
	arr := []int{1,2,3,4,5};
	arrC := make([]int, 2, 5);
	// copy(target, source)
	copy(arrC, arr);
	fmt.Println("copy:------------------")
	fmt.Println(arr);
	fmt.Println(arrC);
}

func testSlice() {
	var arr []int = []int {1,2,3,4,5,6,7,8,9};
	sliced := arr[:3];
	slicedDes := arr[3:];
	fmt.Println(arr);
	fmt.Println(sliced);
	fmt.Println(slicedDes);
}

func displayArraysPro() {
	array := make([]int, 5, 20);
	temp := []int{6,7,8,9};
	fmt.Println(array);
	// 给数组后面追加元素，元素只能为单个元素， temp... 省略号的作用是将temp数组打散后传入append函数
	array = append(array, temp...);
	fmt.Println(array);
	
}

func displayArrays() {
	// 创建一个长度为5的固定数组空间
	var array0 [5]int;
	// 创建一个长度为5的固定数组，并赋初始值为 1,2,3,4,5
	array1 := []int{1,2,3,4,5};
	// 创建一个初始元素为5的数组切片，所有元素初始值为0
	array2 := make([]int, 5);
	// 创建一个初始元素个数为6的数组切片，所有元素初始值为0，共有10个元素的存储空间
	// 若需要使用剩下的4个存储空间，需要通过 append()函数从尾部向 array3中追加元素
	array3 := make([]int, 6, 10);
	
	fmt.Println("array0:", array0)
	fmt.Println("array1:", array1)
	fmt.Println("array2:", array2)
	fmt.Println("array3:", array3)
	lena := len(array3);
	capa := cap(array3);
	fmt.Println("len(array3):", lena, ", cap(array3): ",capa)
	
	for i:=0; i<capa; i++ {
		if(i >= lena) {
			array3 = append(array3, i);
			continue;
		}
		array3[i] = i
	}
	
	fmt.Println("array3:", array3)
	array4 := array3;
	fmt.Println("-----------------");
	fmt.Println("array3:", array3)
	fmt.Println("array4:", array4)
	
	fmt.Println("------------------------------------------------");
}