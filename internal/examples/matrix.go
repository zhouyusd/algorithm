package main

import (
	"fmt"
	"github.com/zhouyusd/algorithm/math"
	"time"
)

func main() {
	mat := math.NewMatrix[int64](2, 2)
	mat[0][0] = 0
	mat[0][1] = 1
	mat[1][0] = 1
	mat[1][1] = 1
	begin := time.Now()
	ans, err := mat.PowMod(1000000000000000, 1000000007)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v, cost: %v\n", ans[0][1]+ans[1][1], time.Since(begin))
}
