package PoS

import (
	"crypto/rand"
	"fmt"
	space "github.com/kwonalbert/spacemint/pos"
	"log"
	"testing"
	"time"
)

var prover *space.Prover = nil
var verifier *space.Verifier = nil

func TestPoS(t *testing.T) {
	seed := make([]byte, 64)
	rand.Read(seed)
	challenges := verifier.SelectChallenges(seed)
	now := time.Now()
	hashes, parents, proofs, pProofs := prover.ProveSpace(challenges)
	fmt.Printf("Prove: %f\n", time.Since(now).Seconds())

	now = time.Now()
	if !verifier.VerifySpace(challenges, hashes, parents, proofs, pProofs) {
		log.Fatal("Verify space failed:", challenges)
	}
	fmt.Printf("Verify: %f\n", time.Since(now).Seconds())
}
