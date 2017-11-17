package main

import (
	"fmt"
)

/**
	不定类型、不定数量参数
*/

//func main() {
//	parameterLength("中国", "长城", "井冈山", "还有那个啥啥啥");
//	
//	parameterAll("日本", 12, 99.0, true);
//	
//	getParameterType("字符串");
//	getParameterType(12);
//	getParameterType(uint(12));
//	getParameterType(false);
//	getParameterType(999999999999999999.0);
//}


/**
不定长度参数，同一类型
*/
func parameterLength(str ...string) {
	fmt.Println("不定长度参数:", str)
}

/**
不定长度参数，不定类型
*/
func parameterAll(args ...interface{}) {
	fmt.Println("不定长度参数，不定类型: ", args)
}

/**
获取参数类型
*/
func getParameterType(arg interface{}) {
	switch arg.(type) {
	case int:
	fmt.Println("arg[", arg, "] is int type.")
	break;
	case string:
	fmt.Println("arg[", arg, "] is string type.")
	break;
	case float32:
	fmt.Println("arg[", arg, "] is float32 type.")
	break;
	case float64:
	fmt.Println("arg[", arg, "] is float64 type.")
	break;
	case bool:
	fmt.Println("arg[", arg, "] is bool type.")
	break;
	case uint:
	fmt.Println("arg[", arg, "] is uint type.")
	break;
	case int32:
	fmt.Println("arg[", arg, "] is int32 type.")
	break;
	case int64:
	fmt.Println("arg[", arg, "] is int64 type.")
	break;
	default:
	fmt.Println("arg[", arg, "] is other type.")
	}
}