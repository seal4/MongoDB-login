package main

import (
	"crypto/sha256"
	"encoding/hex"
)

func hash(password string) string {
	hash := sha256.Sum256([]byte(password))
	output := hex.EncodeToString(hash[:])
	return output
}
