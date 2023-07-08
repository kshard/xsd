//
// Copyright (C) 2023 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/xsd
//

package symbol

import (
	"fmt"
	"sync"

	"github.com/kshard/xsd"
)

// Abstract notion of xsd.Symbol map
type Map interface {
	SymbolOf(string) (xsd.Symbol, error)
	ValueOf(s xsd.Symbol) string
}

type symbolMap struct {
	rw      sync.RWMutex
	hash    Hash
	symbols hashmap
}

// Create instance of in-memory symbols table
func New() Map {
	m := hashmap{}
	f := NewHash(m)

	return &symbolMap{hash: f, symbols: m}
}

// Cast string to xsd.Symbol
func (m *symbolMap) SymbolOf(s string) (xsd.Symbol, error) {
	m.rw.Lock()
	defer m.rw.Unlock()

	hash, err := m.hash.String(s)
	if err != nil {
		return 0, err
	}

	m.symbols[hash] = s

	return xsd.Symbol(hash), nil
}

// Cast xsd.Symbol to string
func (m *symbolMap) ValueOf(s xsd.Symbol) string {
	m.rw.RLock()
	defer m.rw.RUnlock()

	val, has := m.symbols[uint32(s)]
	if !has {
		return fmt.Sprintf(":%d", s)
	}

	return val
}

// ---------------------------------------------------------------

type hashmap map[uint32]string

func (m hashmap) Get(key uint32) (string, error) {
	val, has := m[key]

	if !has {
		return "", nil
	}
	return val, nil
}
