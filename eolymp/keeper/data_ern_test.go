package keeper_test

import (
	"fmt"
	"github.com/eolymp/go-sdk/eolymp/keeper"
	"testing"
)

func TestFormatDataErn(t *testing.T) {
	tt := []struct {
		Data []byte
		ERN  string
	}{
		{Data: []byte("hello/world.txt"), ERN: "ern:data:aGVsbG8vd29ybGQudHh0"},
		{Data: []byte("00/0000a3914ea7"), ERN: "ern:data:MDAvMDAwMGEzOTE0ZWE3"},
	}

	for idx, tc := range tt {
		t.Run(fmt.Sprint(idx), func(t *testing.T) {
			if want, got := tc.ERN, keeper.FormatDataErn(tc.Data).String(); want != got {
				t.Errorf("Formatted ERN does not match expected value: want %v, got %v", want, got)
			}
		})
	}
}
