package math

import "golang.org/x/exp/constraints"

// Gcd 计算a和b的最大公约数
func Gcd[T constraints.Integer](a, b T) (gcd T) {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

// ExtGcd 计算a和b的最大公约数并求解ax+by=gcd(a,b)
func ExtGcd[T constraints.Integer](a, b T) (gcd T, x T, y T) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, y, x = ExtGcd(b, a%b)
	y -= (a / b) * x
	return gcd, x, y
}

// Lcm 计算a和b的最小公倍数
func Lcm[T constraints.Integer](a, b T) (lcm T) {
	if a < b {
		a, b = b, a
	}
	gcd := Gcd(a, b)
	a /= gcd
	return a * b
}
