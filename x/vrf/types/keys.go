package types

const (
	// ModuleName defines the module name
	ModuleName = "vrf"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_vrf"
)

var (
	ParamsKey = []byte("p_vrf")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
