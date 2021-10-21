package main

import (
	"context"
	"net/http"
	"time"
)

func handleSearch(w http.ResponseWriter, req *http.Request){
	// contextのキャンセルファンクションを定義。
	// コンテキストを中断するメソッド
	var(
		ctx context.Context
		cancel context.CancelFunc
	)

	// 取得したリクエストのtimeoutパラメーターからtimeout変数を定義
	timeout, err := time.ParseDuration(req.FormValue("timeout"))

	if err == nil {
		// エラーがなくても、タイムアウトしたらキャンセルするようにしている
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	}else{
		// エラーがあれば、どうしてコンテキストを終了したかを伝えながら終了する
		ctx, cancel = context.WithCancel(context.Background())
	}

	// 処理が一通り終了したらコンテキストを削除する
	defer cancel()
}

func main(){

}