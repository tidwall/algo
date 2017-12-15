package algo

import "unsafe"

//go:linkname memhash64 runtime.memhash64
func memhash64(p unsafe.Pointer, seed uintptr) uintptr

//go:linkname memhash32 runtime.memhash32
func memhash32(p unsafe.Pointer, seed uintptr) uintptr

// Sum64 returns a 64-bit hash for the provided key.
func Sum64(key string) uint64 {
	return uint64(memhash64(unsafe.Pointer(&key), 0))
}

// Sum64Seed returns a 64-bit hash for the provided key.
func Sum64Seed(key string, seed uint64) uint64 {
	return uint64(memhash64(unsafe.Pointer(&key), uintptr(seed)))
}

// Sum64Bytes returns a 64-bit hash for the provided key.
func Sum64Bytes(key []byte) uint64 {
	return uint64(memhash64(unsafe.Pointer(&key), 0))
}

// Sum64BytesSeed returns a 64-bit hash for the provided key.
func Sum64BytesSeed(key []byte, seed uint64) uint64 {
	return uint64(memhash64(unsafe.Pointer(&key), uintptr(seed)))
}

// Sum32 returns a 32-bit hash for the provided key.
func Sum32(key string) uint32 {
	return uint32(memhash32(unsafe.Pointer(&key), 0))
}

// Sum32Seed returns a 32-bit hash for the provided key.
func Sum32Seed(key string, seed uint64) uint32 {
	return uint32(memhash32(unsafe.Pointer(&key), uintptr(seed)))
}

// Sum32Bytes returns a 32-bit hash for the provided key.
func Sum32Bytes(key []byte) uint32 {
	return uint32(memhash32(unsafe.Pointer(&key), 0))
}

// Sum32BytesSeed returns a 32-bit hash for the provided key.
func Sum32BytesSeed(key []byte, seed uint64) uint32 {
	return uint32(memhash32(unsafe.Pointer(&key), uintptr(seed)))
}
