// Copyright 2020 ZetaMesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package tun

import "io"

// DefaultMTU represents the default Maximum Transmission Unit
const DefaultMTU = 1420

// Device represents a layer-three virtual network device
type Device interface {
	Name() string
	io.ReadWriteCloser
}

type generalDevice struct {
	name string
	io.ReadWriteCloser
}

// Name implements the Device interface
func (d *generalDevice) Name() string {
	return d.name
}
