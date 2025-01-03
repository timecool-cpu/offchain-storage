package tmps

import (
	"math/big"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
)

type CommonParams struct {
	g_1 *bn256.G1 `json:"G_1"`
	g_2 *bn256.G2 `json:"G_2"`
	G_T *bn256.GT `json:"g_t"`
	b_1 *bn256.G1 `json:"b_1"`
	b_2 *bn256.G1 `json:"b_2"`
	b1  *big.Int  `json:"B1"`
	b2  *big.Int  `json:"B2"`
	N   int       `json:"n"`
	d   int       `json:"d"`
}

type Sk struct {
	g_1     *bn256.G1   `json:"G_1"`
	g_2     *bn256.G2   `json:"G_2"`
	xi      *big.Int    `json:"xi"`
	g_n1    []*bn256.G1 `json:"g_n1"`
	g_n2    []*bn256.G2 `json:"g_n2"`
	eta     []*big.Int  `json:"eta"`
	bf_z    []*bn256.G1 `json:"bf_z"`
	bf_b_n0 []*bn256.G2 `json:"bf_b_n0"`
	bf_b_n1 []*bn256.G2 `json:"bf_b_n1"`
}

type Pk struct {
	G_1 *G1Alias `json:"G_1"`
	G_2 *G2Alias `json:"G_2"`
	G_T *GTAlias `json:"g_t"`
	B1  *G1Alias `json:"B1"`
	B2  *G1Alias `json:"B2"`
	N   int      `json:"n"`
	D   int      `json:"d"`
	R_0 *G2Alias `json:"R_0"`
	R_1 *G2Alias `json:"R_1"`
}

// GetG1 返回 G_1 的值
func (pk *Pk) GetG1() *bn256.G1 {
	return (*bn256.G1)(pk.G_1)
}

// GetG2 返回 G_2 的值
func (pk *Pk) GetG2() *bn256.G2 {
	return (*bn256.G2)(pk.G_2)
}

type Ek struct {
	W_i  []*big.Int  `json:"w_i"`
	q1_i []*bn256.G2 `json:"q1_i"`
	q2_i []*bn256.G2 `json:"q2_i"`
}

type Vk struct {
	Challenge *BigIntAlias `json:"Challenge"`
	Vk        *G2Alias     `json:"Vk"`
}

type Pi struct {
	Y   *BigIntAlias `json:"Y"`
	Pi1 *G2Alias     `json:"Pi1"`
	Pi2 *G2Alias     `json:"Pi2"`
}

var bf_A []*Polynomial
var bf_B []*Polynomial
var bf_M []*Polynomial
var bf_P *Polynomial
