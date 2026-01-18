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
	"context"
	"encoding/base64"
	"sort"

	"entgo.io/contrib/entgql"
	"github.com/in-toto/archivista/ent"
	"github.com/in-toto/archivista/ent/statement"
)

// PredicateTypes is the resolver for the predicateTypes field.
func (r *queryResolver) PredicateTypes(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, where *PredicateTypeWhereInput) (*PredicateTypeConnection, error) {
	// Build query with filters
	query := r.client.Statement.Query()

	// Apply filters from where input
	if where != nil {
		if where.Value != nil {
			query = query.Where(statement.PredicateEQ(*where.Value))
		}
		if where.ValueHasPrefix != nil {
			query = query.Where(statement.PredicateHasPrefix(*where.ValueHasPrefix))
		}
		if where.ValueHasSuffix != nil {
			query = query.Where(statement.PredicateHasSuffix(*where.ValueHasSuffix))
		}
		if where.ValueContains != nil {
			query = query.Where(statement.PredicateContains(*where.ValueContains))
		}
	}

	// Get distinct predicates with counts using GroupBy
	var results []struct {
		Predicate string `json:"predicate"`
		Count     int    `json:"count"`
	}

	err := query.
		GroupBy(statement.FieldPredicate).
		Aggregate(ent.Count()).
		Scan(ctx, &results)
	if err != nil {
		return nil, err
	}

	// Sort results by predicate value for consistent ordering
	sort.Slice(results, func(i, j int) bool {
		return results[i].Predicate < results[j].Predicate
	})

	// Create all predicate types
	allPredicateTypes := make([]*PredicateType, len(results))
	for i, r := range results {
		allPredicateTypes[i] = &PredicateType{
			Value: r.Predicate,
			Count: r.Count,
		}
	}

	totalCount := len(allPredicateTypes)

	// Apply cursor-based pagination
	startIdx := 0
	endIdx := totalCount

	// Handle 'after' cursor
	if after != nil {
		afterValue, err := decodeCursor(after.ID)
		if err == nil {
			for i, pt := range allPredicateTypes {
				if pt.Value == afterValue {
					startIdx = i + 1
					break
				}
			}
		}
	}

	// Handle 'before' cursor
	if before != nil {
		beforeValue, err := decodeCursor(before.ID)
		if err == nil {
			for i, pt := range allPredicateTypes {
				if pt.Value == beforeValue {
					endIdx = i
					break
				}
			}
		}
	}

	// Apply first/last limits
	if first != nil && *first > 0 {
		if startIdx+*first < endIdx {
			endIdx = startIdx + *first
		}
	}

	if last != nil && *last > 0 {
		if endIdx-*last > startIdx {
			startIdx = endIdx - *last
		}
	}

	// Ensure valid range
	if startIdx < 0 {
		startIdx = 0
	}
	if endIdx > totalCount {
		endIdx = totalCount
	}
	if startIdx > endIdx {
		startIdx = endIdx
	}

	// Get the slice of predicate types for this page
	pagePredicateTypes := allPredicateTypes[startIdx:endIdx]

	// Build edges with cursors
	edges := make([]*PredicateTypeEdge, len(pagePredicateTypes))
	for i, pt := range pagePredicateTypes {
		edges[i] = &PredicateTypeEdge{
			Node:   pt,
			Cursor: entgql.Cursor[string]{ID: encodeCursor(pt.Value)},
		}
	}

	// Build page info
	var startCursor, endCursor *entgql.Cursor[string]
	if len(edges) > 0 {
		startCursor = &edges[0].Cursor
		endCursor = &edges[len(edges)-1].Cursor
	}

	pageInfo := entgql.PageInfo[string]{
		HasNextPage:     endIdx < totalCount,
		HasPreviousPage: startIdx > 0,
		StartCursor:     startCursor,
		EndCursor:       endCursor,
	}

	return &PredicateTypeConnection{
		Edges:      edges,
		PageInfo:   pageInfo,
		TotalCount: totalCount,
	}, nil
}

// encodeCursor encodes a predicate value to a cursor string
func encodeCursor(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}

// decodeCursor decodes a cursor string to a predicate value
func decodeCursor(cursor string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
