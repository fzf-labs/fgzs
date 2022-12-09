package json

import (
	"bytes"
	encodingJson "encoding/json"
	JsonIterator "github.com/json-iterator/go"
	"io"
)

var json = JsonIterator.ConfigCompatibleWithStandardLibrary

// Package json provides json operations wrapping ignoring stdlib or third-party lib json.

// Marshal adapts to json/encoding Marshal API.
//
// Marshal returns the JSON encoding of v, adapts to json/encoding Marshal API
// Refer to https://godoc.org/encoding/json#Marshal for more information.
func Marshal(v interface{}) ([]byte, error) {

	return json.Marshal(v)
}

// MarshalIndent same as json.MarshalIndent. Prefix is not supported.
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

// Unmarshal adapts to json/encoding Unmarshal API
//
// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
// Refer to https://godoc.org/encoding/json#Unmarshal for more information.
func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// UnmarshalUseNumber decodes the json data bytes to target interface using number option.
func UnmarshalUseNumber(data []byte, v interface{}) error {
	decoder := NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	return decoder.Decode(v)
}

// NewEncoder same as json.NewEncoder
func NewEncoder(writer io.Writer) *encodingJson.Encoder {
	return encodingJson.NewEncoder(writer)
}

// NewDecoder adapts to json/stream NewDecoder API.
//
// NewDecoder returns a new decoder that reads from r.
//
// Instead of a json/encoding Decoder, an Decoder is returned
// Refer to https://godoc.org/encoding/json#NewDecoder for more information.
func NewDecoder(reader io.Reader) *encodingJson.Decoder {
	return encodingJson.NewDecoder(reader)
}

// Valid reports whether data is a valid JSON encoding.
func Valid(data []byte) bool {
	return encodingJson.Valid(data)
}
