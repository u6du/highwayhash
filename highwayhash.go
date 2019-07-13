package highwayhash

import (
	"crypto/rand"
	"hash"

	"github.com/minio/highwayhash"
	"github.com/u6du/ex"
)

type Key []byte

var Rand Key
var Zero Key

func init() {
	Zero = make([]byte, 32)
	Rand = make([]byte, 32)
	rand.Read(Rand)
}

func (k Key) New() (h hash.Hash) {
	h, err := highwayhash.New(k)
	ex.Panic(err)
	return h
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
