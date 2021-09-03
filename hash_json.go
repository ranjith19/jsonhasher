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
	return hashRawJson([]byte(jsonString))
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
	case "bool":
		return hashJsonBool(jsonString)
	case "string":
		return hashJsonString(jsonString)
	case "float":
		return hashJsonFloat(jsonString)
	default:
		return hashNil()
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
		khash := hashString(k)
		final := fmt.Sprintf("%s:::%s", *khash, *vhash)
		hash := hashString(final)
		if err != nil {
			return nil, err
		}
		hashes = append(hashes, *hash)

	}
	combined := strings.Join(hashes, "|||")
	return hashString(combined), nil
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
	return hashString(combined), nil
}

func hashJsonBool(jsonString []byte) (*string, error) {
	var v bool
	err := json.Unmarshal(jsonString, &v)
	if err != nil {
		return nil, err
	}
	if v {
		return hashString("$$$TRUE$$$"), nil
	}
	return hashString("$$$FALSE$$$"), nil
}

func hashJsonFloat(jsonString []byte) (*string, error) {
	var v float64
	err := json.Unmarshal(jsonString, &v)
	if err != nil {
		return nil, err
	}
	return hashString(fmt.Sprintf("%f", v)), nil
}

func hashNil() (*string, error) {
	return hashString("$$$NULL$$$"), nil
}

func hashJsonString(jsonString []byte) (*string, error) {
	var v string
	err := json.Unmarshal(jsonString, &v)
	if err != nil {
		return nil, err
	}
	return hashString(v), nil
}

func hashString(s string) *string {
	hasher := sha256.New()
	hasher.Write([]byte(s))
	sha := hex.EncodeToString(hasher.Sum(nil))
	return &sha
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
