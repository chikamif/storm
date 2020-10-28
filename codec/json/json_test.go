package json

import (
	"testing"

	"github.com/chikamif/storm/codec/internal"
)

func TestJSON(t *testing.T) {
	internal.RoundtripTester(t, Codec)
}
