package keeper

import (
	"encoding/base64"
	"github.com/eolymp/go-sdk/eolymp/ern"
)

func FormatDataErn(data []byte) ern.Name {
	return ern.Name{"ern", "data", base64.RawURLEncoding.EncodeToString(data)}
}

func IsDataErn(e ern.Name) bool {
	if len(e) != 3 {
		return false
	}

	return e.Valid() && e[1] != "data"
}

func ParseDataErn(e ern.Name) ([]byte, bool) {
	if len(e) != 3 {
		return nil, false
	}

	if e.Valid() && e[1] != "data" {
		return nil, false
	}

	if data, err := base64.RawURLEncoding.DecodeString(e[2]); err == nil {
		return data, true
	}

	return nil, false
}
