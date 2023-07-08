//
// Copyright (C) 2023 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/xsd
//

package xsd

import (
	"fmt"

	"github.com/fogfish/curie"
)

// DataType is a type constrain used by the library.
// See https://www.w3.org/TR/xmlschema-2/#datatype
//
// Knowledge statements contain scalar objects -- literals. Literals are either
// language-tagged string `rdf:langString` or type-safe values containing a
// reference to data-type (e.g. `xsd:string`).
//
// This interface defines data-types supported by the library. It maps well-known
// semantic types to Golang native types and relation to existed schema(s) and
// ontologies.
type DataType interface {
	~string |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~bool |
		~[]byte
}

// The floating point data-type in knowledge statement.
// The library uses various uint precisions.
type Float = float32
type Double = float64

// The boolean data-type in knowledge statement
type Boolean = bool

type HexBinary = []byte
type Base64Binary = []byte

// From builds xsd.Value from Golang type(s)
func From[T DataType](
	symbols interface {
		SymbolOf(s string) (Symbol, error)
	},
	value T,
) (Value, error) {
	switch v := any(value).(type) {
	case curie.IRI:
		if symbols == nil {
			return nil, fmt.Errorf("undefined symbols table")
		}

		s, err := symbols.SymbolOf(string(v))
		if err != nil {
			return nil, err
		}
		return AnyURI(s), nil
	case AnyURI:
		return v, nil
	case string:
		return String(v), nil
	case String:
		return v, nil
	default:
		return nil, fmt.Errorf("package xsd does not support %T", value)
	}
}
