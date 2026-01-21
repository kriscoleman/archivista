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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncodeCursor(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{
			name:     "simple string",
			value:    "https://in-toto.io/attestation/vulns/v0.1",
			expected: "aHR0cHM6Ly9pbi10b3RvLmlvL2F0dGVzdGF0aW9uL3Z1bG5zL3YwLjE=",
		},
		{
			name:     "empty string",
			value:    "",
			expected: "",
		},
		{
			name:     "special characters",
			value:    "https://example.com/predicate?type=test&version=1",
			expected: "aHR0cHM6Ly9leGFtcGxlLmNvbS9wcmVkaWNhdGU/dHlwZT10ZXN0JnZlcnNpb249MQ==",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := encodeCursor(tt.value)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDecodeCursor(t *testing.T) {
	tests := []struct {
		name     string
		cursor   string
		expected string
		wantErr  bool
	}{
		{
			name:     "simple string",
			cursor:   "aHR0cHM6Ly9pbi10b3RvLmlvL2F0dGVzdGF0aW9uL3Z1bG5zL3YwLjE=",
			expected: "https://in-toto.io/attestation/vulns/v0.1",
			wantErr:  false,
		},
		{
			name:     "empty string",
			cursor:   "",
			expected: "",
			wantErr:  false,
		},
		{
			name:     "special characters",
			cursor:   "aHR0cHM6Ly9leGFtcGxlLmNvbS9wcmVkaWNhdGU/dHlwZT10ZXN0JnZlcnNpb249MQ==",
			expected: "https://example.com/predicate?type=test&version=1",
			wantErr:  false,
		},
		{
			name:     "invalid base64",
			cursor:   "not-valid-base64!@#$",
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := decodeCursor(tt.cursor)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestEncodeDecode_RoundTrip(t *testing.T) {
	values := []string{
		"https://in-toto.io/attestation/vulns/v0.1",
		"https://slsa.dev/provenance/v1",
		"https://witness.dev/attestations/collection/v0.1",
		"https://in-toto.io/Statement/v1",
	}

	for _, value := range values {
		t.Run(value, func(t *testing.T) {
			encoded := encodeCursor(value)
			decoded, err := decodeCursor(encoded)
			require.NoError(t, err)
			assert.Equal(t, value, decoded)
		})
	}
}

func TestPredicateType_Fields(t *testing.T) {
	pt := &PredicateType{
		Value: "https://in-toto.io/attestation/vulns/v0.1",
		Count: 42,
	}

	assert.Equal(t, "https://in-toto.io/attestation/vulns/v0.1", pt.Value)
	assert.Equal(t, 42, pt.Count)
}

func TestPredicateTypeWhereInput_Filters(t *testing.T) {
	tests := []struct {
		name  string
		input PredicateTypeWhereInput
	}{
		{
			name: "exact value filter",
			input: PredicateTypeWhereInput{
				Value: strPtr("https://in-toto.io/attestation/vulns/v0.1"),
			},
		},
		{
			name: "prefix filter",
			input: PredicateTypeWhereInput{
				ValueHasPrefix: strPtr("https://in-toto.io"),
			},
		},
		{
			name: "suffix filter",
			input: PredicateTypeWhereInput{
				ValueHasSuffix: strPtr("/v0.1"),
			},
		},
		{
			name: "contains filter",
			input: PredicateTypeWhereInput{
				ValueContains: strPtr("attestation"),
			},
		},
		{
			name: "multiple filters",
			input: PredicateTypeWhereInput{
				ValueHasPrefix: strPtr("https://"),
				ValueContains:  strPtr("attestation"),
			},
		},
		{
			name:  "no filters",
			input: PredicateTypeWhereInput{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Just verify the struct can be created with various filter combinations
			assert.NotNil(t, &tt.input)
		})
	}
}

func strPtr(s string) *string {
	return &s
}
