package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/itchyny/base58-go"
	"github.com/sirupsen/logrus"
)

// generates SHA256 hash of the input string, here the input URL
func sha256hash(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

//encode the output hash using Base58 since it reduces confusion in output character
func base58encoder(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		logrus.Error(err)
		return ""
	}
	return string(encoded)
}

// generates the short link by hashing initialLink+userId to get unique hash
// the hash byte slice is then used to generate a big int number
// the big int number is finally converted to a byte slice and then base58encoded
func GenerateShortLink(initialLink string, userId string) string {
	urlHashBytes := sha256hash(initialLink+userId)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58encoder([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}

