package mcp

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"building-code-map/backend/engine"
)

const protocolVersion = "2024-11-05"

type Server struct {
	engine engine.Engine
	log    io.Writer
}

func NewServer(authority engine.Engine, log io.Writer) *Server {
	return &Server{engine: authority, log: log}
}

func Serve(ctx context.Context, in io.Reader, out io.Writer, log io.Writer, authority engine.Engine) error {
	return NewServer(authority, log).Serve(ctx, in, out)
}

type request struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type response struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Result  any             `json:"result,omitempty"`
	Error   *protocolError  `json:"error,omitempty"`
}

type protocolError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type tool struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	InputSchema any    `json:"inputSchema"`
}

type toolCallResult struct {
	Content           []map[string]any `json:"content"`
	StructuredContent any              `json:"structuredContent,omitempty"`
	IsError           bool             `json:"isError,omitempty"`
}

func (server *Server) Serve(ctx context.Context, in io.Reader, out io.Writer) error {
	reader := bufio.NewReader(in)
	for {
		if err := ctx.Err(); err != nil {
			return err
		}
		payload, err := readFrame(reader)
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}
		var req request
		if err := json.Unmarshal(payload, &req); err != nil {
			if err := writeResponse(out, response{JSONRPC: "2.0", Error: &protocolError{Code: -32700, Message: "parse error"}}); err != nil {
				return err
			}
			continue
		}
		result, shouldRespond := server.dispatch(ctx, req)
		if !shouldRespond {
			continue
		}
		if err := writeResponse(out, result); err != nil {
			return err
		}
	}
}

func (server *Server) dispatch(ctx context.Context, req request) (response, bool) {
	if req.JSONRPC != "2.0" || req.Method == "" {
		return errorResponse(req.ID, -32600, "invalid request", nil), true
	}
	switch req.Method {
	case "notifications/initialized", "notifications/cancelled":
		return response{}, false
	case "initialize":
		return response{JSONRPC: "2.0", ID: req.ID, Result: map[string]any{
			"protocolVersion": protocolVersion,
			"capabilities":    map[string]any{"tools": map[string]any{}},
			"serverInfo":      map[string]any{"name": "bcm", "version": "0.1.0"},
		}}, true
	case "tools/list":
		return response{JSONRPC: "2.0", ID: req.ID, Result: map[string]any{"tools": toolDefinitions()}}, true
	case "tools/call":
		result := server.callTool(ctx, req.Params)
		return response{JSONRPC: "2.0", ID: req.ID, Result: result}, true
	default:
		return errorResponse(req.ID, -32601, "method not found", nil), true
	}
}

func (server *Server) callTool(ctx context.Context, raw json.RawMessage) toolCallResult {
	var params struct {
		Name      string          `json:"name"`
		Arguments json.RawMessage `json:"arguments"`
	}
	if err := decodeStrict(raw, &params); err != nil || strings.TrimSpace(params.Name) == "" {
		return toolError(engine.EngineError{Code: engine.ErrorInvalidQuery, Message: "tool name and arguments are required", Retryable: false})
	}
	switch params.Name {
	case "building_code_resolve":
		var query engine.Query
		if err := decodeStrict(params.Arguments, &query); err != nil {
			return toolError(engine.EngineError{Code: engine.ErrorInvalidQuery, Message: err.Error(), Retryable: false})
		}
		result, err := server.engine.Resolve(ctx, query)
		if err != nil {
			return toolError(err)
		}
		return toolSuccess(result)
	case "building_code_geocode":
		var input struct {
			Address string `json:"address"`
		}
		if err := decodeStrict(params.Arguments, &input); err != nil {
			return toolError(engine.EngineError{Code: engine.ErrorInvalidQuery, Message: err.Error(), Retryable: false})
		}
		result, err := server.engine.Geocode(ctx, input.Address)
		if err != nil {
			return toolError(err)
		}
		return toolSuccess(result)
	case "building_code_lookup":
		var input struct {
			Point engine.Point `json:"point"`
		}
		if err := decodeStrict(params.Arguments, &input); err != nil {
			return toolError(engine.EngineError{Code: engine.ErrorInvalidQuery, Message: err.Error(), Retryable: false})
		}
		result, err := server.engine.Lookup(ctx, input.Point)
		if err != nil {
			return toolError(err)
		}
		return toolSuccess(result)
	case "building_code_get_bundle":
		return toolSuccess(server.engine.BundleIdentity(ctx))
	case "building_code_get_jurisdiction":
		var input struct {
			ID string `json:"id"`
		}
		if err := decodeStrict(params.Arguments, &input); err != nil {
			return toolError(engine.EngineError{Code: engine.ErrorInvalidQuery, Message: err.Error(), Retryable: false})
		}
		result, err := server.engine.GetJurisdiction(ctx, input.ID)
		if err != nil {
			return toolError(err)
		}
		return toolSuccess(result)
	default:
		return toolError(engine.EngineError{Code: engine.ErrorInvalidQuery, Message: "unknown tool", Details: map[string]any{"name": params.Name}, Retryable: false})
	}
}

func toolDefinitions() []tool {
	return []tool{
		{Name: "building_code_resolve", Description: "Resolve authority and applicability for a point or address.", InputSchema: map[string]any{"type": "object"}},
		{Name: "building_code_geocode", Description: "Geocode an address using the local snapshot.", InputSchema: map[string]any{"type": "object", "required": []string{"address"}}},
		{Name: "building_code_lookup", Description: "Look up geographic context for a coordinate.", InputSchema: map[string]any{"type": "object", "required": []string{"point"}}},
		{Name: "building_code_get_bundle", Description: "Return exact engine and component identity.", InputSchema: map[string]any{"type": "object"}},
		{Name: "building_code_get_jurisdiction", Description: "Return one compiled jurisdiction profile by ID or FIPS.", InputSchema: map[string]any{"type": "object", "required": []string{"id"}}},
	}
}

func toolSuccess(value any) toolCallResult {
	raw, _ := json.Marshal(value)
	return toolCallResult{Content: []map[string]any{{"type": "text", "text": string(raw)}}, StructuredContent: value}
}

func toolError(err error) toolCallResult {
	var engineErr engine.EngineError
	if !errors.As(err, &engineErr) {
		engineErr = engine.EngineError{Code: engine.ErrorInternal, Message: err.Error(), Retryable: false}
	}
	raw, _ := json.Marshal(engineErr)
	return toolCallResult{Content: []map[string]any{{"type": "text", "text": string(raw)}}, StructuredContent: engineErr, IsError: true}
}

func decodeStrict(raw []byte, target any) error {
	decoder := json.NewDecoder(strings.NewReader(string(raw)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(target); err != nil {
		return err
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return fmt.Errorf("request must contain one JSON value")
	}
	return nil
}

func readFrame(reader *bufio.Reader) ([]byte, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		if errors.Is(err, io.EOF) && strings.TrimSpace(line) == "" {
			return nil, io.EOF
		}
		return nil, err
	}
	line = strings.TrimSpace(line)
	if strings.HasPrefix(strings.ToLower(line), "content-length:") {
		var length int
		if _, err := fmt.Sscanf(line, "Content-Length: %d", &length); err != nil || length < 0 {
			return nil, fmt.Errorf("invalid Content-Length")
		}
		if _, err := reader.ReadString('\n'); err != nil {
			return nil, err
		}
		payload := make([]byte, length)
		_, err := io.ReadFull(reader, payload)
		return payload, err
	}
	return []byte(line), nil
}

func writeResponse(out io.Writer, value response) error {
	raw, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = out.Write(append(raw, '\n'))
	return err
}

func errorResponse(id json.RawMessage, code int, message string, data any) response {
	return response{JSONRPC: "2.0", ID: id, Error: &protocolError{Code: code, Message: message, Data: data}}
}
