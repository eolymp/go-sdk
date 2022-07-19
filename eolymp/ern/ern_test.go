package ern_test

import (
	"encoding/json"
	"github.com/eolymp/go-sdk/eolymp/ern"
	"testing"
)

func TestERN_Includes(t *testing.T) {
	tt := []struct {
		ERN      string
		Parent   string
		Includes bool
	}{
		{ERN: "ern:parent:id:child:key", Parent: "ern:parent:id", Includes: true},
		{ERN: "ern:parent:id:child:key:attr:code", Parent: "ern:parent:id", Includes: true},
		{ERN: "ern:parent:id", Parent: "ern:parent:id:child:id"},
		{ERN: "ern:parent:id", Parent: "ern:parent:id"},
		{ERN: "ern:parent:id", Parent: "ern:child:id"},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			e := ern.Parse(tc.ERN)
			p := ern.Parse(tc.Parent)

			if want, got := tc.Includes, e.Includes(p); want != got {
				t.Errorf("Includes does not match: want %v, bot %v", want, got)
			}
		})
	}
}

func TestERN_MarshalJSON(t *testing.T) {
	tt := []struct {
		ERN  string
		JSON string
	}{
		{ERN: "ern:parent:id:child:key", JSON: "\"ern:parent:id:child:key\""},
		{ERN: "ern:parent:id", JSON: "\"ern:parent:id\""},
	}

	for _, tc := range tt {
		t.Run(tc.ERN, func(t *testing.T) {
			e := ern.Parse(tc.ERN)

			d, err := json.Marshal(e)
			if err != nil {
				t.Fatal("An error while marshalling: ", err)
			}

			if want, got := tc.JSON, string(d); want != got {
				t.Errorf("JSON marshal does not match: want %v, bot %v", want, got)
			}
		})
	}
}

func TestERN_UnmarshalJSON(t *testing.T) {
	tt := []struct {
		ERN  string
		JSON string
	}{
		{ERN: "ern:parent:id:child:key", JSON: "\"ern:parent:id:child:key\""},
		{ERN: "ern:parent:id:child:key", JSON: "\"ern:parent:id:child:key\""},
		{ERN: "ern:parent:id", JSON: "\"ern:parent:id\""},
	}

	for _, tc := range tt {
		t.Run(tc.ERN, func(t *testing.T) {
			var e ern.Name

			if err := json.Unmarshal([]byte(tc.JSON), &e); err != nil {
				t.Fatal("An error while marshalling: ", err)
			}

			if want, got := tc.ERN, e.String(); want != got {
				t.Errorf("JSON unmarshal does not match: want %v, bot %v", want, got)
			}
		})
	}
}

func TestERN_MarshalBinary(t *testing.T) {
	tt := []struct {
		ERN    string
		Binary string
	}{
		{ERN: "ern:parent:id:child:key", Binary: "\"ern:parent:id:child:key\""},
		{ERN: "ern:parent:id", Binary: "\"ern:parent:id\""},
	}

	for _, tc := range tt {
		t.Run(tc.ERN, func(t *testing.T) {
			e := ern.Parse(tc.ERN)

			d, err := e.MarshalBinary()
			if err != nil {
				t.Fatal("An error while marshalling: ", err)
			}

			if want, got := tc.Binary, string(d); want != got {
				t.Errorf("Binary marshal does not match: want %v, bot %v", want, got)
			}
		})
	}
}

func TestERN_UnmarshalBinary(t *testing.T) {
	tt := []struct {
		ERN    string
		Binary string
	}{
		{ERN: "ern:parent:id:child:key", Binary: "\"ern:parent:id:child:key\""},
		{ERN: "ern:parent:id:child:key", Binary: "\"ern:parent:id:child:key\""},
		{ERN: "ern:parent:id", Binary: "\"ern:parent:id\""},
	}

	for _, tc := range tt {
		t.Run(tc.ERN, func(t *testing.T) {
			var e ern.Name

			if err := e.UnmarshalBinary([]byte(tc.Binary)); err != nil {
				t.Fatal("An error while marshalling: ", err)
			}

			if want, got := tc.ERN, e.String(); want != got {
				t.Errorf("Binary unmarshal does not match: want %v, bot %v", want, got)
			}
		})
	}
}
