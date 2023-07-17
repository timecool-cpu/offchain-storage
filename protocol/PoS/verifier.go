package PoS

import (
	"encoding/binary"
	//"fmt"
	"github.com/kwonalbert/spacemint/util"
	"golang.org/x/crypto/sha3"
)

type Verifier struct {
	pk   []byte // public key to verify the proof
	beta int    // number of challenges needed
	root []byte // root hash

	graph *Graph
	index int64 // index of the graphy in the family
	size  int64
	pow2  int64
	log2  int64
}

func NewVerifier(pk []byte, index int64, beta int, root []byte) *Verifier {
	size := numXi(index)
	log2 := util.Log2(size) + 1
	pow2 := int64(1 << uint64(log2))
	if (1 << uint64(log2-1)) == size {
		log2--
		pow2 = 1 << uint64(log2)
	}

	graph := &Graph{
		pk:    pk,
		index: index,
		log2:  log2,
		pow2:  pow2,
		size:  size,
	}

	v := Verifier{
		pk:   pk,
		beta: beta,
		root: root,

		graph: graph,
		index: index,
		size:  size,
		pow2:  pow2,
		log2:  log2,
	}
	return &v
}

// SelectChallenges TODO: need to select based on some pseudorandomness/gamma function?
//
//	Note that these challenges are different from those of cryptocurrency
func (v *Verifier) SelectChallenges(seed []byte) []int64 {
	challenges := make([]int64, v.beta*int(v.log2))
	rands := make([]byte, v.beta*int(v.log2)*8)
	sha3.ShakeSum256(rands, seed) //PRNG
	for i := range challenges {
		val, num := binary.Uvarint(rands[i*8 : (i+1)*8])
		if num < 0 {
			panic("Couldn't read PRNG")
		}
		challenges[i] = int64(val % uint64(v.size))
	}
	return challenges
}

/*
*
Verifier contract in Solidity
pragma solidity >=0.4.22 <0.9.0;

	contract Verifier {
	    uint256 constant hashSize = 32; // Assuming SHA-256
	    uint256 pow2;
	    bytes pk;

	    function VerifySpace(int64[] memory challenges, bytes32[][] memory hashes, bytes32[][][] memory parents, bytes32[][][] memory proofs, bytes32[][][][] memory pProofs) public view returns (bool) {
	        for (uint i = 0; i < challenges.length; i++) {
	            bytes memory val = abi.encodePacked(pk, challenges[i] + pow2);
	            for (uint j = 0; j < parents[i].length; j++) {
	                val = abi.encodePacked(val, parents[i][j]);
	            }
	            bytes32 exp = keccak256(val);
	            if (exp != hashes[i][0]) {
	                return false;
	            }
	            if (!Verify(challenges[i], hashes[i], proofs[i])) {
	                return false;
	            }

	            int64[] memory ps = GetParents(challenges[i], index);
	            for (uint j = 0; j < ps.length; j++) {
	                if (!Verify(ps[j], parents[i][j], pProofs[i][j])) {
	                    return false;
	                }
	            }
	        }
	        return true;
	    }
	}
*/
func (v *Verifier) VerifySpace(challenges []int64, hashes [][]byte, parents [][][]byte, proofs [][][]byte, pProofs [][][][]byte) bool {
	for i := range challenges {
		buf := make([]byte, hashSize)
		binary.PutVarint(buf, challenges[i]+v.pow2)
		val := append(v.pk, buf...)
		for _, ph := range parents[i] {
			val = append(val, ph...)
		}
		exp := sha3.Sum256(val)
		for j := range exp {
			if exp[j] != hashes[i][j] {
				return false
			}
		}
		if !v.Verify(challenges[i], hashes[i], proofs[i]) {
			return false
		}

		ps := v.graph.GetParents(challenges[i], v.index)
		for j := range ps {
			if !v.Verify(ps[j], parents[i][j], pProofs[i][j]) {
				return false
			}
		}
	}
	return true
}

func (v *Verifier) Verify(node int64, hash []byte, proof [][]byte) bool {
	curHash := hash
	counter := 0
	for i := node + v.pow2; i > 1; i /= 2 {
		var val []byte
		if i%2 == 0 {
			val = append(curHash, proof[counter]...)
		} else {
			val = append(proof[counter], curHash...)
		}
		hash := sha3.Sum256(val)
		curHash = hash[:]
		counter++
	}
	for i := range v.root {
		if v.root[i] != curHash[i] {
			return false
		}
	}

	return len(v.root) == len(curHash)
}
