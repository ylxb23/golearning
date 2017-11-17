package main

import (
	"fmt"
	"sync"
	"runtime"
)

var counter int = 0;
/** 共享内存的方式进行通信 (不推荐) */
//func main() {
//	testCount();
//}

func testCount() {
	lock := &sync.Mutex{}
	
	for i:=0; i<10; i++ {
		go Count(lock);
	}
	
	for {
		lock.Lock();
		
		c := counter;
		
		lock.Unlock();
		
		runtime.Gosched();
		if c >= 10 {
			break;
		}
	}
}

/**  */
func Count(lock *sync.Mutex) {
	lock.Lock();	// 加锁
	counter ++;
	
	fmt.Println(counter);
	
	lock.Unlock();	// 去除锁
}
