package otp
import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"fmt"
	"math"
	"time"
)

type Config struct {
	Secret      string
	TimeInterval int64
	Digits      int
	Algorithm   string
}

type Options struct {
	Issuer      string
	AccountName string
}

func New(config Config) *Config {
	return &config
}

func (otp *Config) Generate(options Options) (string, error) {
	key, err := base32.StdEncoding.DecodeString(otp.Secret)
	if err != nil {
		return "", err
	}

	timeCount := time.Now().Unix() / otp.TimeInterval

	timeBytes := make([]byte, 8)
	for i := 0; i < 8; i++ {
		timeBytes[7-i] = byte(timeCount >> (i * 8))
	}

	h := hmac.New(sha1.New, key)
	h.Write(timeBytes)
	hmacResult := h.Sum(nil)

	offset := hmacResult[len(hmacResult)-1] & 0x0F

	code := (int(hmacResult[offset])&0x7F) << 24 |
		(int(hmacResult[offset+1]) & 0xFF) << 16 |
		(int(hmacResult[offset+2]) & 0xFF) << 8 |
		(int(hmacResult[offset+3]) & 0xFF)

	otpValue := code % int(math.Pow(10, float64(otp.Digits)))

	return fmt.Sprintf(fmt.Sprintf("%%0%d.d", otp.Digits), otpValue), nil
}
