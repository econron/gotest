package main

import(
	"fmt"
	"os"
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

func file_exists(name string) bool {
	_, err := os.Stat(name)
	if(err == nil) {
		return true
	}
	return false
}

func main(){
	var file_name string = "test.txt"
	
	if(!file_exists(file_name)){
		os.Create(file_name)
	}
	// 型指定してないとエディタ上でエラーになる
	// 変数を定義する場合は型指定して書いてあげる
	file, err :=  os.OpenFile(file_name, os.O_WRONLY | os.O_RDONLY, 0666)
	check(err)
	data := []byte{115, 111, 234, 111, 231}
	write_result, err := file.Write(data)
	check(err)
	fmt.Println("wrote %data bytes\n", write_result)
	defer file.Close()
}