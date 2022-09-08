package main

import (
	"fmt"
	"github.com/zhouyusd/algorithm/math"
)

func main() {
	mat := math.NewMatrix[int](2, 2)
	fmt.Println(mat.XDim(), mat.YDim())
	fmt.Printf("%v\n", mat)
	mat1 := mat.Mul()
	fmt.Printf("%v\n", mat1)
}
