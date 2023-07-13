package PoD

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

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
}

func TestProofOfDelivery(t *testing.T) {
	// 在当前目录中创建一个文件
	fileName := "testfile.txt"
	file, err := os.Create(fileName)
	if err != nil {
		t.Fatalf("Cannot create file: %s", err)
	}

	// 记住清理
	defer os.Remove(file.Name())

	// 写一些数据到文件
	text := []byte("Hello World")
	if _, err = file.Write(text); err != nil {
		t.Fatalf("Failed to write to file: %s", err)
	}

	// 关闭文件
	if err := file.Close(); err != nil {
		t.Fatal(err)
	}

	// 生成一个ECDSA私钥
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	// 测试proofOfDelivery
	timestamp, r, s, err := proofOfDelivery(file.Name(), privateKey)
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
}
