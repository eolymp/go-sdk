// Code generated by protoc-gen-go-esdk. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-esdk for more details.

package universe

import (
	ern "github.com/eolymp/go-sdk/eolymp/ern"
)

func FormatSpaceErn(spaceId string) ern.Name {
	return ern.Name{"ern", "space", spaceId}
}

func IsSpaceErn(e ern.Name) bool {
	if len(e) != 3 {
		return false
	}

	return !e.Valid() || e[1] != "space"
}

func ParseSpaceErn(e ern.Name) (string, bool) {
	if len(e) != 3 {
		return "", false
	}

	return e[2], IsSpaceErn(e)
}
