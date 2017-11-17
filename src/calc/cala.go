package main

import 	"os"
import 	"fmt"
import 	"simplemath"
import 	"strconv"

var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ...");
	fmt.Println("\nThe commands are:\n\tadd\tAddition of two values.\n\tsqrt\tSquare");
}

var Logger = func(words string) {
	fmt.Println(words);
}

func main() {
	args := os.Args	// 获取参数
	
	if args == nil || len(args) < 3 {
		Usage();
		return;
	}
	for i:=0; i<len(args); i++ {
		fmt.Println("args[", i, "]: ", args[i])
	}
	
	switch args[1] {
	case "add":
		if len(args) != 4 {
			fmt.Println("USAGE: calc add <Integer1> <Integer2>");
			return;
		}
		v1, err1 := strconv.Atoi(args[2]);
		v2, err2 := strconv.Atoi(args[3]);
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: calc sqart <Integer>");
			return;
		}
		ret := simplemath.Add(v1, v2);
		fmt.Println("Result: ", ret);
	case "sqrt":
		if len(args) != 3 {
			fmt.Println("USAGE: calc sqart <Integer>");
			return;
		}
		v, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("USAGE: calc sqart <Integer>");
			return;
		}
		ret := simplemath.Sqrt(v);
		fmt.Println("Result: ", ret);
	default:
		Usage()
	}
}
