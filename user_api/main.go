// user-api/main.go

package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	//各フィールドをjson形式で表示
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

//usersというハンドラ関数
func users(w http.ResponseWriter, req *http.Request) {
	//wからヘッダーを取得、そのヘッダー情報についてのレスポンスの形式を指定（Content-Type：application/json）
	w.Header().Set("Content-Type", "application/json")
	user := User{
		FirstName: "John",
		LastName:  "Doe",
	}
	var users []User
	//スライス自体を保持する変数に追加したい要素を格納
	users = append(users, user)
	//json形式で構造体Userのフィールドを8002番ポートのwebサーバーに表示させる
	json.NewEncoder(w).Encode(users)
}

func main() {
	//第二引数にハンドラ関数を渡す
	http.HandleFunc("/users", users)
	//第二引数にはハンドラ型かハンドラ関数、又はnilを渡す
	//nilの場合はDefaultServeMuxを使う
	http.ListenAndServe(":8002", nil)
}
