package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/setCookie", SetCookie)
	http.HandleFunc("/getCookie", GetCookie)
	http.ListenAndServe(":8080", nil)

	http.ListenAndServe(":8080", nil)
}

//设置Cookie

func SetCookie(w http.ResponseWriter, r *http.Request) {
	//构造Cookie
	cookie := http.Cookie{
		Name:   "Name",
		Value:  "hanru",
		Path:   "/",
		MaxAge: 60,
	}
	http.SetCookie(w, &cookie)
	w.Write([]byte("write cookie ok"))
}

//获取Cookie
func GetCookie(w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		fmt.Fprint(w, cookie.Name)
	}
}
