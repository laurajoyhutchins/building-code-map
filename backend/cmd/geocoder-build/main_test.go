package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunRejectsMissingRequiredFlags(t *testing.T) {
	var stderr bytes.Buffer
	if code := run(nil, &stderr); code != 2 {
		t.Fatalf("code=%d stderr=%q", code, stderr.String())
	}
	if !strings.Contains(stderr.String(), "output") {
		t.Fatalf("stderr=%q", stderr.String())
	}
}
