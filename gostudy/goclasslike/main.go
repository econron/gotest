package main

import "fmt"

type Person struct{
	FirstName string
	LastName string
}

func(p *Person) getName() string{
	return p.FirstName + ", " + p.LastName
}

func main(){
	// 構造体への代入をするとき、改行したらエラーになった。
	person := &Person{"Taro", "Yamada"}
	fmt.Println(person.getName())
}