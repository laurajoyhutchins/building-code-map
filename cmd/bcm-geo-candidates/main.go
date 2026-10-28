// Command bcm-geo-candidates checks an immutable, private candidate-observation
// export against the public geocoder's source-only intake boundary.
// It never produces an accepted jurisdiction, legal claim, or project verdict.
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/laurajoyhutchins/building-code-map/geocoder"
)

func main() {
	input := flag.String("input", "", "path to private source-only geographic observation JSONL")
	sha := flag.String("sha256", "", "trusted immutable SHA-256 of exact input file")
	scope := flag.String("scope", "", "expected US state scope, e.g. US-CO")
	receipt := flag.String("source-receipt", "", "expected immutable publisher observation receipt ID")
	rows := flag.Int("rows", 0, "expected complete source observation count")
	address := flag.String("address", "", "optional civic address, produces source-name candidates only")
	place := flag.String("place-code", "", "optional five-digit PLACE code, no accepted government")
	flag.Parse()
	if flag.NArg() != 0 || *input == "" || (*address != "" && *place != "") {
		die(errors.New("required input, or mutually exclusive optional lookup, is invalid"))
	}
	file, err := os.Open(*input)
	if err != nil {
		die(err)
	}
	defer file.Close()
	idx, err := geocoder.ReadPlaceCandidates(file, geocoder.PlaceCandidateSource{
		SHA256: *sha, ScopeID: *scope, SourceReceiptID: *receipt, ExpectedRows: *rows,
	})
	if err != nil {
		die(err)
	}
	output := struct {
		Summary geocoder.PlaceCandidateSummary `json:"summary"`
		Lookup  *geocoder.PlaceCandidateMatch  `json:"lookup,omitempty"`
	}{Summary: idx.Summary()}
	switch {
	case *address != "":
		parsed, err := geocoder.ParseAddress(*address)
		if err != nil {
			die(err)
		}
		value := idx.ForAddress(parsed)
		output.Lookup = &value
	case *place != "":
		if len(*place) != 5 || strings.Trim(*place, "0123456789") != "" {
			die(errors.New("place-code must be exactly five digits"))
		}
		value := idx.ByPlaceCode(*place)
		output.Lookup = &value
	}
	if err := json.NewEncoder(os.Stdout).Encode(output); err != nil {
		die(err)
	}
}

func die(err error) {
	fmt.Fprintln(os.Stderr, "source-only candidate intake rejected: "+err.Error())
	os.Exit(1)
}
