package sum

import "testing"

func TestSum(t *testing.T){
	t.Run("collection of numbers.", func(t *testing.T){
		numbers := []int{1,2,3,4,5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d but real is %d", got, want, got)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1,2,3}
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d but real is not %d", got, want, got)
		}
	})
}