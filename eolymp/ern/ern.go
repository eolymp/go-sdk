package ern

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

type Name []string

func Parse(ern string) Name {
	if ern == "" {
		return nil
	}

	return strings.Split(ern, ":")
}

func Valid(ern string) bool {
	return Parse(ern).Valid()
}

func Parent(ern string) Name {
	return Parse(ern).Parent()
}

func (e Name) String() string {
	return strings.Join(e, ":")
}

func (e Name) Valid() bool {
	return e[0] == "ern" && len(e)%2 == 1
}

func (e Name) Parent() Name {
	if !e.Valid() || len(e) <= 3 {
		return nil
	}

	return e[0 : len(e)-2]
}

func (e Name) Match(m Name) bool {
	return e.String() == m.String()
}

func (e Name) Includes(parent Name) bool {
	if len(e) <= len(parent) {
		return false
	}

	return e[0:len(parent)].String() == parent.String()
}

func (e Name) MarshalJSON() ([]byte, error) {
	if e == nil {
		return json.Marshal(nil)
	}

	return json.Marshal(e.String())
}

func (e *Name) UnmarshalJSON(data []byte) error {
	var str *string

	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	if str == nil {
		*e = nil
	} else {
		*e = Parse(*str)
	}

	return nil
}

func (e *Name) Scan(src interface{}) error {
	switch v := src.(type) {
	case nil:
		*e = nil
	case []uint8:
		*e = Parse(string(v))
	case string:
		*e = Parse(v)
	default:
		return fmt.Errorf("unexpected type for registration form: %T", src)
	}

	return nil
}

func (e Name) Value() (driver.Value, error) {
	if e == nil {
		return nil, nil
	}

	if !e.Valid() {
		return nil, MalformedErr{}
	}

	return e.String(), nil
}

func (e Name) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

func (e *Name) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, e)
}
