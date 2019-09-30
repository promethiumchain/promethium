// Copyright 2018 The promethium Authors
// This file is part of the promethium library.
//
// The promethium library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The promethium library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the promethium library. If not, see <http://www.gnu.org/licenses/>.

package accounts

import (
	"testing"
)

func TestURLParsing(t *testing.T) {
	url, err := parseURL("https://promethium.org")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if url.Scheme != "https" {
		t.Errorf("expected: %v, got: %v", "https", url.Scheme)
	}
	if url.Path != "promethium.org" {
		t.Errorf("expected: %v, got: %v", "promethium.org", url.Path)
	}

	_, err = parseURL("promethium.org")
	if err == nil {
		t.Error("expected err, got: nil")
	}
}

func TestURLString(t *testing.T) {
	url := URL{Scheme: "https", Path: "promethium.org"}
	if url.String() != "https://promethium.org" {
		t.Errorf("expected: %v, got: %v", "https://promethium.org", url.String())
	}

	url = URL{Scheme: "", Path: "promethium.org"}
	if url.String() != "promethium.org" {
		t.Errorf("expected: %v, got: %v", "promethium.org", url.String())
	}
}

func TestURLMarshalJSON(t *testing.T) {
	url := URL{Scheme: "https", Path: "promethium.org"}
	json, err := url.MarshalJSON()
	if err != nil {
		t.Errorf("unexpcted error: %v", err)
	}
	if string(json) != "\"https://promethium.org\"" {
		t.Errorf("expected: %v, got: %v", "\"https://promethium.org\"", string(json))
	}
}

func TestURLUnmarshalJSON(t *testing.T) {
	url := &URL{}
	err := url.UnmarshalJSON([]byte("\"https://promethium.org\""))
	if err != nil {
		t.Errorf("unexpcted error: %v", err)
	}
	if url.Scheme != "https" {
		t.Errorf("expected: %v, got: %v", "https", url.Scheme)
	}
	if url.Path != "promethium.org" {
		t.Errorf("expected: %v, got: %v", "https", url.Path)
	}
}

func TestURLComparison(t *testing.T) {
	tests := []struct {
		urlA   URL
		urlB   URL
		expect int
	}{
		{URL{"https", "promethium.org"}, URL{"https", "promethium.org"}, 0},
		{URL{"http", "promethium.org"}, URL{"https", "promethium.org"}, -1},
		{URL{"https", "promethium.org/a"}, URL{"https", "promethium.org"}, 1},
		{URL{"https", "abc.org"}, URL{"https", "promethium.org"}, -1},
	}

	for i, tt := range tests {
		result := tt.urlA.Cmp(tt.urlB)
		if result != tt.expect {
			t.Errorf("test %d: cmp mismatch: expected: %d, got: %d", i, tt.expect, result)
		}
	}
}
