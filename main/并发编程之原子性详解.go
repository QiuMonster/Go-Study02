package main

import (
	"fmt"
	"sync/atomic"
)
//原子性操作的详解
func go_atomic_detail() {
	var i int32 = 100
	//原子性添加
	atomic.AddInt32(&i,1)
	fmt.Println(i)
	atomic.CompareAndSwapInt32(&i,101,102) //1 资源地址 2 原来的资源值 3 新的资源值
	fmt.Println(i) //修改成功

	//原子性的读取资源数据
	atomic.LoadInt32(&i)
	fmt.Println(i)

	//原子性的写入资源数据 算是资源的直接更新
	atomic.StoreInt32(&i,200)
	fmt.Println(i)
	//原子性的值的交换
	atomic.SwapInt32(&i,300)
	fmt.Println(i)
}
