package main

import (
	"fmt"
	"log"
	rpc2 "net/rpc"
)

type Params struct {
	Width, Height int
}

func main() {
	rpc, err := rpc2.Dial("tcp", "127.0.0.1:2001")
	if err != nil {
		log.Fatal(err)
	}
	ret := 10
	//调用RPC服务
	for i := 0; i < 20; i++ {
		err2 := rpc.Call("Rect.Area", Params{50, 100}, &ret)
		if err2 != nil {
			log.Fatal(err2)
		}
		fmt.Println(ret)
	}

	//调用RPC服务
	err3 := rpc.Call("Rect.Perimeter", Params{50, 100}, &ret)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println(ret)
}
