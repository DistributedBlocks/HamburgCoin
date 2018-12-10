package main

import (
	_ "net/http/pprof"

	"github.com/skycoin/skycoin/src/skycoin"
	"github.com/skycoin/skycoin/src/util/logging"
	"github.com/skycoin/skycoin/src/visor"
)

var (
	// Version of the node. Can be set by -ldflags
	Version = "0.24.1"
	// Commit ID. Can be set by -ldflags
	Commit = ""
	// Branch name. Can be set by -ldflags
	Branch = ""
	// ConfigMode (possible values are "", "STANDALONE_CLIENT").
	// This is used to change the default configuration.
	// Can be set by -ldflags
	ConfigMode = ""

	logger = logging.MustGetLogger("main")

	// GenesisSignatureStr hex string of genesis signature
	GenesisSignatureStr = "54c45d88a81dc0da07c0a770035019d7bd63916a65845c7ac86c21e37a10bbe02ca36dbc2348a2db200626e4adbdd808edb80c7087158cee4f7337346a2ef6e400"
	// GenesisAddressStr genesis address string
	GenesisAddressStr = "jqtz1AGJEHpxLjKvCFEzraYsY7oZH3JQnk"
	// BlockchainPubkeyStr pubic key string
	BlockchainPubkeyStr = "02626de86a9e732e7c6a654c612bcf5227106679fbddd4219d09f84c18cac45983"
	// BlockchainSeckeyStr empty private key string
	BlockchainSeckeyStr = ""

	// GenesisTimestamp genesis block create unix time
	GenesisTimestamp uint64 = 1543925416
	// GenesisCoinVolume represents the coin capacity
	GenesisCoinVolume uint64 = 500e12

	// DefaultConnections the default trust node addresses
	DefaultConnections = []string{
	"104.248.129.193:33002",
	"104.248.129.193:20100",
	"104.248.129.193:33001",
	"104.248.131.27:8100",
	"178.128.201.134:8100",
}
)

func main() {
	// get node config
	nodeConfig := skycoin.NewNodeConfig(ConfigMode, skycoin.NodeParameters{
		GenesisSignatureStr: GenesisSignatureStr,
		GenesisAddressStr:   GenesisAddressStr,
		GenesisCoinVolume:   GenesisCoinVolume,
		GenesisTimestamp:    GenesisTimestamp,
		BlockchainPubkeyStr: BlockchainPubkeyStr,
		BlockchainSeckeyStr: BlockchainSeckeyStr,
		DefaultConnections:  DefaultConnections,
		PeerListURL:         "https://hamburgcoin.com/peers.txt",
		Port:                8100,
		WebInterfacePort:    8120,
		DataDirectory:       "$HOME/.hamburgcoin",
		ProfileCPUFile:      "skycoin.prof",
	})

	// create a new fiber coin instance
	coin := skycoin.NewCoin(
		skycoin.Config{
			Node: *nodeConfig,
			Build: visor.BuildInfo{
				Version: Version,
				Commit:  Commit,
				Branch:  Branch,
			},
		},
		logger,
	)

	// parse config values
	coin.ParseConfig()

	// run fiber coin node
	coin.Run()
}
