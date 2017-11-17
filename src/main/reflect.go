package main

import (
	"fmt"
	"reflect"
)

/** go 映射 */
type Dog struct {
	Name string
	LifeExpectance int
}

func (b *Dog) Fly() {
	fmt.Println("I am flying.")
}

func testReflect() {
	sparrow := &Dog{"Sparrow", 3}
	s := reflect.ValueOf(sparrow).Elem()
	typeOfT := s.Type();
	for i:=0; i<s.NumField(); i++ {
		f := s.Field(i);
		fmt.Printf("%d: %s %s = %v \n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
