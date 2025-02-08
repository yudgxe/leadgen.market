package hasher

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
)

var ErrUnsupportedAlg error = errors.New("unsupported hash alg")

type HashAlg string

const (
	HashAlgMD5    HashAlg = "MD5"
	HashAlgSHA256 HashAlg = "SHA256"
)

func GetHash(alg HashAlg, text string) (string, error) {
	hasher, err := getHashFuncBy(alg)
	if err != nil {
		return "", err
	}
	return hasher(text)
}

func IsValidAlg(alg HashAlg) error {
	switch strings.ToUpper(string(alg)) {
	case string(HashAlgMD5):
		return nil
	case string(HashAlgSHA256):
		return nil
	}

	return ErrUnsupportedAlg
}

func getMD5Hash(in string) (string, error) {
	hash := md5.Sum([]byte(in))

	return hex.EncodeToString(hash[:]), nil
}

func getSHA256Hash(in string) (string, error) {
	hasher := sha256.New()
	if _, err := hasher.Write([]byte(in)); err != nil {
		return "", nil
	}

	hash := hasher.Sum(nil)

	return hex.EncodeToString(hash), nil
}

func getHashFuncBy(alg HashAlg) (func(string) (string, error), error) {
	switch HashAlg(strings.ToUpper(string(alg))) {
	case HashAlgMD5:
		return getMD5Hash, nil
	case HashAlgSHA256:
		return getSHA256Hash, nil
	}

	return nil, ErrUnsupportedAlg
}
