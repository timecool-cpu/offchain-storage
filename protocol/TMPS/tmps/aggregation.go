package tmps

import (
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"math/big"
)

func Aggregation(pis []*Pi, vks []*Vk) (*Pi, *Vk) {
	//生成pis长度的随机系数
	r := GenerateRandomBigIntArray(Prime, len(pis))
	y := new(big.Int)
	pi1 := new(bn256.G2)
	pi2 := new(bn256.G2)
	vk_a := new(bn256.G2)

	for i := 0; i < len(pis); i++ {
		y.Add(y, new(big.Int).Mul(r[i], pis[i].y))
		pi1.Add(pi1, new(bn256.G2).ScalarMult(pis[i].pi1, r[i]))
		pi2.Add(pi2, new(bn256.G2).ScalarMult(pis[i].pi2, r[i]))
		vk_a.Add(vk_a, new(bn256.G2).ScalarMult(vks[i].vk, r[i]))
	}
	pi := Pi{
		y:   y,
		pi1: pi1,
		pi2: pi2,
	}

	vk1 := Vk{challenge: vks[0].challenge,
		vk: vk_a,
	}
	return &pi, &vk1
}
