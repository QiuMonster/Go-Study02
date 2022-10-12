package main

import (
	"fmt"
	"sync"
	"time"
)

var i int = 100
var wt sync.WaitGroup

var lock sync.Mutex //给资源进行加锁操作

//资源使用时 是先加锁后解锁 
func add() {
	defer wg.Done()
	lock.Lock()
	i += 1
	fmt.Printf("i++: %v \n", i)
	time.Sleep(time.Millisecond*10)
	
	lock.Unlock()
}
func sub() {
	defer wg.Done()
	lock.Lock()
	i -= 1
	fmt.Printf("i--: %v\n", i)
	time.Sleep(time.Millisecond*10)
	lock.Unlock()
}
func go_mutex() {
	//正常的输出使用资源
	// for i := 0; i < 100; i++ {
	// 	add()
	// 	sub()
	// }
	//使用协程时 可能发生资源访问出错问题
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go add()
		wg.Add(1)
		go sub()
	}
	//注意需要使主协程进行等待
	wg.Wait()
	fmt.Println("end..",i)
}
