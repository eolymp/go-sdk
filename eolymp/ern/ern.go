package ern

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

type ERN []string

func Parse(ern string) ERN {
	return strings.Split(ern, ":")
}

func Valid(ern string) bool {
	return Parse(ern).Valid()
}

func Parent(ern string) ERN {
	return Parse(ern).Parent()
}

func (e ERN) String() string {
	return strings.Join(e, ":")
}

func (e ERN) Valid() bool {
	return e[0] == "ern" && len(e)%2 == 1
}

func (e ERN) Parent() ERN {
	if !e.Valid() || len(e) <= 3 {
		return nil
	}

	return e[0 : len(e)-2]
}

func (e ERN) Match(m ERN) bool {
	return e.String() == m.String()
}

func (e ERN) Includes(parent ERN) bool {
	if len(e) <= len(parent) {
		return false
	}

	return e[0:len(parent)].String() == parent.String()
}

func (e ERN) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

func (e *ERN) UnmarshalJSON(data []byte) error {
	var str string

	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	*e = Parse(str)

	return nil
}

func (e *ERN) Scan(src interface{}) error {
	switch v := src.(type) {
	case nil:
	case []uint8:
		*e = Parse(string(v))
	case string:
		*e = Parse(v)
	default:
		return fmt.Errorf("unexpected type for registration form: %T", src)
	}

	return nil
}

func (e ERN) Value() (driver.Value, error) {
	return e.String(), nil
}

func (e ERN) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

func (e *ERN) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, e)
}
