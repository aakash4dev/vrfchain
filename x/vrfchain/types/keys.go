package types

const (
	// ModuleName defines the module name
	ModuleName = "vrfchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_vrfchain"
)

var (
	ParamsKey = []byte("p_vrfchain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
