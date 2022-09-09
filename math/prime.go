package math

import "math"

// Prime 判断n是否为素数 O(sqrt(n))
func Prime(n int64) bool {
	if n < 2 {
		return false
	}
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// EulerSieve 欧拉筛法 O(n)
func EulerSieve(n int64) (primes []int64, vis []bool) {
	primes = make([]int64, 0, int64(math.Ceil(float64(n)/math.Log(float64(n)))))
	vis = make([]bool, n+1)
	vis[0], vis[1] = true, true
	for i := int64(2); i <= n; i++ {
		if !vis[i] {
			primes = append(primes, i)
		}
		for j := 0; j < len(primes) && i*primes[j] <= n; j++ {
			vis[i*primes[j]] = true
			if i%primes[j] == 0 {
				break
			}
		}
	}
	return primes, vis
}

// EratosthenesSieve 埃氏筛法 O(nlog(log(n)))
func EratosthenesSieve(n int64) (primes []int64, vis []bool) {
	primes = make([]int64, 0, int64(math.Ceil(float64(n)/math.Log(float64(n)))))
	vis = make([]bool, n+1)
	vis[0], vis[1] = true, true
	for i := int64(2); i <= n; i++ {
		if !vis[i] {
			vis[i] = true
			primes = append(primes, i)
			for j := i * i; j <= n; j += i {
				vis[j] = true
			}
		}
	}
	return primes, vis
}
