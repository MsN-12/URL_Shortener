package shortener

import (
	"crypto/sha256"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"math/big"
)

func GenerateShortURL(originalURL string) string {
	urlHashBytes := sha256of(originalURL)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58of([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}
func sha256of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}
func base58of(bytes []byte) string {
	return base58.Encode(bytes)
}
