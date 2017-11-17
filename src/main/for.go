package main

import (
	"fmt"
)

//func main() {
//	var result bool = testRoop();
//	fmt.Println(result)
//}

func testRoop()(ok bool) {
	sum := 0;
	
	JLPrintSum: {
		fmt.Println("label PrintSum: ", sum);
		return false;
	}
	i := 0;
	for {
		for {
			if sum > 10 {
				goto JLPrintSum
			}
			sum++
			fmt.Println(sum);
		}
		i++;
		if i>3 {
			break;
		}
		fmt.Println("ok");
	}
	fmt.Println(sum);
	return true;
}