// Package bcastconsumer is a bounded, independent Go read-only consumer of
// BCAST 0.2. It never imports provider compilers or infers code applicability.
package bcastconsumer

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"
)

const maxBytes = 16 << 20

var shaID = regexp.MustCompile("^[0-9a-f]{64}$")
var family = regexp.MustCompile("^[a-z0-9]+(?:-[a-z0-9]+)*$")
var namespace = regexp.MustCompile("^(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?)(?:\\.[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?)+$")

type Package struct {
	ID            string
	PublicationID string
	Objects       []map[string]any
	raw           map[string]any
}

func (p *Package) Object(id string) (map[string]any, bool) {
	for _, o := range p.Objects {
		if o["object_id"] == id {
			return o, true
		}
	}
	return nil, false
}
func (p *Package) Children(parent string) []map[string]any {
	out := []map[string]any{}
	for _, o := range p.Objects {
		if o["parent_id"] == parent {
			out = append(out, o)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i]["object_id"].(string) < out[j]["object_id"].(string) })
	return out
}

type Client struct {
	origin string
	HTTP   *http.Client
}

func New(origin string) (*Client, error) {
	u, err := url.Parse(origin)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Host == "" || u.User != nil || u.RawQuery != "" || u.Fragment != "" || (u.Path != "" && u.Path != "/") {
		return nil, errors.New("BCAST endpoint must be one HTTP(S) origin")
	}
	return &Client{origin: strings.TrimSuffix(origin, "/"), HTTP: &http.Client{Timeout: 15 * time.Second, CheckRedirect: func(_ *http.Request, _ []*http.Request) error { return http.ErrUseLastResponse }}}, nil
}
func id(prefix string, data any) (string, error) {
	var b bytes.Buffer
	if err := canonical(&b, data); err != nil {
		return "", err
	}
	sum := sha256.Sum256(b.Bytes())
	return prefix + ":sha256:" + hex.EncodeToString(sum[:]), nil
}
func canonical(b *bytes.Buffer, v any) error {
	switch x := v.(type) {
	case string:
		b.WriteByte('"')
		for _, r := range x {
			switch r {
			case '"':
				b.WriteString("\\\"")
			case '\\':
				b.WriteString("\\\\")
			case '\b':
				b.WriteString("\\b")
			case '\f':
				b.WriteString("\\f")
			case '\n':
				b.WriteString("\\n")
			case '\r':
				b.WriteString("\\r")
			case '\t':
				b.WriteString("\\t")
			default:
				if r < 32 {
					fmt.Fprintf(b, "\\u%04x", r)
				} else {
					b.WriteRune(r)
				}
			}
		}
		b.WriteByte('"')
		return nil
	case []any:
		b.WriteByte('[')
		for i, z := range x {
			if i > 0 {
				b.WriteByte(',')
			}
			if e := canonical(b, z); e != nil {
				return e
			}
		}
		b.WriteByte(']')
		return nil
	case map[string]any:
		keys := make([]string, 0, len(x))
		for k := range x {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		b.WriteByte('{')
		for i, k := range keys {
			if i > 0 {
				b.WriteByte(',')
			}
			if e := canonical(b, k); e != nil {
				return e
			}
			b.WriteByte(':')
			if e := canonical(b, x[k]); e != nil {
				return e
			}
		}
		b.WriteByte('}')
		return nil
	default:
		return fmt.Errorf("invalid canonical BCAST data type %T", v)
	}
}
func expectID(value, prefix string) bool {
	return strings.HasPrefix(value, prefix+":sha256:") && shaID.MatchString(strings.TrimPrefix(value, prefix+":sha256:"))
}
func asString(m map[string]any, key string) string { s, _ := m[key].(string); return s }
func only(m map[string]any, allowed ...string) bool {
	valid := map[string]bool{}
	for _, k := range allowed {
		valid[k] = true
	}
	for k := range m {
		if !valid[k] {
			return false
		}
	}
	return true
}
func validObjectKind(s string) bool {
	switch s {
	case "document", "chapter", "appendix", "section", "subsection", "paragraph", "list_item", "definition_entry", "table", "figure", "note", "unsupported":
		return true
	}
	return false
}
func validStatus(s string) bool { return s == "resolved" || s == "unresolved" || s == "unsupported" }

// JSON Schema maxLength counts Unicode code points, not UTF-8 bytes.
// The public BCAST package schema owns these metadata limits.
func validPublicText(s string, maxCharacters int) bool {
	return s != "" && utf8.ValidString(s) && utf8.RuneCountInString(s) <= maxCharacters
}
func noDuplicates(data []byte) error {
	d := json.NewDecoder(bytes.NewReader(data))
	var walk func() error
	walk = func() error {
		t, e := d.Token()
		if e != nil {
			return e
		}
		start, ok := t.(json.Delim)
		if !ok {
			return nil
		}
		switch start {
		case '{':
			seen := map[string]bool{}
			for d.More() {
				key, e := d.Token()
				if e != nil {
					return e
				}
				k, ok := key.(string)
				if !ok || seen[k] {
					return errors.New("duplicate or invalid BCAST JSON key")
				}
				seen[k] = true
				if e = walk(); e != nil {
					return e
				}
			}
		case '[':
			for d.More() {
				if e := walk(); e != nil {
					return e
				}
			}
		default:
			return errors.New("bad JSON delimiter")
		}
		_, e = d.Token()
		return e
	}
	if e := walk(); e != nil {
		return e
	}
	if _, e := d.Token(); e != io.EOF {
		return errors.New("trailing BCAST JSON")
	}
	return nil
}
func decode(data []byte) (any, error) {
	if err := noDuplicates(data); err != nil {
		return nil, err
	}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	var payload any
	if e := dec.Decode(&payload); e != nil {
		return nil, e
	}
	if _, e := dec.Token(); e != io.EOF {
		return nil, errors.New("trailing BCAST data")
	}
	// Reject null, numbers and booleans, which the 0.2 package/API cannot carry.
	var atoms func(any) error
	atoms = func(v any) error {
		switch x := v.(type) {
		case string:
			return nil
		case []any:
			for _, z := range x {
				if e := atoms(z); e != nil {
					return e
				}
			}
			return nil
		case map[string]any:
			for _, z := range x {
				if e := atoms(z); e != nil {
					return e
				}
			}
			return nil
		default:
			return errors.New("BCAST JSON contains a non-contract atom")
		}
	}
	return payload, atoms(payload)
}
func validate(raw map[string]any, requested string) (*Package, error) {
	bad := func(s string) (*Package, error) { return nil, errors.New("invalid BCAST " + s) }
	if !only(raw, "schema_version", "package_id", "package_version", "publication", "objects") || asString(raw, "schema_version") != "bcast.package/0.2.0" || !expectID(asString(raw, "package_id"), "bcastpkg") || asString(raw, "package_id") != requested {
		return bad("package header or requested identity")
	}
	if v := asString(raw, "package_version"); !validPublicText(v, 128) {
		return bad("package version")
	}
	pub, ok := raw["publication"].(map[string]any)
	if !ok || !only(pub, "publication_id", "namespace", "family", "edition", "revision", "title", "issuing_body") {
		return bad("publication fields")
	}
	for _, key := range []string{"revision", "title", "issuing_body"} {
		if value, present := pub[key]; present {
			v, ok := value.(string)
			if !ok || v == "" {
				return bad("invalid optional publication field")
			}
			max := map[string]int{"revision": 128, "title": 512, "issuing_body": 256}[key]
			if !validPublicText(v, max) {
				return bad("long publication field")
			}
		}
	}
	ns, fam, edition := asString(pub, "namespace"), asString(pub, "family"), asString(pub, "edition")
	if !namespace.MatchString(ns) || len(ns) > 253 || !family.MatchString(fam) || len(fam) > 256 || edition == "" || len(edition) > 128 || strings.TrimSpace(edition) != edition {
		return bad("publication coordinates")
	}
	// Both edition and revision are canonical publication coordinates. An
	// ASCII revision with surrounding whitespace is not a distinct lawful
	// coordinate, even if a caller correctly recomputes its hash.
	if rev := asString(pub, "revision"); rev != "" && strings.TrimSpace(rev) != rev {
		return bad("noncanonical publication revision")
	}
	for _, r := range edition + asString(pub, "revision") {
		if r > 127 {
			return bad("non-ASCII edition/revision pending NFC implementation")
		}
	}
	fields := map[string]any{"namespace": ns, "family": fam, "edition": edition}
	if rev, ok := pub["revision"]; ok {
		fields["revision"] = rev
	}
	pubID, e := id("bcastpub", fields)
	if e != nil {
		return nil, e
	}
	if pubID != asString(pub, "publication_id") {
		return bad("publication digest")
	}
	objects, ok := raw["objects"].([]any)
	if !ok || len(objects) == 0 {
		return bad("objects")
	}
	byID := map[string]map[string]any{}
	coordinates := map[string]bool{}
	list := make([]map[string]any, 0, len(objects))
	for _, element := range objects {
		obj, ok := element.(map[string]any)
		if !ok || !only(obj, "object_id", "kind", "locator", "parent_id", "label", "status", "content", "provenance") {
			return bad("object shape")
		}
		kind, loc, oid := asString(obj, "kind"), asString(obj, "locator"), asString(obj, "object_id")
		if !validObjectKind(kind) || !validPublicText(loc, 512) || !validStatus(asString(obj, "status")) || !expectID(oid, "bcastobj") {
			return bad("object coordinates/status")
		}
		for _, key := range []string{"parent_id", "label"} {
			if value, exists := obj[key]; exists {
				v, ok := value.(string)
				if !ok || v == "" {
					return bad("invalid optional object field")
				}
				if key == "label" && !validPublicText(v, 1024) {
					return bad("long object label")
				}
			}
		}
		if parent, exists := obj["parent_id"]; exists {
			v, ok := parent.(string)
			if !ok || !expectID(v, "bcastobj") {
				return bad("parent ID")
			}
		}
		if content, exists := obj["content"]; exists {
			c, ok := content.(map[string]any)
			if !ok || !only(c, "media_type", "text") || asString(c, "media_type") != "text/plain" {
				return bad("content")
			}
			if _, ok := c["text"].(string); !ok {
				return bad("content text")
			}
		}
		if provenance, exists := obj["provenance"]; exists {
			p, ok := provenance.(map[string]any)
			if !ok || !only(p, "citation", "source_artifact_sha256") || !validPublicText(asString(p, "citation"), 2048) {
				return bad("provenance")
			}
			if v, exists := p["source_artifact_sha256"]; exists {
				digest, ok := v.(string)
				if !ok || !shaID.MatchString(digest) {
					return bad("source hash")
				}
			}
		}
		expected, e := id("bcastobj", map[string]any{"publication_id": pubID, "kind": kind, "locator": loc})
		if e != nil {
			return nil, e
		}
		if expected != oid || byID[oid] != nil || coordinates[kind+"\x00"+loc] {
			return bad("object digest or duplicate")
		}
		byID[oid] = obj
		coordinates[kind+"\x00"+loc] = true
		list = append(list, obj)
	}
	state := map[string]int{}
	for start := range byID {
		if state[start] != 0 {
			continue
		}
		path := []string{}
		cursor := start
		for cursor != "" && state[cursor] == 0 {
			o, ok := byID[cursor]
			if !ok {
				return bad("missing parent")
			}
			state[cursor] = 1
			path = append(path, cursor)
			cursor = asString(o, "parent_id")
		}
		if cursor != "" && state[cursor] == 1 {
			return bad("parent cycle")
		}
		for _, k := range path {
			state[k] = 2
		}
	}
	sorted := append([]map[string]any(nil), list...)
	sort.Slice(sorted, func(i, j int) bool { return asString(sorted[i], "object_id") < asString(sorted[j], "object_id") })
	toHash := map[string]any{}
	for k, v := range raw {
		if k != "package_id" {
			toHash[k] = v
		}
	}
	arr := make([]any, len(sorted))
	for i, v := range sorted {
		arr[i] = v
	}
	toHash["objects"] = arr
	actual, e := id("bcastpkg", toHash)
	if e != nil {
		return nil, e
	}
	if actual != requested {
		return bad("materialization SHA-256")
	}
	return &Package{ID: requested, PublicationID: pubID, Objects: list, raw: raw}, nil
}
func (c *Client) read(ctx context.Context, path string) (any, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.origin+path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	r, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("BCAST GET returned HTTP %d", r.StatusCode)
	}
	if !strings.HasPrefix(strings.ToLower(r.Header.Get("Content-Type")), "application/json") {
		return nil, errors.New("BCAST expected application/json")
	}
	body, e := io.ReadAll(io.LimitReader(r.Body, maxBytes+1))
	if e != nil {
		return nil, e
	}
	if len(body) > maxBytes {
		return nil, errors.New("BCAST HTTP payload too large")
	}
	return decode(body)
}
func (c *Client) Fetch(ctx context.Context, requested string) (*Package, error) {
	if !expectID(requested, "bcastpkg") {
		return nil, errors.New("invalid requested BCAST package ID")
	}
	payload, e := c.read(ctx, "/packages/"+requested)
	if e != nil {
		return nil, e
	}
	raw, ok := payload.(map[string]any)
	if !ok {
		return nil, errors.New("BCAST package must be object")
	}
	return validate(raw, requested)
}

// VerifyNavigation checks the public standard metadata, object and child
// endpoints against a fully hash-verified package, including nested children.
func (c *Client) VerifyNavigation(ctx context.Context, p *Package) error {
	meta, e := c.read(ctx, "/packages/"+p.ID+"/metadata")
	if e != nil {
		return e
	}
	m, ok := meta.(map[string]any)
	if !ok || !only(m, "schema_version", "package_id", "package_version", "publication") {
		return errors.New("BCAST metadata shape mismatch")
	}
	expected := map[string]any{}
	for _, k := range []string{"schema_version", "package_id", "package_version", "publication"} {
		expected[k] = p.raw[k]
	}
	if !equal(meta, expected) {
		return errors.New("BCAST metadata differs from full package")
	}
	for _, obj := range p.Objects {
		oid := asString(obj, "object_id")
		path := "/packages/" + p.ID + "/objects/" + oid
		observed, e := c.read(ctx, path)
		if e != nil {
			return e
		}
		if !equal(observed, obj) {
			return errors.New("BCAST HTTP object substituted")
		}
		children, e := c.read(ctx, path+"/children")
		if e != nil {
			return e
		}
		expectedChildren := p.Children(oid)
		arr := make([]any, len(expectedChildren))
		for i, v := range expectedChildren {
			arr[i] = v
		}
		if !equal(children, arr) {
			return errors.New("BCAST HTTP child navigation differs")
		}
	}
	return nil
}
func equal(a, b any) bool {
	var x, y bytes.Buffer
	if canonical(&x, a) != nil || canonical(&y, b) != nil {
		return false
	}
	return bytes.Equal(x.Bytes(), y.Bytes())
}
