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

	"entgo.io/contrib/entgql"
	"github.com/google/uuid"
)

// PageInfo is the resolver for the pageInfo field.
// This converts our custom PredicateTypePageInfo to the expected ent PageInfo type.
func (r *predicateTypeConnectionResolver) PageInfo(ctx context.Context, obj *PredicateTypeConnection) (*entgql.PageInfo[uuid.UUID], error) {
	// Convert our custom cursors to ent cursors
	var startCursor, endCursor *entgql.Cursor[uuid.UUID]

	if obj.PageInfo.StartCursor != nil {
		startCursor = &entgql.Cursor[uuid.UUID]{
			Value: obj.PageInfo.StartCursor.Value,
		}
	}

	if obj.PageInfo.EndCursor != nil {
		endCursor = &entgql.Cursor[uuid.UUID]{
			Value: obj.PageInfo.EndCursor.Value,
		}
	}

	return &entgql.PageInfo[uuid.UUID]{
		HasNextPage:     obj.PageInfo.HasNextPage,
		HasPreviousPage: obj.PageInfo.HasPreviousPage,
		StartCursor:     startCursor,
		EndCursor:       endCursor,
	}, nil
}

// PredicateTypes is the resolver for the predicateTypes field.
// Note: The generated signature uses ent Cursor types, but we delegate to
// predicateTypes() helper which uses custom PredicateTypeCursor types.
func (r *queryResolver) PredicateTypes(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, where *PredicateTypeWhereInput) (*PredicateTypeConnection, error) {
	// Convert ent cursors to PredicateTypeCursor
	// Note: This is a workaround because gqlgen generates UUID cursor types
	// but we need string cursors for predicate values
	var afterCursor, beforeCursor *PredicateTypeCursor

	if after != nil {
		// Extract cursor value from the ent.Value field if available
		if v, ok := after.Value.(string); ok {
			afterCursor = &PredicateTypeCursor{Value: v}
		}
	}

	if before != nil {
		// Extract cursor value from the ent.Value field if available
		if v, ok := before.Value.(string); ok {
			beforeCursor = &PredicateTypeCursor{Value: v}
		}
	}

	return r.predicateTypes(ctx, afterCursor, first, beforeCursor, last, where)
}

// PredicateTypeConnection returns PredicateTypeConnectionResolver implementation.
func (r *Resolver) PredicateTypeConnection() PredicateTypeConnectionResolver {
	return &predicateTypeConnectionResolver{r}
}

type predicateTypeConnectionResolver struct{ *Resolver }
