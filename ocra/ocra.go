package ocra

import (
	"github.com/hasan-kilici/otp/utils"
	"crypto/hmac"
	"encoding/base32"
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Config struct {
	Suite       string
	SecretKey   string
	Counter     string
	Question    string
	Password    string
	SessionInfo string
	TimeStamp   string
	Algorithm   string
}

type Options struct {
	Issuer      string
	AccountName string
}

type OCRA struct {
	Config Config
}

func New(config Config) *OCRA {
	return &OCRA{Config: config}
}

func (o *OCRA) Generate(options Options) (string, error) {
	decodedKey, err := base32.StdEncoding.DecodeString(strings.ToUpper(o.Config.SecretKey))
	if err != nil {
		return "", errors.New("invalid secret key")
	}

	hashFunc, hashSize := utils.GetHashFunction(o.Config.Algorithm)
	if hashFunc == nil {
		return "", errors.New("unsupported hash algorithm")
	}

	message, err := o.buildMessage()
	if err != nil {
		return "", err
	}

	hmacResult := hmac.New(hashFunc, decodedKey)
	hmacResult.Write(message)
	hash := hmacResult.Sum(nil)

	return truncate(hash, hashSize), nil
}

func (o *OCRA) buildMessage() ([]byte, error) {
	var message []byte

	message = append(message, []byte(o.Config.Suite)...)

	if len(o.Config.Counter) > 0 {
		counterVal, err := strconv.ParseInt(o.Config.Counter, 10, 64)
		if err != nil {
			return nil, errors.New("invalid counter")
		}
		counterBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(counterBytes, uint64(counterVal))
		message = append(message, counterBytes...)
	}

	message = append(message, []byte(o.Config.Question)...)

	if len(o.Config.Password) > 0 {
		message = append(message, []byte(o.Config.Password)...)
	}

	if len(o.Config.SessionInfo) > 0 {
		message = append(message, []byte(o.Config.SessionInfo)...)
	}

	if len(o.Config.TimeStamp) > 0 {
		message = append(message, []byte(o.Config.TimeStamp)...)
	}

	return message, nil
}

func truncate(hash []byte, hashSize int) string {
	offset := hash[len(hash)-1] & 0xF
	binaryCode := (int(hash[offset]&0x7F) << 24) |
		(int(hash[offset+1]&0xFF) << 16) |
		(int(hash[offset+2]&0xFF) << 8) |
		(int(hash[offset+3]) & 0xFF)

	otp := binaryCode % 1000000
	return fmt.Sprintf("%06d", otp)
}
