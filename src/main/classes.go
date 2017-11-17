package main

import (
	"fmt"
	"errors"
)

//func main() {
//	testStructSimple();
//	testInitStruct();
//}

func testInitStruct() {
	fmt.Println("------------------------------")
	book1 := new(Book);
	book2 := &Book{};
	book3 := &Book{"书名", "作者", 99, "ISBN"};
	book4 := &Book{name:"Go语言编程", author:"Who", price: 23.00}
	
	fmt.Println(*book1, *book2, *book3, *book4)
}

func testStructSimple() {
	var b1, b2 Book = Book{"大型网站技术架构", "李智慧", 49.00, "978-7-121-21200-0"}, Book{"大型网站技术架构", "李智慧", 49.00, "978-7-121-21200-0"};
	b3 := Book{"大型网站技术架构2", "李智慧", 99.00, "978-7-121-21200-2"};
	ba,err1 := b1.bookAdd(b2);
	bb,err2 := b1.bookAdd(b3);
	fmt.Println(ba, err1)
	fmt.Println(bb, err2)
	
	var a Integer = 9;
	fmt.Println(a.Less(13))
	fmt.Println("Init:\t a=", a);
	fmt.Println("Add: \t a=", a.Add(91))
	fmt.Println("Added:\t a=", a);
	fmt.Println("Sub: \t a=", a.Sub(91))
	fmt.Println("Subed:\t a=", a);
}

type Book struct {
	name string
	author string
	price float64
	isbn string
}


func (b1 Book) bookAdd(b2 Book) (b Book, err error) {
	if b1.isbn != b2.isbn || b1.name != b2.name || b1.author != b2.author {
		err = errors.New("书目不一样,不能相加")
		return b, err
	}
	b.author = b1.author
	b.isbn = b1.isbn
	b.name = b1.name
	b.price = b1.price + b2.price
	return b, err;
}

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b;
}

func (a *Integer) Add(b Integer) Integer {
	*a = *a + b;
	return *a;
}

func (a Integer) Sub(b Integer) Integer {
	a = a - b;
	return a;
}