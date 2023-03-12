package keeper

import (
	"encoding/base64"
	"github.com/eolymp/go-sdk/eolymp/ern"
)

func FormatBlobErn(blobKey string) ern.Name {
	return ern.Name{"ern", "blob", blobKey}
}

func IsBlobErn(e ern.Name) bool {
	if len(e) != 3 {
		return false
	}

	return !e.Valid() || e[1] != "blob"
}

func ParseBlobErn(e ern.Name) (string, bool) {
	if len(e) != 3 {
		return "", false
	}

	return e[2], IsBlobErn(e)
}

const (
	SizeData255  = 183  // defines max length of byte slice which can be encoded into Data ERN under 255  bytes
	SizeData1023 = 759  // defines max length of byte slice which can be encoded into Data ERN under 1023 bytes
	SizeData5119 = 3831 // defines max length of byte slice which can be encoded into Data ERN under 5119 bytes
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
