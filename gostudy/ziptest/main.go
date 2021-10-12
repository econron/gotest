package main

import(
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

func main(){
	// test.zipファイルを開ける
	// ReadCloserとerrorを返す
	r, err := zip.OpenReader("test.zip")
	// errorがある場合、ログにエラー内容を吐き出す
	if err != nil {
		log.Fatal(err)
	}
	// 処理の最後に必ずこの処理を実行する
	// 具体的には、ファイルを閉じる処理
	defer r.Close()

	// foreachの代わり。回数を指定しない場合、_で数値宣言を放棄する
	for _, f := range r.File{
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			// ドキュメントによれば下記の通り。だからCLIでメッセージがきた
			// Fatal is equivalent to Print() followed by a call to os.Exit(1).
			log.Fatal(err)
		}
		f, err := io.CopyN(os.Stdout, rc, 68)
		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println(f)
	}

}