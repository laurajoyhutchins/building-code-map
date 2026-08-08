package mcp

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"strings"
	"testing"

	"building-code-map/backend/engine"
	"building-code-map/backend/regulatory"
)

func TestServeInitializesListsBoundedToolsAndReturnsStructuredEngineError(t *testing.T) {
	input := strings.NewReader(`{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}
{"jsonrpc":"2.0","id":2,"method":"tools/list","params":{}}
{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"building_code_resolve","arguments":{}}}
`)
	var output bytes.Buffer
	if err := Serve(context.Background(), input, &output, io.Discard, fakeEngine{resolveErr: engine.EngineError{Code: engine.ErrorInvalidQuery, Message: "date required"}}); err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(output.String()), "\n")
	if len(lines) != 3 {
		t.Fatalf("responses=%q", lines)
	}
	var initialize map[string]any
	if err := json.Unmarshal([]byte(lines[0]), &initialize); err != nil {
		t.Fatal(err)
	}
	if initialize["result"].(map[string]any)["protocolVersion"] != protocolVersion {
		t.Fatalf("initialize=%#v", initialize)
	}
	var list map[string]any
	if err := json.Unmarshal([]byte(lines[1]), &list); err != nil {
		t.Fatal(err)
	}
	if len(list["result"].(map[string]any)["tools"].([]any)) != 5 {
		t.Fatalf("tools=%#v", list)
	}
	var call map[string]any
	if err := json.Unmarshal([]byte(lines[2]), &call); err != nil {
		t.Fatal(err)
	}
	result := call["result"].(map[string]any)
	if result["isError"] != true || result["structuredContent"].(map[string]any)["code"] != string(engine.ErrorInvalidQuery) {
		t.Fatalf("tool error=%#v", result)
	}
}

func TestServePropagatesCancellationToEngine(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	input := strings.NewReader(`{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"building_code_resolve","arguments":{"point":{"longitude":0,"latitude":0},"applicability_date":"2026-08-06"}}}
`)
	server := NewServer(fakeEngine{waitForCancel: true}, io.Discard)
	result := make(chan error, 1)
	go func() {
		result <- server.Serve(ctx, input, io.Discard)
	}()
	cancel()
	if err := <-result; !errors.Is(err, context.Canceled) {
		t.Fatalf("Serve() error=%v", err)
	}
}

type fakeEngine struct {
	resolveErr    error
	waitForCancel bool
}

func (fake fakeEngine) Resolve(ctx context.Context, _ engine.Query) (engine.Result, error) {
	if fake.waitForCancel {
		<-ctx.Done()
		return engine.Result{}, ctx.Err()
	}
	return engine.Result{}, fake.resolveErr
}
func (fake fakeEngine) Geocode(context.Context, string) (engine.GeocodeResult, error) {
	return engine.GeocodeResult{}, nil
}
func (fake fakeEngine) Lookup(context.Context, engine.Point) (engine.LookupResult, error) {
	return engine.LookupResult{}, nil
}
func (fake fakeEngine) Readiness(context.Context) engine.Readiness { return engine.Readiness{} }
func (fake fakeEngine) BundleIdentity(context.Context) engine.BundleIdentity {
	return engine.BundleIdentity{}
}
func (fake fakeEngine) GetJurisdiction(context.Context, string) (regulatory.StateProfile, error) {
	return regulatory.StateProfile{}, nil
}
