package tmps

import (
	"math/big"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
)

func Prove(c *big.Int, ek *Ek) *Pi {
	mul := big.NewInt(1)

	y := big.NewInt(0)
	pi1 := new(bn256.G2).ScalarBaseMult(big.NewInt(0)) // 初始化为零值
	pi2 := new(bn256.G2).ScalarBaseMult(big.NewInt(0)) // 初始化为零值

	// 计算 y = Σ(mul * ek.W_i[i])
	for i := len(ek.W_i) - 1; i >= 0; i-- {
		y.Add(y, new(big.Int).Mul(mul, ek.W_i[i]))
		mul.Mul(mul, c)
	}

	// 重置mul为1，并计算 pi1 和 pi2
	mul.SetInt64(1)
	for i := len(ek.q1_i) - 1; i >= 0; i-- {
		tmp := new(bn256.G2).ScalarMult(ek.q1_i[i], mul)
		pi1.Add(pi1, tmp)
		mul.Mul(mul, c)
	}

	// 重置mul为1，并计算 pi2
	mul.SetInt64(1)
	for i := len(ek.q2_i) - 1; i >= 0; i-- {
		tmp := new(bn256.G2).ScalarMult(ek.q2_i[i], mul)
		pi2.Add(pi2, tmp)
		mul.Mul(mul, c)
	}

	// 打印调试信息
	PrintG2(pi1, "pi1")
	PrintG2(pi2, "pi2")

	// 构造并返回 Pi 结构
	pi := &Pi{
		y:   y,
		pi1: pi1,
		pi2: pi2,
	}

	return pi
}
