package ndjson

import (
	"bufio"
	"bytes"
	"encoding/json"
	"reflect"
)

// Unmarshal parses the data that is encoded Newline delimited JSON format
// and stores the result in the value pointed to by v.
func Unmarshal(data []byte, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		if rv.Elem().Kind() != reflect.Slice {
			return &InvalidUnmarshalError{reflect.TypeOf(v)}
		}
	} else {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}

	st := reflect.TypeOf(v).Elem()
	vt := st.Elem()
	// FIXME: Specify the suitable length and capacity.
	vs := reflect.MakeSlice(st, 0, 0)
	s := bufio.NewScanner(bytes.NewReader(data))
	for s.Scan() {
		u := reflect.New(vt)
		if err := json.Unmarshal(s.Bytes(), u.Interface()); err != nil {
			return err
		}
		vs = reflect.Append(vs, reflect.Indirect(u))
	}
	if err := s.Err(); err != nil {
		return err
	}
	rv.Elem().Set(vs)
	return nil
}

// InvalidUnmarshalError describes an invalid argument passed to Unmarshal.
// (The argument to Unmarshal must be a non-nil pointer to a slice.)
type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "ndjson: Unmarshal(nil)"
	}
	if e.Type.Kind() == reflect.Ptr {
		if e.Type.Elem().Kind() != reflect.Slice {
			return "ndjson: Unmarshal(non-slice-pointer " + e.Type.Elem().String() + ")"
		}
	}
	return "ndjson: Unmarshal(non-pointer " + e.Type.String() + ")"
}
