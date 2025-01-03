package tmps

import (
	"math/big"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
)

type CommonParams struct {
	g_1 *bn256.G1 `json:"g_1"`
	g_2 *bn256.G2 `json:"g_2"`
	G_T *bn256.GT `json:"g_t"`
	b_1 *bn256.G1 `json:"b_1"`
	b_2 *bn256.G1 `json:"b_2"`
	b1  *big.Int  `json:"b1"`
	b2  *big.Int  `json:"b2"`
	N   int       `json:"n"`
	d   int       `json:"d"`
}

type Sk struct {
	g_1     *bn256.G1   `json:"g_1"`
	g_2     *bn256.G2   `json:"g_2"`
	xi      *big.Int    `json:"xi"`
	g_n1    []*bn256.G1 `json:"g_n1"`
	g_n2    []*bn256.G2 `json:"g_n2"`
	eta     []*big.Int  `json:"eta"`
	bf_z    []*bn256.G1 `json:"bf_z"`
	bf_b_n0 []*bn256.G2 `json:"bf_b_n0"`
	bf_b_n1 []*bn256.G2 `json:"bf_b_n1"`
}

type Pk struct {
	g_1 *bn256.G1 `json:"g_1"`
	g_2 *bn256.G2 `json:"g_2"`
	G_T *bn256.GT `json:"g_t"`
	b1  *bn256.G1 `json:"b1"`
	b2  *bn256.G1 `json:"b2"`
	N   int       `json:"n"`
	d   int       `json:"d"`
	r_0 *bn256.G2 `json:"r_0"`
	r_1 *bn256.G2 `json:"r_1"`
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
	W_i  []*big.Int  `json:"w_i"`
	q1_i []*bn256.G2 `json:"q1_i"`
	q2_i []*bn256.G2 `json:"q2_i"`
}

type Vk struct {
	challenge *big.Int  `json:"challenge"`
	vk        *bn256.G2 `json:"vk"`
}

type Pi struct {
	y   *big.Int  `json:"y"`
	pi1 *bn256.G2 `json:"pi1"`
	pi2 *bn256.G2 `json:"pi2"`
}

var bf_A []*Polynomial
var bf_B []*Polynomial
var bf_M []*Polynomial
var bf_P *Polynomial
