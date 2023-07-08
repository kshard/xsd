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
)

type Value interface{ XSDType() XSDType }

type XSDType int

// XSD Types codes
const (
	XSD_NIL XSDType = iota
	XSD_SYMBOL
	XSD_ANYURI
	XSD_STRING
)

// The data type is symbol
type Symbol uint32

func (v Symbol) XSDType() XSDType { return XSD_SYMBOL }

// The data type represents Internationalized Resource Identifier.
// Used to uniquely identify concept, objects, etc.
type AnyURI Symbol

func (v AnyURI) XSDType() XSDType { return XSD_ANYURI }

// The string data-type represents character strings in knowledge statements.
// The language strings are annotated with corresponding language tag.
type String string

func (v String) XSDType() XSDType { return XSD_STRING }
func (v String) String() string   { return strconv.Quote(string(v)) }

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
