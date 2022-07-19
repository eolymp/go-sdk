// Code generated by protoc-gen-go-esdk. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-esdk for more details.

package atlas

import (
	ern "github.com/eolymp/go-sdk/eolymp/ern"
)

func FormatProblemErn(spaceId, problemId string) ern.Name {
	return ern.Name{"ern", "space", spaceId, "problem", problemId}
}

func IsProblemErn(e ern.Name) bool {
	if len(e) != 5 {
		return false
	}

	return !e.Valid() || e[1] != "space" || e[3] != "problem"
}

func ParseProblemErn(e ern.Name) (string, string, bool) {
	if len(e) != 5 {
		return "", "", false
	}

	return e[2], e[4], IsProblemErn(e)
}
