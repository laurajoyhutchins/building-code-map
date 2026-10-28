# Candidate municipal sources in the public geocoder

BCM's canonical `geocoder` now accepts a **trusted, content-pinned candidate-only
observation feed** through `ReadPlaceCandidates`. It is the first narrow consumer
for private jurisdiction-data exports. The private producer owns original
publisher bytes, JSON Schema + ODCS source contracts, rights receipts,
W3C PROV source graph and migration closure.

This public library and the optional `bcm-geo-candidates` offline probe contain
**no maintained private corpus**. All checked-in tests are synthetic. The
consumer checks an independently pinned SHA-256, source scope, receipt ID and
complete row count, then strictly parses geographic-observation schema v1
`government_identity` records. It rejects duplicate JSON keys, unrecognized
fields including `government_entity_id`, unapproved review states, missing
provenance, source-row alias replacement, and candidate-to-reviewed promotion.

`ByPlaceCode` and `ForAddress(ParsedAddress)` are **source candidates only**.
In particular, `ForAddress` reuses civic-address parsing and compares the
supplied city name with publisher municipality names: **there is no
point-in-polygon or dated legal boundary test**. Multiple same-name source
records are returned as ambiguous, not deduplicated. Even a unique place-code
match cannot establish a legal government identity or project authority.

Use the probe from a private checkout only, with independently obtained
commit/content pins:

```text
go run ./cmd/bcm-geo-candidates \
  -input /private/observations.jsonl \
  -sha256 EXPECTED_SOURCE_SHA256_FROM_TRUSTED_LOCK \
  -scope US-CO \
  -source-receipt EXPECTED_RECEIPT_ID \
  -rows EXPECTED_COMPLETE_ROWS \
  -address '100 Example St, Exampletown, CO'
```

Output contains a `summary` and optional `lookup` with
`legal_verdict_authorized=false` and `authority_status=not_established`.
It is neither a deployed private provider nor a `Resolution`; it does **not**
call or replace `engine.ProjectVerifier`. Real source-derived public-safe
verification is performed in private CI against this commit-pinned consumer
with the complete migrated source feed; that independent consumer check is
not proof of RSO legal evidence or BCM #90 address-level geographic authority.

The full production cutover still requires a separately reviewed dated
site-to-jurisdiction authority observation, polygon/topology source custody,
RSO adopted-code evidence, and a successful call through BCM's one
`POST /v1/project-code-basis` verdict path.
