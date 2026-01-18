// Copyright 2022 The Archivista Contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package archivista

import "entgo.io/contrib/entgql"

// PredicateType represents a unique predicate type from in-toto statements.
type PredicateType struct {
	// Value is the predicate type URI (e.g., "https://in-toto.io/attestation/vulns/v0.1").
	Value string `json:"value"`
	// Count is the number of statements using this predicate type.
	Count int `json:"count"`
}

// PredicateTypeConnection is the connection containing PredicateType edges.
type PredicateTypeConnection struct {
	// Edges contains the PredicateType edges.
	Edges []*PredicateTypeEdge `json:"edges"`
	// PageInfo contains pagination information.
	PageInfo entgql.PageInfo[string] `json:"pageInfo"`
	// TotalCount is the total number of unique predicate types.
	TotalCount int `json:"totalCount"`
}

// PredicateTypeEdge is an edge in a PredicateType connection.
type PredicateTypeEdge struct {
	// Node is the PredicateType at the end of the edge.
	Node *PredicateType `json:"node"`
	// Cursor is a cursor for pagination.
	Cursor entgql.Cursor[string] `json:"cursor"`
}

// PredicateTypeWhereInput is used for filtering PredicateType results.
type PredicateTypeWhereInput struct {
	// Value filters by exact predicate value.
	Value *string `json:"value,omitempty"`
	// ValueHasPrefix filters by predicate prefix.
	ValueHasPrefix *string `json:"valueHasPrefix,omitempty"`
	// ValueHasSuffix filters by predicate suffix.
	ValueHasSuffix *string `json:"valueHasSuffix,omitempty"`
	// ValueContains filters by predicate containing substring.
	ValueContains *string `json:"valueContains,omitempty"`
}
