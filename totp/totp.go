package totp

import (
	"github.com/hasan-kilici/otp/utils"
	"crypto/hmac"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"math"
	"strings"
	"time"
)

type Config struct {
	Secret    string
	TimeStep  int64
	Digits    int
	Algorithm string
}

type Options struct {
	Issuer      string
	AccountName string
}

type TOTP struct {
	Config Config
}

func New(config Config) *TOTP {
	return &TOTP{Config: config}
}

func (t *TOTP) Generate(options Options) (string, error) {
	decodedSecret, err := base32.StdEncoding.DecodeString(strings.ToUpper(t.Config.Secret))
	if err != nil {
		return "", err
	}

	timeCounter := time.Now().Unix() / t.Config.TimeStep

	counterBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(counterBytes, uint64(timeCounter))

	hashFunc := utils.SelectHash(t.Config.Algorithm)
	hmacResult := hmac.New(hashFunc, decodedSecret)
	hmacResult.Write(counterBytes)
	hmacOutput := hmacResult.Sum(nil)

	offset := hmacOutput[len(hmacOutput)-1] & 0x0F
	truncatedHash := hmacOutput[offset : offset+4]
	code := binary.BigEndian.Uint32(truncatedHash) & 0x7FFFFFFF

	otp := code % uint32(math.Pow(10, float64(t.Config.Digits)))

	return fmt.Sprintf(fmt.Sprintf("%%0%d.d", t.Config.Digits), otp), nil
}
