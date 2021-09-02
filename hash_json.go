package jsonhasher

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func HashJsonString(jsonString string) (*string, error) {
	h, err := hashRawJson([]byte(jsonString))
	if err != nil {
		return nil, err
	}
	return h, nil
}

func hashRawJson(jsonString []byte) (*string, error) {
	jtype, err := determineType(jsonString)
	if err != nil {
		return nil, err
	}
	switch *jtype {
	case "dict":
		return hashJsonDict(jsonString)
	case "list":
		return hashJsonList(jsonString)
	default:
		return hashBaseVal(jsonString)
	}
}

func hashJsonDict(jsonString []byte) (*string, error) {
	kv, err := getAsKeyValPair(jsonString)
	if err != nil {
		return nil, err
	}
	keys := lisSortedtKeys(&kv)
	hashes := make([]string, len(keys))
	for _, k := range keys {
		v := kv[k]
		b, err := v.MarshalJSON()
		if err != nil {
			return nil, err
		}
		vhash, err := hashRawJson(b)
		if err != nil {
			return nil, err
		}
		khash, err := hashBaseVal([]byte(k))
		if err != nil {
			return nil, err
		}
		final := fmt.Sprintf("%s:::%s", *khash, *vhash)
		hash, err := hashBaseVal([]byte(final))
		if err != nil {
			return nil, err
		}
		hashes = append(hashes, *hash)

	}
	combined := strings.Join(hashes, "|||")
	return hashBaseVal([]byte(combined))
}

func hashJsonList(jsonString []byte) (*string, error) {
	parsed := make([]json.RawMessage, 0)
	e := json.Unmarshal(jsonString, &parsed)
	if e != nil {
		return nil, e
	}
	hashes := make([]string, len(parsed))
	for i, v := range parsed {
		b, err := v.MarshalJSON()
		if err != nil {
			return nil, err
		}
		h, err := hashRawJson(b)
		if err != nil {
			return nil, err
		}
		hashes[i] = *h
	}
	combined := strings.Join(hashes, "|||")
	return hashBaseVal([]byte(combined))
}

func hashBaseVal(jsonString []byte) (*string, error) {
	hasher := sha256.New()
	hasher.Write([]byte(jsonString))
	sha := hex.EncodeToString(hasher.Sum(nil))
	return &sha, nil
}

func getAsKeyValPair(jsonString []byte) (map[string]json.RawMessage, error) {
	c := make(map[string]json.RawMessage)

	// unmarschal JSON
	e := json.Unmarshal(jsonString, &c)

	// panic on error
	if e != nil {
		return nil, e
	}

	return c, nil
}

func lisSortedtKeys(c *map[string]json.RawMessage) []string {

	// a string slice to hold the keys
	k := make([]string, 0)

	// copy c's keys into k
	for s, _ := range *c {
		k = append(k, s)
	}
	sort.Strings(k)
	return k
}

func determineType(jsonString []byte) (*string, error) {
	var v interface{}
	err := json.Unmarshal(jsonString, &v)
	if err != nil {
		return nil, err
	}
	determinedType := "nil"
	switch v.(type) {
	case map[string]interface{}:
		determinedType = "dict"
	case []interface{}:
		determinedType = "list"
	case bool:
		determinedType = "bool"
	case float64:
		determinedType = "float"
	case string:
		determinedType = "string"
	default:
		determinedType = "nil"
	}
	return &determinedType, err
}
