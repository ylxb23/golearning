package main

import (
	logger "log"
)

//func main() {
//	testLogger();
//}

func testLogger() {
	// 执行完后 os.Exit(1);
//	logger.Fatal("This is Fatal.");
//	logger.Fatalf("This is Fatalf, %s.", "呵呵");
//	logger.Fatalln("This is Fatalln.");

	logger.Println("This is Println.")
	
	// 执行完后打印堆栈，并 os.Exit(1)	
	logger.Panic("This is Panic.")
	logger.Panicf("This is Panicf, %s.", "呵呵");
	logger.Panicln("This is Panicln.");
}

type MyLogger struct {
	Level int
}
type T struct {
	*MyLogger
	Name string
	*logger.Logger
}