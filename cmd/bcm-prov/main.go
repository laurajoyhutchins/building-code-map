// bcm-prov projects one canonical BCM evidence claim as PROV-O JSON-LD.
// It never makes an evidence admission or a project-verification decision.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/laurajoyhutchins/building-code-map/engine"
)

const maxClaimBytes = 1 << 20

func run(input io.Reader, output io.Writer) error {
	content, err := io.ReadAll(io.LimitReader(input, maxClaimBytes+1))
	if err != nil {
		return fmt.Errorf("read claim: %w", err)
	}
	if len(content) > maxClaimBytes {
		return fmt.Errorf("claim input exceeds 1 MiB")
	}
	decoder := json.NewDecoder(bytes.NewReader(content))
	decoder.DisallowUnknownFields()
	var claim engine.EvidenceClaim
	if err := decoder.Decode(&claim); err != nil {
		return fmt.Errorf("decode evidence claim: %w", err)
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return fmt.Errorf("input must contain exactly one claim")
	}
	graph, err := engine.ProjectClaimPROV(claim)
	if err != nil {
		return fmt.Errorf("project evidence: %w", err)
	}
	doc, err := graph.JSONLD()
	if err != nil {
		return fmt.Errorf("export PROV-O: %w", err)
	}
	_, err = output.Write(doc)
	return err
}

func main() {
	if err := run(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
