/*******************************************************************************
 * Copyright 2020 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package status

import "github.com/project-alvarium/go-sdk/pkg/annotator/provenance"

const (
	Success Value = iota
	PublisherError
	NotFound
	Exists
	Unknown
)

// New is a factory function that returns an initialized Contract.
func New(provenance provenance.Contract, value Value) *Contract {
	return &Contract{
		Provenance: provenance,
		Value:      value,
	}
}
