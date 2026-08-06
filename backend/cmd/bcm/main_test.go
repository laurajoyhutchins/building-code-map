package main

import (
	"path/filepath"
	"testing"

	"building-code-map/backend/bundle"
)

func TestRegulatoryCatalogRoot(t *testing.T) {
	tests := []struct {
		name      string
		component bundle.Component
		path      string
		want      string
	}{
		{
			name:      "recursive catalog directory",
			component: bundle.Component{Recursive: true},
			path:      filepath.Join("bundle", "data", "regulatory"),
			want:      filepath.Join("bundle", "data", "regulatory"),
		},
		{
			name:      "legacy catalog file",
			component: bundle.Component{},
			path:      filepath.Join("bundle", "data", "regulatory", "colorado.json"),
			want:      filepath.Join("bundle", "data", "regulatory"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := regulatoryCatalogRoot(test.component, test.path); got != test.want {
				t.Fatalf("regulatoryCatalogRoot() = %q, want %q", got, test.want)
			}
		})
	}
}
