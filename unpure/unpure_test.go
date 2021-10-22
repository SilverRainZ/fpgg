package unpure

import (
	"fmt"
	"strconv"

	"github.com/SilverRainZ/fpgg/data"
	"github.com/SilverRainZ/fpgg/util"
)

func ExampleMax() {
	i := data.FromSlice([]int{1, 2, 3, 4})
	v := MaxOfOrdered(i.Iter())
	fmt.Println("max:", v.Must())
	// Output:
	// max: 4
}

func ExampleReverse() {
	i := data.FromSlice([]int{1, 2, 3, 4})
	s := Reverse(i.Iter())
	fmt.Println(List(s))
	// Output:
	// [4 3 2 1]
}

func ExampleFilter() {
	i := data.FromSlice([]int{1, 2, 0, 3, 0, 4})
	s := Filter(util.NonZero[int], i.Iter())
	fmt.Println(List(s))
	// Output:
	// [1 2 3 4]
}

func ExampleMap() {
	f := func(v int) string { return strconv.Itoa(v) }
	i := data.FromSlice([]int{1, 2, 3, 4})
	s := Map(f, i.Iter())
	fmt.Println(List(s))
	// Output:
	// [1 2 3 4]
}

func ExampleTake() {
	i := data.FromSlice([]int{1, 2, 3, 4})
	s := Take(2, i.Iter())
	fmt.Println(List(s))
	// Output:
	// [1 2]
}

func ExampleConcat() {
	i := data.MatrixFromValue([][]int{[]int{1, 2}, []int{3, 4}})
	s := Concat(i.Iter())
	fmt.Println(List(s))
	// Output:
	// [1 2 3 4]
}
