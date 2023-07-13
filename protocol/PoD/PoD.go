package PoD

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"math/big"
	"time"
)

func proofOfDelivery(filePath string, privateKey *ecdsa.PrivateKey) (timestamp time.Time, signatureR *big.Int, signatureS *big.Int, err error) {
	//// 解析私钥
	//block, _ := pem.Decode([]byte(privateKeyPem))
	//if block == nil {
	//	return time.Time{}, nil, nil, fmt.Errorf("failed to parse private key")
	//}
	//privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	//if err != nil {
	//	return time.Time{}, nil, nil, fmt.Errorf("failed to parse private key: %v", err)
	//}

	// 读取文件内容
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return time.Time{}, nil, nil, fmt.Errorf("failed to read file: %v", err)
	}

	// 计算文件的SHA-256哈希
	hash := sha256.Sum256(content)

	// 生成时间戳
	timestamp = time.Now()

	// 签名哈希
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return time.Time{}, nil, nil, fmt.Errorf("failed to sign hash: %v", err)
	}

	return timestamp, r, s, nil
}
