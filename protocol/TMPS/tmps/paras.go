package tmps

import (
	"math/big"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
)

type CommonParams struct {
	g_1 *bn256.G1
	g_2 *bn256.G2
	G_T *bn256.GT
	b_1 *bn256.G1
	b_2 *bn256.G1
	b1  *big.Int
	b2  *big.Int
	N   int
	d   int
}

type Sk struct {
	g_1     *bn256.G1
	g_2     *bn256.G2
	xi      *big.Int
	g_n1    []*bn256.G1
	g_n2    []*bn256.G2
	eta     []*big.Int
	bf_z    []*bn256.G1
	bf_b_n0 []*bn256.G2
	bf_b_n1 []*bn256.G2
}

type Pk struct {
	g_1 *bn256.G1
	g_2 *bn256.G2
	G_T *bn256.GT
	b1  *bn256.G1
	b2  *bn256.G1
	N   int
	d   int
	r_0 *bn256.G2
	r_1 *bn256.G2
}

// GetG1 返回 g_1 的值
func (pk *Pk) GetG1() *bn256.G1 {
	return pk.g_1
}

// GetG2 返回 g_2 的值
func (pk *Pk) GetG2() *bn256.G2 {
	return pk.g_2
}

type Ek struct {
	W_i  []*big.Int
	q1_i []*bn256.G2
	q2_i []*bn256.G2
}

type Vk struct {
	challenge *big.Int
	vk        *bn256.G2
}

type Pi struct {
	y   *big.Int
	pi1 *bn256.G2
	pi2 *bn256.G2
}

var bf_A []*Polynomial
var bf_B []*Polynomial
var bf_M []*Polynomial
var bf_P *Polynomial
