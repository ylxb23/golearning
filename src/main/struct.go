package main

import (
	logger "log"
)

//func main() {
//	apple := new(Apple);
//	apple.test();
//	apple.Name = "Me is Apple";
//	logger.Println(apple.Name);
//	apple.Color = "Red";
//	logger.Println(apple.Color);
//	
////	apple.Age = 13;
////	logger.Println(apple.Age);
//	apple.whoami();
////	var a int = 13;
////	apple.Age = a;
////	logger.Println("equire: ", apple.Age == apple.OtherType.Age);
//}

type OtherType struct {
	Age int;
}
func (o *OtherType) whoami() {
	logger.Println("This is other type.")
}

type Fruit struct {
	Name string;
}

type Apple struct {
	Fruit;
	Color string;
	*OtherType;
}

func (f *Fruit) test() {
	logger.Println("test fruit.");
}

func (a *Apple) test() {
	a.Fruit.test();
	logger.Println("------------------- apple -------------------");
}