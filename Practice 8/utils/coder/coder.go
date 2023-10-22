package coder

import (
	"crypto/sha512"
	"encoding/hex"
)

func Sha512Hash(input []byte) string {
	hasher := sha512.New()
	hasher.Write(input)
	return hex.EncodeToString(hasher.Sum(nil))
}
