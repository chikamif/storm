package gob

import (
	"testing"

	"github.com/chikamif/storm/codec/internal"
)

func TestGob(t *testing.T) {
	internal.RoundtripTester(t, Codec)
}
