package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	var src = []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	var srcData = make([]byte, hex.DecodedLen(len(src)))

	n, err := hex.Decode(srcData, src)
	if err != nil {
		panic("Couldn't decode hex string")
	}

	var dstData = make([]byte, base64.StdEncoding.EncodedLen(len(srcData)))

	base64.StdEncoding.Encode(dstData, srcData[:n])
	fmt.Println(string(dstData))
}
