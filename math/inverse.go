package math

// Inverse 计算1-n模p下的逆元
// p是素数
func Inverse(n, p int64) (inv []int64) {
	inv = make([]int64, n+1)
	inv[0] = 0
	inv[1] = 1
	for i := int64(2); i <= n; i++ {
		inv[i] = (p - p/i) * inv[p%i] % p
	}
	return inv
}

// Inverse1 费马小定理求逆元 a*a^(p-2)=1(mod p)
// p是素数
func Inverse1(a, p int64) (inv int64) {
	return PowMod(a, p-2, p)
}

// Inverse2 拓展欧几里得算法求逆元
// 要求gcd(a,p)=1，否则返回 0 和 ErrInverse
func Inverse2(a, p int64) (inv int64, err error) {
	gcd, x, _ := ExtGcd(a, p)
	if gcd != 1 {
		return 0, ErrInverse
	}
	return (p + (x % p)) % p, nil
}
