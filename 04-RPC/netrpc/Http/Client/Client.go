package main

import (
	"fmt"
	"log"
	"net/rpc"
)

//客户端

type Params struct {
	Width, Height int
}

func main() {
	rpc, err := rpc.DialHTTP("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}
	ret := 0
	//调用远程方法
	//注意第三个参数是指针类型
	err2 := rpc.Call("Rect.Area", Params{50, 100}, &ret)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(ret)
	err3 := rpc.Call("Rect.Perimeter", Params{50, 100}, &ret)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println(ret)

	/*
	   //	异步方式
	   	divCall := client.Go("Arith.Divide", Params{50, 100}, &ret, nil)
	   	<-divCall.Done
	   	fmt.Println(ret)
	*/

}
