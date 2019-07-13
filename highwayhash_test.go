package highwayhash

import (
	"testing"
)

func TestHighwayhash(t *testing.T) {
	t.Logf("Sum64 a %x", Zero.Sum([]byte("a")))
	t.Logf("Sum64 a %x", Zero.Sum([]byte("a")))
	t.Logf("Sum128 a %x", Rand.Sum([]byte("a")))
	t.Logf("Sum128 a %x", Rand.Sum([]byte("a")))
}
