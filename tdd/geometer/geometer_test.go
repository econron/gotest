package geometer

import (
	"testing"
)

func TestArea(t *testing.T){

	checkArea := func(t *testing.T, shape Shape, want float64){
		t.Helper()
		expected := shape.Area()
		if expected != want {
			t.Errorf("expected %.2f, want %.2f, but %.2f", expected, want, expected)
		}
	}

	t.Run("rectangles", func(t *testing.T) {        
		rectangle := Rectangle{12, 6}        
		checkArea(t, rectangle, 72.0)    
	})

	t.Run("circles", func(t *testing.T) {        
		circle := Circle{10}        
		checkArea(t, circle, 314.1592653589793)    
	})

	// t.Run("float64 args test", func(t *testing.T){
	// 	rectangle := Rectangle{10.0, 10.0}
	// 	expected := rectangle.Area()
	// 	want := 100.0

	// 	if expected != want {
	// 		t.Errorf("expected %.2f, want %.2f, but %.2f", expected, want, expected)
	// 	}
	// })

	// t.Run("circle", func(t *testing.T){
	// 	circle := Circle{10}
	// 	expected := circle.Area()
	// 	want := 314.1592653589793

	// 	if expected != want {
	// 		t.Errorf("expected %.2f, want %.2f, but %.2f", expected, want, expected)
	// 	}
	// })
}