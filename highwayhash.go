package highwayhash

import (
	"crypto/rand"

	"github.com/minio/highwayhash"
)

type Key []byte

var Rand Key
var Zero Key

func init() {
	Zero = make([]byte, 32)
	Rand = make([]byte, 32)
	rand.Read(Rand)
}

func (k Key) Sum64(data []byte) uint64 {
	return highwayhash.Sum64(data, k)
}

func (k Key) Sum128(data []byte) [16]byte {
	return highwayhash.Sum128(data, k)
}

func (k Key) Sum(data []byte) [32]byte {
	return highwayhash.Sum(data, k)
}
