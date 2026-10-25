package geocoder

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

// The source records are deliberately fictional and may be published. Runtime
// lookup must never infer a governing jurisdiction from a point alone.
func syntheticGeocoder(t *testing.T) *SQLiteService {
	t.Helper()
	dir := t.TempDir()
	points := "source_record_id,address_number,street,city,state,postal_code,longitude,latitude,matched_address\n" +
		"a,101,EXAMPLE ST,SAMPLE CITY,CO,80202,-104.9000,39.7000,101 Example St Sample City CO 80202\n" +
		"b,100,EXAMPLE ST,SAMPLE CITY,CO,80202,-104.9100,39.7100,100 Example St Sample City CO 80202\n" +
		"c,100,EXAMPLE ST,SAMPLE CITY,CO,80202,-104.9200,39.7200,100 Example St Sample City CO 80202\n"
	ranges := "source_record_id,from_number,to_number,parity,street,city,state,postal_code,from_longitude,from_latitude,to_longitude,to_latitude\n" +
		"range-a,200,298,E,EXAMPLE ST,SAMPLE CITY,CO,80202,-104.9000,39.7000,-104.8000,39.7000\n"
	pointFile := filepath.Join(dir, "synthetic-points.csv")
	rangeFile := filepath.Join(dir, "synthetic-ranges.csv")
	if err := os.WriteFile(pointFile, []byte(points), 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(rangeFile, []byte(ranges), 0600); err != nil {
		t.Fatal(err)
	}
	snapshot := filepath.Join(dir, "snapshot.sqlite")
	if err := BuildSnapshot(BuildOptions{
		OutputPath:       snapshot,
		AddressPointsCSV: pointFile,
		StreetRangesCSV:  rangeFile,
		SourceName:       "synthetic-only",
		SourceVintage:    "2026-10-25",
	}); err != nil {
		t.Fatal(err)
	}
	service, err := Open(snapshot)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = service.Close() })
	return service
}

func TestCanonicalBCMSyntheticGeocoderPreservesPrecisionAndSource(t *testing.T) {
	service := syntheticGeocoder(t)
	for _, tc := range []struct {
		address string
		status Status
		precision Precision
		sourceID string
	}{
		{"101 Example Street, Sample City, CO 80202", StatusMatched, PrecisionAddressPoint, "a"},
		{"100 Example St, Sample City, CO 80202", StatusAmbiguous, "", ""},
		{"220 Example St, Sample City, CO 80202", StatusMatched, PrecisionInterpolated, "range-a"},
		{"300 Example St, Sample City, CO 80202", StatusNotFound, "", ""},
		{"101 Example St, Other City, CO 80202", StatusNotFound, "", ""},
		{"101 Example St, Sample City, CO 99999", StatusNotFound, "", ""},
	} {
		t.Run(tc.address,func(t *testing.T){
			out,err:=service.Geocode(context.Background(),Query{Address:tc.address})
			if err!=nil { t.Fatal(err) }
			if out.Status!=tc.status { t.Fatalf("got %s, want %s: %#v",out.Status,tc.status,out) }
			if tc.status!=StatusMatched {
				if out.Selected!=nil {t.Fatalf("selected incompatible/ambiguous point: %#v",out)}
				if tc.status==StatusAmbiguous && len(out.Candidates)!=2 {t.Fatalf("lost disambiguation: %#v",out)}
				return
			}
			if out.Selected==nil || out.Selected.Precision!=tc.precision || out.Selected.SourceRecordID!=tc.sourceID {
				t.Fatalf("source identity/precision lost: %#v",out)
			}
			if out.Selected.Source!="synthetic-only" || out.Selected.SourceVintage!="2026-10-25" || out.Selected.RankingPolicyVersion=="" {
				t.Fatalf("unproven source/ranking: %#v",out.Selected)
			}
			if tc.precision==PrecisionInterpolated {
				if out.Selected.Interpolation==nil || out.Selected.Interpolation.SourceRangeID=="" || len(out.Warnings)==0 {
					t.Fatalf("street-range approximation was presented as exact: %#v",out)
				}
			}else if out.Selected.Interpolation!=nil {
				t.Fatalf("address point mislabeled as interpolation: %#v",out)
			}
		})
	}
}

func TestPublicGeocoderRejectsBadInputsAndCorruptSnapshot(t *testing.T){
	if err:=BuildSnapshot(BuildOptions{OutputPath:filepath.Join(t.TempDir(),"x.sqlite")});err==nil {
		t.Fatal("builder accepted empty source")
	}
	path:=filepath.Join(t.TempDir(),"corrupt.sqlite")
	if err:=os.WriteFile(path,[]byte("not a SQLite database"),0600);err!=nil{t.Fatal(err)}
	if svc,err:=Open(path);err==nil {
		_ = svc.Close()
		t.Fatal("unverified SQLite snapshot opened")
	}
}
