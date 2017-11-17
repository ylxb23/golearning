package main

import (
	"fmt"
	"errors"
)

//func main() {
//	var a, b int32 = 2,3;
//	sum, err := Add(a, b);
//	fmt.Println(sum, err)
//	
//	sum1, err1 := AddAll(1,2,3,4,"hehehe",6,7);
//	fmt.Println(sum1, err1)
//}

func Add(a, b int32) (sum int32, err error) {
	if a < 0 || b < 0 {
		err = errors.New("参数不能为负值.")
		return
	}
	return a+b, nil
}

func AddAll(args ...interface{}) (sum int, err error) {
	fmt.Println("-------------------------------------");
	sum = 0
	for _, arg := range args {
		switch arg.(type) {
			case int:
				sum += 1
			default :
				err = errors.New("参数类型错误,只能为 int 类型")
				
		}
	}
	return
}