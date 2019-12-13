package gulden

import (
	"blockbook/bchain/coins/btc"
	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
)

// magic numbers
const (
	MainnetMagic wire.BitcoinNet = 0xe0f7fefc
)

// chain parameters
var (
	MainNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic
	MainNetParams.PubKeyHashAddrID = []byte{38}
	MainNetParams.ScriptHashAddrID = []byte{98}
}

// GuldenParser handle
type GuldenParser struct {
	*btc.BitcoinParser
}

// NewGuldenParser returns new GuldenParser instance
func NewGuldenParser(params *chaincfg.Params, c *btc.Configuration) *GuldenParser {
	return &GuldenParser{BitcoinParser: btc.NewBitcoinParser(params, c)}
}

// GetChainParams contains network parameters for the main Gulden network,
// and the test Gulden network
func GetChainParams(chain string) *chaincfg.Params {
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err != nil {
			panic(err)
		}
	}

	return &MainNetParams
}
