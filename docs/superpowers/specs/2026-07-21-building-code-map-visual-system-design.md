# Building Code Map Shared Visual System Design

**Date:** 2026-07-21  
**Status:** Design approved in conversation; awaiting written-spec review  
**Scope:** Building Code Map React frontend only

## Summary

Building Code Map will adopt the colors, typography, and interface language of the Laura J Hutchins LLC website while remaining a purpose-built GIS and regulatory-resolution application.

The result should feel like a sibling product, not a copied marketing page. The application will use the LLC website's phthalo blue, green, pink, mint-white, and near-black palette; editorial serif and monospace typography; compact window frames; two-pixel borders; small radii; hard offset shadows; grid-backed canvases; and first-class light and dark themes. The map and operational controls will remain denser and more restrained than the main website.

The approved direction is a **map-native sibling** with moderate interface chrome, quiet basemaps, restrained pink accents, and a mobile task flow that prioritizes the authority result after a search.

## Goals

1. Establish a recognizable shared design language across the LLC website and Building Code Map.
2. Preserve the map application's professional GIS usability and information density.
3. Make authority, code, amendment, uncertainty, and source evidence easier to scan.
4. Provide complete, accessible light and dark themes.
5. Improve mobile use by prioritizing the regulatory answer and its evidence over persistent map chrome.
6. Keep the redesign isolated from jurisdiction-resolution, data-loading, and map-selection behavior.

## Non-goals

- Rewriting the location-to-authority resolution logic.
- Changing backend APIs or data contracts.
- Replacing MapLibre or the current basemap provider.
- Creating a shared cross-repository component package during this change.
- Reproducing the LLC website layout exactly.
- Adding decorative animation, simulated desktop behavior, or novelty controls that reduce GIS clarity.
- Changing the project's conservative treatment of unresolved, ambiguous, or conflicting evidence.

## Approved design decisions

### Product relationship

Building Code Map is a map-native sibling of the main LLC website. It shares visual tokens, typography, border and shadow language, title-bar treatment, theme behavior, and interaction details. Its layout and density remain specific to a regulatory GIS workspace.

### Chrome density

Use moderate window-frame chrome. Panels should be visibly framed and titled, but the central map must not feel boxed in by excessive decoration. The application should look deliberate and distinctive without becoming a desktop-interface pastiche.

### Typography

Use the LLC website's editorial serif stack for:

- authority names;
- jurisdiction names;
- adopted-code and amendment titles;
- major result summaries;
- explanatory headings.

Use the monospace UI stack for:

- navigation;
- form labels and controls;
- coordinates;
- layer names;
- source identifiers;
- timestamps;
- status metadata;
- technical attributes.

Body explanations may use the editorial body serif where reading comfort benefits, while compact operational copy may remain monospace.

### Color roles

Blue and green are the primary map and data colors. Pink is deliberately restrained and is used for:

- selected evidence;
- review-required states;
- exceptional emphasis;
- subtle selected-feature cues.

Pink must not replace semantic success, warning, conflict, or error colors. Color is never the only indicator of state.

### Map treatment

Keep the basemap visually subdued so jurisdiction boundaries and selections dominate. Streets, labels, attribution, and context must remain legible, but boundary overlays should carry the primary visual weight.

### Mobile behavior

Before a lookup, the query is primary and the map is secondary. After a lookup, the authority result and evidence move ahead of the map. The map is collapsed by default behind a clear "View jurisdiction on map" control and expands without losing the resolved point, selected boundaries, or result context.

## Visual system

### Theme tokens

Create a dedicated theme layer based on the LLC website's existing token model. It should expose semantic tokens rather than coupling components to raw color values.

Core shared tokens:

- canvas and alternate canvas;
- base, raised, and muted surfaces;
- primary and soft ink;
- strong and soft lines;
- blue and blue-strong;
- green and green-soft;
- pink and pink-soft;
- focus;
- selection;
- hard and soft shadows;
- title-bar gradient endpoints;
- canvas grid and glow values.

Map-specific aliases:

- `--boundary-primary`;
- `--boundary-secondary`;
- `--boundary-tertiary`;
- `--boundary-selected`;
- `--point-resolved`;
- `--evidence-selected`;
- `--evidence-review`;
- `--map-overlay-surface`;
- `--map-overlay-ink`.

Map aliases may resolve differently in light and dark themes to preserve contrast.

### Typography tokens

Use three font roles:

- display/editorial serif;
- body serif;
- monospace UI.

The application should use smaller type scales than the editorial site. Large marketing-page heading sizes are inappropriate for the map workspace. Product and panel titles should remain compact enough to preserve working area.

### Shape and elevation

- Primary panel borders: two pixels.
- Default corner radius: small, approximately two to five pixels.
- Hard shadow: visible offset shadow based on the strong line color.
- Soft shadow: optional secondary depth for larger framed workspaces.
- Pills: limited to true state chips or compact tags; operational controls should usually be rectangular.
- Glassmorphism and large rounded SaaS cards are removed.

### Canvas

Use the shared grid-backed canvas and restrained blue/green glows outside the map. The grid should fade where it would compete with map detail or text. In dark mode, the canvas must be genuinely dark rather than a tinted midtone.

## Desktop workspace

Desktop preserves a three-part GIS workspace.

### Application header

The header is compact and contains:

- product mark and `BUILDING CODE MAP` name;
- optional `JURISDICTION EXPLORER` descriptor;
- current data/readiness state;
- theme toggle;
- a small link to the main LLC website.

The current three large statistic cards are replaced with a single compact status strip for layer count, cached record count, and refresh state.

### Left workspace: Query and layers

The left column contains:

- address or coordinate query;
- resolution controls;
- data readiness and refresh state;
- layer toggles.

Suggested panel titles:

- `QUERY.ADDRESS` or `QUERY.LOCATION`;
- `DATA.STATUS`;
- `MAP.LAYERS`.

Labels and controls use monospace type. Descriptive guidance may use the body serif. Layer toggles remain dense and keyboard accessible.

### Center workspace: Map view

The center column remains the visual anchor.

Requirements:

- shared framed title-bar treatment around the workspace;
- minimal decorative elements inside the map canvas;
- quiet basemap;
- blue and green jurisdiction boundaries;
- strongest contrast for the selected boundary;
- subtle pink cue for the current selection or associated evidence;
- preserved OpenStreetMap attribution;
- map overlays that remain readable in both themes.

The map frame may be titled `MAP.VIEW` with compact context such as the current state, selected layer, or cursor coordinates.

### Right workspace: Authority record

The right column contains:

- selected jurisdiction;
- authority chain;
- adopted codes;
- amendments;
- resolution status and uncertainty;
- evidence and provenance;
- expandable technical attributes.

Suggested title: `AUTHORITY.RECORD`.

Authority names, code titles, and result summaries use editorial serif. Source IDs, dates, feature IDs, and evidence metadata use monospace.

Evidence sections may expand vertically without forcing a map rerender or resetting selection.

## Tablet workspace

At tablet widths, use two columns:

- query and layers on the left;
- map on the right;
- authority record below, spanning the available width.

The layout transition must preserve:

- selected feature;
- resolved location;
- active layers;
- expanded evidence sections where practical;
- meaningful scroll position.

The authority record should not be hidden behind a drawer at tablet widths unless viewport height makes the standard layout unusable.

## Mobile task flow

### Before lookup

Order:

1. compact product header;
2. query panel;
3. optional recent or example locations;
4. collapsed map panel;
5. layer controls in a drawer or disclosure.

The query must be immediately usable without scrolling through map controls.

### After lookup

Order:

1. authority result;
2. adopted codes and amendments;
3. confidence, ambiguity, conflict, or unresolved conditions;
4. source evidence grouped by authority or document;
5. collapsed map control;
6. expandable technical details.

The map control reads `View jurisdiction on map` when collapsed and has an equally clear close action when expanded.

Expanding the map must preserve:

- the resolved point;
- selected jurisdiction boundaries;
- current layer state;
- evidence context;
- return scroll position.

Closing the map returns the user to the same result context rather than the top of the page.

### Mobile panel treatment

Desktop title bars become thinner section headers. Touch targets remain practical even though the interface looks compact. Technical metadata is available through disclosures rather than occupying the default result view.

## Component boundaries

The visual redesign must not merge domain responsibilities.

### Existing responsibilities retained

- `App`: loading, application state, feature selection, and composition.
- `MapStage`: map rendering and map interaction.
- `ResolutionPanel`: query and resolution behavior.
- `LayerToggleList`: layer selection.
- `FeatureInspector`: selected-record and evidence display.
- API and type modules: transport and domain contracts.

### Presentation components to add or refine

Small presentation units may include:

- `AppHeader`;
- `ThemeToggle`;
- `WorkspaceFrame`;
- `WorkspaceTitlebar`;
- `CompactStatusStrip`;
- `MobileSection` or accessible disclosure wrapper;
- `MobileMapDisclosure`;
- reusable semantic status treatment.

These components should have narrow, testable interfaces and no direct API-fetching responsibility.

### CSS organization

Use a layered stylesheet structure with clear separation among:

- reset;
- shared and map-specific tokens;
- base typography and focus behavior;
- reusable frames, title bars, controls, and status treatments;
- application layout;
- map-specific styles;
- resolution and evidence styles;
- responsive rules;
- utilities.

The existing `styles.css` and `resolution.css` may be retained and reorganized, or split into focused CSS files, provided import order is explicit and style ownership is clear.

A cross-repository package is intentionally deferred. The initial implementation copies and adapts stable token values and patterns with attribution in project documentation, avoiding premature shared-package coupling.

## Theme behavior

- Support explicit light and dark themes.
- On first visit, respect the system color-scheme preference.
- Persist the user's explicit choice in local storage.
- Apply the theme at the document root before or during initial render to minimize flash.
- The theme control must have an accessible name, visible focus, and state conveyed beyond color.
- Tests must not depend on the host machine's theme preference.

## Map styling and contrast

MapLibre layer colors should be derived from semantic theme values through a small adapter or centralized style definition. Components should not scatter unrelated hard-coded colors.

Requirements:

- selected geometry remains visible over light and dark basemaps;
- overlapping jurisdiction families remain distinguishable;
- fill opacity does not obscure important labels;
- line width and pattern may supplement color where overlap is likely;
- hover and selected states are visually distinct;
- map controls and overlays meet contrast expectations;
- attribution remains visible and readable;
- pink is not the sole distinction between operational layer families.

## State and error treatment

The interface must provide consistent framed states for:

- initial loading;
- runtime snapshot unavailable;
- successful resolution;
- partial resolution;
- local record required;
- ambiguous resolution;
- conflicting evidence;
- insufficient evidence;
- API or data-loading failure.

Every state includes a textual label. Warning and error states use semantic wording and border/icon treatment in addition to color. The application continues to show uncertainty rather than infer unsupported answers.

## Accessibility

- Visible high-contrast focus treatment based on the shared focus token.
- Keyboard operation for query, layers, evidence disclosures, theme toggle, and mobile map disclosure.
- Correct `aria-expanded` and control relationships for collapsible sections.
- Focus is retained or moved intentionally when panels open and close.
- Status is not communicated solely through color.
- Text remains readable at browser zoom.
- Touch controls remain large enough for reliable mobile use.
- Reduced-motion preferences disable nonessential transitions.
- Theme combinations preserve readable text, controls, map overlays, and focus indicators.

## Responsive state preservation

Responsive changes must not reset domain state. Resizing or rotating the device should preserve:

- active query result;
- selected feature;
- active layer set;
- theme;
- map center and zoom where practical;
- expanded mobile map state;
- evidence disclosure state where practical.

Layout reflow should be CSS-led wherever possible. JavaScript viewport branching is allowed only when behavior genuinely differs, such as the mobile result ordering and collapsed-map interaction.

## Testing and verification

### Component and behavior tests

Add tests for:

- first-visit system-theme selection;
- persisted explicit theme selection;
- theme-toggle accessibility and state;
- mobile pre-query ordering;
- mobile post-query result ordering;
- mobile map disclosure expansion and collapse;
- preserved selected result when toggling the map;
- semantic status labels for resolved, ambiguous, conflicting, and unavailable states;
- no loss of existing resolution and layer-toggle behavior.

### Visual regression coverage

Capture and review at least:

- desktop light;
- desktop dark;
- tablet light;
- tablet dark;
- mobile before lookup;
- mobile after resolved lookup;
- mobile after ambiguous lookup;
- mobile after conflicting lookup;
- mobile with map expanded from a result;
- unavailable-snapshot state in both themes.

Representative fixtures must be deterministic and must not depend on live external services.

### Accessibility checks

Verify:

- keyboard path through header, query, layers, result, evidence, and map disclosure;
- visible focus in both themes;
- disclosure semantics;
- text and control contrast;
- status meaning without color;
- reduced-motion behavior.

### Baseline repository checks

The implementation must pass:

```bash
pnpm check
pnpm lint
pnpm format:check
pnpm test
pnpm build
```

Backend behavior is not expected to change, but the existing backend checks should still be run before merge when the full repository verification contract requires them.

## Acceptance criteria

The work is complete when:

1. Building Code Map visually reads as a sibling of the LLC website in both light and dark modes.
2. Large rounded glass cards and diffuse SaaS styling have been replaced by compact framed workspaces.
3. Desktop retains a usable three-part query/map/record workspace.
4. Tablet uses the approved two-column layout with the authority record below.
5. Mobile prioritizes the query before lookup and the authority result after lookup.
6. The mobile map is collapsed after a result and expands without losing selection or context.
7. Blue and green dominate map/data visualization; pink remains restrained.
8. The basemap stays quiet while labels and attribution remain readable.
9. Authority and code content uses editorial serif; operational metadata and controls use monospace.
10. Theme preference persists and first visit respects system preference.
11. Loading, unavailable, resolved, ambiguous, conflicting, and insufficient-evidence states remain explicit and accessible.
12. Existing resolution, layer selection, map selection, and evidence behavior has no functional regression.
13. Required automated checks and responsive visual review pass.

## Implementation sequence

The implementation plan should proceed in this order:

1. introduce and test theme state and semantic tokens;
2. add reusable frame, title-bar, header, theme-toggle, and status-strip presentation components;
3. restyle desktop layout and panels without changing domain behavior;
4. centralize MapLibre semantic colors and validate both themes;
5. implement tablet reflow;
6. implement the mobile pre-query and post-query task flow with map disclosure;
7. refine state, evidence, and technical-detail treatments;
8. add behavior, accessibility, and visual regression coverage;
9. run the complete verification suite and review all required viewport/theme captures.

## Risks and mitigations

### Map legibility in dark mode

Dark UI tokens can reduce map-label or boundary contrast. Mitigate with theme-specific map aliases, controlled fill opacity, and visual fixtures covering overlapping boundaries.

### Decorative chrome consuming workspace

Window frames can reduce usable map area. Mitigate by keeping title bars compact, avoiding nested frames, and using thinner mobile section headers.

### Responsive remounting or state loss

Separate mobile/desktop trees can reset map or result state. Prefer shared state above layout branches and CSS reflow; test resize and disclosure interactions.

### Cross-repository style drift

Copying tokens can diverge over time. Document the source and adaptation, keep semantic names aligned where useful, and reconsider a shared package only after both applications demonstrate stable repeated needs.

### Pink becoming semantically ambiguous

Restrict pink to selection and review emphasis. Preserve explicit semantic labels and distinct warning/error treatments.

## Deferred follow-up

After this implementation has stabilized, evaluate whether the LLC website and Building Code Map have enough repeated, stable primitives to justify a small shared design-token package. That decision is outside this change and must not block delivery.
