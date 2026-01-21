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

import (
	"encoding/base64"
	"fmt"
	"io"
)

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
	PageInfo PredicateTypePageInfo `json:"pageInfo"`
	// TotalCount is the total number of unique predicate types.
	TotalCount int `json:"totalCount"`
}

// PredicateTypeEdge is an edge in a PredicateType connection.
type PredicateTypeEdge struct {
	// Node is the PredicateType at the end of the edge.
	Node *PredicateType `json:"node"`
	// Cursor is a cursor for pagination.
	Cursor PredicateTypeCursor `json:"cursor"`
}

// PredicateTypePageInfo contains pagination information for PredicateTypeConnection.
type PredicateTypePageInfo struct {
	// HasNextPage indicates if there are more items when paginating forward.
	HasNextPage bool `json:"hasNextPage"`
	// HasPreviousPage indicates if there are more items when paginating backward.
	HasPreviousPage bool `json:"hasPreviousPage"`
	// StartCursor is the cursor of the first item in the list.
	StartCursor *PredicateTypeCursor `json:"startCursor,omitempty"`
	// EndCursor is the cursor of the last item in the list.
	EndCursor *PredicateTypeCursor `json:"endCursor,omitempty"`
}

// PredicateTypeCursor is a cursor for pagination based on predicate string values.
type PredicateTypeCursor struct {
	Value string
}

// MarshalGQL implements the graphql.Marshaler interface.
func (c PredicateTypeCursor) MarshalGQL(w io.Writer) {
	encoded := encodeCursor(c.Value)
	fmt.Fprintf(w, "%q", encoded)
}

// UnmarshalGQL implements the graphql.Unmarshaler interface.
func (c *PredicateTypeCursor) UnmarshalGQL(v interface{}) error {
	cursor, ok := v.(string)
	if !ok {
		return fmt.Errorf("cursor must be a string")
	}
	decoded, err := decodeCursor(cursor)
	if err != nil {
		return fmt.Errorf("invalid cursor: %w", err)
	}
	c.Value = decoded
	return nil
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

// encodeCursor encodes a string value into a base64-encoded cursor.
func encodeCursor(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}

// decodeCursor decodes a base64-encoded cursor back into a string value.
func decodeCursor(cursor string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
