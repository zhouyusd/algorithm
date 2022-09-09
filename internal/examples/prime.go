package main

import (
	"fmt"
	"github.com/zhouyusd/algorithm/math"
)

func main() {
	primes, vis := math.EulerSieve(1000)
	fmt.Println(primes)
	for i := 0; i < len(vis); i++ {
		if !vis[i] {
			fmt.Print(i, " ")
		}
	}
}
