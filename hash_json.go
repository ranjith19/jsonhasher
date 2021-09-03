package jsonhasher

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
)

const (
	sha_1   = iota
	sha_256 = iota
	sha_512 = iota
)

func HashJsonString(jsonString string) (*string, error) {
	return hashJsonString(jsonString, sha_256)
}

func HashJsonStringSha1(jsonString string) (*string, error) {
	return hashJsonString(jsonString, sha_1)
}

func HashJsonStringSha256(jsonString string) (*string, error) {
	return hashJsonString(jsonString, sha_256)
}

func HashJsonStringSha512(jsonString string) (*string, error) {
	return hashJsonString(jsonString, sha_512)
}

func hashJsonString(jsonString string, shaType uint) (*string, error) {
	var v interface{}
	err := json.Unmarshal([]byte(jsonString), &v)
	if err != nil {
		return nil, err
	}
	cdoc, _ := json.Marshal(v)
	switch shaType {
	case sha_1:
		return createSha1(cdoc), nil
	case sha_256:
		return createSha256(cdoc), nil
	case sha_512:
		return createSha512(cdoc), nil
	default:
		return createSha256(cdoc), nil
	}
}

func createSha1(b []byte) *string {
	hasher := sha1.New()
	hasher.Write(b)
	sha := hex.EncodeToString(hasher.Sum(nil))
	return &sha
}

func createSha256(b []byte) *string {
	hasher := sha256.New()
	hasher.Write(b)
	sha := hex.EncodeToString(hasher.Sum(nil))
	return &sha
}

func createSha512(b []byte) *string {
	hasher := sha512.New()
	hasher.Write(b)
	sha := hex.EncodeToString(hasher.Sum(nil))
	return &sha
}
