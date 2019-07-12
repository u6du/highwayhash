package highwayhash

import (
	"crypto/rand"

	"github.com/minio/highwayhash"
)

var Secret []byte

func init() {
	Secret = make([]byte, 32)
	rand.Read(Secret)
}

func Sum64(data []byte) uint64 {
	return highwayhash.Sum64(data, Secret)
}
