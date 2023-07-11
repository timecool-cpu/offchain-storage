package main

//
//
//// dirty example to use go-orbitdb eventlog
//
//import (
//	orbitdb "berty.tech/go-orbit-db"
//	"berty.tech/go-orbit-db/iface"
//	"context"
//	"fmt"
//	icore "github.com/ipfs/interface-go-ipfs-core"
//	ifGoIpfsCore "github.com/ipfs/interface-go-ipfs-core"
//	"github.com/ipfs/kubo/config"
//	"github.com/ipfs/kubo/core"
//	"github.com/ipfs/kubo/core/coreapi"
//	"github.com/ipfs/kubo/core/node/libp2p"
//	"github.com/ipfs/kubo/plugin/loader" // This package is needed so that all the preloaded plugins are loaded automatically
//	"github.com/ipfs/kubo/repo/fsrepo"
//	"io"
//	"os"
//	"path/filepath"
//	"sync"
//	"time"
//)
//
//var flagExp = true //flag.Bool("experimental", false, "enable experimental features")
//
//// Creates an IPFS node and returns its coreAPI
//func createNode(ctx context.Context, repoPath string) (*core.IpfsNode, error) {
//	// Open the repo
//	repo, err := fsrepo.Open(repoPath)
//	if err != nil {
//		return nil, err
//	}
//
//	// Construct the node
//
//	nodeOptions := &core.BuildCfg{
//		Online:  true,
//		Routing: libp2p.DHTOption, // This option sets the node to be a full DHT node (both fetching and storing DHT Records)
//		// Routing: libp2p.DHTClientOption, // This option sets the node to be a client DHT node (only fetching records)
//		Repo: repo,
//		ExtraOpts: map[string]bool{
//			"pubsub": true,
//		},
//	}
//
//	return core.NewNode(ctx, nodeOptions)
//}
//
//func setupPlugins(externalPluginsPath string) error {
//	// Load any external plugins if available on externalPluginsPath
//	plugins, err := loader.NewPluginLoader(filepath.Join(externalPluginsPath, "plugins"))
//	if err != nil {
//		return fmt.Errorf("error loading plugins: %s", err)
//	}
//
//	// Load preloaded and external plugins
//	if err := plugins.Initialize(); err != nil {
//		return fmt.Errorf("error initializing plugins: %s", err)
//	}
//
//	if err := plugins.Inject(); err != nil {
//		return fmt.Errorf("error initializing plugins: %s", err)
//	}
//
//	return nil
//}
//
//func createTempRepo() (string, error) {
//	repoPath, err := os.MkdirTemp("", "ipfs-shell")
//	if err != nil {
//		return "", fmt.Errorf("failed to get temp dir: %s", err)
//	}
//
//	// Create a config with default options and a 2048 bit key
//	cfg, err := config.Init(io.Discard, 2048)
//	if err != nil {
//		return "", err
//	}
//
//	// When creating the repository, you can define custom settings on the repository, such as enabling experimental
//	// features (See experimental-features.md) or customizing the gateway endpoint.
//	// To do such things, you should modify the variable `cfg`. For example:
//	if flagExp {
//		// https://github.com/ipfs/kubo/blob/master/docs/experimental-features.md#ipfs-filestore
//		cfg.Experimental.FilestoreEnabled = true
//		// https://github.com/ipfs/kubo/blob/master/docs/experimental-features.md#ipfs-urlstore
//		cfg.Experimental.UrlstoreEnabled = true
//		// https://github.com/ipfs/kubo/blob/master/docs/experimental-features.md#ipfs-p2p
//		cfg.Experimental.Libp2pStreamMounting = true
//		// https://github.com/ipfs/kubo/blob/master/docs/experimental-features.md#p2p-http-proxy
//		cfg.Experimental.P2pHttpProxy = true
//		// See also: https://github.com/ipfs/kubo/blob/master/docs/config.md
//		// And: https://github.com/ipfs/kubo/blob/master/docs/experimental-features.md
//	}
//
//	// Create the repo with the config
//	err = fsrepo.Init(repoPath, cfg)
//	if err != nil {
//		return "", fmt.Errorf("failed to init ephemeral node: %s", err)
//	}
//
//	return repoPath, nil
//}
//
//var loadPluginsOnce sync.Once
//
//// Spawns a node to be used just for this run (i.e. creates a tmp repo)
//func spawnEphemeral(ctx context.Context) (icore.CoreAPI, *core.IpfsNode, error) {
//	var onceErr error
//	loadPluginsOnce.Do(func() {
//		onceErr = setupPlugins("")
//	})
//	if onceErr != nil {
//		return nil, nil, onceErr
//	}
//
//	// Create a Temporary Repo
//	repoPath, err := createTempRepo()
//	if err != nil {
//		return nil, nil, fmt.Errorf("failed to create temp repo: %s", err)
//	}
//
//	node, err := createNode(ctx, repoPath)
//	if err != nil {
//		return nil, nil, err
//	}
//
//	api, err := coreapi.NewCoreAPI(node)
//
//	return api, node, err
//}
//
//func main() {
//	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Second)
//	defer cancel()
//
//	//var IPFSNode *core.IpfsNode
//	var IPFSCoreAPI ifGoIpfsCore.CoreAPI
//	var err error
//	var orbit iface.OrbitDB
//
//	/*
//	   _, IPFSCoreAPI, err = createNode(ctx, "/home/flowpoint/.ipfs")
//	   if err != nil {
//	       fmt.Printf("%+v",err)
//	       defer cancel()
//	   }
//	*/
//
//	IPFSCoreAPI, _, err = spawnEphemeral(ctx)
//	if err != nil {
//		panic(fmt.Errorf("failed to spawn ephemeral node: %s", err))
//	}
//
//	orbit, err = orbitdb.NewOrbitDB(
//		ctx,
//		IPFSCoreAPI,
//		&orbitdb.NewOrbitDBOptions{})
//
//	if err != nil {
//		fmt.Printf("%+v", err)
//		defer cancel()
//	}
//
//	//var logstoreprovider OrbitDBLogStoreProvider
//	var db orbitdb.EventLogStore
//
//	db, err = orbit.Log(
//		ctx,
//		"log db",
//		nil)
//
//	if err != nil {
//		fmt.Printf("%+v", err)
//		defer cancel()
//	}
//
//	op, err := db.Add(ctx, []byte{1, 2})
//	if err != nil {
//		fmt.Printf("%+v", err)
//		defer cancel()
//	}
//
//	//fmt.Printf("%+v", db.DBName())
//	fmt.Printf("%+v", op.GetEntry().GetHash())
//	//logstoreprovider)
//
//	op, err = db.Get(ctx, op.GetEntry().GetHash())
//	if err != nil {
//		fmt.Printf("%+v", err)
//		defer cancel()
//	}
//	fmt.Printf("%+v", op)
//
//	//orbit.close()
//}
//
//// might need to run
//// sysctl -w net.core.rmem_max=2500000
//// because of ipfs quic stuff
//
//// usefull:
//// https://github.com/ipfs/kubo/blob/master/docs/examples/kubo-as-a-library/main.go
//// https://github.com/berty/go-orbit-db/issues/5
//// https://github.com/berty/go-orbit-db/blob/e9658b4e3af28dd8123b11fad123ba7a27e18bf5/orbitdb.go#L86
