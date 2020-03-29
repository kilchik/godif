package main

import "testing"

func TestIsSkippable(t *testing.T) {
	for fname, expected := range map[string]bool {
		"path/to/file.go": false,
		"path/to/file_test.go": true,
		"path/to/file_mock.go": true,
		"./pkg/card/card.pb.go": true,
		"pkg/card/card.pb.goclay.go": true,
	} {
		got := isSkippable(fname)
		if got != expected {
			t.Fatalf("%q: expected %v; got: %v", fname, expected, got)
		}
	}
}
