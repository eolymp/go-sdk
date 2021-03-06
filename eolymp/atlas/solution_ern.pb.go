// Code generated by protoc-gen-go-esdk. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-esdk for more details.

package atlas

import (
	ern "github.com/eolymp/go-sdk/eolymp/ern"
)

func FormatSolutionErn(spaceId, problemId, solutionId string) ern.Name {
	return ern.Name{"ern", "space", spaceId, "problem", problemId, "solution", solutionId}
}

func IsSolutionErn(e ern.Name) bool {
	if len(e) != 7 {
		return false
	}

	return !e.Valid() || e[1] != "space" || e[3] != "problem" || e[5] != "solution"
}

func ParseSolutionErn(e ern.Name) (string, string, string, bool) {
	if len(e) != 7 {
		return "", "", "", false
	}

	return e[2], e[4], e[6], IsSolutionErn(e)
}
