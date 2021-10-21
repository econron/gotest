package main

import(
	"fmt"
)

// 整数値型の値が入った配列をストリームする
// 配列が空になるとチャンネルを閉じる
func gen(nums ...int) <- chan int{
	out := make(chan int)
	go func(){
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// 上記と同じ感じで二乗する処理を実行する
func Square(nums ...int) <- chan int{
	square := make(chan int)
	go func(){
		for _, n := range nums{
			square <- n * n
		}
		close(square)
	}()
	return square
}

func main(){
	c := gen(2,3)
	out := Square(c)

	fmt.Println(<-out)
	fmt.Println(<-out)
}