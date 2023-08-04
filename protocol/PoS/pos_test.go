package PoS

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"testing"
	"time"
)

// exp* gets setup in test.go
var prover *Prover = nil
var verifier *Verifier = nil
var pk []byte
var index int64 = 3
var size int64 = 0
var beta int = 30
var graphDir string = "Xi"
var name string = "G"

func TestPoS(t *testing.T) {
	//seed := make([]byte, 64)
	seed := []byte{211, 235, 5, 101, 140, 84, 244, 89, 197, 165, 155, 171, 154, 60, 31, 164, 252, 224, 49, 106, 204, 108, 193, 179, 190, 143, 100, 159, 1, 189, 143, 25, 26, 125, 93, 194, 72, 134, 216, 160, 106, 132, 150, 200, 207, 24, 148, 82, 3, 165, 208, 207, 24, 228, 193, 94, 55, 21, 62, 79, 189, 180, 232, 23}
	fmt.Println("seed: ", seed)
	challenges := verifier.SelectChallenges(seed)
	now := time.Now()
	hashes, parents, proofs, pProofs := prover.ProveSpace(challenges)
	fmt.Printf("Prove: %f\n", time.Since(now).Seconds())
	fmt.Println("challenges: ", challenges)
	fmt.Println("hashes: ", hashes)
	fmt.Println("parents: ", parents)
	fmt.Println("proofs: ", proofs)
	fmt.Println("pProofs: ", pProofs)

	fmt.Println("verifier pk: ", verifier.pk)
	fmt.Println("verifier index: ", verifier.index)
	fmt.Println("verifier beta: ", verifier.beta)
	fmt.Println("verifier root: ", verifier.root)

	now = time.Now()
	if !verifier.VerifySpace(challenges, hashes, parents, proofs, pProofs) {
		log.Fatal("Verify space failed:", challenges)
	}
	fmt.Printf("Verify: %f\n", time.Since(now).Seconds())
}

func TestMain(m *testing.M) {
	size = numXi(index)
	pk = []byte{1}

	runtime.GOMAXPROCS(runtime.NumCPU())

	id := flag.Int("index", 1, "graph index")
	flag.Parse()
	index = int64(*id)

	graphDir = fmt.Sprintf("%s%d", graphDir, *id)
	//os.RemoveAll(graphDir)

	now := time.Now()
	prover = NewProver(pk, index, name, graphDir)
	fmt.Printf("%d. Graph gen: %fs\n", index, time.Since(now).Seconds())

	now = time.Now()
	commit := prover.Init()
	fmt.Printf("%d. Graph commit: %fs\n", index, time.Since(now).Seconds())

	root := commit.Commit
	verifier = NewVerifier(pk, index, beta, root)
	code := m.Run()
	os.Exit(code)
}

func TestHash(t *testing.T) {
	sum256 := sha256.Sum256([]byte("hello"))
	fmt.Println(sum256)
}
