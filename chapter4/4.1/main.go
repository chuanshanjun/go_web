package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {

}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method) // 获取请求方法
	if r.Method == http.MethodGet {
		t, _ := template.ParseFiles("chapter4/4.1/login.html")
		log.Println(t.Execute(w, nil))
	} else {
		err := r.ParseForm() // 解析url传递的参数，对于POST则响应包的主体(request body)
		if err != nil {
			// handle error
			log.Fatal("ParseForm", err)
		}

		// 请求的是登陆数据，那么执行登陆的逻辑判断
		fmt.Println("username", r.Form["username"])
		fmt.Println("password", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)       // 设置访问的路由
	http.HandleFunc("/login", login)         // 设置监听的端口
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
