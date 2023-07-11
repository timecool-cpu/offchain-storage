package main

//
//import (
//	"context"
//	"fmt"
//	"github.com/libp2p/go-libp2p"
//	"github.com/libp2p/go-libp2p-core/crypto"
//	"github.com/libp2p/go-libp2p-core/host"
//	"github.com/libp2p/go-libp2p-core/network"
//	ma "github.com/multiformats/go-multiaddr"
//)
//
//// 创建主机
//func makeBasicHost(listenPort int, secio bool, randseed int64) (host.Host, error) {
//	priv, _, err := crypto.GenerateKeyPair(crypto.RSA, 2048)
//	if err != nil {
//		return nil, err
//	}
//
//	opts := []libp2p.Option{
//		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", listenPort)),
//		libp2p.Identity(priv),
//		libp2p.DisableRelay(),
//	}
//
//	basicHost, err := libp2p.New(opts...)
//	if err != nil {
//		return nil, err
//	}
//
//	hostAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/ipfs/%s", basicHost.ID().Pretty()))
//
//	fmt.Printf("I am %s\n", hostAddr)
//	fmt.Printf("Now run \"./echo -l %d -d %s\" on a different terminal\n", listenPort+1, hostAddr)
//
//	return basicHost, nil
//}
//
//// 处理流
//func handleStream(s network.Stream) {
//	fmt.Println("Got a new stream!")
//
//	// 创建一个缓冲区，将接收的数据读入
//	buf := make([]byte, 512)
//	_, err := s.Read(buf)
//	if err != nil {
//		return
//	}
//
//	// 输出接收到的消息
//	fmt.Println("Received message:", string(buf))
//
//	// 关闭流
//	s.Close()
//}
//
//func main() {
//	// 创建一个新的 libp2p Host
//	_, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	h, err := makeBasicHost(5001, false, 10)
//	if err != nil {
//		fmt.Println("Error creating host: ", err)
//	}
//
//	h.SetStreamHandler("/p2p/1.0.0", handleStream)
//
//	select {} // 阻止 main 结束
//}
