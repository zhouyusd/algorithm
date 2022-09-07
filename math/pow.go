package math

// Pow 计算a^b O(log n)
func Pow(a, b int64) (pow int64) {
	pow = 1
	for b != 0 {
		if b&1 == 1 {
			pow *= a
		}
		a *= a
		b >>= 1
	}
	return pow
}

// Pow1 计算a^b O(n)
func Pow1(a, b int64) (pow int64) {
	pow = 1
	for i := int64(0); i < b; i++ {
		pow *= a
	}
	return pow
}

// PowMod 计算a^b%mod O(log n)
func PowMod(a, b, mod int64) (powMod int64) {
	powMod = 1
	a %= mod
	for b != 0 {
		if b&1 == 1 {
			powMod = powMod * a % mod
		}
		a = a * a % mod
		b >>= 1
	}
	return powMod % mod
}

// PowMod1 计算a^b%mod O(n)
func PowMod1(a, b, mod int64) (powMod int64) {
	powMod = 1
	a %= mod
	for i := int64(0); i < b; i++ {
		powMod = powMod * a % mod
	}
	return powMod % mod
}
