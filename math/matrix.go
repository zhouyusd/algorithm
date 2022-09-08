package math

type Matrix[T Number] [][]T

// DiagMatrix 创建主对角线填充 x 的方阵
func DiagMatrix[T Number](n int, x T) (mat Matrix[T]) {
	mat = NewMatrix[T](n, n)
	for i := 0; i < n; i++ {
		mat[i][i] = x
	}
	return mat
}

// NewMatrix 创建 xDim 行 yDim 列的矩阵
func NewMatrix[T Number](xDim, yDim int) (mat Matrix[T]) {
	mat = make(Matrix[T], xDim)
	for i := 0; i < xDim; i++ {
		mat[i] = make([]T, yDim)
	}
	return mat
}

// Int64 返回类型 int64 的矩阵
func (mat Matrix[T]) Int64() (rt Matrix[int64]) {
	xDim, yDim := mat.XDim(), mat.YDim()
	rt = NewMatrix[int64](xDim, yDim)
	for i := 0; i < xDim; i++ {
		for j := 0; j < yDim; j++ {
			rt[i][j] = int64(mat[i][j])
		}
	}
	return rt
}

// XDim 返回矩阵的 xDim 大小
func (mat Matrix[T]) XDim() (xDim int) {
	return len(mat)
}

// YDim 返回矩阵的 yDim 大小
func (mat Matrix[T]) YDim() (yDim int) {
	if mat.XDim() == 0 {
		return 0
	}
	return len(mat[0])
}

// Mul 矩阵乘法
func (mat Matrix[T]) Mul(oth Matrix[T]) (mul Matrix[T]) {
	xDim, yDim, zDim := mat.XDim(), mat.YDim(), oth.YDim()
	mul = NewMatrix[T](xDim, zDim)
	for i := 0; i < xDim; i++ {
		for j := 0; j < yDim; j++ {
			mul[i][j] = 0
			for k := 0; k < zDim; k++ {
				mul[i][j] += mat[i][k] * oth[k][j]
			}
		}
	}
	return mul
}

// MulMod 矩阵乘法模运算
func (mat Matrix[T]) MulMod(oth Matrix[T], mod int64) (mulMod Matrix[int64]) {
	xDim, yDim, zDim := mat.XDim(), mat.YDim(), oth.YDim()
	mulMod = NewMatrix[int64](xDim, zDim)
	for i := 0; i < xDim; i++ {
		for j := 0; j < yDim; j++ {
			mulMod[i][j] = 0
			for k := 0; k < zDim; k++ {
				mulMod[i][j] += int64(mat[i][k]) * int64(oth[k][j]) % mod
			}
		}
	}
	return mulMod
}

// Pow 计算矩阵 mat^n
// 若 mat 不是方阵，返回 nil 和 ErrSquareMatrix
func (mat Matrix[T]) Pow(n int) (pow Matrix[T], err error) {
	xDim, yDim := mat.XDim(), mat.YDim()
	if xDim != yDim {
		return nil, ErrSquareMatrix
	}
	x := mat
	pow = DiagMatrix[T](xDim, 1)
	for n != 0 {
		if n&1 == 1 {
			pow = pow.Mul(x)
		}
		x = x.Mul(x)
		n >>= 1
	}
	return pow, nil
}

// PowMod 计算矩阵 mat^n%mod
// 若 mat 不是方阵，返回 nil 和 ErrSquareMatrix
func (mat Matrix[T]) PowMod(n int, mod int64) (powMod Matrix[int64], err error) {
	xDim, yDim := mat.XDim(), mat.YDim()
	if xDim != yDim {
		return nil, ErrSquareMatrix
	}
	x := mat.Int64()
	powMod = DiagMatrix[int64](xDim, 1)
	for n != 0 {
		if n&1 == 1 {
			powMod = powMod.MulMod(x, mod)
		}
		x = x.MulMod(x, mod)
		n >>= 1
	}
	return powMod, nil
}
