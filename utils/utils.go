package utils

import(
	"hash"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
)

func SelectHash(algorithm string) func() hash.Hash {
	switch algorithm {
	case AlgorithmSHA256:
		return sha256.New
	case AlgorithmSHA512:
		return sha512.New
	default:
		return sha1.New
	}
}

func GetHashFunction(algorithm string) (func() hash.Hash, int) {
	switch algorithm {
	case AlgorithmSHA256:
		return sha256.New, 256
	case AlgorithmSHA512:
		return sha512.New, 512
	default:
		return sha256.New, 256
	}
}

