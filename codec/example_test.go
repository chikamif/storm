package codec_test

import (
	"fmt"

	"github.com/chikamif/storm"
	"github.com/chikamif/storm/codec/gob"
	"github.com/chikamif/storm/codec/json"
)

func Example() {
	// The examples below show how to set up all the codecs shipped with Storm.
	// Proper error handling left out to make it simple.
	var gobDb, _ = storm.Open("gob.db", storm.Codec(gob.Codec))
	var jsonDb, _ = storm.Open("json.db", storm.Codec(json.Codec))
	fmt.Printf("%T\n", gobDb.Codec())
	fmt.Printf("%T\n", jsonDb.Codec())

	// Output:
	// *gob.gobCodec
	// *json.jsonCodec
}
