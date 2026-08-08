package engine

type ErrorCode string

const (
	ErrorInvalidQuery                 ErrorCode = "invalid_query"
	ErrorInvalidCoordinates           ErrorCode = "invalid_coordinates"
	ErrorAddressNotFound              ErrorCode = "address_not_found"
	ErrorAddressAmbiguous             ErrorCode = "address_ambiguous"
	ErrorBoundaryAmbiguous            ErrorCode = "boundary_ambiguous"
	ErrorOutsideSupportedCoverage     ErrorCode = "outside_supported_coverage"
	ErrorRegulatoryCatalogUnavailable ErrorCode = "regulatory_catalog_unavailable"
	ErrorRegulatoryProfileMissing     ErrorCode = "regulatory_profile_missing"
	ErrorDataBundleInvalid            ErrorCode = "data_bundle_invalid"
	ErrorInternal                     ErrorCode = "internal_error"
)

type EngineError struct {
	Code      ErrorCode      `json:"code"`
	Message   string         `json:"message"`
	Details   map[string]any `json:"details,omitempty"`
	Retryable bool           `json:"retryable"`
}

func (err EngineError) Error() string { return err.Message }
