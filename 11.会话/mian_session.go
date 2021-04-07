package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

// Sessions
// gorilla/sessions为自定义session后端提供cookie和文件系统session以及基础结构。
// 主要功能是：
// 简单的API：将其用作设置签名（以及可选的加密）cookie的简便方法。
// 内置的后端可将session存储在cookie或文件系统中。
// Flash消息：一直持续读取的session值。
// 切换session持久性（又称“记住我”）和设置其他属性的便捷方法。
// 旋转身份验证和加密密钥的机制。
// 每个请求有多个session，即使使用不同的后端也是如此。
// 自定义session后端的接口和基础结构：可以使用通用API检索并批量保存来自不同商店的session。

var stroe = sessions.NewCookieStore([]byte("something-very-secret"))

func main() {
	http.HandleFunc("/save", saveSession)
	http.HandleFunc("/get", getSession)
	http.HandleFunc("/delete", deleteSession)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
	}
}

func saveSession(w http.ResponseWriter, r *http.Request) {
	session, err := stroe.Get(r, "user")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["username"] = "ZM"
	session.Values["password"] = "123456"

	if err := session.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getSession(w http.ResponseWriter, r *http.Request) {
	session, err := stroe.Get(r, "user")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(session)
}

func deleteSession(w http.ResponseWriter, r *http.Request) {
	session, err := stroe.Get(r, "user")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 删除
	// 将session的最大存储时间设置为小于零的数即为删除
	session.Options.MaxAge = -1
	if err := session.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


