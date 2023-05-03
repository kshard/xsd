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
	"sync"
)

type Symbol uint32

var (
	m        = &sync.Mutex{}
	symbolID = xsd_USER_DEFINED
	bySymbol = make(map[Symbol]string)
	byString = make(map[string]Symbol)
)

// Symbol types
const (
	XSD_NIL Symbol = iota
	XSD_ANYURI
	XSD_STRING
	xsd_USER_DEFINED
)

func ToSymbol(name string) Symbol {
	m.Lock()
	symb, has := byString[name]
	if !has {
		symb = symbolID
		bySymbol[symb] = name
		byString[name] = symb
		symbolID++
	}
	m.Unlock()

	return symb
}

func (s Symbol) String() string {
	m.Lock()
	name, has := bySymbol[s]
	m.Unlock()

	if has {
		return string(name)
	}

	return fmt.Sprintf(":%d", s)
}
