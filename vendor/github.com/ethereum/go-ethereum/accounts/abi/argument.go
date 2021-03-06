// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package abi

import (
	"encoding/json"
	"fmt"
)

// Argument holds the name of the argument and the corresponding type.
// Types are used when packing and testing arguments.
type Argument struct {
	Name    string
	Type    Type
	Indexed bool // indexed is only used by events
}

// Type only used for unmarshalling json argument types
type unmarshalArg struct {
	Name       string
	Type       string
	Components []unmarshalArg // used for tuples/structs
	Indexed    bool
}

func (a *Argument) UnmarshalJSON(data []byte) error {
	var extarg unmarshalArg
	err := json.Unmarshal(data, &extarg)
	if err != nil {
		return fmt.Errorf("argument json err: %v", err)
	}

	if len(extarg.Components) > 0 {
		a.Type, err = ParseStructType(extarg.Type, extarg.Components...)
	} else {
		a.Type, err = NewType(extarg.Type)
	}

	if err != nil {
		return err
	}

	a.Name = extarg.Name
	a.Indexed = extarg.Indexed

	return nil
}
