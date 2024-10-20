package hotp

import (
	"github.com/hasan-kilici/otp/utils"
	"crypto/hmac"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"math"
	"strings"
)

type Config struct {
	Secret    string
	Counter   uint64
	Digits    int
	Algorithm string
}

type Options struct {
	Issuer      string
	AccountName string
}

type HOTP struct {
	Config Config
}

func New(config Config) *HOTP {
	return &HOTP{Config: config}
}

func (h *HOTP) Generate(options Options) (string, error) {
	decodedSecret, err := base32.StdEncoding.DecodeString(strings.ToUpper(h.Config.Secret))
	if err != nil {
		return "", err
	}

	counterBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(counterBytes, h.Config.Counter)

	hashFunc := utils.SelectHash(h.Config.Algorithm)
	hmacResult := hmac.New(hashFunc, decodedSecret)
	hmacResult.Write(counterBytes)
	hmacOutput := hmacResult.Sum(nil)

	offset := hmacOutput[len(hmacOutput)-1] & 0x0F
	truncatedHash := hmacOutput[offset : offset+4]

	code := binary.BigEndian.Uint32(truncatedHash) & 0x7FFFFFFF
	otp := code % uint32(math.Pow(10, float64(h.Config.Digits)))

	return fmt.Sprintf(fmt.Sprintf("%%0%dd", h.Config.Digits), otp), nil
}
