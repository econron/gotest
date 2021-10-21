package geometer

import(
	"math"
)

// 図形を包括するinterfaceを実装する
type Shape interface {
	Area() float64
}

// 四角形について明示的な宣言を実行する
type Rectangle struct {
	width float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// 円の面積を取得するメソッドを書く
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
