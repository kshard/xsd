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

type Map struct {
	rw      sync.RWMutex
	hash    Hash
	symbols map[uint32]string
}

func New() *Map {
	m := &Map{symbols: make(map[uint32]string)}
	m.hash = NewHash(m)
	return m
}

func (m *Map) Get(key uint32) (string, error) {
	val, has := m.symbols[key]

	if !has {
		return "", nil
	}
	return val, nil
}

func (m *Map) SymbolOf(s string) (xsd.Symbol, error) {
	m.rw.Lock()
	defer m.rw.Unlock()

	hash, err := m.hash.String(s)
	if err != nil {
		return 0, err
	}

	m.symbols[hash] = s

	return xsd.Symbol(hash), nil
}

func (m *Map) ValueOf(s xsd.Symbol) string {
	m.rw.RLock()
	defer m.rw.RUnlock()

	val, has := m.symbols[uint32(s)]
	if !has {
		return fmt.Sprintf(":%d", s)
	}

	return val
}
