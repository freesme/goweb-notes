package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

//访问http://localhost:8080/name=hanru&age=30

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("----------------------------")
	r.ParseForm() //解析参数
	//控制台打印信息
	fmt.Println(r.Form) //信息是输出到服务器的打印信息
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println(r.Form["url_ long"])

	for k, v := range r.Form {
		fmt.Println("Key:", k)
		fmt.Println("val：", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello GoWeb")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
