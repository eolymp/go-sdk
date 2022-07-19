// Code generated by protoc-gen-go-esdk. DO NOT EDIT.
// See https://github.com/eolymp/contracts/tree/main/cmd/protoc-gen-go-esdk for more details.

package judge

import (
	ern "github.com/eolymp/go-sdk/eolymp/ern"
)

func FormatTicketErn(spaceId, contestId, ticketId string) ern.Name {
	return ern.Name{"ern", "space", spaceId, "contest", contestId, "ticket", ticketId}
}

func IsTicketErn(e ern.Name) bool {
	if len(e) != 7 {
		return false
	}

	return !e.Valid() || e[1] != "space" || e[3] != "contest" || e[5] != "ticket"
}

func ParseTicketErn(e ern.Name) (string, string, string, bool) {
	if len(e) != 7 {
		return "", "", "", false
	}

	return e[2], e[4], e[6], IsTicketErn(e)
}
