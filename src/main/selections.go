package main

import (
	"fmt"
)

/*
流程控制语句:
	选择：
	循环：
	跳转：
	条件语句：对应的关键字为：if/else/else if
	选择语句：对应的关键字为：switch/case/select(channel)
	循环语句：对应的关键字为：for/range
	跳转语句：对应的关键字为：goto/break/continue/fullthrough
*/

//func main() {
//	result := ifSelection(3)
//	fmt.Println(result)
//
//	fmt.Println(switchStatment(8))
//	fmt.Println(switchProStatment(8))
//	
//	fmt.Println("\n----------------");
//	for1Statment();
//	fmt.Println("\n----------------");
//	for2Statment();
//	fmt.Println("\n----------------");
//	for3Statment();
//	fmt.Println("\n----------------");
//	jumpStatment();
//	fmt.Println("\n----------------");
//	gotoStatment();
//}

/**
if selection statment
*/
func ifSelection(x int) bool {
	if x == 3 {
		return true
	} else {
		return false
	}
}

func switchStatment(x int) int {
	var res int = 0
	switch x {
	case 0:
		res = 0
		break
	case 1, 2, 3, 4, 5:
		res = 1
		break
	case 6, 7, 8, 9:
		res = 9
		break
	default:
		res = -1
	}
	return res
}

/**
switch 后面的表达式可以不需要
*/
func switchProStatment(x int) int {
	switch {
	case x < 0:
		return -1;
		
	case x == 0:
		return 0;
		
	case x > 0:
		return 1;
	default:
	return 999;
	}
}

/**
for statment
*/
func for1Statment() {
	i := 9;
	for {
		if(i < 0) {
			break;
		}
		fmt.Print(i, "|");
		i=i-1
	}
}

/**
for statment
*/
func for2Statment() {
	for i:=9; i>0; i-- {
		fmt.Print(i, "|");
	}
}

/**
for statment
*/
func for3Statment() {
	array:= []int{1,2,3,4,5,6,7,8,9};
	// i==index, j==value
	for i,j := range array {
		fmt.Print(i, ",", j, "|");
	}
}

/**
jump statment
*/
func jumpStatment() {
	JumpHere:
	for i:=0; i<99; i++ {
		for j:=0; j<5; j++ {
			if(i > 2) {
				break JumpHere;
			}
			fmt.Print(i,",",j,"|")
		}
		fmt.Print("\n");
	}
	fmt.Println("\n Jumped out of here");
}

/**
	goto statment
*/
func gotoStatment() {
	for i:=0; i<99; i++ {
		if i>5 {
			goto BigThan5;
		}
	}
	// 这一行将被跳过，不执行
	fmt.Println("middle of goto statment function.");
	BigThan5: 
	fmt.Println("This is label BigThan5 here.");
}