# Building Code Map Shared Visual System Design

**Date:** 2026-07-21  
**Status:** Design approved in conversation; awaiting written-spec review  
**Scope:** Building Code Map React frontend only

## Summary

Building Code Map will adopt the colors, typography, and interface language of the Laura J Hutchins LLC website while remaining a purpose-built GIS and regulatory-resolution application.

The approved direction is a **map-native sibling**: phthalo blue, green, pink, mint-white, and near-black colors; editorial serif and monospace typography; compact framed workspaces; two-pixel borders; small radii; hard offset shadows; a grid-backed canvas; and first-class light and dark themes. The map and operational controls remain denser and more restrained than the main website.

## Goals

1. Establish a recognizable shared design language across the LLC website and Building Code Map.
2. Preserve professional GIS usability and information density.
3. Make authority, code, amendment, uncertainty, and source evidence easier to scan.
4. Provide complete, accessible light and dark themes.
5. Make mobile use answer-first after a lookup.
6. Keep visual changes isolated from jurisdiction-resolution, data-loading, and map-selection logic.

## Non-goals

- Rewriting location-to-authority resolution logic.
- Changing backend APIs or data contracts.
- Replacing MapLibre or the current basemap provider.
- Creating a cross-repository component or token package in this change.
- Reproducing the LLC website layout exactly.
- Adding decorative animation or simulated-desktop behavior that reduces GIS clarity.
- Changing the conservative treatment of unresolved, ambiguous, or conflicting evidence.
- Adding recent-location or example-location product features as part of the redesign.

## Approved design decisions

### Product relationship

Building Code Map is a map-native sibling of the LLC website. It shares visual tokens, typography, border and shadow language, title-bar treatment, theme behavior, and interaction details. Its layout and density remain specific to a regulatory GIS workspace.

### Chrome density

Use moderate window-frame chrome. Panels are visibly framed and titled, but the central map is not nested inside additional decorative frames. Desktop title bars are compact; mobile section headers are thinner.

### Typography

Use the LLC website's editorial serif stack for authority names, jurisdiction names, adopted-code and amendment titles, result summaries, explanatory headings, and body explanations.

Use the monospace UI stack for navigation, form labels and controls, coordinates, layer names, source identifiers, timestamps, status metadata, technical attributes, and compact operational copy.

### Color roles

Blue and green are the primary map and data colors. Pink is restricted to selected evidence, review-required states, exceptional emphasis, and a subtle selected-feature cue.

Pink does not replace semantic success, warning, conflict, or error treatments. Every state has a text label and a non-color cue.

### Map treatment

Keep the basemap visually subdued so jurisdiction boundaries and selections dominate. Streets, labels, geographic context, and attribution remain legible.

### Mobile behavior

Before a lookup, the query is primary and the map is secondary. After a lookup, the authority result and evidence move ahead of the map. The map is collapsed by default behind a `View jurisdiction on map` control and expands without losing the resolved point, selected boundaries, active layers, result context, or return scroll position.

## Visual system

### Theme tokens

Create `src/theme.css` as the authoritative token layer. Components consume semantic tokens rather than raw colors.

Shared tokens cover:

- canvas and alternate canvas;
- base, raised, and muted surfaces;
- primary and soft ink;
- strong and soft lines;
- blue and blue-strong;
- green and green-soft;
- pink and pink-soft;
- focus and selection;
- hard and soft shadows;
- title-bar gradient endpoints;
- canvas grid and glow values.

Map-specific aliases are:

- `--boundary-primary`;
- `--boundary-secondary`;
- `--boundary-tertiary`;
- `--boundary-selected`;
- `--point-resolved`;
- `--evidence-selected`;
- `--evidence-review`;
- `--map-overlay-surface`;
- `--map-overlay-ink`.

Map aliases have explicit light and dark values chosen for contrast.

### Typography tokens

Use three font roles:

- display/editorial serif;
- body serif;
- monospace UI.

The map application uses a compact scale. Product and panel titles remain materially smaller than the main website's editorial hero headings.

### Shape and elevation

- `--radius-sm: 2px`.
- `--radius-md: 5px`.
- Primary panel borders are two pixels.
- Framed workspaces use the hard offset shadow; large workspaces may also use the soft shadow.
- Pills are limited to true state chips and compact tags.
- Operational controls are rectangular.
- Glassmorphism, backdrop blur, large radii, and diffuse SaaS cards are removed.

### Canvas

Use the shared grid-backed canvas and restrained blue/green glows outside the map. The grid fades beneath dense content and does not overlay the map. Dark mode uses near-black canvas and surfaces rather than tinted midtones.

## Desktop workspace

Desktop preserves a three-part GIS workspace.

### Application header

The compact header contains:

- product mark;
- `BUILDING CODE MAP`;
- `JURISDICTION EXPLORER`;
- current data/readiness state;
- theme toggle;
- a `Laura J Hutchins LLC` link.

The current three statistic cards are replaced with one status strip containing layer count, cached record count, and refresh state.

### Left workspace

The left column contains these exact framed sections:

- `QUERY.LOCATION`: address or coordinate query and resolution controls;
- `DATA.STATUS`: readiness and refresh state;
- `MAP.LAYERS`: layer toggles.

Labels and controls use monospace. Explanatory guidance uses the body serif. Layer toggles remain dense and keyboard accessible.

### Center workspace

The center column is `MAP.VIEW` and remains the visual anchor.

Requirements:

- one compact shared title bar around the workspace;
- minimal decoration inside the map canvas;
- quiet basemap;
- blue and green jurisdiction boundaries;
- strongest contrast for the selected boundary;
- subtle pink cue linking the selection to associated evidence;
- preserved OpenStreetMap attribution;
- readable overlays in both themes.

The title bar shows the selected layer and resolved or cursor coordinates when those values are available.

### Right workspace

The right column is `AUTHORITY.RECORD` and contains:

- selected jurisdiction;
- authority chain;
- adopted codes;
- amendments;
- resolution status and uncertainty;
- evidence and provenance;
- expandable technical attributes.

Authority names, code titles, and result summaries use editorial serif. Source IDs, dates, feature IDs, and evidence metadata use monospace. Expanding evidence does not rerender the map or reset selection.

## Tablet workspace

At tablet widths, use two columns:

- query, status, and layers on the left;
- map on the right;
- authority record below both columns at full width.

The transition preserves selected feature, resolved location, active layers, expanded evidence sections, map center and zoom, and scroll position. The authority record is not placed in a drawer at tablet widths.

## Mobile task flow

### Before lookup

Order:

1. compact product header;
2. `QUERY.LOCATION`;
3. collapsed `MAP.VIEW`;
4. `MAP.LAYERS` disclosure;
5. `DATA.STATUS` disclosure.

The query is immediately usable without scrolling through map controls.

### After lookup

Order:

1. `AUTHORITY.RESULT` summary;
2. adopted codes and amendments;
3. confidence, ambiguity, conflict, or unresolved conditions;
4. source evidence grouped by authority or document;
5. collapsed `MAP.VIEW` control;
6. expandable technical details;
7. layer and data-status disclosures.

The collapsed control reads `View jurisdiction on map`; the expanded control reads `Close jurisdiction map`.

Expanding and closing the map preserves the resolved point, selected jurisdiction boundaries, active layers, evidence context, map center and zoom, and the user's return scroll position.

### Mobile panel treatment

Desktop title bars become thinner section headers. Technical metadata is hidden in accessible disclosures by default. Interactive targets have a minimum 44-by-44 CSS-pixel hit area even when the visible control is compact.

## Component boundaries

### Existing responsibilities retained

- `App`: loading, application state, feature selection, and composition.
- `MapStage`: map rendering and map interaction.
- `ResolutionPanel`: query and resolution behavior.
- `LayerToggleList`: layer selection.
- `FeatureInspector`: selected-record and evidence display.
- API and type modules: transport and domain contracts.

### Presentation components

Add or refine these narrow presentation units:

- `AppHeader`;
- `ThemeToggle`;
- `WorkspaceFrame`;
- `WorkspaceTitlebar`;
- `CompactStatusStrip`;
- `MobileSection`;
- `MobileMapDisclosure`;
- semantic status component or helper.

These components do not fetch APIs and do not own regulatory or map-selection state.

### CSS organization

Use this explicit ownership:

- `src/theme.css`: shared and map-specific tokens plus light/dark theme definitions;
- `src/styles.css`: reset, base typography, reusable frames and controls, application layout, map workspace styles, responsive rules, and utilities;
- `src/resolution.css`: query, result, authority, evidence, and resolution-state styles.

`src/main.tsx` imports these files in the order `theme.css`, MapLibre CSS, `styles.css`, `resolution.css`.

The implementation copies and adapts stable LLC token values with a source note in the design documentation. A shared package is deferred.

## Theme behavior

- Support explicit light and dark themes.
- On first visit, use `prefers-color-scheme`.
- Persist the user's explicit choice in local storage.
- Apply the theme at the document root before React paints to minimize flash.
- Give the theme control an accessible name, visible focus, and textual state.
- Tests control media-query and storage state explicitly.

## Map styling and contrast

MapLibre layer colors come from a centralized semantic style adapter. Components do not scatter unrelated hard-coded colors.

Requirements:

- selected geometry is visible over both themes;
- overlapping jurisdiction families remain distinguishable;
- fill opacity does not obscure important labels;
- line width or dash pattern supplements color where layers overlap;
- hover and selected states differ visibly;
- map controls and overlays remain readable;
- attribution remains visible;
- pink is not the sole distinction between operational layer families.

## State and error treatment

Provide consistent framed treatments for:

- initial loading;
- runtime snapshot unavailable;
- successful resolution;
- partial resolution;
- local record required;
- ambiguous resolution;
- conflicting evidence;
- insufficient evidence;
- API or data-loading failure.

Every state includes a textual label. Warning and error states use semantic wording plus border and icon treatment. The application continues to expose uncertainty rather than infer unsupported answers.

## Accessibility

- High-contrast focus treatment uses the shared focus token.
- Query, layers, evidence disclosures, theme toggle, and mobile map disclosure are keyboard operable.
- Disclosures use correct `aria-expanded` and control relationships.
- Opening and closing panels retains or intentionally moves focus.
- Status is not communicated solely through color.
- Text remains readable at browser zoom.
- Mobile hit areas meet the 44-by-44 CSS-pixel requirement.
- Reduced-motion preferences disable nonessential transitions.
- Both themes preserve readable text, controls, map overlays, and focus indicators.

## Responsive state preservation

Resizing or rotating preserves:

- active query result;
- selected feature;
- active layer set;
- theme;
- map center and zoom;
- expanded mobile map state;
- evidence disclosure state;
- meaningful scroll position.

Use CSS reflow for layout. JavaScript viewport behavior is limited to the approved mobile result ordering, disclosures, and collapsed-map interaction. Shared domain state remains above responsive presentation branches so components do not reset during reflow.

## Testing and verification

### Component and behavior tests

Add tests for:

- first-visit system-theme selection;
- persisted explicit theme selection;
- theme-toggle accessibility and state;
- mobile pre-query ordering;
- mobile post-query result ordering;
- mobile map expansion and collapse;
- preserved result, selection, layers, map viewport, and scroll context when toggling the map;
- semantic status labels for resolved, ambiguous, conflicting, insufficient-evidence, and unavailable states;
- retained resolution, layer-toggle, and map-selection behavior.

### Visual regression coverage

Capture and review:

- desktop light and dark;
- tablet light and dark;
- mobile before lookup;
- mobile after resolved lookup;
- mobile after ambiguous lookup;
- mobile after conflicting lookup;
- mobile after insufficient-evidence lookup;
- mobile with the result map expanded;
- unavailable-snapshot state in light and dark.

Fixtures are deterministic and do not use live external services.

### Accessibility checks

Verify:

- keyboard path through header, query, layers, result, evidence, and map disclosure;
- visible focus in both themes;
- disclosure semantics and focus behavior;
- text and control contrast;
- status meaning without color;
- reduced-motion behavior;
- mobile hit-area dimensions.

### Baseline repository checks

The implementation passes:

```bash
pnpm check
pnpm lint
pnpm format:check
pnpm test
pnpm build
```

Backend behavior does not change. Before merge, also run the repository's backend test and lint commands required by the full verification contract.

## Acceptance criteria

The work is complete when:

1. Building Code Map reads as a sibling of the LLC website in both themes.
2. Large rounded glass cards and diffuse SaaS styling are replaced by compact framed workspaces.
3. Desktop retains a usable three-part query/map/record workspace.
4. Tablet uses the two-column layout with the authority record below.
5. Mobile prioritizes the query before lookup and the authority result after lookup.
6. The mobile map is collapsed after a result and expands without losing state or context.
7. Blue and green dominate map/data visualization; pink remains restrained.
8. The basemap stays quiet while labels and attribution remain readable.
9. Authority and code content uses editorial serif; operational metadata and controls use monospace.
10. Theme preference persists and first visit respects system preference.
11. Loading, unavailable, resolved, partial, local-record-required, ambiguous, conflicting, insufficient-evidence, and failure states are explicit and accessible.
12. Existing resolution, layer selection, map selection, and evidence behavior has no regression.
13. Automated checks, accessibility checks, and required responsive visual review pass.

## Implementation sequence

1. Introduce and test theme state and semantic tokens.
2. Add frame, title-bar, header, theme-toggle, and status-strip presentation components.
3. Restyle desktop layout and panels without changing domain behavior.
4. Centralize MapLibre semantic colors and validate both themes.
5. Implement tablet reflow.
6. Implement the mobile pre-query and post-query flow with map disclosure.
7. Refine state, evidence, and technical-detail treatments.
8. Add behavior, accessibility, and visual-regression coverage.
9. Run the complete verification suite and review all viewport/theme captures.

## Risks and mitigations

### Map legibility in dark mode

Use theme-specific map aliases, controlled fill opacity, line patterns, and visual fixtures covering overlapping boundaries.

### Decorative chrome consuming workspace

Keep title bars compact, avoid nested frames, and use thinner mobile section headers.

### Responsive remounting or state loss

Keep shared state above responsive branches, prefer CSS reflow, and test resize, rotation, and map-disclosure interactions.

### Cross-repository style drift

Document the source and adaptation, align semantic names where useful, and reconsider a shared package only after both applications demonstrate stable repeated needs.

### Pink becoming semantically ambiguous

Restrict pink to selection and review emphasis while retaining explicit labels and distinct success, warning, and error treatments.

## Deferred follow-up

After this implementation stabilizes, evaluate whether the LLC website and Building Code Map have enough repeated, stable primitives to justify a small shared design-token package. That decision does not block this work.
