package main

import (
	"fmt"
	"time"
)

func f(str string){
	for i := 0; i < 3; i++ {
		fmt.Println(str, ":" ,i)
	}
}

func main(){
	f("direct")
	// goで並行処理を実行したい場合、go hogehogeと書くこと
	go f ("goroutine")

	// つまりこれはgoで無名関数をコールしているだけ
	go func(msg string){
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}

// RESULT::::::
// direct : 0
// direct : 1
// direct : 2
// going
// goroutine : 0
// goroutine : 1
// goroutine : 2
// done