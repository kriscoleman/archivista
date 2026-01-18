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
	"fmt"
	"strconv"
	"sync"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// PredicateType object resolvers

var predicateTypeImplementors = []string{"PredicateType"}

func (ec *executionContext) _PredicateType(ctx context.Context, sel ast.SelectionSet, obj *PredicateType) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, predicateTypeImplementors)

	out := graphql.NewFieldSet(fields)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("PredicateType")
		case "value":
			out.Values[i] = ec._PredicateType_value(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "count":
			out.Values[i] = ec._PredicateType_count(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}
	return out
}

func (ec *executionContext) _PredicateType_value(ctx context.Context, field graphql.CollectedField, obj *PredicateType) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_PredicateType_value(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	res := obj.Value
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_PredicateType_value(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "PredicateType",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, fmt.Errorf("no field named %q was found under type String", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _PredicateType_count(ctx context.Context, field graphql.CollectedField, obj *PredicateType) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_PredicateType_count(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	res := obj.Count
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_PredicateType_count(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "PredicateType",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, fmt.Errorf("no field named %q was found under type Int", field.Name)
		},
	}
	return fc, nil
}

// PredicateTypeConnection object resolvers

var predicateTypeConnectionImplementors = []string{"PredicateTypeConnection"}

func (ec *executionContext) _PredicateTypeConnection(ctx context.Context, sel ast.SelectionSet, obj *PredicateTypeConnection) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, predicateTypeConnectionImplementors)

	out := graphql.NewFieldSet(fields)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("PredicateTypeConnection")
		case "edges":
			out.Values[i] = ec._PredicateTypeConnection_edges(ctx, field, obj)
		case "pageInfo":
			out.Values[i] = ec._PredicateTypeConnection_pageInfo(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "totalCount":
			out.Values[i] = ec._PredicateTypeConnection_totalCount(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}
	return out
}

func (ec *executionContext) _PredicateTypeConnection_edges(ctx context.Context, field graphql.CollectedField, obj *PredicateTypeConnection) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_PredicateTypeConnection_edges(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	res := obj.Edges
	fc.Result = res
	return ec.marshalOPredicateTypeEdge2ᚕᚖgithubᚗcomᚋinᚑtotoᚋarchivistaᚐPredicateTypeEdge(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_PredicateTypeConnection_edges(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "PredicateTypeConnection",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "node":
				return ec.fieldContext_PredicateTypeEdge_node(ctx, field)
			case "cursor":
				return ec.fieldContext_PredicateTypeEdge_cursor(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type PredicateTypeEdge", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _PredicateTypeConnection_pageInfo(ctx context.Context, field graphql.CollectedField, obj *PredicateTypeConnection) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_PredicateTypeConnection_pageInfo(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	res := obj.PageInfo
	fc.Result = res
	return ec.marshalNPageInfo2entgoᚗioᚋcontribᚋentgqlᚐPageInfo(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_PredicateTypeConnection_pageInfo(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "PredicateTypeConnection",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "hasNextPage":
				return ec.fieldContext_PageInfo_hasNextPage(ctx, field)
			case "hasPreviousPage":
				return ec.fieldContext_PageInfo_hasPreviousPage(ctx, field)
			case "startCursor":
				return ec.fieldContext_PageInfo_startCursor(ctx, field)
			case "endCursor":
				return ec.fieldContext_PageInfo_endCursor(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type PageInfo", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _PredicateTypeConnection_totalCount(ctx context.Context, field graphql.CollectedField, obj *PredicateTypeConnection) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_PredicateTypeConnection_totalCount(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	res := obj.TotalCount
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_PredicateTypeConnection_totalCount(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "PredicateTypeConnection",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, fmt.Errorf("no field named %q was found under type Int", field.Name)
		},
	}
	return fc, nil
}

// PredicateTypeEdge object resolvers

var predicateTypeEdgeImplementors = []string{"PredicateTypeEdge"}

func (ec *executionContext) _PredicateTypeEdge(ctx context.Context, sel ast.SelectionSet, obj *PredicateTypeEdge) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, predicateTypeEdgeImplementors)

	out := graphql.NewFieldSet(fields)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("PredicateTypeEdge")
		case "node":
			out.Values[i] = ec._PredicateTypeEdge_node(ctx, field, obj)
		case "cursor":
			out.Values[i] = ec._PredicateTypeEdge_cursor(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}
	return out
}

func (ec *executionContext) _PredicateTypeEdge_node(ctx context.Context, field graphql.CollectedField, obj *PredicateTypeEdge) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_PredicateTypeEdge_node(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	res := obj.Node
	fc.Result = res
	return ec.marshalOPredicateType2ᚖgithubᚗcomᚋinᚑtotoᚋarchivistaᚐPredicateType(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_PredicateTypeEdge_node(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "PredicateTypeEdge",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "value":
				return ec.fieldContext_PredicateType_value(ctx, field)
			case "count":
				return ec.fieldContext_PredicateType_count(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type PredicateType", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _PredicateTypeEdge_cursor(ctx context.Context, field graphql.CollectedField, obj *PredicateTypeEdge) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_PredicateTypeEdge_cursor(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	res := obj.Cursor
	fc.Result = res
	return ec.marshalNCursor2entgoᚗioᚋcontribᚋentgqlᚐCursor(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_PredicateTypeEdge_cursor(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "PredicateTypeEdge",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, fmt.Errorf("no field named %q was found under type Cursor", field.Name)
		},
	}
	return fc, nil
}

// Marshalers

func (ec *executionContext) marshalNPredicateTypeConnection2githubᚗcomᚋinᚑtotoᚋarchivistaᚐPredicateTypeConnection(ctx context.Context, sel ast.SelectionSet, v PredicateTypeConnection) graphql.Marshaler {
	return ec._PredicateTypeConnection(ctx, sel, &v)
}

func (ec *executionContext) marshalNPredicateTypeConnection2ᚖgithubᚗcomᚋinᚑtotoᚋarchivistaᚐPredicateTypeConnection(ctx context.Context, sel ast.SelectionSet, v *PredicateTypeConnection) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._PredicateTypeConnection(ctx, sel, v)
}

func (ec *executionContext) marshalOPredicateType2ᚖgithubᚗcomᚋinᚑtotoᚋarchivistaᚐPredicateType(ctx context.Context, sel ast.SelectionSet, v *PredicateType) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._PredicateType(ctx, sel, v)
}

func (ec *executionContext) marshalOPredicateTypeEdge2ᚕᚖgithubᚗcomᚋinᚑtotoᚋarchivistaᚐPredicateTypeEdge(ctx context.Context, sel ast.SelectionSet, v []*PredicateTypeEdge) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalOPredicateTypeEdge2ᚖgithubᚗcomᚋinᚑtotoᚋarchivistaᚐPredicateTypeEdge(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}
	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalOPredicateTypeEdge2ᚖgithubᚗcomᚋinᚑtotoᚋarchivistaᚐPredicateTypeEdge(ctx context.Context, sel ast.SelectionSet, v *PredicateTypeEdge) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._PredicateTypeEdge(ctx, sel, v)
}

// Unmarshalers

func (ec *executionContext) unmarshalOPredicateTypeWhereInput2ᚖgithubᚗcomᚋinᚑtotoᚋarchivistaᚐPredicateTypeWhereInput(ctx context.Context, v any) (*PredicateTypeWhereInput, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalInputPredicateTypeWhereInput(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalInputPredicateTypeWhereInput(ctx context.Context, obj any) (PredicateTypeWhereInput, error) {
	var it PredicateTypeWhereInput
	asMap := map[string]any{}
	for k, v := range obj.(map[string]any) {
		asMap[k] = v
	}

	for k, v := range asMap {
		switch k {
		case "value":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("value"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.Value = data
		case "valueHasPrefix":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("valueHasPrefix"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.ValueHasPrefix = data
		case "valueHasSuffix":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("valueHasSuffix"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.ValueHasSuffix = data
		case "valueContains":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("valueContains"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.ValueContains = data
		}
	}
	return it, nil
}

func (ec *executionContext) unmarshalOCursor2ᚖentgoᚗioᚋcontribᚋentgqlᚐCursorᚄstring(ctx context.Context, v any) (*entgql.Cursor[string], error) {
	if v == nil {
		return nil, nil
	}
	var res = new(entgql.Cursor[string])
	err := res.UnmarshalGQL(v)
	return res, graphql.ErrorOnPath(ctx, err)
}
