package main

import (
	"fmt"
)

/**
	不确定参数类型的参数：args ...interface{}, eg: fmt.Println(a ...interface{});
*/

type IFly interface {
	Fly();
	Sing(word string);
}

type Bird struct {
}

func (b *Bird) Fly() {
	fmt.Println("bird fly.");
}
func (b *Bird) Sing(word string) {
	fmt.Println(word);
}

func playInterface() {
	var fly IFly = new(Bird);
	fly.Fly();
	fly.Sing("Singing a song.");
}

func sum(values []int, resultChan chan int) {
	sum := 0;
	for _, value:=range values {
		sum += value;
	}
	resultChan <- sum // 将计算结果发送到 channel 中
}

func playSum() {
	values := []int{1,2,3,4,5,6,7,8,9,10};
	resultChan := make(chan int, 2);
	go sum(values[:len(values)/2], resultChan);
	go sum(values[len(values)/2:], resultChan);
	
	sum1, sum2 := <- resultChan, <- resultChan	// 接收结果
	fmt.Println("Result: ", sum1, sum2, sum1+sum2);
}

// ---------------------------------------------

type File struct {
	// ...
}
func (f *File) Read(buf []byte) (n int, err error) {
	fmt.Println("This is File Read.")
	return -1, nil;
}
func (f *File) Write(buf []byte) (n int, err error) {
	fmt.Println("This is File Write.")
	return -1, nil;
}
func (f *File) Seek(offset int64, whence int) (pos int64, err error) {
	fmt.Println("This is File Seek.")
	return int64(-1), nil;
}
func (f *File) Close() error {
	fmt.Println("This is File Close.")
	return nil;
}

type IFile interface {
	Read(buf []byte) (n int, err error);
	Write(buf []byte) (n int, err error);
	Seek(offset int64, whence int) (pos int64, err error);
	Close() error;
}
type IRead interface {
	Read(buf []byte) (n int, err error);
}
type IWrite interface {
	Write(buf []byte) (n int, err error);
}
type ISeek interface {
	Seek(offset int64, whence int) (pos int64, err error);
}
type IClose interface {
	Close() error;
}


//func main() {
//	playSum();
//	
//	fmt.Println("-------------------------------------");
//	
//	var file1 IFile = new(File);
//	var file2 IRead = new(File);
//	var file3 IWrite = new(File);
//	var file4 ISeek = new(File);
//	var file5 IClose = new(File);
//	
//	file1.Read(nil);
//	file2.Read(nil);
//	file3.Write(nil);
//	file4.Seek(1, 1);
//	file5.Close();
//}