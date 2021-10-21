package main

import(
	"fmt"
)

// 整数型の配列と整数値を取得し、整数値型の配列を返す
func appendInt(x []int, y int) []int{
	// 整数値型の配列を定義する
	var z []int
	// zの大きさを１大くする
	zlen := len(x) + 1
	// 初期化した配列のサイズが
	// 引数で取得した整数値配列のサイズより小さい時
	// 初期化した配列の大きさを引数の配列と等しくする
	if zlen <= cap(x) {
		z = x[:zlen]
	}else{
		// もしサイズが等しいなら
		zcap := zlen
		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func main(){
	var runes []rune
	for _, r := range "Hello 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
}

// ['H' 'e' 'l' 'l' 'o' ' ' '世' '界']