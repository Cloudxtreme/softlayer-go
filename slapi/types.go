package slapi

//go:generate go run ../cmd/sl-generate-types/generate.go ../cmd/sl-generate-types/types.go
//go:generate gofmt -s -w types

import (
	"fmt"
	"strconv"
)

// Float64 is a float type that deals with some of the oddities when unmarshalling from the SLAPI
type Float64 float64

// UnmarshalJSON statisied the json.Unmarshaler interface
func (f *Float64) UnmarshalJSON(data []byte) error {

	// Attempt parsing the float normally
	v, err := strconv.ParseFloat(string(data), 64)

	// Attempt parsing the float as a string
	if err != nil {
		if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
			return fmt.Errorf("malformed data")
		}

		v, err = strconv.ParseFloat(string(data[1:len(data)-1]), 64)
		if err != nil {
			return err
		}
	}
	*f = Float64(v)
	return nil
}
