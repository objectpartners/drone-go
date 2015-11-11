package drone

import (
	"encoding/json"
	"strconv"
)

// StringSlice representes a string or an array of strings.
type StringSlice struct {
	parts []string
}

func (e *StringSlice) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}

	p := make([]string, 0, 1)
	if err := json.Unmarshal(b, &p); err != nil {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		p = append(p, s)
	}

	e.parts = p
	return nil
}

func (e *StringSlice) Slice() []string {
	if e == nil {
		return nil
	}
	return e.parts
}

// StringInt representes a string or an integer value.
type StringInt struct {
	value string
}

func (e *StringInt) UnmarshalJSON(b []byte) error {
	var num int
	err := json.Unmarshal(b, &num)
	if err == nil {
		e.value = strconv.Itoa(num)
		return nil
	}
	return json.Unmarshal(b, &e.value)
}

func (e StringInt) String() string {
	return e.value
}
