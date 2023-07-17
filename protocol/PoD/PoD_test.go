package PoD

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

func TestSol(T *testing.T) {
	// 生成一个新的私钥
	//privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	privateKey, err := crypto.HexToECDSA("af96c14f7494e2c452691207287791ba45849c8bd3baa84efb6ed91d823995da")
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	// 要签名的消息
	message := "Hello, world!"

	// 计算消息的 Keccak-256 哈希
	hash := crypto.Keccak256Hash([]byte(message))

	// 签名哈希
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatalf("Failed to sign hash: %v", err)
	}

	// 签名是一个65字节的数组，其中前32字节是r，接下来的32字节是s，最后一个字节是v
	r := new(big.Int).SetBytes(signature[:32])
	s := new(big.Int).SetBytes(signature[32:64])
	v := signature[64]

	// 如果v是0或1，需要加上27，以符合以太坊的签名标准
	if v < 27 {
		v += 27
	}

	fmt.Printf("Message: %s\n", message)
	fmt.Printf("Hash: %s\n", hash.Hex())
	fmt.Printf("Signature: %x\n", signature)
	fmt.Printf("R: %x\n", r)
	fmt.Printf("S: %x\n", s)

	// Set up client
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	// Set up sender address
	privateKey, err = crypto.HexToECDSA("af96c14f7494e2c452691207287791ba45849c8bd3baa84efb6ed91d823995da")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Set up contract instance
	contractAddress := common.HexToAddress("0x8C6E3b3Cb647831cB136BCCDb775B0D14D7B8Eaa") //部署的地址
	instance, err := NewPoD(contractAddress, client)                                     // 修改这里，使用你刚生成的包的函数
	if err != nil {
		log.Fatal(err)
	}

	// Set up auth
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// Set identifier and hash of hash
	identifier := "unique_identifier"
	hashOfHash := crypto.Keccak256Hash(hash.Bytes())

	// Call the SetHash function
	log.Printf("hashofHash: %x\n", hashOfHash.Bytes())
	log.Printf("fromAddress: %x\n", fromAddress.Bytes())
	tx, err := instance.SetData(auth, identifier, hashOfHash, fromAddress)
	if err != nil {
		log.Fatalf("Failed to call SetHash: %v", err)
	}
	log.Printf("Transaction sent: %s\n", tx.Hash().Hex())

	// Wait for the transaction to be mined
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatalf("Failed to mine transaction: %v", err)
	}

	rBytes := r.Bytes()
	sBytes := s.Bytes()

	var rArray [32]byte
	var sArray [32]byte

	copy(rArray[:], rBytes[len(rBytes)-32:])
	copy(sArray[:], sBytes[len(sBytes)-32:])
	log.Printf("rArray:0x %x\n", rArray)
	log.Printf("sArray:0x %x\n", sArray)
	log.Printf("hash:0x %x\n", hash.Bytes())

	// Call the Verify function
	//isValid, err := instance.VerifySignature(&bind.CallOpts{}, identifier, hash, v, rArray, sArray)
	//if err != nil {
	//	log.Fatalf("Failed to call Verify: %v", err)
	//}
	//fmt.Printf("Verification result: %v\n", isValid)
	addr, err := instance.GetAddr(&bind.CallOpts{}, identifier, hash, v, rArray, sArray)
	if err != nil {
		return
	}
	fmt.Printf("addr: %x\n", addr)
}

//func TestSol(T *testing.T) {
//	// 生成一个新的私钥
//	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
//	if err != nil {
//		log.Fatalf("Failed to generate private key: %v", err)
//	}
//
//	// 要签名的消息
//	message := "Hello, world!"
//
//	// 计算消息的 Keccak-256 哈希
//	hash := crypto.Keccak256Hash([]byte(message))
//
//	// 签名哈希
//	signature, err := crypto.Sign(hash.Bytes(), privateKey)
//	if err != nil {
//		log.Fatalf("Failed to sign hash: %v", err)
//	}
//
//	// 签名是一个65字节的数组，其中前32字节是r，接下来的32字节是s，最后一个字节是v
//	r := new(big.Int).SetBytes(signature[:32])
//	s := new(big.Int).SetBytes(signature[32:64])
//	v := signature[64]
//
//	// 如果v是0或1，需要加上27，以符合以太坊的签名标准
//	if v < 27 {
//		v += 27
//	}
//
//	fmt.Printf("Message: %s\n", message)
//	fmt.Printf("Hash: %s\n", hash.Hex())
//	fmt.Printf("Signature: %x\n", signature)
//	fmt.Printf("R: %x\n", r)
//	fmt.Printf("S: %x\n", s)
//
//	// Set up client
//	client, err := ethclient.Dial("http://localhost:8545")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Set up sender address
//	privateKey, err = crypto.HexToECDSA("af96c14f7494e2c452691207287791ba45849c8bd3baa84efb6ed91d823995da")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	publicKey := privateKey.Public()
//	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
//	if !ok {
//		log.Fatal("error casting public key to ECDSA")
//	}
//
//	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
//
//	// Set up contract instance
//	contractAddress := common.HexToAddress("0x0B0Fc8cD1a5ab9993246d187cd8eA23CB596D100") //部署的地址
//	instance, err := NewPoD(contractAddress, client)                                     // 修改这里，使用你刚生成的包的函数
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Set up auth
//	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	gasPrice, err := client.SuggestGasPrice(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	auth := bind.NewKeyedTransactor(privateKey)
//	auth.Nonce = big.NewInt(int64(nonce))
//	auth.Value = big.NewInt(0)     // in wei
//	auth.GasLimit = uint64(300000) // in units
//	auth.GasPrice = gasPrice
//
//	rBytes := r.Bytes()
//	sBytes := s.Bytes()
//
//	var rArray [32]byte
//	var sArray [32]byte
//
//	copy(rArray[:], rBytes[len(rBytes)-32:])
//	copy(sArray[:], sBytes[len(sBytes)-32:])
//
//	address, err := instance.VerifySignature(&bind.CallOpts{}, hash, uint8(v), rArray, sArray)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("Recovered address:", address.Hex())
//
//}

func TestProofOfDelivery_tmpfile(t *testing.T) {
	// 创建临时文件
	tempFile, err := ioutil.TempFile(os.TempDir(), "prefix")
	if err != nil {
		t.Fatalf("Cannot create temporary file: %s", err)
	}

	// 记住清理
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {

		}
	}(tempFile.Name())

	// 写一些数据到文件
	text := []byte("Hello World")
	if _, err = tempFile.Write(text); err != nil {
		t.Fatalf("Failed to write to temporary file: %s", err)
	}

	// 关闭文件
	if err := tempFile.Close(); err != nil {
		t.Fatal(err)
	}

	// 生成一个ECDSA私钥
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	// 测试proofOfDelivery
	timestamp, r, s, err := proofOfDelivery(tempFile.Name(), privateKey)
	if err != nil {
		t.Fatalf("Failed to generate proof of delivery: %v", err)
	}

	// 检查返回的时间戳是否近期
	now := time.Now()
	if now.Sub(timestamp) > time.Second {
		t.Fatalf("Timestamp is not recent: %v", timestamp)
	}

	// 检查签名是否非空
	if r.Sign() <= 0 || s.Sign() <= 0 {
		t.Fatalf("Invalid signature: r: %v, s: %v", r, s)
	}

	// 检查签名是否有效
	// 通过调用智能合约
}
