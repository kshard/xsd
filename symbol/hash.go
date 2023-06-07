//
// Copyright (C) 2023 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/xsd
//

package symbol

import (
	"errors"
	"fmt"
	"hash"
	"hash/fnv"
	"strconv"
)

const maxLoop = 10

type Getter interface {
	Get(uint32) (string, error)
}

type Hash struct {
	hash   hash.Hash64
	getter Getter
}

func NewHash(getter Getter) Hash {
	return Hash{
		getter: getter,
		hash:   fnv.New64a(),
	}
}

func (h Hash) String(s string) (uint32, error) {
	h.hash.Reset()
	h.hash.Write([]byte(s))
	hash := h.hash.Sum64()

	lo := uint32(hash)
	hi := uint32(hash >> 32)

	for attempt := 0; attempt < maxLoop; attempt++ {
		val, err := h.getter.Get(lo)
		if err != nil && isNotFound(err) {
			return lo, nil
		}
		if err != nil {
			return 0, err
		}
		if val == "" || val == s {
			return lo, nil
		}

		h.hash.Write([]byte(strconv.Itoa(int(hash))))
		hash = hash ^ h.hash.Sum64()
		lo = ((lo << 16) | (lo >> 16)) ^ hi
		hi = uint32(hash >> 32)
	}

	return 0, fmt.Errorf("hash collision of value: %s", s)
}

func isNotFound(err error) bool {
	var e interface{ NotFound() string }

	if ok := errors.As(err, &e); !ok {
		return false
	}

	return e.NotFound() != ""
}
