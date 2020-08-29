package ndjson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

// Marshal returns the Newline delimited JSON encoding of v.
func Marshal(v interface{}) ([]byte, error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Slice {
		return nil, &InvalidMarshalError{rv.Kind()}
	}

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	for i := 0; i < rv.Len(); i++ {
		if err := encoder.Encode(rv.Index(i).Interface()); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

// InvalidMarshalError describes an invalid argument passed to Marshal.
// (The argument to Marshal must be a slice.)
type InvalidMarshalError struct {
	Kind reflect.Kind
}

func (e *InvalidMarshalError) Error() string {
	return fmt.Sprintf("ndjson: Marshal(non-slice %s)", e.Kind.String())
}
