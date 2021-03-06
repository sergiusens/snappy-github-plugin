// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2014-2015 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

// Package pkg packages odds and ends that apply to packages in more than one
// layer (to more than one of deb, click, snap, and part).
package pkg

import (
	"encoding/json"
)

// Type represents the kind of snap (app, core, frameworks, oem)
type Type string

// The various types of snap parts we support
const (
	TypeApp       Type = "app"
	TypeCore      Type = "core"
	TypeFramework Type = "framework"
	TypeOem       Type = "oem"
)

// MarshalJSON returns *m as the JSON encoding of m.
func (m Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(m))
}

// UnmarshalJSON sets *m to a copy of data.
func (m *Type) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	// this is a workaround as the store sends "application" but snappy uses
	// "app" for TypeApp
	if str == "application" {
		*m = TypeApp
	} else {
		*m = Type(str)
	}

	return nil
}
