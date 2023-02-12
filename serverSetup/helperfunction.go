package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"strings"
)

func provideRandomAlphabetCharacter() string {
	//in ASCII : 65:A, 90:Z
	var end int = 90 + 1
	var start int = 65
	var result big.Int
	tempBigInt, _ := rand.Int(rand.Reader, big.NewInt(int64(end-start)))
	moduloResult := result.Mod(tempBigInt, big.NewInt(int64(end-start)))
	charCode := int64(start) + moduloResult.Int64()
	return string(rune(charCode)) // to suppress warning https://stackoverflow.com/questions/67733156/convert-int-to-ascii-more-than-one-character-in-rune-literal
}

func provideRandomNumericalCharacter() string {
	//in ASCII : 48:0, 57:9
	var end int = 57 + 1
	var start int = 48
	var result big.Int
	tempBigInt, _ := rand.Int(rand.Reader, big.NewInt(int64(end-start)))
	moduloResult := result.Mod(tempBigInt, big.NewInt(int64(end-start)))
	charCode := int64(start) + moduloResult.Int64()
	return string(rune(charCode)) // to suppress warning https://stackoverflow.com/questions/67733156/convert-int-to-ascii-more-than-one-character-in-rune-literal
}

func createRandomNumericalString(length int) string {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteString(provideRandomNumericalCharacter())
	}
	return sb.String()
}

func createRandomString(length int) string {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteString(provideRandomAlphabetCharacter())
	}
	return sb.String()
}

func createHashString(message string) string {
	var dat []byte = []byte(message)
	h := sha256.New()
	h.Write(dat)
	byteS := h.Sum(nil)
	return hex.EncodeToString(byteS)
}
