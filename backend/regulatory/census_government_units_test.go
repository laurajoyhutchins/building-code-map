package regulatory

import "testing"

func TestCensusCountyMapsToCountyEquivalentWithoutInventingGeometry(t *testing.T) {
	record := CensusGovernmentUnitRecord{
		CensusID:                 "100001",
		UnitName:                 "COUNTY OF AUTAUGA",
		UnitType:                 "1 - COUNTY",
		PoliticalCodeDescription: "COUNTY",
		State:                    "AL",
		FIPSState:                "01",
		FIPSCounty:               "001",
		FIPSPlace:                "99001",
		Population:               "60342",
		PopulationSourceYear:     "2023",
		Active:                   "Y",
	}

	entity, err := GovernmentalEntityFromCensusUnit(record, "src:census-government-units-2025")
	if err != nil {
		t.Fatal(err)
	}
	if entity.EntityID != "gov:us-al:county_equivalent:census-government-unit:100001" {
		t.Fatalf("entity id=%q", entity.EntityID)
	}
	if entity.EntityType != EntityTypeCountyEquivalent {
		t.Fatalf("entity type=%q", entity.EntityType)
	}
	if entity.Classification != ClassificationUnresolved {
		t.Fatalf("classification=%q", entity.Classification)
	}
	if len(entity.Geographies) != 0 {
		t.Fatalf("Census government-unit FIPS fields are not geometry: %#v", entity.Geographies)
	}
	if entity.HistoricalGeographyStatus != HistoricalGeographyUnavailable {
		t.Fatalf("historical geography status=%q", entity.HistoricalGeographyStatus)
	}
	if entity.Population == nil || entity.Population.Count != 60342 || entity.Population.SourceYear != "2023" {
		t.Fatalf("population=%#v", entity.Population)
	}
	if len(entity.SourceClassifications) != 1 || entity.SourceClassifications[0].UnitType != "1 - COUNTY" {
		t.Fatalf("source classifications=%#v", entity.SourceClassifications)
	}
	wantIdentifiers := map[string]string{
		"census-government-unit-pid6": "100001",
		"government-units-fips-state": "01",
		"government-units-fips-county": "001",
		"government-units-fips-place": "99001",
	}
	for _, identifier := range entity.ExternalIdentifiers {
		if want, ok := wantIdentifiers[identifier.Scheme]; ok {
			if identifier.Value != want {
				t.Fatalf("identifier %s=%q want %q", identifier.Scheme, identifier.Value, want)
			}
			delete(wantIdentifiers, identifier.Scheme)
		}
	}
	if len(wantIdentifiers) != 0 {
		t.Fatalf("missing identifiers=%#v", wantIdentifiers)
	}
}

func TestCensusCountyClassifiedAsCityMapsToIndependentCity(t *testing.T) {
	entity, err := GovernmentalEntityFromCensusUnit(CensusGovernmentUnitRecord{
		CensusID:                 "va-city",
		UnitName:                 "CITY OF EXAMPLE",
		UnitType:                 "1 - COUNTY",
		PoliticalCodeDescription: "CITY",
		State:                    "VA",
		FIPSState:                "51",
		FIPSCounty:               "760",
		Active:                   "Y",
	}, "src:census-government-units-2025")
	if err != nil {
		t.Fatal(err)
	}
	if entity.EntityType != EntityTypeIndependentCity {
		t.Fatalf("entity type=%q", entity.EntityType)
	}
}

func TestCensusConsolidatedGovernmentDoesNotCollapseToMunicipality(t *testing.T) {
	entity, err := GovernmentalEntityFromCensusUnit(CensusGovernmentUnitRecord{
		CensusID:                 "denver",
		UnitName:                 "CITY AND COUNTY OF DENVER",
		UnitType:                 "2 - MUNICIPAL",
		PoliticalCodeDescription: "CITY AND COUNTY",
		State:                    "CO",
		FIPSState:                "08",
		FIPSCounty:               "031",
		FIPSPlace:                "20000",
		Active:                   "Y",
	}, "src:census-government-units-2025")
	if err != nil {
		t.Fatal(err)
	}
	if entity.EntityType != EntityTypeConsolidatedGovernment {
		t.Fatalf("entity type=%q", entity.EntityType)
	}
}

func TestCensusTownshipMapsToMinorCivilDivision(t *testing.T) {
	entity, err := GovernmentalEntityFromCensusUnit(CensusGovernmentUnitRecord{
		CensusID:                 "ct-town",
		UnitName:                 "TOWN OF AVON",
		UnitType:                 "3 - TOWNSHIP",
		PoliticalCodeDescription: "TOWN",
		State:                    "CT",
		FIPSState:                "09",
		FIPSCounty:               "110",
		FIPSPlace:                "02060",
		Active:                   "Y",
	}, "src:census-government-units-2025")
	if err != nil {
		t.Fatal(err)
	}
	if entity.EntityType != EntityTypeMinorCivilDivision {
		t.Fatalf("entity type=%q", entity.EntityType)
	}
}

func TestCensusInactiveUnitIsExplicitlyInactiveRatherThanUnresolved(t *testing.T) {
	entity, err := GovernmentalEntityFromCensusUnit(CensusGovernmentUnitRecord{
		CensusID:                 "inactive-town",
		UnitName:                 "FORMER TOWN",
		UnitType:                 "2 - MUNICIPAL",
		PoliticalCodeDescription: "TOWN",
		State:                    "AL",
		FIPSState:                "01",
		Active:                   "N",
	}, "src:census-government-units-2025")
	if err != nil {
		t.Fatal(err)
	}
	if entity.LegalStatus != LegalStatusInactive || entity.Classification != ClassificationInactive {
		t.Fatalf("legal=%q classification=%q", entity.LegalStatus, entity.Classification)
	}
}

func TestCensusUnitRejectsUnknownTypeAndInvalidActivity(t *testing.T) {
	base := CensusGovernmentUnitRecord{
		CensusID:  "fixture",
		UnitName:  "EXAMPLE",
		UnitType:  "2 - MUNICIPAL",
		State:     "AL",
		FIPSState: "01",
		Active:    "Y",
	}
	badType := base
	badType.UnitType = "9 - UNKNOWN"
	if _, err := GovernmentalEntityFromCensusUnit(badType, "src:fixture"); err == nil {
		t.Fatal("expected unknown unit type failure")
	}
	badActivity := base
	badActivity.Active = "?"
	if _, err := GovernmentalEntityFromCensusUnit(badActivity, "src:fixture"); err == nil {
		t.Fatal("expected invalid active flag failure")
	}
}
