//
// Copyright (C) 2023 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/xsd
//

package xsd

import (
	"strconv"

	"github.com/fogfish/curie"
)

type Value interface{ XSDType() Symbol }

// The data type represents Internationalized Resource Identifier.
// Used to uniquely identify concept, objects, etc.
type AnyURI Symbol

func (v AnyURI) XSDType() Symbol { return XSD_ANYURI }
func (v AnyURI) String() string  { return Symbol(v).String() }

func (v AnyURI) MarshalBinary() ([]byte, error) {
	s := Symbol(v).String()
	return []byte(s), nil
}

func (v *AnyURI) UnmarshalBinary(data []byte) error {
	s := ToSymbol(string(data))
	*v = AnyURI(s)
	return nil
}

func ToAnyURI(iri curie.IRI) AnyURI { return AnyURI(ToSymbol(string(iri))) }

// The string data-type represents character strings in knowledge statements.
// The language strings are annotated with corresponding language tag.
type String string

func (v String) XSDType() Symbol { return XSD_STRING }
func (v String) String() string  { return strconv.Quote(string(v)) }

// The Integer data-type in knowledge statement.
// The library uses various int precision data-types to represent decimal values.
// type XSDInteger = int
// type Byte = int8
// type Short = int16
// type Int = int32
// type Long = int64
// type NonNegativeInteger = uint
// type UnsignedByte = uint8
// type UnsignedShort = uint16
// type UnsignedInt = uint32
// type UnsignedLong = uint64

// const XSD_INTEGER = curie.IRI("xsd:integer")

// type Integer struct{ Value XSDInteger }

// func (Integer) XSDType() curie.IRI { return XSD_INTEGER }
