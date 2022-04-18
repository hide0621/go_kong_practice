// client-api/main.go

package main

import (
	//各フィールドをjson形式で表示
	"encoding/json"
	"net/http"
)

type Client struct {
	CompanyName string `json:"companyName"`
	Email       string `json:"email"`
}

//clientsというハンドラ関数
func clients(w http.ResponseWriter, req *http.Request) {
	//wからヘッダーを取得、そのヘッダー情報についてのレスポンスの形式を指定（Content-Type：application/json）
	w.Header().Set("Content-Type", "application/json")
	user := Client{
		CompanyName: "John Inc.",
		Email:       "john@example.com",
	}
	var clients []Client
	//スライス自体を保持する変数に追加したい要素を格納
	clients = append(clients, user)
	//json形式で構造体Clientのフィールドを8003番ポートのwebサーバーに表示させる
	json.NewEncoder(w).Encode(clients)
}

func main() {
	//第二引数にハンドラ関数を渡す
	http.HandleFunc("/clients", clients)
	//第二引数にはハンドラ型かハンドラ関数、又はnilを渡す
	//nilの場合はDefaultServeMuxを使う
	http.ListenAndServe(":8003", nil)
}
