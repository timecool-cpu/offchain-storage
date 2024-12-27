package tmps

import (
	"fmt"
	tmps "github.com/timecool-cpu/offchain-storage/blockchain-crypto/random"
	"testing"
)

func TestFileToPolynomial(t *testing.T) {
	poly := &Polynomial{}
	poly.prime = tmps.Prime
	poly, _ = poly.FileToPolynomial("/Users/bytedance/GolandProjects/offchain-storage/testdata/fuel.xlsx")
	fmt.Println(poly)
	poly.PolynomialToFile("test.xlsx")
}
