package keeper

const (
	// ModuleName defines the module name
	ModuleName = "observer"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_observer"
)

// Keys for store prefixes
// Items are stored with the following key: values
//
// - 0x01<eth_tx_hash (32 bytes)><reporter_addr>: Report
// - 0x02: Params
var (
	reportsKeyPrefix = []byte{0x01}
	paramsKey        = []byte{0x02}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func reportsKey(ethTxHash []byte) []byte {
	return append(reportsKeyPrefix, ethTxHash...)
}

func reportsKeyWithObserver(ethTxHash []byte, reporterAddr string) []byte {
	prefix := reportsKey(ethTxHash)
	return append(prefix, []byte(reporterAddr)...)
}
