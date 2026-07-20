# Populate a State Report Template

## Goal

Turn [`reports/_template.md`](../../reports/_template.md) into a state-specific report that is source-backed, structurally complete, and safe to mark `draft` or `partially_verified`.

## Prerequisites

- You know the target state name, postal abbreviation, and report filename slug.
- You have the official state sources needed to support at least the core authority and code-adoption fields.
- You are prepared to leave unresolved items explicit instead of guessing.

## Steps

1. Copy [`reports/_template.md`](../../reports/_template.md) to the target state file in `reports/`, using the lowercase hyphenated state slug such as `new-york.md` or `south-carolina.md`.
2. Replace the frontmatter placeholders first:
   - Set `state.state_id`, `state.name`, and `state.abbreviation`.
   - Set `report.report_id` to `state-report:usa-xx`.
   - Set `report.last_updated` to today's date.
   - Keep `report.status` conservative: use `draft` until the body is source-backed, then upgrade to `partially_verified` only after a validation pass.
   - Leave `last_verified`, `reviewed_by`, confidence values, and open-question counts explicit instead of inventing certainty.
3. Build the source registry before filling narrative sections:
   - Add every official statute, regulation, agency page, register notice, or amendment publication to section 8.
   - Mint stable `src:usa-xx:...` identifiers and use them consistently everywhere else in the report.
   - Add caveats when a source is a courtesy copy, unofficial HTML, OCR-heavy PDF, or otherwise imperfect.
4. Populate section 2 with the authority structure:
   - Identify the primary statewide building-code authority first.
   - Fill specialized authorities only where a distinct board, agency, or program is actually supported by sources.
   - Use `unknown`, `unresolved`, or `none` where the authority was not verified in this pass.
5. Populate section 3 with current statewide code adoptions:
   - Fill the adoption matrix row by row for each code family that the state actually administers.
   - Keep adoption date, effective date, operative date, and mandatory date separate. Do not collapse them into one field.
   - If a code family was not verified, keep the row explicit rather than deleting it.
   - Add normalized adoption records only for facts you can support from the cited sources.
6. Fill sections 4 through 7 only to the level the evidence supports:
   - Capture date rules and transition logic when the state distinguishes permit-date, issuance-date, grace-period, or concurrency rules.
   - Record state amendments and local-amendment rules only when the legal basis or publication path is actually known.
   - Keep local enforcement separate from local amendment authority; they are not the same question.
   - Use plain-English notes to explain hybrid or unusual authority models.
7. Write the executive summary after the detailed sections:
   - Summarize the authority model, statewide status, local enforcement model, local amendment posture, and pending changes from the evidence already captured below.
   - Set `Production readiness` conservatively.
   - Use the key-findings table to surface the most important supported conclusions with confidence and source IDs.
8. Finish sections 9 through 11 as a QA pass:
   - Record field-level verification only for the highest-value fields you actually checked.
   - Mark validation rows `pass` or `fail` based on the current file, not the intended future state.
   - Add open issues for missing fire-code authority, unresolved local-amendment scope, missing dates, or similar evidence gaps.

## Result

The state report is no longer a stub: placeholders are replaced, source IDs resolve, unresolved items stay explicit, and the file is ready for a narrow validation pass.

## Validation

Run these checks before treating the file as ready:

1. Search for leftover template markers:

```powershell
rg -n "\{\{|TODO|TBD|<null>|unknown pending|not verified in this pass" reports\<state>.md
```

2. Spot-check the file head and tail to catch duplicated stub content or partial replacements.
3. Confirm every `src:...` cited in the body exists in section 8.
4. Confirm date fields use ISO format and that impossible sequences are not introduced.
5. Confirm unresolved items are labeled explicitly instead of being silently omitted.
