# Building Code Map Shared Visual System Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Restyle and recompose the Building Code Map frontend as a light/dark, map-native sibling of the Laura J Hutchins LLC website while preserving all regulatory, boundary-selection, and API behavior.

**Architecture:** Keep `App` as the owner of fetched metadata, selected boundaries, theme, and the active resolution result. Split query submission from authority-result presentation so the same result can occupy the desktop authority column and become the first mobile section after lookup. Add small presentation components for theme, framed workspaces, status, and disclosures; keep one mounted MapLibre instance and drive its colors from semantic CSS variables through a centralized adapter.

**Tech Stack:** React 18, TypeScript 5.5 strict mode, Vite 5, Vitest 2, Testing Library, MapLibre GL JS 5, Playwright Chromium, CSS custom properties and cascade layers.

## Global Constraints

- Use Node.js 22 and pnpm 10.12.1-compatible tooling.
- Do not change backend endpoints, response contracts, regulatory normalization, or boundary-selection semantics.
- Keep MapLibre and the attributed OpenStreetMap raster source.
- Preserve visible, readable OpenStreetMap attribution in both themes.
- Do not create a cross-repository component or token package in this change.
- Use the editorial serif for authority/code content and monospace for controls, labels, coordinates, source IDs, dates, and technical metadata.
- Blue and green are the primary map/data colors; pink is limited to selection, review-required states, and exceptional emphasis.
- Color must never be the only indicator of loading, success, warning, ambiguity, conflict, or error.
- Mobile order before lookup is query, collapsed map, authority/technical content.
- Mobile order after lookup is authority result and evidence, collapsed map, query, technical details.
- The MapLibre component must remain mounted when mobile disclosures open, close, or sections reorder.
- Respect system theme on first visit, persist an explicit choice, and apply the root theme before React renders.
- Required baseline commands remain `pnpm check`, `pnpm lint`, `pnpm format:check`, `pnpm test`, and `pnpm build`.

---

## File Structure

### New files

- `vitest.config.ts` — jsdom test environment and shared setup.
- `src/test/setup.ts` — Testing Library cleanup and browser API defaults.
- `src/test/fixtures.ts` — deterministic layer, boundary, refresh, and resolution fixtures.
- `src/lib/theme.ts` — theme types, storage key, initial resolution, and DOM application.
- `src/hooks/useTheme.ts` — React theme state and persistence.
- `src/lib/mapTheme.ts` — semantic CSS-variable to MapLibre paint adapter.
- `src/components/ThemeToggle.tsx` — accessible theme control.
- `src/components/AppHeader.tsx` — product identity, status strip, optional main-site link, and theme control.
- `src/components/CompactStatusStrip.tsx` — dense layer/record/refresh summary.
- `src/components/WorkspaceFrame.tsx` — shared window frame and title bar.
- `src/components/AuthorityRecord.tsx` — authority, adoption, rule, warning, and evidence presentation.
- `src/components/FeatureSelectionList.tsx` — cached boundary selection list extracted from `MapStage`.
- `src/components/BackendContractDetails.tsx` — compact technical endpoint reference extracted from `MapStage`.
- `src/components/MobileMapDisclosure.tsx` — mobile-only map collapse without unmounting children.
- `src/components/ResponsiveDisclosure.tsx` — disclosure that is collapsible on mobile and always visible on larger screens.
- `src/hooks/useMediaQuery.ts` — responsive behavior hook that tracks the mobile breakpoint without remounting content.
- `src/lib/workspaceOrder.ts` — keyed pre-query and post-query section ordering.
- `src/theme.css` — shared light/dark tokens and canvas treatment.
- `src/map.css` — MapLibre canvas, controls, overlays, and boundary workspace styling.
- `src/components/*.test.tsx`, `src/lib/*.test.ts`, `src/hooks/*.test.tsx` — focused behavior tests.
- `playwright.visual.config.ts` — deterministic visual-regression configuration.
- `tsconfig.visual.json` — strict typecheck for Playwright configuration and visual tests.
- `tests/visual/building-code-map.visual.spec.ts` — responsive/theme/state screenshot matrix.
- `docs/how-to/review-frontend-visuals.md` — local visual review workflow.

### Modified files

- `package.json`, `pnpm-lock.yaml` — test dependencies and scripts.
- `eslint.config.js` — Node globals and lint coverage for Playwright files.
- `index.html` — correct product metadata and pre-React theme bootstrap.
- `src/main.tsx` — ordered CSS imports.
- `src/App.tsx` — lifted resolution/theme state and keyed workspace composition.
- `src/components/ResolutionPanel.tsx` — query-only controlled result callback.
- `src/components/MapStage.tsx` — map-only responsibility and theme updates.
- `src/components/FeatureInspector.tsx` — frameless technical record content.
- `src/components/LayerToggleList.tsx` — frameless layer-control content.
- `src/components/StatusBanner.tsx` — compact semantic data-status content.
- `src/styles.css` — base, reusable frame/control, and desktop/tablet layout rules.
- `src/resolution.css` — query, authority result, evidence, status, and mobile ordering rules.
- `.github/workflows/ci.yml` — deterministic Chromium visual-regression job.
- `README.md` — theme and visual verification commands.

### Removed files

None. Keep `src/styles.css` and `src/resolution.css`; reorganize them rather than introducing unnecessary churn.

---

### Task 1: Establish the component-test harness and deterministic fixtures

**Files:**
- Modify: `package.json`
- Modify: `pnpm-lock.yaml`
- Create: `vitest.config.ts`
- Create: `src/test/setup.ts`
- Create: `src/test/fixtures.ts`
- Test: existing `src/**/*.test.ts`

**Interfaces:**
- Produces: `demoLayers`, `demoRefreshStatus`, `demoBoundaryFeatures`, `resolvedResult`, `ambiguousResult`, and `conflictingResult`.
- Produces: a jsdom Vitest environment with `window.matchMedia`, `ResizeObserver`, and Testing Library cleanup.
- Consumes: existing types from `src/types.ts`.

- [ ] **Step 1: Install the browser-component test dependencies**

Run:

```bash
pnpm add -D @testing-library/react @testing-library/user-event jsdom
```

Expected: `package.json` and `pnpm-lock.yaml` change; installation exits with status 0.

- [ ] **Step 2: Create the Vitest configuration**

Create `vitest.config.ts`:

```ts
import { defineConfig } from "vitest/config";

export default defineConfig({
  test: {
    environment: "jsdom",
    globals: true,
    setupFiles: ["./src/test/setup.ts"],
    restoreMocks: true,
    clearMocks: true,
  },
});
```

- [ ] **Step 3: Create shared browser test setup**

Create `src/test/setup.ts`:

```ts
import { cleanup } from "@testing-library/react";
import { afterEach, beforeEach, vi } from "vitest";

class TestResizeObserver {
  observe() {}
  unobserve() {}
  disconnect() {}
}

Object.defineProperty(window, "ResizeObserver", {
  configurable: true,
  value: TestResizeObserver,
});

Object.defineProperty(window, "matchMedia", {
  configurable: true,
  writable: true,
  value: vi.fn().mockImplementation((query: string) => ({
    matches: false,
    media: query,
    onchange: null,
    addEventListener: vi.fn(),
    removeEventListener: vi.fn(),
    addListener: vi.fn(),
    removeListener: vi.fn(),
    dispatchEvent: vi.fn(),
  })),
});

beforeEach(() => {
  window.localStorage.clear();
  document.documentElement.removeAttribute("data-theme");
});

afterEach(() => {
  cleanup();
});
```

- [ ] **Step 4: Create deterministic typed fixtures**

Create `src/test/fixtures.ts`:

```ts
import type {
  BoundaryFeatureRecord,
  LayerFamilyDefinition,
  RefreshStatus,
  ResolutionResult,
} from "../types";

export const demoLayers: LayerFamilyDefinition[] = [
  {
    key: "states",
    label: "States",
    martinLayerId: "tigerweb.states",
    description: "State boundaries mirrored from TIGERweb.",
    defaultEnabled: true,
  },
  {
    key: "counties",
    label: "Counties",
    martinLayerId: "tigerweb.counties",
    description: "County boundaries used for jurisdiction lookups.",
    defaultEnabled: true,
  },
  {
    key: "municipalities",
    label: "Municipalities",
    martinLayerId: "tigerweb.municipalities",
    description: "Incorporated places and municipal limits.",
    defaultEnabled: true,
  },
  {
    key: "special_areas",
    label: "Special areas",
    martinLayerId: "tigerweb.special_areas",
    description: "Military installations and other special land-use areas.",
    defaultEnabled: false,
  },
  {
    key: "tribal_areas",
    label: "American Indian areas",
    martinLayerId: "tigerweb.tribal_areas",
    description: "Tribal boundaries and related areas.",
    defaultEnabled: false,
  },
  {
    key: "neris_jurisdictions",
    label: "NERIS jurisdictions",
    martinLayerId: "neris.department_jurisdictions",
    description: "Department jurisdiction polygons from NERIS.",
    defaultEnabled: false,
  },
];

export const demoRefreshStatus: RefreshStatus = {
  status: "ok",
  latestSuccessfulRefresh: "2026-07-20T18:00:00Z",
  latestAttempt: "2026-07-20T18:00:00Z",
  nextScheduledRefresh: "2026-07-21T18:00:00Z",
  message: "Cached boundary snapshot is ready.",
};

export const demoBoundaryFeatures: BoundaryFeatureRecord[] = [
  {
    layerFamily: "states",
    featureId: "08",
    title: "Colorado",
    subtitle: "State boundary",
    sourceId: "STATEFP=08",
    geometryLabel: "Polygon",
    geometrySource: "tigerweb_live",
    lastSyncedAt: "2026-07-20T18:00:00Z",
    attributes: {
      NAME: "Colorado",
      STATEFP: "08",
    },
    geometry: {
      type: "Polygon",
      coordinates: [
        [
          [-109.05, 37],
          [-102.05, 37],
          [-102.05, 41],
          [-109.05, 41],
          [-109.05, 37],
        ],
      ],
    },
  },
];

export const resolvedResult: ResolutionResult = {
  schemaVersion: "0.4",
  generatedAt: "2026-07-20T18:05:00Z",
  profileId: "co",
  profileLastVerified: "2026-07-20",
  geography: {
    stateId: "co",
    stateFips: "08",
    stateName: "Colorado",
    county: {
      layerFamily: "counties",
      featureId: "031",
      name: "Denver County",
      sourceId: "COUNTYFP=031",
    },
    municipality: {
      layerFamily: "municipalities",
      featureId: "20000",
      name: "Denver",
      sourceId: "PLACEFP=20000",
    },
    incorporated: true,
    specialAreas: [],
    tribalAreas: [],
    fireJurisdictions: [],
  },
  codeFamily: "building",
  projectType: "ordinary",
  applicabilityDate: "2026-07-21",
  status: "resolved",
  policyBasis: {
    status: "resolved",
    requiredLocalRecords: [],
    warnings: [],
    sourceIds: ["co-state-code"],
    verification: { status: "verified", confidence: 0.96 },
  },
  authorityCandidates: [
    {
      kind: "municipality",
      authorityId: "denver",
      name: "City and County of Denver",
      roles: ["adoption", "enforcement"],
      sourceIds: ["denver-code"],
      verification: { status: "verified", confidence: 0.94 },
    },
  ],
  authorityPath: [
    {
      id: "co-to-denver",
      fromId: "co",
      relationship: "delegates_local_adoption",
      to: "denver",
      scope: ["building"],
      summary: "Local adoption and enforcement apply within Denver.",
      sourceIds: ["co-state-code", "denver-code"],
      verification: { status: "verified", confidence: 0.92 },
    },
  ],
  adoptions: [
    {
      id: "denver-ibc-2024",
      codeFamily: "building",
      status: "effective",
      stateCodeName: "2024 Denver Building Code",
      enforcementModel: "local",
      dates: { effective_date: "2025-01-01" },
      sourceIds: ["denver-code"],
      verification: { status: "verified", confidence: 0.95 },
    },
  ],
  applicableRules: [
    {
      id: "denver-local-amendments",
      kind: "amendment",
      codeFamily: "building",
      summary: "Denver amendments apply to the adopted model code.",
      sourceIds: ["denver-code"],
      verification: { status: "verified", confidence: 0.9 },
    },
  ],
  supportingClaims: [
    {
      id: "claim-denver-authority",
      subjectId: "denver",
      field: "adoption_authority",
      status: "supported",
      value: "local",
      sourceIds: ["denver-code"],
      verification: { status: "verified", confidence: 0.94 },
    },
  ],
  requiredLocalRecords: [],
  warnings: [],
  evidence: [
    {
      id: "denver-code",
      title: "Denver Building and Fire Code",
      url: "https://example.invalid/denver-code",
      kind: "official_code",
      accessedAt: "2026-07-20",
      lastCheckedAt: "2026-07-20",
      availability: "available",
    },
  ],
};

export const ambiguousResult: ResolutionResult = {
  ...resolvedResult,
  status: "ambiguous",
  authorityCandidates: [
    ...resolvedResult.authorityCandidates,
    {
      kind: "special_district",
      authorityId: "metro-district",
      name: "Example Metropolitan District",
      roles: ["possible_enforcement"],
      sourceIds: ["district-record"],
      verification: { status: "needs_review", confidence: 0.55 },
    },
  ],
  warnings: ["A local record is required to confirm the enforcing authority."],
  requiredLocalRecords: ["Current local enforcement agreement"],
};

export const conflictingResult: ResolutionResult = {
  ...resolvedResult,
  status: "conflicting",
  supportingClaims: [
    {
      id: "claim-conflict-a",
      subjectId: "denver",
      field: "effective_code",
      status: "conflicting",
      value: "2021 IBC",
      conflictGroup: "effective-code",
      sourceIds: ["source-a"],
      verification: { status: "needs_review", confidence: 0.5 },
    },
    {
      id: "claim-conflict-b",
      subjectId: "denver",
      field: "effective_code",
      status: "conflicting",
      value: "2024 IBC",
      conflictGroup: "effective-code",
      sourceIds: ["source-b"],
      verification: { status: "needs_review", confidence: 0.5 },
    },
  ],
  warnings: ["Official sources disagree about the currently effective edition."],
};
```

- [ ] **Step 5: Run the existing unit suite in jsdom**

Run:

```bash
pnpm test
```

Expected: all pre-existing tests pass. Fix only environment assumptions introduced by jsdom; do not change production behavior.

- [ ] **Step 6: Commit the harness**

```bash
git add package.json pnpm-lock.yaml vitest.config.ts src/test/setup.ts src/test/fixtures.ts
git commit -m "test: add deterministic frontend component harness"
```

---

### Task 2: Implement theme resolution, persistence, and accessible switching

**Files:**
- Create: `src/lib/theme.ts`
- Create: `src/lib/theme.test.ts`
- Create: `src/hooks/useTheme.ts`
- Create: `src/hooks/useTheme.test.tsx`
- Create: `src/components/ThemeToggle.tsx`
- Create: `src/components/ThemeToggle.test.tsx`
- Modify: `index.html`

**Interfaces:**
- Produces: `ThemeName = "light" | "dark"`.
- Produces: `resolveInitialTheme(storedTheme, prefersDark)`.
- Produces: `applyTheme(theme, root?)`.
- Produces: `useTheme(): { theme; setTheme; toggleTheme }`.
- Produces: `<ThemeToggle theme onToggle />`.

- [ ] **Step 1: Write failing pure theme tests**

Create `src/lib/theme.test.ts`:

```ts
import { describe, expect, it } from "vitest";
import { applyTheme, resolveInitialTheme } from "./theme";

describe("resolveInitialTheme", () => {
  it("uses an explicit stored theme before the system preference", () => {
    expect(resolveInitialTheme("light", true)).toBe("light");
    expect(resolveInitialTheme("dark", false)).toBe("dark");
  });

  it("falls back to the system preference for missing or invalid storage", () => {
    expect(resolveInitialTheme(null, true)).toBe("dark");
    expect(resolveInitialTheme("sepia", false)).toBe("light");
  });
});

describe("applyTheme", () => {
  it("writes theme and color scheme to the document root", () => {
    applyTheme("dark");
    expect(document.documentElement.dataset.theme).toBe("dark");
    expect(document.documentElement.style.colorScheme).toBe("dark");
  });
});
```

Run:

```bash
pnpm vitest run src/lib/theme.test.ts
```

Expected: FAIL because `src/lib/theme.ts` does not exist.

- [ ] **Step 2: Implement theme utilities**

Create `src/lib/theme.ts`:

```ts
export type ThemeName = "light" | "dark";

export const THEME_STORAGE_KEY = "building-code-map-theme";

export function isThemeName(value: string | null): value is ThemeName {
  return value === "light" || value === "dark";
}

export function resolveInitialTheme(
  storedTheme: string | null,
  prefersDark: boolean,
): ThemeName {
  if (isThemeName(storedTheme)) {
    return storedTheme;
  }

  return prefersDark ? "dark" : "light";
}

export function applyTheme(
  theme: ThemeName,
  root: HTMLElement = document.documentElement,
): void {
  root.dataset.theme = theme;
  root.style.colorScheme = theme;
}
```

Run:

```bash
pnpm vitest run src/lib/theme.test.ts
```

Expected: PASS.

- [ ] **Step 3: Write a failing hook persistence test**

Create `src/hooks/useTheme.test.tsx`:

```tsx
import { act, renderHook } from "@testing-library/react";
import { describe, expect, it } from "vitest";
import { THEME_STORAGE_KEY } from "../lib/theme";
import { useTheme } from "./useTheme";

describe("useTheme", () => {
  it("uses the bootstrapped root theme and persists an explicit toggle", () => {
    document.documentElement.dataset.theme = "dark";

    const { result } = renderHook(() => useTheme());

    expect(result.current.theme).toBe("dark");
    expect(window.localStorage.getItem(THEME_STORAGE_KEY)).toBeNull();

    act(() => {
      result.current.toggleTheme();
    });

    expect(result.current.theme).toBe("light");
    expect(document.documentElement.dataset.theme).toBe("light");
    expect(window.localStorage.getItem(THEME_STORAGE_KEY)).toBe("light");
  });
});
```

Run:

```bash
pnpm vitest run src/hooks/useTheme.test.tsx
```

Expected: FAIL because `useTheme` does not exist.

- [ ] **Step 4: Implement the hook**

Create `src/hooks/useTheme.ts`:

```ts
import { useCallback, useEffect, useState } from "react";
import {
  applyTheme,
  isThemeName,
  resolveInitialTheme,
  THEME_STORAGE_KEY,
  type ThemeName,
} from "../lib/theme";

function readInitialTheme(): ThemeName {
  const bootstrappedTheme = document.documentElement.dataset.theme ?? null;
  if (isThemeName(bootstrappedTheme)) {
    return bootstrappedTheme;
  }

  return resolveInitialTheme(
    window.localStorage.getItem(THEME_STORAGE_KEY),
    window.matchMedia("(prefers-color-scheme: dark)").matches,
  );
}

export function useTheme(): {
  theme: ThemeName;
  setTheme: (theme: ThemeName) => void;
  toggleTheme: () => void;
} {
  const [theme, setThemeState] = useState<ThemeName>(readInitialTheme);

  const setTheme = useCallback((nextTheme: ThemeName) => {
    applyTheme(nextTheme);
    window.localStorage.setItem(THEME_STORAGE_KEY, nextTheme);
    setThemeState(nextTheme);
  }, []);

  const toggleTheme = useCallback(() => {
    setTheme(theme === "light" ? "dark" : "light");
  }, [setTheme, theme]);

  useEffect(() => {
    applyTheme(theme);
  }, [theme]);

  return { theme, setTheme, toggleTheme };
}
```

Run:

```bash
pnpm vitest run src/hooks/useTheme.test.tsx
```

Expected: PASS.

- [ ] **Step 5: Write a failing accessible toggle test**

Create `src/components/ThemeToggle.test.tsx`:

```tsx
import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { describe, expect, it, vi } from "vitest";
import { ThemeToggle } from "./ThemeToggle";

describe("ThemeToggle", () => {
  it("announces the current theme and invokes the toggle callback", async () => {
    const user = userEvent.setup();
    const onToggle = vi.fn();

    render(<ThemeToggle theme="dark" onToggle={onToggle} />);

    const button = screen.getByRole("button", { name: "Switch to light theme" });
    expect(button.getAttribute("aria-pressed")).toBe("true");

    await user.click(button);
    expect(onToggle).toHaveBeenCalledTimes(1);
  });
});
```

Run:

```bash
pnpm vitest run src/components/ThemeToggle.test.tsx
```

Expected: FAIL because `ThemeToggle` does not exist.

- [ ] **Step 6: Implement the toggle**

Create `src/components/ThemeToggle.tsx`:

```tsx
import type { ThemeName } from "../lib/theme";

interface ThemeToggleProps {
  theme: ThemeName;
  onToggle: () => void;
}

export function ThemeToggle({ theme, onToggle }: ThemeToggleProps): JSX.Element {
  const nextTheme = theme === "dark" ? "light" : "dark";

  return (
    <button
      className="theme-toggle"
      type="button"
      aria-label={`Switch to ${nextTheme} theme`}
      aria-pressed={theme === "dark"}
      onClick={onToggle}
    >
      <span className="theme-toggle__track" aria-hidden="true">
        <span className="theme-toggle__thumb" />
      </span>
      <span className="theme-toggle__text">{theme}</span>
    </button>
  );
}
```

Run:

```bash
pnpm vitest run src/components/ThemeToggle.test.tsx
```

Expected: PASS.

- [ ] **Step 7: Bootstrap the theme before React and correct page metadata**

Replace the `<head>` contents in `index.html` with:

```html
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <meta name="color-scheme" content="light dark" />
  <title>Building Code Map</title>
  <meta
    name="description"
    content="A source-backed jurisdiction and adopted-code explorer for local building-code research."
  />
  <script>
    (() => {
      const storageKey = "building-code-map-theme";
      const storedTheme = window.localStorage.getItem(storageKey);
      const theme =
        storedTheme === "light" || storedTheme === "dark"
          ? storedTheme
          : window.matchMedia("(prefers-color-scheme: dark)").matches
            ? "dark"
            : "light";
      document.documentElement.dataset.theme = theme;
      document.documentElement.style.colorScheme = theme;
    })();
  </script>
</head>
```

- [ ] **Step 8: Run the focused and full tests**

Run:

```bash
pnpm vitest run src/lib/theme.test.ts src/hooks/useTheme.test.tsx src/components/ThemeToggle.test.tsx
pnpm test
```

Expected: all tests PASS.

- [ ] **Step 9: Commit theme behavior**

```bash
git add index.html src/lib/theme.ts src/lib/theme.test.ts src/hooks/useTheme.ts src/hooks/useTheme.test.tsx src/components/ThemeToggle.tsx src/components/ThemeToggle.test.tsx
git commit -m "feat: add persistent light and dark themes"
```

---

### Task 3: Add the shared token layer and reusable workspace chrome

**Files:**
- Create: `src/theme.css`
- Create: `src/components/WorkspaceFrame.tsx`
- Create: `src/components/WorkspaceFrame.test.tsx`
- Create: `src/components/CompactStatusStrip.tsx`
- Create: `src/components/CompactStatusStrip.test.tsx`
- Create: `src/components/AppHeader.tsx`
- Create: `src/components/AppHeader.test.tsx`
- Modify: `src/main.tsx`
- Modify: `src/styles.css`

**Interfaces:**
- Produces: `<WorkspaceFrame title eyebrow? actions? className? id? children />`.
- Produces: `<CompactStatusStrip layerCount recordCount refreshStatus />`.
- Produces: `<AppHeader theme onToggleTheme layerCount recordCount refreshStatus mainSiteUrl? />`.
- Consumes: `ThemeName` and `RefreshStatus`.

- [ ] **Step 1: Write failing frame and header tests**

Create `src/components/WorkspaceFrame.test.tsx`:

```tsx
import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";
import { WorkspaceFrame } from "./WorkspaceFrame";

describe("WorkspaceFrame", () => {
  it("connects the title bar to the framed content", () => {
    render(
      <WorkspaceFrame id="query-frame" title="QUERY.LOCATION" eyebrow="Input">
        <p>Query content</p>
      </WorkspaceFrame>,
    );

    const region = screen.getByRole("region", { name: "QUERY.LOCATION" });
    expect(region.id).toBe("query-frame");
    expect(screen.getByText("Input")).toBeTruthy();
    expect(screen.getByText("Query content")).toBeTruthy();
  });
});
```

Create `src/components/AppHeader.test.tsx`:

```tsx
import { render, screen } from "@testing-library/react";
import { describe, expect, it, vi } from "vitest";
import { demoRefreshStatus } from "../test/fixtures";
import { AppHeader } from "./AppHeader";

describe("AppHeader", () => {
  it("renders product identity, compact status, optional site link, and theme control", () => {
    render(
      <AppHeader
        theme="light"
        onToggleTheme={vi.fn()}
        layerCount={6}
        recordCount={24}
        refreshStatus={demoRefreshStatus}
        mainSiteUrl="https://example.invalid"
      />,
    );

    expect(screen.getByRole("heading", { name: "Building Code Map" })).toBeTruthy();
    expect(screen.getByText("6 layer families")).toBeTruthy();
    expect(screen.getByText("24 cached records")).toBeTruthy();
    expect(screen.getByRole("link", { name: "Main website" }).getAttribute("href")).toBe(
      "https://example.invalid",
    );
    expect(screen.getByRole("button", { name: "Switch to dark theme" })).toBeTruthy();
  });
});
```

Run:

```bash
pnpm vitest run src/components/WorkspaceFrame.test.tsx src/components/AppHeader.test.tsx
```

Expected: FAIL because the components do not exist.

- [ ] **Step 2: Implement the reusable frame**

Create `src/components/WorkspaceFrame.tsx`:

```tsx
import type { ReactNode } from "react";

interface WorkspaceFrameProps {
  id?: string;
  title: string;
  eyebrow?: string;
  actions?: ReactNode;
  className?: string;
  children: ReactNode;
}

export function WorkspaceFrame({
  id,
  title,
  eyebrow,
  actions,
  className = "",
  children,
}: WorkspaceFrameProps): JSX.Element {
  const titleId = `${id ?? title.toLowerCase().replace(/[^a-z0-9]+/g, "-")}-title`;

  return (
    <section
      id={id}
      className={`workspace-frame ${className}`.trim()}
      aria-labelledby={titleId}
    >
      <div className="workspace-titlebar">
        <div className="workspace-titlebar__copy">
          {eyebrow ? <span>{eyebrow}</span> : null}
          <h2 id={titleId}>{title}</h2>
        </div>
        {actions ? <div className="workspace-titlebar__actions">{actions}</div> : null}
      </div>
      <div className="workspace-frame__body">{children}</div>
    </section>
  );
}
```

- [ ] **Step 3: Implement compact status and header components**

Create `src/components/CompactStatusStrip.tsx`:

```tsx
import type { RefreshStatus } from "../types";

interface CompactStatusStripProps {
  layerCount: number;
  recordCount: number;
  refreshStatus: RefreshStatus | null;
}

export function CompactStatusStrip({
  layerCount,
  recordCount,
  refreshStatus,
}: CompactStatusStripProps): JSX.Element {
  const status = refreshStatus?.status ?? "warning";
  const statusLabel =
    status === "ok" ? "Data ready" : status === "error" ? "Data unavailable" : "Data attention";

  return (
    <div className="compact-status" aria-label="Application data status">
      <span>{layerCount} layer families</span>
      <span>{recordCount} cached records</span>
      <span className={`semantic-status semantic-status--${status}`}>{statusLabel}</span>
    </div>
  );
}
```

Create `src/components/AppHeader.tsx`:

```tsx
import type { ThemeName } from "../lib/theme";
import type { RefreshStatus } from "../types";
import { CompactStatusStrip } from "./CompactStatusStrip";
import { ThemeToggle } from "./ThemeToggle";

interface AppHeaderProps {
  theme: ThemeName;
  onToggleTheme: () => void;
  layerCount: number;
  recordCount: number;
  refreshStatus: RefreshStatus | null;
  mainSiteUrl?: string;
}

export function AppHeader({
  theme,
  onToggleTheme,
  layerCount,
  recordCount,
  refreshStatus,
  mainSiteUrl,
}: AppHeaderProps): JSX.Element {
  return (
    <header className="app-header">
      <div className="app-brand">
        <span className="app-brand__mark" aria-hidden="true">
          BCM
        </span>
        <div>
          <p>Jurisdiction explorer</p>
          <h1>Building Code Map</h1>
        </div>
      </div>

      <CompactStatusStrip
        layerCount={layerCount}
        recordCount={recordCount}
        refreshStatus={refreshStatus}
      />

      <div className="app-header__actions">
        {mainSiteUrl ? (
          <a href={mainSiteUrl} rel="noreferrer">
            Main website
          </a>
        ) : null}
        <ThemeToggle theme={theme} onToggle={onToggleTheme} />
      </div>
    </header>
  );
}
```

Create `src/components/CompactStatusStrip.test.tsx`:

```tsx
import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";
import { demoRefreshStatus } from "../test/fixtures";
import { CompactStatusStrip } from "./CompactStatusStrip";

describe("CompactStatusStrip", () => {
  it("uses text as well as a semantic class for status", () => {
    render(
      <CompactStatusStrip layerCount={6} recordCount={24} refreshStatus={demoRefreshStatus} />,
    );

    expect(screen.getByText("Data ready").className).toContain("semantic-status--ok");
  });
});
```

Run:

```bash
pnpm vitest run src/components/WorkspaceFrame.test.tsx src/components/CompactStatusStrip.test.tsx src/components/AppHeader.test.tsx
```

Expected: PASS.

- [ ] **Step 4: Create the approved light/dark token layer**

Create `src/theme.css`:

```css
@layer tokens, reset, base, components, layout, map, resolution, responsive;

@layer tokens {
  :root {
    --font-display: "Iowan Old Style", "Palatino Linotype", Palatino, Georgia, serif;
    --font-body: Georgia, "Times New Roman", serif;
    --font-ui: ui-monospace, "Cascadia Mono", "SFMono-Regular", Consolas, monospace;
    --space-1: 0.375rem;
    --space-2: 0.625rem;
    --space-3: 0.875rem;
    --space-4: 1.125rem;
    --space-5: 1.5rem;
    --space-6: 2rem;
    --radius-sm: 2px;
    --radius-md: 5px;
    --transition: 160ms ease;
  }

  :root,
  :root[data-theme="light"] {
    color-scheme: light;
    --canvas: #f4fbf8;
    --canvas-alt: #e6f5ef;
    --surface: #ffffff;
    --surface-raised: #f9fffc;
    --surface-muted: #e9f6f1;
    --ink: #071c22;
    --ink-soft: #3d5b62;
    --line: #103f45;
    --line-soft: #a8c7bf;
    --blue: #123b91;
    --blue-strong: #09276e;
    --green: #0b6b51;
    --green-soft: #8ee0c0;
    --pink: #e99fc3;
    --pink-soft: #f7d5e5;
    --focus: #004ee8;
    --success: #0b6b51;
    --warning: #8a5300;
    --danger: #9a2738;
    --titlebar-start: #c4f0e1;
    --titlebar-end: #9ec9ef;
    --shadow-hard: 5px 5px 0 #17383c;
    --shadow-soft: 0 18px 48px rgb(20 67 70 / 0.14);
    --selection: #b7e5d4;
    --canvas-grid: rgb(16 63 69 / 0.1);
    --canvas-glow: rgb(13 107 81 / 0.14);
    --canvas-glow-alt: rgb(18 59 145 / 0.11);
    --map-overlay-surface: rgb(249 255 252 / 0.92);
    --map-overlay-ink: #071c22;
    --boundary-states-fill: rgb(18 59 145 / 0.14);
    --boundary-states-line: #123b91;
    --boundary-counties-fill: rgb(11 107 81 / 0.12);
    --boundary-counties-line: #0b6b51;
    --boundary-municipalities-fill: rgb(18 59 145 / 0.09);
    --boundary-municipalities-line: #315aa9;
    --boundary-special-fill: rgb(11 107 81 / 0.08);
    --boundary-special-line: #438f79;
    --boundary-tribal-fill: rgb(18 59 145 / 0.08);
    --boundary-tribal-line: #6278b6;
    --boundary-neris-fill: rgb(11 107 81 / 0.1);
    --boundary-neris-line: #24755f;
    --boundary-selected: #09276e;
    --boundary-selected-halo: #e99fc3;
  }

  :root[data-theme="dark"] {
    color-scheme: dark;
    --canvas: #050b0d;
    --canvas-alt: #081315;
    --surface: #0a1518;
    --surface-raised: #0d1d21;
    --surface-muted: #10262a;
    --ink: #effff9;
    --ink-soft: #a8cbc2;
    --line: #62d7b8;
    --line-soft: #315f58;
    --blue: #83aaff;
    --blue-strong: #b7ccff;
    --green: #60dfba;
    --green-soft: #173f36;
    --pink: #f3a7cf;
    --pink-soft: #402438;
    --focus: #a9c3ff;
    --success: #60dfba;
    --warning: #f0bd68;
    --danger: #ff9aad;
    --titlebar-start: #0d4c43;
    --titlebar-end: #152f60;
    --shadow-hard: 5px 5px 0 #000000;
    --shadow-soft: 0 24px 64px rgb(0 0 0 / 0.48);
    --selection: #174f43;
    --canvas-grid: rgb(98 215 184 / 0.11);
    --canvas-glow: rgb(30 141 113 / 0.16);
    --canvas-glow-alt: rgb(50 91 176 / 0.16);
    --map-overlay-surface: rgb(5 11 13 / 0.88);
    --map-overlay-ink: #effff9;
    --boundary-states-fill: rgb(131 170 255 / 0.16);
    --boundary-states-line: #83aaff;
    --boundary-counties-fill: rgb(96 223 186 / 0.14);
    --boundary-counties-line: #60dfba;
    --boundary-municipalities-fill: rgb(131 170 255 / 0.1);
    --boundary-municipalities-line: #a0bbff;
    --boundary-special-fill: rgb(96 223 186 / 0.09);
    --boundary-special-line: #91e9d0;
    --boundary-tribal-fill: rgb(131 170 255 / 0.09);
    --boundary-tribal-line: #c0d1ff;
    --boundary-neris-fill: rgb(96 223 186 / 0.11);
    --boundary-neris-line: #b0f2df;
    --boundary-selected: #effff9;
    --boundary-selected-halo: #f3a7cf;
  }
}
```

- [ ] **Step 5: Replace the base/chrome portion of `src/styles.css`**

At the beginning of `src/styles.css`, replace the existing `:root`, reset, body, header, panel, and status-card rules through the current `.panel__header` block with:

```css
@layer reset, base, components, layout, responsive;

@layer reset {
  *,
  *::before,
  *::after {
    box-sizing: border-box;
  }

  html,
  body,
  #root {
    min-width: 320px;
    min-height: 100%;
  }

  body,
  h1,
  h2,
  h3,
  p,
  dl {
    margin: 0;
  }

  button,
  input,
  select,
  a {
    font: inherit;
  }
}

@layer base {
  body {
    min-height: 100vh;
    overflow-x: hidden;
    color: var(--ink);
    background:
      radial-gradient(circle at 10% 12%, var(--canvas-glow), transparent 30rem),
      radial-gradient(circle at 88% 20%, var(--canvas-glow-alt), transparent 28rem),
      linear-gradient(var(--canvas-grid) 1px, transparent 1px),
      linear-gradient(90deg, var(--canvas-grid) 1px, transparent 1px),
      var(--canvas);
    background-size: auto, auto, 32px 32px, 32px 32px, auto;
    font-family: var(--font-body);
    line-height: 1.55;
  }

  ::selection {
    color: var(--ink);
    background: var(--selection);
  }

  :where(a, button, input, select, summary):focus-visible {
    outline: 3px solid var(--focus);
    outline-offset: 3px;
  }

  h1,
  h2,
  h3 {
    color: var(--ink);
    font-family: var(--font-display);
    font-weight: 600;
    letter-spacing: -0.025em;
  }

  a {
    color: var(--blue);
    text-decoration-thickness: 0.09em;
    text-underline-offset: 0.2em;
  }

  a:hover {
    color: var(--green);
  }
}

@layer components {
  .app-shell {
    width: min(100%, 118rem);
    min-height: 100vh;
    margin-inline: auto;
    padding: var(--space-5);
  }

  .app-header {
    display: grid;
    grid-template-columns: minmax(15rem, auto) minmax(0, 1fr) auto;
    align-items: center;
    gap: var(--space-5);
    padding-bottom: var(--space-4);
    border-bottom: 2px solid var(--line);
  }

  .app-brand {
    display: flex;
    align-items: center;
    gap: var(--space-3);
  }

  .app-brand__mark {
    display: grid;
    width: 2.8rem;
    aspect-ratio: 1;
    place-items: center;
    border: 2px solid var(--line);
    border-radius: var(--radius-sm);
    background: linear-gradient(135deg, var(--green-soft), var(--pink-soft));
    box-shadow: 3px 3px 0 var(--line);
    font-family: var(--font-ui);
    font-size: 0.7rem;
    font-weight: 800;
  }

  .app-brand p,
  .workspace-titlebar__copy span {
    color: var(--ink-soft);
    font-family: var(--font-ui);
    font-size: 0.68rem;
    letter-spacing: 0.08em;
    text-transform: uppercase;
  }

  .app-brand h1 {
    font-size: clamp(1.45rem, 2.4vw, 2.15rem);
    line-height: 1;
  }

  .app-header__actions {
    display: flex;
    align-items: center;
    gap: var(--space-3);
    font-family: var(--font-ui);
    font-size: 0.78rem;
  }

  .compact-status {
    display: flex;
    min-width: 0;
    justify-content: center;
    gap: var(--space-3);
    color: var(--ink-soft);
    font-family: var(--font-ui);
    font-size: 0.72rem;
  }

  .compact-status > span {
    padding-inline: var(--space-3);
    border-left: 1px solid var(--line-soft);
  }

  .semantic-status {
    font-weight: 700;
  }

  .semantic-status--ok {
    color: var(--success);
  }

  .semantic-status--warning {
    color: var(--warning);
  }

  .semantic-status--error {
    color: var(--danger);
  }

  .theme-toggle {
    display: inline-flex;
    min-height: 2.5rem;
    align-items: center;
    gap: var(--space-2);
    padding: 0.45rem 0.62rem;
    border: 2px solid var(--line);
    border-radius: var(--radius-sm);
    background: var(--surface);
    box-shadow: 2px 2px 0 var(--line);
    color: var(--ink);
    cursor: pointer;
  }

  .theme-toggle__track {
    position: relative;
    width: 2.15rem;
    height: 1.15rem;
    border: 1px solid var(--line);
    border-radius: 999px;
    background: var(--surface-muted);
  }

  .theme-toggle__thumb {
    position: absolute;
    top: 2px;
    left: 2px;
    width: 0.72rem;
    aspect-ratio: 1;
    border-radius: 50%;
    background: var(--blue);
    transition: transform var(--transition);
  }

  :root[data-theme="dark"] .theme-toggle__thumb {
    background: var(--green);
    transform: translateX(0.96rem);
  }

  .theme-toggle__text {
    min-width: 4ch;
    font-family: var(--font-ui);
    text-transform: uppercase;
  }

  .workspace-frame {
    min-width: 0;
    overflow: clip;
    border: 2px solid var(--line);
    border-radius: var(--radius-md);
    background: var(--surface);
    box-shadow: var(--shadow-hard), var(--shadow-soft);
  }

  .workspace-titlebar {
    display: flex;
    min-height: 2.45rem;
    align-items: center;
    justify-content: space-between;
    gap: var(--space-3);
    padding: 0.55rem 0.75rem;
    border-bottom: 2px solid var(--line);
    background: linear-gradient(90deg, var(--titlebar-start), var(--titlebar-end));
  }

  .workspace-titlebar__copy {
    display: flex;
    min-width: 0;
    align-items: baseline;
    gap: var(--space-2);
  }

  .workspace-titlebar h2 {
    overflow: hidden;
    font-family: var(--font-ui);
    font-size: 0.78rem;
    font-weight: 800;
    letter-spacing: 0.04em;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  .workspace-titlebar__actions {
    flex: none;
    font-family: var(--font-ui);
    font-size: 0.72rem;
  }

  .workspace-frame__body {
    padding: var(--space-4);
  }
}
```

Keep later feature-list, metadata, inspector, and responsive rules for now; subsequent tasks replace their panel assumptions.

- [ ] **Step 6: Import the token layer first**

Replace the CSS imports in `src/main.tsx` with:

```ts
import "maplibre-gl/dist/maplibre-gl.css";
import "./theme.css";
import "./styles.css";
import "./resolution.css";
import "./map.css";
```

Create an empty `src/map.css` containing:

```css
@layer map;
```

This keeps the import graph valid until Task 6 moves map rules.

- [ ] **Step 7: Run tests, typecheck, and formatting**

Run:

```bash
pnpm test
pnpm check
pnpm format
pnpm format:check
```

Expected: all commands PASS.

- [ ] **Step 8: Commit shared chrome**

```bash
git add src/theme.css src/map.css src/main.tsx src/styles.css src/components/WorkspaceFrame.tsx src/components/WorkspaceFrame.test.tsx src/components/CompactStatusStrip.tsx src/components/CompactStatusStrip.test.tsx src/components/AppHeader.tsx src/components/AppHeader.test.tsx
git commit -m "feat: add shared map-native workspace chrome"
```

---

### Task 4: Lift resolution state and separate query from authority presentation

**Files:**
- Modify: `src/components/ResolutionPanel.tsx`
- Create: `src/components/ResolutionPanel.test.tsx`
- Create: `src/components/AuthorityRecord.tsx`
- Create: `src/components/AuthorityRecord.test.tsx`
- Modify: `src/App.tsx`
- Test: `src/App.test.tsx`

**Interfaces:**
- `ResolutionPanel` consumes `onResultChange(result: ResolutionResult | null)`.
- `ResolutionPanel` retains coordinate/form/loading/error behavior and no longer renders the result.
- `AuthorityRecord` consumes `result: ResolutionResult | null`.
- `App` owns `resolutionResult: ResolutionResult | null`.

- [ ] **Step 1: Write failing query-callback and authority-state tests**

Create `src/components/ResolutionPanel.test.tsx`:

```tsx
import { render, screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { describe, expect, it, vi } from "vitest";
import { fetchResolution } from "../lib/api";
import { resolvedResult } from "../test/fixtures";
import { ResolutionPanel } from "./ResolutionPanel";

vi.mock("../lib/api", () => ({
  fetchResolution: vi.fn(),
}));

describe("ResolutionPanel", () => {
  it("reports a successful result to its owner without rendering authority output", async () => {
    const user = userEvent.setup();
    const onResultChange = vi.fn();
    vi.mocked(fetchResolution).mockResolvedValue(resolvedResult);

    render(<ResolutionPanel onResultChange={onResultChange} />);

    await user.click(screen.getByRole("button", { name: "Resolve location" }));

    await waitFor(() => {
      expect(onResultChange).toHaveBeenCalledWith(resolvedResult);
    });
    expect(screen.queryByText("City and County of Denver")).toBeNull();
  });
});
```

Create `src/components/AuthorityRecord.test.tsx`:

```tsx
import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";
import { ambiguousResult, conflictingResult, resolvedResult } from "../test/fixtures";
import { AuthorityRecord } from "./AuthorityRecord";

describe("AuthorityRecord", () => {
  it("renders authority, adoption, and source-backed evidence", () => {
    render(<AuthorityRecord result={resolvedResult} />);

    expect(screen.getByRole("heading", { name: "City and County of Denver" })).toBeTruthy();
    expect(screen.getByText("2024 Denver Building Code")).toBeTruthy();
    expect(screen.getByRole("link", { name: "Denver Building and Fire Code" })).toBeTruthy();
    expect(screen.getByText("Resolved")).toBeTruthy();
  });

  it("states ambiguity and conflict in text", () => {
    const { rerender } = render(<AuthorityRecord result={ambiguousResult} />);
    expect(screen.getByText("Ambiguous")).toBeTruthy();
    expect(screen.getByText(/local record is required/i)).toBeTruthy();

    rerender(<AuthorityRecord result={conflictingResult} />);
    expect(screen.getByText("Conflicting evidence")).toBeTruthy();
    expect(screen.getByText(/official sources disagree/i)).toBeTruthy();
  });
});
```

Run:

```bash
pnpm vitest run src/components/ResolutionPanel.test.tsx src/components/AuthorityRecord.test.tsx
```

Expected: FAIL because the controlled interface and `AuthorityRecord` do not exist.

- [ ] **Step 2: Convert `ResolutionPanel` to query-only behavior**

In `src/components/ResolutionPanel.tsx`:

1. Change the React import to:

```ts
import { useState, type FormEvent } from "react";
```

2. Change the component declaration to:

```tsx
interface ResolutionPanelProps {
  onResultChange: (result: ResolutionResult | null) => void;
}

export function ResolutionPanel({ onResultChange }: ResolutionPanelProps): JSX.Element {
```

3. Remove the local `result` state declaration.

4. After numeric validation succeeds and before setting `isLoading`, clear any stale result:

```ts
onResultChange(null);
setIsLoading(true);
setError(null);
```

5. In `handleSubmit`, replace success and failure result assignments with:

```ts
const nextResult = await fetchResolution({
  longitude: nextLongitude,
  latitude: nextLatitude,
  codeFamily,
  projectType: projectType || undefined,
  applicabilityDate,
});
onResultChange(nextResult);
```

and:

```ts
onResultChange(null);
setError(requestError instanceof Error ? requestError.message : "Resolution failed.");
```

6. Replace the returned outer element and remove result rendering:

```tsx
return (
  <div className="resolution-panel">
    <p className="workspace-intro">
      Enter a coordinate and applicability date to identify the controlling state policy, likely
      authorities, evidence, and records that still require local verification.
    </p>

    <form className="resolution-form" onSubmit={handleSubmit}>
      <div className="resolution-form__coordinates">
        <label>
          <span>Longitude</span>
          <input
            inputMode="decimal"
            value={longitude}
            onChange={(event) => setLongitude(event.target.value)}
            required
          />
        </label>
        <label>
          <span>Latitude</span>
          <input
            inputMode="decimal"
            value={latitude}
            onChange={(event) => setLatitude(event.target.value)}
            required
          />
        </label>
      </div>
      <label>
        <span>Code family</span>
        <select value={codeFamily} onChange={(event) => setCodeFamily(event.target.value)}>
          {codeFamilies.map(([value, label]) => (
            <option key={value} value={value}>
              {label}
            </option>
          ))}
        </select>
      </label>
      <label>
        <span>Project context</span>
        <select value={projectType} onChange={(event) => setProjectType(event.target.value)}>
          {projectTypes.map(([value, label]) => (
            <option key={value || "ordinary"} value={value}>
              {label}
            </option>
          ))}
        </select>
      </label>
      <label>
        <span>Applicability date</span>
        <input
          type="date"
          value={applicabilityDate}
          onChange={(event) => setApplicabilityDate(event.target.value)}
          required
        />
      </label>
      <button className="resolution-form__submit" type="submit" disabled={isLoading}>
        {isLoading ? "Resolving…" : "Resolve location"}
      </button>
    </form>

    {error ? (
      <p className="resolution-message resolution-message--error" role="alert">
        {error}
      </p>
    ) : null}
  </div>
);
```

Delete the old `ResolutionSummary`, `ResultSection`, and `formatClaimValue` functions from this file.

- [ ] **Step 3: Implement `AuthorityRecord`**

Create `src/components/AuthorityRecord.tsx`:

```tsx
import type { ReactNode } from "react";
import type { ResolutionClaim, ResolutionResult, ResolutionStatus } from "../types";

const statusLabels: Record<ResolutionStatus, string> = {
  resolved: "Resolved",
  partially_resolved: "Partially resolved",
  local_record_required: "Local record required",
  ambiguous: "Ambiguous",
  conflicting: "Conflicting evidence",
  insufficient_evidence: "Insufficient evidence",
};

function formatClaimValue(claim: ResolutionClaim): string {
  if (claim.value === undefined || claim.value === null) {
    return claim.conflictGroup ? `Conflict group: ${claim.conflictGroup}` : "No normalized value";
  }

  const value = typeof claim.value === "string" ? claim.value : JSON.stringify(claim.value);
  return claim.conflictGroup ? `${value} · conflict group ${claim.conflictGroup}` : value;
}

function ResultSection({
  title,
  children,
}: {
  title: string;
  children: ReactNode;
}): JSX.Element {
  return (
    <section className="authority-section">
      <h3>{title}</h3>
      {children}
    </section>
  );
}

export function AuthorityRecord({
  result,
}: {
  result: ResolutionResult | null;
}): JSX.Element {
  if (!result) {
    return (
      <div className="authority-empty">
        <h3>No authority result yet</h3>
        <p>Run a location query to assemble authority, adoption, amendment, and evidence records.</p>
      </div>
    );
  }

  const place =
    result.geography.municipality?.name ??
    result.geography.county?.name ??
    result.geography.stateName ??
    "Matched location";
  const primaryAuthority = result.authorityCandidates[0];

  return (
    <div className="authority-record" aria-live="polite">
      <div className="authority-record__summary">
        <div>
          <span className="authority-record__place">{place}</span>
          <h3>{primaryAuthority?.name ?? "Authority requires local confirmation"}</h3>
          <p>
            {result.codeFamily ? result.codeFamily.replace(/_/g, " ") : "All code families"}
            {result.applicabilityDate ? ` · applicable ${result.applicabilityDate}` : ""}
          </p>
        </div>
        <span className={`resolution-status resolution-status--${result.status}`}>
          {statusLabels[result.status]}
        </span>
      </div>

      <ResultSection title="Authority candidates">
        {result.authorityCandidates.length > 0 ? (
          <ul className="authority-list">
            {result.authorityCandidates.map((candidate) => (
              <li key={`${candidate.kind}:${candidate.authorityId ?? candidate.name}`}>
                <strong>{candidate.name}</strong>
                <span>{candidate.roles.join(", ") || candidate.kind.replace(/_/g, " ")}</span>
              </li>
            ))}
          </ul>
        ) : (
          <p>No supported authority candidate is available.</p>
        )}
      </ResultSection>

      {result.adoptions.length > 0 ? (
        <ResultSection title="Adopted codes">
          <ul className="authority-list">
            {result.adoptions.map((adoption) => (
              <li key={adoption.id}>
                <strong>{adoption.stateCodeName}</strong>
                <span>
                  {adoption.codeFamily.replace(/_/g, " ")}
                  {adoption.dates.effective_date
                    ? ` · effective ${adoption.dates.effective_date}`
                    : ""}
                </span>
              </li>
            ))}
          </ul>
        </ResultSection>
      ) : null}

      {result.applicableRules.length > 0 ? (
        <ResultSection title="Applicable rules and amendments">
          <ul className="authority-list">
            {result.applicableRules.map((rule) => (
              <li key={rule.id}>
                <strong>{rule.kind.replace(/_/g, " ")}</strong>
                <span>{rule.summary}</span>
              </li>
            ))}
          </ul>
        </ResultSection>
      ) : null}

      {result.authorityPath.length > 0 ? (
        <ResultSection title="Authority path">
          <ul className="authority-list">
            {result.authorityPath.map((relationship) => (
              <li key={relationship.id}>
                <strong>
                  {relationship.fromId} → {relationship.to}
                </strong>
                <span>{relationship.summary ?? relationship.relationship.replace(/_/g, " ")}</span>
              </li>
            ))}
          </ul>
        </ResultSection>
      ) : null}

      {result.supportingClaims.length > 0 ? (
        <ResultSection title="Supporting claims">
          <ul className="authority-list">
            {result.supportingClaims.map((claim) => (
              <li key={claim.id}>
                <strong>
                  {claim.field} · {claim.status}
                </strong>
                <span>{formatClaimValue(claim)}</span>
              </li>
            ))}
          </ul>
        </ResultSection>
      ) : null}

      {result.requiredLocalRecords.length > 0 ? (
        <ResultSection title="Still verify locally">
          <ul className="authority-list">
            {result.requiredLocalRecords.map((record) => (
              <li key={record}>
                <strong>{record}</strong>
                <span>Obtain the current controlling local record before relying on the result.</span>
              </li>
            ))}
          </ul>
        </ResultSection>
      ) : null}

      {result.warnings.length > 0 ? (
        <div className="resolution-message resolution-message--warning" role="status">
          {result.warnings.map((warning) => (
            <p key={warning}>{warning}</p>
          ))}
        </div>
      ) : null}

      <details className="resolution-evidence" open={result.status !== "resolved"}>
        <summary>Source evidence ({result.evidence.length})</summary>
        <ul>
          {result.evidence.map((source) => (
            <li key={source.id}>
              <a href={source.url} target="_blank" rel="noreferrer">
                {source.title}
              </a>
              <span>
                {source.kind.replace(/_/g, " ")} · accessed {source.accessedAt}
                {source.lastCheckedAt ? ` · checked ${source.lastCheckedAt}` : ""}
                {source.availability ? ` · ${source.availability}` : ""}
              </span>
              {source.caveat ? <p>{source.caveat}</p> : null}
            </li>
          ))}
        </ul>
      </details>
    </div>
  );
}
```

- [ ] **Step 4: Lift resolution state into `App`**

In `src/App.tsx`:

1. Import `AuthorityRecord`, `useTheme`, `ResolutionResult`, and `AppHeader`.
2. Add state:

```ts
const [resolutionResult, setResolutionResult] = useState<ResolutionResult | null>(null);
const { theme, toggleTheme } = useTheme();
```

3. Pass the controlled callback:

```tsx
<ResolutionPanel onResultChange={setResolutionResult} />
```

4. Render the authority content in the right-side area:

```tsx
<AuthorityRecord result={resolutionResult} />
```

5. Replace the current header with:

```tsx
<AppHeader
  theme={theme}
  onToggleTheme={toggleTheme}
  layerCount={layerRegistry.length}
  recordCount={boundaryFeatures.length}
  refreshStatus={refreshStatus}
  mainSiteUrl={import.meta.env.VITE_MAIN_SITE_URL?.trim() || undefined}
/>
```

At this step, keep the current column wrappers so the feature compiles; Task 5 performs the full workspace recompose.

- [ ] **Step 5: Add an App ownership regression test**

Create `src/App.test.tsx`:

```tsx
import { render, screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { beforeEach, describe, expect, it, vi } from "vitest";
import App from "./App";
import {
  fetchBoundaryFeatures,
  fetchLayers,
  fetchRefreshStatus,
  fetchResolution,
  getApiBaseUrl,
} from "./lib/api";
import {
  demoBoundaryFeatures,
  demoLayers,
  demoRefreshStatus,
  resolvedResult,
} from "./test/fixtures";

vi.mock("./lib/api", () => ({
  fetchLayers: vi.fn(),
  fetchRefreshStatus: vi.fn(),
  fetchBoundaryFeatures: vi.fn(),
  fetchResolution: vi.fn(),
  getApiBaseUrl: vi.fn(),
}));

vi.mock("./components/MapStage", () => ({
  MapStage: () => <div data-testid="map-stage">Map stage</div>,
}));

describe("App", () => {
  beforeEach(() => {
    vi.mocked(fetchLayers).mockResolvedValue(demoLayers);
    vi.mocked(fetchRefreshStatus).mockResolvedValue(demoRefreshStatus);
    vi.mocked(fetchBoundaryFeatures).mockResolvedValue(demoBoundaryFeatures);
    vi.mocked(fetchResolution).mockResolvedValue(resolvedResult);
    vi.mocked(getApiBaseUrl).mockReturnValue("/api");
  });

  it("owns the active authority result after the query resolves", async () => {
    const user = userEvent.setup();
    render(<App />);

    await waitFor(() => {
      expect(screen.getByText("Colorado")).toBeTruthy();
    });

    await user.click(screen.getByRole("button", { name: "Resolve location" }));

    await waitFor(() => {
      expect(screen.getByRole("heading", { name: "City and County of Denver" })).toBeTruthy();
    });
  });
});
```

Run:

```bash
pnpm vitest run src/components/ResolutionPanel.test.tsx src/components/AuthorityRecord.test.tsx src/App.test.tsx
```

Expected: PASS.

- [ ] **Step 6: Commit the resolution split**

```bash
git add src/components/ResolutionPanel.tsx src/components/ResolutionPanel.test.tsx src/components/AuthorityRecord.tsx src/components/AuthorityRecord.test.tsx src/App.tsx src/App.test.tsx
git commit -m "refactor: separate authority results from location queries"
```

---

### Task 5: Recompose the desktop and tablet workspace without nested panels

**Files:**
- Create: `src/components/FeatureSelectionList.tsx`
- Create: `src/components/BackendContractDetails.tsx`
- Modify: `src/components/MapStage.tsx`
- Modify: `src/components/FeatureInspector.tsx`
- Modify: `src/components/LayerToggleList.tsx`
- Modify: `src/components/StatusBanner.tsx`
- Modify: `src/App.tsx`
- Modify: `src/styles.css`
- Test: `src/App.test.tsx`

**Interfaces:**
- `MapStage` becomes map-only and no longer renders cached feature cards or endpoint documentation.
- `FeatureSelectionList` consumes `featureSummaries`, `selectedFeatureId`, and `onSelectFeature`.
- `BackendContractDetails` is a static technical disclosure body.
- Existing domain props and callbacks remain unchanged.

- [ ] **Step 1: Extract cached feature selection**

Create `src/components/FeatureSelectionList.tsx`:

```tsx
import type { FeatureSummary } from "../types";

interface FeatureSelectionListProps {
  features: FeatureSummary[];
  selectedFeatureId: string | null;
  onSelectFeature: (featureId: string) => void;
}

export function FeatureSelectionList({
  features,
  selectedFeatureId,
  onSelectFeature,
}: FeatureSelectionListProps): JSX.Element {
  return (
    <div className="feature-list">
      {features.map((feature) => {
        const active = selectedFeatureId === feature.featureId;

        return (
          <button
            className={`feature-card ${active ? "is-active" : ""}`}
            key={`${feature.layerFamily}-${feature.featureId}`}
            onClick={() => onSelectFeature(feature.featureId)}
            type="button"
            aria-pressed={active}
          >
            <strong>{feature.title}</strong>
            <span>{feature.subtitle}</span>
            <small>
              {feature.layerFamily} · {feature.sourceId}
            </small>
          </button>
        );
      })}
    </div>
  );
}
```

Create `src/components/BackendContractDetails.tsx`:

```tsx
export function BackendContractDetails(): JSX.Element {
  return (
    <dl className="meta-grid meta-grid--stack">
      <div>
        <dt>Health</dt>
        <dd>/health</dd>
      </div>
      <div>
        <dt>Readiness</dt>
        <dd>/ready</dd>
      </div>
      <div>
        <dt>Layer registry</dt>
        <dd>/layers</dd>
      </div>
      <div>
        <dt>Feature lookup</dt>
        <dd>{"/features/{layer_family}/{feature_id}"}</dd>
      </div>
      <div>
        <dt>Boundary list</dt>
        <dd>/boundaries</dd>
      </div>
      <div>
        <dt>Resolution</dt>
        <dd>/resolve</dd>
      </div>
      <div>
        <dt>Refresh status</dt>
        <dd>/refresh/status</dd>
      </div>
    </dl>
  );
}
```

- [ ] **Step 2: Make `MapStage` map-only**

In `src/components/MapStage.tsx`:

1. Remove the `featureSummaries` prop and all feature-list/backend-contract JSX.
2. Replace the return with:

```tsx
return (
  <div className="map-stage">
    <div className="map-stage__topline">
      <div>
        <p className="eyebrow">OpenStreetMap basemap</p>
        <h3>TIGERweb and NERIS boundaries</h3>
        <p className="lede">
          Cached boundaries remain selectable while full source and authority records stay visible
          beside the map.
        </p>
      </div>
      <span className={`semantic-status semantic-status--${refreshStatus.status}`}>
        {refreshStatus.status === "ok" ? "Mirror healthy" : "Mirror attention needed"}
      </span>
    </div>

    <div className="map-canvas">
      <div className="map-canvas__surface">
        <div ref={mapContainerRef} className="map-canvas__map" />

        <div className="map-canvas__overlay">
          <span className="map-canvas__label">
            {mapStatus === "ready"
              ? "Boundary map ready"
              : mapStatus === "error"
                ? "Map unavailable"
                : "Loading map"}
          </span>
          <p className="map-canvas__hint">
            {mapStatus === "error"
              ? (mapError ?? "The map surface could not be initialized.")
              : selectedFeature
                ? `Selected: ${selectedFeature.title}`
                : "Select a boundary to inspect its source record."}
          </p>
        </div>
      </div>
    </div>
  </div>
);
```

- [ ] **Step 3: Remove legacy panel wrappers from leaf components**

Replace the outer wrappers as follows:

`LayerToggleList.tsx` returns:

```tsx
<div className="toggle-list">
  {layers.map((layer) => {
    const active = Boolean(enabledLayers[layer.key]);

    return (
      <label className={`toggle-row ${active ? "is-active" : ""}`} key={layer.key}>
        <span className="toggle-row__copy">
          <strong>{layer.label}</strong>
          <span>{layer.description}</span>
        </span>
        <input
          type="checkbox"
          checked={active}
          onChange={() => onToggle(layer.key)}
          aria-label={`Toggle ${layer.label}`}
        />
      </label>
    );
  })}
</div>
```

`StatusBanner.tsx` returns one compact block for the unavailable state:

```tsx
<div className="data-status data-status--warning" role="status">
  <div>
    <strong>{error ?? "Loading cached boundary data."}</strong>
    <p>Waiting for the local runtime snapshot to load.</p>
  </div>
  <dl>
    <dt>API</dt>
    <dd>{apiBaseUrl}</dd>
  </dl>
</div>
```

and one for a loaded state:

```tsx
<div className={`data-status data-status--${refreshStatus.status}`} role="status">
  <div>
    <strong>{error ?? refreshStatus.message}</strong>
    <p>
      Latest success: {refreshStatus.latestSuccessfulRefresh ?? "none"} · Next scheduled:{" "}
      {refreshStatus.nextScheduledRefresh}
    </p>
  </div>
  <dl>
    <dt>API</dt>
    <dd>{apiBaseUrl}</dd>
  </dl>
</div>
```

`FeatureInspector.tsx` keeps its loading, error, populated, and empty branches, but each branch returns `<div className="feature-inspector">` instead of `<aside className="panel inspector">`. Replace the repeated panel header with a short intro:

```tsx
<p className="workspace-intro">
  Complete raw boundary fields keyed by the selected feature ID.
</p>
```

- [ ] **Step 4: Recompose `App` with explicit grid areas**

Replace the `main` subtree in `src/App.tsx` with:

```tsx
<main className="app-workspace">
  <div className="workspace-column workspace-column--query">
    <WorkspaceFrame id="query-location" title="QUERY.LOCATION" eyebrow="Input">
      <ResolutionPanel onResultChange={setResolutionResult} />
    </WorkspaceFrame>

    <WorkspaceFrame id="data-status" title="DATA.STATUS" eyebrow="Runtime">
      <StatusBanner refreshStatus={refreshStatus} apiBaseUrl={apiBaseUrl} error={loadError} />
    </WorkspaceFrame>

    <WorkspaceFrame id="map-layers" title="MAP.LAYERS" eyebrow="Visibility">
      <LayerToggleList
        layers={layerRegistry}
        enabledLayers={enabledLayers}
        onToggle={handleToggleLayer}
      />
    </WorkspaceFrame>
  </div>

  <WorkspaceFrame
    id="map-view"
    title="MAP.VIEW"
    eyebrow="Geography"
    className="workspace-frame--map"
  >
    <MapStage
      layers={layerRegistry}
      selectedFeature={selectedFeature}
      enabledLayers={enabledLayers}
      onSelectFeature={setSelectedFeatureId}
      refreshStatus={refreshStatus ?? loadingRefreshStatus}
      boundaryFeatures={boundaryFeatures}
    />
  </WorkspaceFrame>

  <div className="workspace-column workspace-column--record">
    <WorkspaceFrame id="authority-record" title="AUTHORITY.RECORD" eyebrow="Result">
      <AuthorityRecord result={resolutionResult} />
    </WorkspaceFrame>

    <WorkspaceFrame id="technical-record" title="TECHNICAL.RECORD" eyebrow="Evidence">
      <FeatureInspector
        feature={selectedFeature}
        selectedFeature={selectedSummary}
        isLoading={isLoading}
        error={loadError}
      />

      <details className="technical-disclosure">
        <summary>Cached boundary choices</summary>
        <FeatureSelectionList
          features={boundaryFeatures.map(
            ({ layerFamily, featureId, title, subtitle, sourceId }) => ({
              layerFamily,
              featureId,
              title,
              subtitle,
              sourceId,
            }),
          )}
          selectedFeatureId={selectedFeatureId}
          onSelectFeature={setSelectedFeatureId}
        />
      </details>

      <details className="technical-disclosure">
        <summary>Backend contract</summary>
        <BackendContractDetails />
      </details>
    </WorkspaceFrame>
  </div>
</main>
```

Add imports for `WorkspaceFrame`, `FeatureSelectionList`, and `BackendContractDetails`. Remove the now-unused `featureSummaries` prop passed to `MapStage`.

- [ ] **Step 5: Replace desktop and tablet layout rules**

Add to the `layout` and `responsive` layers in `src/styles.css`:

```css
@layer layout {
  .app-workspace {
    display: grid;
    grid-template-columns: minmax(17rem, 20rem) minmax(34rem, 1fr) minmax(20rem, 25rem);
    gap: var(--space-5);
    align-items: start;
    padding-block: var(--space-5);
  }

  .workspace-column {
    display: grid;
    gap: var(--space-5);
    min-width: 0;
  }

  .workspace-column--record {
    position: sticky;
    top: var(--space-4);
  }

  .workspace-frame--map {
    min-height: calc(100vh - 8.5rem);
  }

  .workspace-frame--map .workspace-frame__body {
    padding: var(--space-3);
  }

  .workspace-intro {
    margin-bottom: var(--space-4);
    color: var(--ink-soft);
    font-size: 0.92rem;
  }

  .technical-disclosure {
    margin-top: var(--space-4);
    border-top: 1px solid var(--line-soft);
    padding-top: var(--space-3);
  }

  .technical-disclosure summary {
    cursor: pointer;
    color: var(--blue);
    font-family: var(--font-ui);
    font-size: 0.78rem;
    font-weight: 800;
    text-transform: uppercase;
  }
}

@layer responsive {
  @media (max-width: 82rem) {
    .app-workspace {
      grid-template-columns: minmax(17rem, 20rem) minmax(0, 1fr);
    }

    .workspace-column--record {
      position: static;
      grid-column: 1 / -1;
      grid-template-columns: repeat(2, minmax(0, 1fr));
    }
  }

  @media (max-width: 58rem) {
    .app-header {
      grid-template-columns: 1fr auto;
    }

    .compact-status {
      grid-column: 1 / -1;
      grid-row: 2;
      justify-content: flex-start;
      flex-wrap: wrap;
    }

    .app-workspace {
      grid-template-columns: 1fr;
    }

    .workspace-column--record {
      grid-column: auto;
      grid-template-columns: 1fr;
    }

    .workspace-frame--map {
      min-height: auto;
    }
  }
}
```

Delete the old `.app-grid`, `.sidebar`, `.panel`, `.panel--*`, and sticky `.inspector` layout rules after all references are gone.

- [ ] **Step 6: Strengthen the App layout test**

Append to `src/App.test.tsx`:

```tsx
it("renders one framed query, map, authority, and technical workspace", async () => {
  render(<App />);

  await waitFor(() => {
    expect(screen.getByRole("region", { name: "QUERY.LOCATION" })).toBeTruthy();
  });

  expect(screen.getByRole("region", { name: "MAP.VIEW" })).toBeTruthy();
  expect(screen.getByRole("region", { name: "AUTHORITY.RECORD" })).toBeTruthy();
  expect(screen.getByRole("region", { name: "TECHNICAL.RECORD" })).toBeTruthy();
  expect(screen.getAllByText("Backend contract")).toHaveLength(1);
});
```

Run:

```bash
pnpm vitest run src/App.test.tsx
pnpm check
pnpm test
```

Expected: PASS.

- [ ] **Step 7: Commit the workspace recompose**

```bash
git add src/App.tsx src/App.test.tsx src/components/MapStage.tsx src/components/FeatureInspector.tsx src/components/LayerToggleList.tsx src/components/StatusBanner.tsx src/components/FeatureSelectionList.tsx src/components/BackendContractDetails.tsx src/styles.css
git commit -m "feat: recompose the regulatory GIS workspace"
```

---

### Task 6: Centralize theme-aware MapLibre paint and quiet the basemap

**Files:**
- Create: `src/lib/mapTheme.ts`
- Create: `src/lib/mapTheme.test.ts`
- Modify: `src/components/MapStage.tsx`
- Create: `src/map.css`
- Test: existing `src/lib/mapData.test.ts`

**Interfaces:**
- Produces: `MapThemePalette`.
- Produces: `readMapThemePalette(root?)`.
- `MapStage` consumes `theme: ThemeName`.
- Map layer IDs remain stable so click selection and visibility updates do not regress.

- [ ] **Step 1: Write the failing palette adapter test**

Create `src/lib/mapTheme.test.ts`:

```ts
import { describe, expect, it } from "vitest";
import { readMapThemePalette } from "./mapTheme";

describe("readMapThemePalette", () => {
  it("reads semantic boundary and basemap values from the root token layer", () => {
    const root = document.documentElement;
    const tokens: Record<string, string> = {
      "--boundary-states-fill": "rgba(1, 2, 3, 0.2)",
      "--boundary-states-line": "#010203",
      "--boundary-counties-fill": "rgba(2, 3, 4, 0.2)",
      "--boundary-counties-line": "#020304",
      "--boundary-municipalities-fill": "rgba(3, 4, 5, 0.2)",
      "--boundary-municipalities-line": "#030405",
      "--boundary-special-fill": "rgba(4, 5, 6, 0.2)",
      "--boundary-special-line": "#040506",
      "--boundary-tribal-fill": "rgba(5, 6, 7, 0.2)",
      "--boundary-tribal-line": "#050607",
      "--boundary-neris-fill": "rgba(6, 7, 8, 0.2)",
      "--boundary-neris-line": "#060708",
      "--boundary-selected": "#ffffff",
      "--boundary-selected-halo": "#ff00aa",
      "--map-raster-opacity": "0.66",
      "--map-raster-saturation": "-0.8",
      "--map-raster-contrast": "0.1",
    };
    for (const [name, value] of Object.entries(tokens)) {
      root.style.setProperty(name, value);
    }

    const palette = readMapThemePalette(root);

    expect(palette.families.states.fillColor).toBe("rgba(1, 2, 3, 0.2)");
    expect(palette.selectedColor).toBe("#ffffff");
    expect(palette.selectedHaloColor).toBe("#ff00aa");
    expect(palette.rasterOpacity).toBe(0.66);
    expect(palette.rasterSaturation).toBe(-0.8);
  });
});
```

Run:

```bash
pnpm vitest run src/lib/mapTheme.test.ts
```

Expected: FAIL because `mapTheme.ts` does not exist.

- [ ] **Step 2: Implement the CSS-variable adapter**

Create `src/lib/mapTheme.ts`:

```ts
import type { LayerFamilyKey } from "../types";

interface BoundaryPaint {
  fillColor: string;
  lineColor: string;
}

export interface MapThemePalette {
  families: Record<LayerFamilyKey, BoundaryPaint>;
  selectedColor: string;
  selectedHaloColor: string;
  rasterOpacity: number;
  rasterSaturation: number;
  rasterContrast: number;
}

function requiredToken(styles: CSSStyleDeclaration, name: string): string {
  const value = styles.getPropertyValue(name).trim();
  if (!value) {
    throw new Error(`Missing required map theme token ${name}`);
  }
  return value;
}

function numericToken(styles: CSSStyleDeclaration, name: string): number {
  const value = Number(requiredToken(styles, name));
  if (!Number.isFinite(value)) {
    throw new Error(`Map theme token ${name} must be numeric`);
  }
  return value;
}

export function readMapThemePalette(
  root: HTMLElement = document.documentElement,
): MapThemePalette {
  const styles = getComputedStyle(root);

  return {
    families: {
      states: {
        fillColor: requiredToken(styles, "--boundary-states-fill"),
        lineColor: requiredToken(styles, "--boundary-states-line"),
      },
      counties: {
        fillColor: requiredToken(styles, "--boundary-counties-fill"),
        lineColor: requiredToken(styles, "--boundary-counties-line"),
      },
      municipalities: {
        fillColor: requiredToken(styles, "--boundary-municipalities-fill"),
        lineColor: requiredToken(styles, "--boundary-municipalities-line"),
      },
      special_areas: {
        fillColor: requiredToken(styles, "--boundary-special-fill"),
        lineColor: requiredToken(styles, "--boundary-special-line"),
      },
      tribal_areas: {
        fillColor: requiredToken(styles, "--boundary-tribal-fill"),
        lineColor: requiredToken(styles, "--boundary-tribal-line"),
      },
      neris_jurisdictions: {
        fillColor: requiredToken(styles, "--boundary-neris-fill"),
        lineColor: requiredToken(styles, "--boundary-neris-line"),
      },
    },
    selectedColor: requiredToken(styles, "--boundary-selected"),
    selectedHaloColor: requiredToken(styles, "--boundary-selected-halo"),
    rasterOpacity: numericToken(styles, "--map-raster-opacity"),
    rasterSaturation: numericToken(styles, "--map-raster-saturation"),
    rasterContrast: numericToken(styles, "--map-raster-contrast"),
  };
}
```

Add these tokens to each theme block in `src/theme.css`:

```css
--map-raster-opacity: 0.7;
--map-raster-saturation: -0.72;
--map-raster-contrast: -0.08;
```

for light and:

```css
--map-raster-opacity: 0.48;
--map-raster-saturation: -0.86;
--map-raster-contrast: 0.12;
```

for dark.

Run:

```bash
pnpm vitest run src/lib/mapTheme.test.ts
```

Expected: PASS.

- [ ] **Step 3: Replace hard-coded MapLibre paints**

In `src/components/MapStage.tsx`:

1. Import:

```ts
import type { ThemeName } from "../lib/theme";
import { readMapThemePalette } from "../lib/mapTheme";
```

2. Add `theme: ThemeName` to `MapStageProps` and destructuring.
3. In `src/App.tsx`, add `theme={theme}` to the `MapStage` call.
4. Delete `familyStyling` and `getLayerPaints`.
5. In `createMapStyle`, read the palette and set raster paint:

```ts
function createMapStyle(): StyleSpecification {
  const palette = readMapThemePalette();

  return {
    version: 8,
    sources: {
      osm: {
        type: "raster",
        tiles: ["https://tile.openstreetmap.org/{z}/{x}/{y}.png"],
        tileSize: 256,
        attribution: "© OpenStreetMap contributors",
      },
    },
    layers: [
      {
        id: "osm-basemap",
        type: "raster",
        source: "osm",
        paint: {
          "raster-opacity": palette.rasterOpacity,
          "raster-saturation": palette.rasterSaturation,
          "raster-contrast": palette.rasterContrast,
        },
      },
    ],
  };
}
```

6. During map load, use:

```ts
const palette = readMapThemePalette();
const paints = palette.families[layer.key];
```

7. Add a selected halo layer before the selected core layer:

```ts
map.addLayer({
  id: `${sourceId}-${layer.key}-selected-halo`,
  type: "line",
  source: sourceId,
  filter: [
    "all",
    ["==", ["get", "layerFamily"], layer.key],
    ["==", ["get", "selected"], true],
  ],
  paint: {
    "line-color": palette.selectedHaloColor,
    "line-width": 6,
    "line-opacity": 0.75,
  },
  layout: { visibility },
});
```

8. Use `palette.selectedColor` and `line-width: 3` for the selected core layer.
9. Include `selected-halo` in visibility updates.
10. Add a theme effect:

```ts
useEffect(() => {
  const map = mapRef.current;
  if (!map || mapStatus !== "ready") {
    return;
  }

  const palette = readMapThemePalette();
  map.setPaintProperty("osm-basemap", "raster-opacity", palette.rasterOpacity);
  map.setPaintProperty("osm-basemap", "raster-saturation", palette.rasterSaturation);
  map.setPaintProperty("osm-basemap", "raster-contrast", palette.rasterContrast);

  for (const layer of layers) {
    const paints = palette.families[layer.key];
    map.setPaintProperty(`${sourceId}-${layer.key}-fill`, "fill-color", paints.fillColor);
    map.setPaintProperty(`${sourceId}-${layer.key}-fill`, "fill-outline-color", paints.lineColor);
    map.setPaintProperty(`${sourceId}-${layer.key}-line`, "line-color", paints.lineColor);
    map.setPaintProperty(
      `${sourceId}-${layer.key}-selected-halo`,
      "line-color",
      palette.selectedHaloColor,
    );
    map.setPaintProperty(`${sourceId}-${layer.key}-selected`, "line-color", palette.selectedColor);
  }
}, [layers, mapStatus, theme]);
```

- [ ] **Step 4: Move map-specific CSS into `src/map.css`**

Replace `src/map.css` with:

```css
@layer map {
  .map-stage {
    display: grid;
    gap: var(--space-3);
  }

  .map-stage__topline {
    display: flex;
    align-items: start;
    justify-content: space-between;
    gap: var(--space-4);
  }

  .map-stage__topline h3 {
    font-size: clamp(1.35rem, 2vw, 2rem);
  }

  .eyebrow {
    margin-bottom: var(--space-1);
    color: var(--ink-soft);
    font-family: var(--font-ui);
    font-size: 0.68rem;
    font-weight: 800;
    letter-spacing: 0.12em;
    text-transform: uppercase;
  }

  .lede,
  .map-canvas__hint {
    color: var(--ink-soft);
    font-size: 0.9rem;
  }

  .map-canvas {
    position: relative;
    min-height: 42rem;
    overflow: hidden;
    border: 2px solid var(--line);
    border-radius: var(--radius-sm);
    background: var(--canvas-alt);
  }

  .map-canvas__surface,
  .map-canvas__map {
    position: absolute;
    inset: 0;
  }

  .map-canvas__overlay {
    position: absolute;
    top: var(--space-3);
    left: var(--space-3);
    z-index: 2;
    display: grid;
    max-width: min(25rem, calc(100% - 2 * var(--space-3)));
    gap: var(--space-2);
    padding: var(--space-3);
    border: 1px solid var(--line);
    background: var(--map-overlay-surface);
    box-shadow: 3px 3px 0 var(--line);
    color: var(--map-overlay-ink);
    backdrop-filter: blur(8px);
  }

  .map-canvas__label {
    font-family: var(--font-ui);
    font-size: 0.72rem;
    font-weight: 800;
    letter-spacing: 0.08em;
    text-transform: uppercase;
  }

  .maplibregl-ctrl-group {
    border: 1px solid var(--line);
    border-radius: var(--radius-sm);
    background: var(--surface);
    box-shadow: 2px 2px 0 var(--line);
  }

  .maplibregl-ctrl-group button {
    background-color: var(--surface);
  }

  .maplibregl-ctrl-attrib {
    color: var(--ink);
    background: var(--map-overlay-surface);
    font-family: var(--font-ui);
    font-size: 0.65rem;
  }

  .maplibregl-ctrl-attrib a {
    color: var(--blue);
  }

  @media (max-width: 58rem) {
    .map-canvas {
      min-height: 28rem;
    }
  }

  @media (max-width: 40rem) {
    .map-stage__topline {
      display: grid;
    }

    .map-canvas {
      min-height: 23rem;
    }
  }
}
```

Delete the moved `.map-stage*`, `.map-canvas*`, `.layer-pill`, and MapLibre-related rules from `src/styles.css`.

- [ ] **Step 5: Verify map data and theme behavior**

Run:

```bash
pnpm vitest run src/lib/mapTheme.test.ts src/lib/mapData.test.ts
pnpm check
pnpm test
pnpm build
```

Expected: all commands PASS; the production build includes the OSM attribution string.

- [ ] **Step 6: Commit map theming**

```bash
git add src/lib/mapTheme.ts src/lib/mapTheme.test.ts src/components/MapStage.tsx src/theme.css src/map.css src/styles.css
git commit -m "feat: theme MapLibre from semantic design tokens"
```

---

### Task 7: Implement keyed mobile ordering and disclosures without map remounts

**Files:**
- Create: `src/lib/workspaceOrder.ts`
- Create: `src/lib/workspaceOrder.test.ts`
- Create: `src/components/MobileMapDisclosure.tsx`
- Create: `src/components/MobileMapDisclosure.test.tsx`
- Create: `src/components/ResponsiveDisclosure.tsx`
- Create: `src/hooks/useMediaQuery.ts`
- Modify: `src/App.tsx`
- Modify: `src/styles.css`
- Modify: `src/resolution.css`
- Test: `src/App.test.tsx`

**Interfaces:**
- Produces: `WorkspaceSectionKey = "query" | "map" | "authority" | "technical"`.
- Produces: `getWorkspaceOrder(hasResolution)`.
- `MobileMapDisclosure` keeps children mounted and exposes `aria-expanded`.
- `ResponsiveDisclosure` keeps layer content mounted and is always visually expanded above the mobile breakpoint.
- `useMediaQuery("(max-width: 58rem)")` controls the `hidden` attribute only; it never conditionally renders the map subtree.

- [ ] **Step 1: Write failing ordering tests**

Create `src/lib/workspaceOrder.test.ts`:

```ts
import { describe, expect, it } from "vitest";
import { getWorkspaceOrder } from "./workspaceOrder";

describe("getWorkspaceOrder", () => {
  it("prioritizes query before a result exists", () => {
    expect(getWorkspaceOrder(false)).toEqual(["query", "map", "authority", "technical"]);
  });

  it("prioritizes the authority answer and evidence after resolution", () => {
    expect(getWorkspaceOrder(true)).toEqual(["authority", "map", "query", "technical"]);
  });
});
```

Run:

```bash
pnpm vitest run src/lib/workspaceOrder.test.ts
```

Expected: FAIL because the module does not exist.

- [ ] **Step 2: Implement keyed ordering**

Create `src/lib/workspaceOrder.ts`:

```ts
export type WorkspaceSectionKey = "query" | "map" | "authority" | "technical";

const PRE_QUERY_ORDER: WorkspaceSectionKey[] = ["query", "map", "authority", "technical"];
const POST_QUERY_ORDER: WorkspaceSectionKey[] = ["authority", "map", "query", "technical"];

export function getWorkspaceOrder(hasResolution: boolean): WorkspaceSectionKey[] {
  return [...(hasResolution ? POST_QUERY_ORDER : PRE_QUERY_ORDER)];
}
```

Run:

```bash
pnpm vitest run src/lib/workspaceOrder.test.ts
```

Expected: PASS.

- [ ] **Step 3: Write the failing map disclosure preservation test**

Create `src/components/MobileMapDisclosure.test.tsx`:

```tsx
import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { useState } from "react";
import { describe, expect, it, vi } from "vitest";
import { MobileMapDisclosure } from "./MobileMapDisclosure";

vi.mock("../hooks/useMediaQuery", () => ({
  useMediaQuery: () => true,
}));

function StatefulMapProbe(): JSX.Element {
  const [count, setCount] = useState(0);
  return (
    <button type="button" onClick={() => setCount((value) => value + 1)}>
      Map count {count}
    </button>
  );
}

describe("MobileMapDisclosure", () => {
  it("collapses visually without unmounting map state", async () => {
    const user = userEvent.setup();

    render(
      <MobileMapDisclosure hasResolution>
        <StatefulMapProbe />
      </MobileMapDisclosure>,
    );

    const mapProbe = screen.getByRole("button", { name: "Map count 0" });
    await user.click(mapProbe);
    expect(screen.getByRole("button", { name: "Map count 1" })).toBeTruthy();

    const toggle = screen.getByRole("button", { name: "View jurisdiction on map" });
    expect(toggle.getAttribute("aria-expanded")).toBe("false");

    await user.click(toggle);
    expect(screen.getByRole("button", { name: "Close jurisdiction map" })).toBeTruthy();

    await user.click(screen.getByRole("button", { name: "Close jurisdiction map" }));
    expect(screen.getByRole("button", { name: "Map count 1" })).toBeTruthy();
  });
});
```

Run:

```bash
pnpm vitest run src/components/MobileMapDisclosure.test.tsx
```

Expected: FAIL because the component does not exist.

- [ ] **Step 4: Implement mobile and layer disclosures**

Create `src/hooks/useMediaQuery.ts`:

```ts
import { useEffect, useState } from "react";

export function useMediaQuery(query: string): boolean {
  const [matches, setMatches] = useState(() => window.matchMedia(query).matches);

  useEffect(() => {
    const mediaQuery = window.matchMedia(query);
    const update = (event: MediaQueryListEvent) => setMatches(event.matches);

    setMatches(mediaQuery.matches);
    mediaQuery.addEventListener("change", update);
    return () => mediaQuery.removeEventListener("change", update);
  }, [query]);

  return matches;
}
```

Create `src/components/MobileMapDisclosure.tsx`:

```tsx
import { useEffect, useId, useState, type ReactNode } from "react";
import { useMediaQuery } from "../hooks/useMediaQuery";

interface MobileMapDisclosureProps {
  hasResolution: boolean;
  children: ReactNode;
}

export function MobileMapDisclosure({
  hasResolution,
  children,
}: MobileMapDisclosureProps): JSX.Element {
  const contentId = useId();
  const isMobile = useMediaQuery("(max-width: 58rem)");
  const [expanded, setExpanded] = useState(false);
  const contentVisible = !isMobile || expanded;

  useEffect(() => {
    setExpanded(false);
  }, [hasResolution]);

  useEffect(() => {
    if (!contentVisible) {
      return;
    }

    const frame = window.requestAnimationFrame(() => {
      window.dispatchEvent(new Event("resize"));
    });
    return () => window.cancelAnimationFrame(frame);
  }, [contentVisible]);

  const openLabel = hasResolution ? "View jurisdiction on map" : "Open map explorer";
  const closeLabel = hasResolution ? "Close jurisdiction map" : "Close map explorer";

  return (
    <div className={`mobile-map-disclosure ${expanded ? "is-expanded" : "is-collapsed"}`}>
      {isMobile ? (
        <button
          className="mobile-map-disclosure__toggle"
          type="button"
          aria-controls={contentId}
          aria-expanded={expanded}
          onClick={() => setExpanded((current) => !current)}
        >
          {expanded ? closeLabel : openLabel}
        </button>
      ) : null}
      <div id={contentId} className="mobile-map-disclosure__content" hidden={!contentVisible}>
        {children}
      </div>
    </div>
  );
}
```

Create `src/components/ResponsiveDisclosure.tsx`:

```tsx
import { useId, useState, type ReactNode } from "react";
import { useMediaQuery } from "../hooks/useMediaQuery";

interface ResponsiveDisclosureProps {
  label: string;
  children: ReactNode;
}

export function ResponsiveDisclosure({
  label,
  children,
}: ResponsiveDisclosureProps): JSX.Element {
  const contentId = useId();
  const isMobile = useMediaQuery("(max-width: 58rem)");
  const [expanded, setExpanded] = useState(false);
  const contentVisible = !isMobile || expanded;

  return (
    <div className={`responsive-disclosure ${expanded ? "is-expanded" : "is-collapsed"}`}>
      {isMobile ? (
        <button
          className="responsive-disclosure__toggle"
          type="button"
          aria-controls={contentId}
          aria-expanded={expanded}
          onClick={() => setExpanded((current) => !current)}
        >
          {expanded ? `Hide ${label}` : `Show ${label}`}
        </button>
      ) : null}
      <div id={contentId} className="responsive-disclosure__content" hidden={!contentVisible}>
        {children}
      </div>
    </div>
  );
}
```

Run:

```bash
pnpm vitest run src/components/MobileMapDisclosure.test.tsx
```

Expected: PASS.

- [ ] **Step 5: Render keyed sections in the approved mobile order**

In `src/App.tsx`, build the four section nodes before `return`:

```tsx
const workspaceSections = {
  query: (
    <div className="workspace-column workspace-column--query" key="query">
      <WorkspaceFrame id="query-location" title="QUERY.LOCATION" eyebrow="Input">
        <ResolutionPanel onResultChange={setResolutionResult} />
      </WorkspaceFrame>

      <WorkspaceFrame id="data-status" title="DATA.STATUS" eyebrow="Runtime">
        <StatusBanner refreshStatus={refreshStatus} apiBaseUrl={apiBaseUrl} error={loadError} />
      </WorkspaceFrame>

      <WorkspaceFrame id="map-layers" title="MAP.LAYERS" eyebrow="Visibility">
        <ResponsiveDisclosure label="map layers">
          <LayerToggleList
            layers={layerRegistry}
            enabledLayers={enabledLayers}
            onToggle={handleToggleLayer}
          />
        </ResponsiveDisclosure>
      </WorkspaceFrame>
    </div>
  ),
  map: (
    <div className="workspace-map" key="map">
      <MobileMapDisclosure hasResolution={resolutionResult !== null}>
        <WorkspaceFrame
          id="map-view"
          title="MAP.VIEW"
          eyebrow="Geography"
          className="workspace-frame--map"
        >
          <MapStage
            layers={layerRegistry}
            selectedFeature={selectedFeature}
            enabledLayers={enabledLayers}
            onSelectFeature={setSelectedFeatureId}
            refreshStatus={refreshStatus ?? loadingRefreshStatus}
            boundaryFeatures={boundaryFeatures}
            theme={theme}
          />
        </WorkspaceFrame>
      </MobileMapDisclosure>
    </div>
  ),
  authority: (
    <div className="workspace-authority" key="authority">
      <WorkspaceFrame id="authority-record" title="AUTHORITY.RECORD" eyebrow="Result">
        <AuthorityRecord result={resolutionResult} />
      </WorkspaceFrame>
    </div>
  ),
  technical: (
    <div className="workspace-technical" key="technical">
      <WorkspaceFrame id="technical-record" title="TECHNICAL.RECORD" eyebrow="Evidence">
        <FeatureInspector
          feature={selectedFeature}
          selectedFeature={selectedSummary}
          isLoading={isLoading}
          error={loadError}
        />
        <details className="technical-disclosure">
          <summary>Cached boundary choices</summary>
          <FeatureSelectionList
            features={boundaryFeatures.map(
              ({ layerFamily, featureId, title, subtitle, sourceId }) => ({
                layerFamily,
                featureId,
                title,
                subtitle,
                sourceId,
              }),
            )}
            selectedFeatureId={selectedFeatureId}
            onSelectFeature={setSelectedFeatureId}
          />
        </details>
        <details className="technical-disclosure">
          <summary>Backend contract</summary>
          <BackendContractDetails />
        </details>
      </WorkspaceFrame>
    </div>
  ),
} satisfies Record<WorkspaceSectionKey, JSX.Element>;

const workspaceOrder = getWorkspaceOrder(resolutionResult !== null);
```

Replace the current `main` with:

```tsx
<main
  className="app-workspace"
  data-has-resolution={resolutionResult !== null ? "true" : "false"}
>
  {workspaceOrder.map((sectionKey) => workspaceSections[sectionKey])}
</main>
```

Import `getWorkspaceOrder` and `WorkspaceSectionKey`.

- [ ] **Step 6: Add desktop grid areas and mobile disclosure behavior**

Replace the Task 5 `.app-workspace` grid rules with:

```css
@layer layout {
  .app-workspace {
    display: grid;
    grid-template-areas:
      "query map authority"
      "query map technical";
    grid-template-columns: minmax(17rem, 20rem) minmax(34rem, 1fr) minmax(20rem, 25rem);
    gap: var(--space-5);
    align-items: start;
    padding-block: var(--space-5);
  }

  .workspace-column--query {
    grid-area: query;
  }

  .workspace-map {
    grid-area: map;
    min-width: 0;
  }

  .workspace-authority {
    grid-area: authority;
    min-width: 0;
  }

  .workspace-technical {
    grid-area: technical;
    min-width: 0;
  }

  .mobile-map-disclosure__content,
  .responsive-disclosure__content {
    min-width: 0;
  }
}
```

Use these responsive rules:

```css
@layer responsive {
  @media (max-width: 82rem) {
    .app-workspace {
      grid-template-areas:
        "query map"
        "authority authority"
        "technical technical";
      grid-template-columns: minmax(17rem, 20rem) minmax(0, 1fr);
    }
  }

  @media (max-width: 58rem) {
    .app-workspace {
      display: flex;
      flex-direction: column;
    }

    .app-workspace > * {
      width: 100%;
    }

    .mobile-map-disclosure__toggle,
    .responsive-disclosure__toggle {
      width: 100%;
      min-height: 2.75rem;
      border: 2px solid var(--line);
      border-radius: var(--radius-sm);
      background: var(--surface);
      box-shadow: 2px 2px 0 var(--line);
      color: var(--blue);
      font-family: var(--font-ui);
      font-size: 0.78rem;
      font-weight: 800;
      cursor: pointer;
    }

    .mobile-map-disclosure__content,
    .responsive-disclosure__content {
      margin-top: var(--space-3);
    }

    .workspace-titlebar {
      min-height: 2.1rem;
      padding-block: 0.4rem;
    }

    .workspace-frame {
      box-shadow: 3px 3px 0 var(--line);
    }
  }
}
```

The `hidden` attribute removes collapsed content from the accessibility tree without unmounting React children. `MobileMapDisclosure` dispatches a resize event when the map becomes visible so MapLibre recalculates its canvas.

- [ ] **Step 7: Add an App ordering regression test**

Append to `src/App.test.tsx`:

```tsx
it("moves the authority section ahead of the map and query after resolution", async () => {
  const user = userEvent.setup();
  const { container } = render(<App />);

  await waitFor(() => {
    expect(screen.getByRole("region", { name: "QUERY.LOCATION" })).toBeTruthy();
  });

  const main = container.querySelector("main");
  expect(main?.dataset.hasResolution).toBe("false");
  expect(
    Array.from(main?.children ?? []).map((element) => element.getAttribute("class")),
  ).toEqual([
    "workspace-column workspace-column--query",
    "workspace-map",
    "workspace-authority",
    "workspace-technical",
  ]);

  await user.click(screen.getByRole("button", { name: "Resolve location" }));

  await waitFor(() => {
    expect(main?.dataset.hasResolution).toBe("true");
  });
  expect(
    Array.from(main?.children ?? []).map((element) => element.getAttribute("class")),
  ).toEqual([
    "workspace-authority",
    "workspace-map",
    "workspace-column workspace-column--query",
    "workspace-technical",
  ]);
});
```

Run:

```bash
pnpm vitest run src/lib/workspaceOrder.test.ts src/components/MobileMapDisclosure.test.tsx src/App.test.tsx
pnpm test
pnpm check
```

Expected: PASS.

- [ ] **Step 8: Commit responsive behavior**

```bash
git add src/lib/workspaceOrder.ts src/lib/workspaceOrder.test.ts src/hooks/useMediaQuery.ts src/components/MobileMapDisclosure.tsx src/components/MobileMapDisclosure.test.tsx src/components/ResponsiveDisclosure.tsx src/App.tsx src/App.test.tsx src/styles.css src/resolution.css
git commit -m "feat: add result-first mobile workflow"
```

---

### Task 8: Finish authority, status, control, and technical-detail styling

**Files:**
- Modify: `src/resolution.css`
- Modify: `src/styles.css`
- Modify: `src/components/AuthorityRecord.test.tsx`
- Modify: `src/components/StatusBanner.tsx`
- Modify: `src/components/FeatureInspector.tsx`

**Interfaces:**
- Keeps all status labels from `AuthorityRecord`.
- Uses semantic classes `resolution-status--<status>` and `data-status--<status>`.
- Preserves visible text for every status.
- Uses serif hierarchy for authority/code content and monospace for operational metadata.

- [ ] **Step 1: Add semantic status assertions**

Append to `src/components/AuthorityRecord.test.tsx`:

```tsx
it("does not rely on color alone for non-resolved states", () => {
  const { rerender } = render(<AuthorityRecord result={ambiguousResult} />);

  const ambiguous = screen.getByText("Ambiguous");
  expect(ambiguous.className).toContain("resolution-status--ambiguous");

  rerender(<AuthorityRecord result={conflictingResult} />);
  const conflicting = screen.getByText("Conflicting evidence");
  expect(conflicting.className).toContain("resolution-status--conflicting");
});
```

Run:

```bash
pnpm vitest run src/components/AuthorityRecord.test.tsx
```

Expected: PASS before styling; this locks required text/class contracts.

- [ ] **Step 2: Replace `src/resolution.css` with the approved visual language**

Replace the entire file with:

```css
@layer resolution {
  .resolution-panel,
  .resolution-form,
  .resolution-form label,
  .authority-record,
  .authority-section,
  .authority-list,
  .resolution-evidence ul {
    display: grid;
    gap: var(--space-3);
  }

  .resolution-form__coordinates {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: var(--space-3);
  }

  .resolution-form label > span,
  .meta-grid dt,
  .attribute-row__key,
  .authority-record__place {
    color: var(--ink-soft);
    font-family: var(--font-ui);
    font-size: 0.7rem;
    font-weight: 800;
    letter-spacing: 0.07em;
    text-transform: uppercase;
  }

  .resolution-form input,
  .resolution-form select {
    width: 100%;
    min-height: 2.65rem;
    padding: 0.6rem 0.7rem;
    border: 2px solid var(--line);
    border-radius: var(--radius-sm);
    background: var(--surface);
    color: var(--ink);
    font-family: var(--font-ui);
  }

  .resolution-form__submit {
    min-height: 2.8rem;
    padding: 0.65rem 0.9rem;
    border: 2px solid var(--line);
    border-radius: var(--radius-sm);
    background: var(--blue);
    box-shadow: 3px 3px 0 var(--line);
    color: #ffffff;
    font-family: var(--font-ui);
    font-weight: 800;
    cursor: pointer;
  }

  :root[data-theme="dark"] .resolution-form__submit {
    color: #050b0d;
    background: var(--green);
  }

  .resolution-form__submit:disabled {
    cursor: wait;
    opacity: 0.65;
  }

  .authority-empty {
    display: grid;
    gap: var(--space-2);
    color: var(--ink-soft);
  }

  .authority-empty h3,
  .authority-record__summary h3,
  .authority-list strong {
    font-family: var(--font-display);
  }

  .authority-record__summary {
    display: flex;
    align-items: start;
    justify-content: space-between;
    gap: var(--space-4);
    padding-bottom: var(--space-4);
    border-bottom: 2px solid var(--line);
  }

  .authority-record__summary h3 {
    margin-top: var(--space-1);
    font-size: clamp(1.5rem, 2.5vw, 2.35rem);
    line-height: 1.05;
  }

  .authority-record__summary p,
  .authority-list span,
  .resolution-evidence span,
  .resolution-evidence p {
    color: var(--ink-soft);
    font-size: 0.88rem;
  }

  .resolution-status {
    flex: none;
    padding: 0.4rem 0.55rem;
    border: 1px solid currentColor;
    border-radius: var(--radius-sm);
    font-family: var(--font-ui);
    font-size: 0.7rem;
    font-weight: 800;
    text-transform: uppercase;
  }

  .resolution-status--resolved,
  .resolution-status--partially_resolved {
    color: var(--success);
    background: var(--green-soft);
  }

  .resolution-status--local_record_required,
  .resolution-status--ambiguous {
    color: var(--warning);
    background: color-mix(in srgb, var(--warning) 12%, var(--surface));
  }

  .resolution-status--conflicting,
  .resolution-status--insufficient_evidence {
    color: var(--danger);
    background: color-mix(in srgb, var(--danger) 12%, var(--surface));
  }

  .authority-section {
    padding-top: var(--space-4);
    border-top: 1px solid var(--line-soft);
  }

  .authority-section h3 {
    font-family: var(--font-ui);
    font-size: 0.75rem;
    font-weight: 800;
    letter-spacing: 0.08em;
    text-transform: uppercase;
  }

  .authority-list,
  .resolution-evidence ul {
    margin: 0;
    padding: 0;
    list-style: none;
  }

  .authority-list li,
  .resolution-evidence li {
    display: grid;
    gap: var(--space-1);
    padding: var(--space-3);
    border: 1px solid var(--line-soft);
    background: var(--surface-raised);
  }

  .authority-list strong {
    font-size: 1.02rem;
  }

  .resolution-message {
    padding: var(--space-3);
    border: 2px solid currentColor;
    border-radius: var(--radius-sm);
    font-size: 0.88rem;
  }

  .resolution-message--warning {
    color: var(--warning);
    background: color-mix(in srgb, var(--warning) 10%, var(--surface));
  }

  .resolution-message--error {
    color: var(--danger);
    background: color-mix(in srgb, var(--danger) 10%, var(--surface));
  }

  .resolution-evidence {
    padding-top: var(--space-4);
    border-top: 1px solid var(--line-soft);
  }

  .resolution-evidence summary {
    cursor: pointer;
    color: var(--blue);
    font-family: var(--font-ui);
    font-size: 0.78rem;
    font-weight: 800;
    text-transform: uppercase;
  }

  .resolution-evidence[open] summary {
    margin-bottom: var(--space-3);
  }

  .resolution-evidence li {
    border-left: 4px solid var(--pink);
  }

  .resolution-evidence a {
    overflow-wrap: anywhere;
    font-family: var(--font-display);
    font-weight: 600;
  }

  @media (max-width: 40rem) {
    .resolution-form__coordinates {
      grid-template-columns: 1fr;
    }

    .authority-record__summary {
      display: grid;
    }

    .resolution-status {
      justify-self: start;
    }
  }
}
```

- [ ] **Step 3: Finish operational component styles in `src/styles.css`**

Add to `@layer components`:

```css
.toggle-list,
.feature-list,
.attribute-table,
.meta-grid {
  display: grid;
  gap: var(--space-2);
}

.toggle-row,
.feature-card {
  display: flex;
  width: 100%;
  align-items: center;
  justify-content: space-between;
  gap: var(--space-3);
  padding: var(--space-3);
  border: 1px solid var(--line-soft);
  border-radius: var(--radius-sm);
  background: var(--surface-raised);
  color: var(--ink);
  text-align: left;
}

.toggle-row.is-active,
.feature-card.is-active {
  border-color: var(--blue);
  box-shadow: inset 4px 0 0 var(--green);
}

.toggle-row__copy {
  display: grid;
  gap: var(--space-1);
}

.toggle-row__copy strong,
.feature-card strong {
  font-family: var(--font-ui);
  font-size: 0.82rem;
}

.toggle-row__copy span,
.feature-card span,
.feature-card small {
  color: var(--ink-soft);
  font-size: 0.78rem;
}

.toggle-row input {
  width: 1.15rem;
  height: 1.15rem;
  accent-color: var(--green);
}

.feature-card {
  display: grid;
  cursor: pointer;
}

.data-status {
  display: flex;
  align-items: start;
  justify-content: space-between;
  gap: var(--space-4);
  padding-left: var(--space-3);
  border-left: 4px solid currentColor;
}

.data-status--ok {
  color: var(--success);
}

.data-status--warning {
  color: var(--warning);
}

.data-status--error {
  color: var(--danger);
}

.data-status p,
.data-status dd {
  color: var(--ink-soft);
}

.data-status dl {
  flex: none;
  font-family: var(--font-ui);
  font-size: 0.7rem;
  text-align: right;
}

.data-status dt {
  text-transform: uppercase;
}

.data-status dd {
  margin: 0;
  overflow-wrap: anywhere;
}

.feature-inspector {
  display: grid;
  gap: var(--space-4);
}

.inspector__summary {
  display: grid;
  gap: var(--space-4);
  padding-bottom: var(--space-4);
  border-bottom: 1px solid var(--line-soft);
}

.inspector__summary h3 {
  font-size: 1.4rem;
}

.meta-grid {
  grid-template-columns: repeat(2, minmax(0, 1fr));
}

.meta-grid--stack {
  grid-template-columns: 1fr;
}

.meta-grid dd {
  margin: 0;
  overflow-wrap: anywhere;
  font-family: var(--font-ui);
  font-size: 0.78rem;
}

.attribute-row {
  display: grid;
  grid-template-columns: minmax(0, 1fr) minmax(0, 1.2fr);
  gap: var(--space-3);
  padding-block: var(--space-2);
  border-top: 1px solid var(--line-soft);
}

.attribute-row__value {
  overflow-wrap: anywhere;
  font-family: var(--font-ui);
  font-size: 0.78rem;
}

.is-muted {
  color: var(--ink-soft);
}

@media (max-width: 40rem) {
  .data-status {
    display: grid;
  }

  .data-status dl {
    text-align: left;
  }

  .meta-grid,
  .attribute-row {
    grid-template-columns: 1fr;
  }
}
```

Remove obsolete rounded-card, glassmorphism, pill, and old status-banner rules.

- [ ] **Step 4: Run behavior, accessibility-contract, and build checks**

Run:

```bash
pnpm vitest run src/components/AuthorityRecord.test.tsx src/components/ResolutionPanel.test.tsx src/App.test.tsx
pnpm lint
pnpm check
pnpm test
pnpm build
pnpm format
pnpm format:check
```

Expected: all commands PASS.

- [ ] **Step 5: Commit the finished component language**

```bash
git add src/resolution.css src/styles.css src/components/AuthorityRecord.test.tsx src/components/StatusBanner.tsx src/components/FeatureInspector.tsx
git commit -m "feat: finish source-backed authority visual language"
```

---

### Task 9: Add deterministic visual regression coverage and CI enforcement

**Files:**
- Modify: `package.json`
- Modify: `pnpm-lock.yaml`
- Create: `playwright.visual.config.ts`
- Create: `tsconfig.visual.json`
- Create: `tests/visual/building-code-map.visual.spec.ts`
- Modify: `eslint.config.js`
- Create: `docs/how-to/review-frontend-visuals.md`
- Modify: `.github/workflows/ci.yml`
- Modify: `README.md`
- Create: Playwright snapshot files generated by the approved test matrix.

**Interfaces:**
- Produces: `pnpm test:visual`.
- Produces: `pnpm test:visual:update`.
- Uses only intercepted deterministic API and tile responses.
- Covers desktop light/dark, tablet light/dark, mobile pre-query, resolved, ambiguous, conflicting, expanded map, and unavailable snapshot.

- [ ] **Step 1: Install Playwright and add scripts**

Run:

```bash
pnpm add -D @playwright/test
```

Add these scripts to `package.json`:

```json
"check:visual": "tsc --noEmit -p tsconfig.visual.json",
"test:visual": "playwright test --config playwright.visual.config.ts",
"test:visual:update": "playwright test --config playwright.visual.config.ts --update-snapshots"
```

Replace the existing lint script with:

```json
"lint": "eslint src tests playwright.visual.config.ts"
```

Extend both Prettier scripts so they include `"playwright.visual.config.ts"`, `"tsconfig.visual.json"`, and `"tests/**/*.{ts,tsx}"` in addition to the existing paths.

- [ ] **Step 2: Add strict typing and lint globals for visual tests**

Create `tsconfig.visual.json`:

```json
{
  "extends": "./tsconfig.json",
  "compilerOptions": {
    "types": ["node", "@playwright/test"]
  },
  "include": ["playwright.visual.config.ts", "tests/visual/**/*.ts"]
}
```

Add this block before `prettier` in `eslint.config.js`:

```js
{
  files: ["tests/**/*.ts", "playwright.visual.config.ts"],
  languageOptions: {
    globals: globals.node,
  },
  rules: {
    "@typescript-eslint/consistent-type-imports": "error",
    "@typescript-eslint/no-unused-vars": [
      "error",
      {
        argsIgnorePattern: "^_",
        varsIgnorePattern: "^_",
      },
    ],
  },
},
```

Run:

```bash
pnpm check:visual
pnpm lint
```

Expected: `check:visual` fails only because the Playwright files do not exist yet; lint passes for the files currently present.

- [ ] **Step 3: Create the deterministic Playwright configuration**

Create `playwright.visual.config.ts`:

```ts
import { defineConfig, devices } from "@playwright/test";

export default defineConfig({
  testDir: "./tests/visual",
  outputDir: ".test-results/visual",
  snapshotDir: "./tests/visual/__snapshots__",
  fullyParallel: false,
  retries: 0,
  reporter: "list",
  use: {
    baseURL: "http://127.0.0.1:4173",
    colorScheme: "light",
    locale: "en-US",
    timezoneId: "America/Denver",
    reducedMotion: "reduce",
    screenshot: "only-on-failure",
    trace: "retain-on-failure",
    launchOptions: {
      args: ["--use-angle=swiftshader", "--enable-webgl"],
    },
  },
  projects: [
    {
      name: "chromium",
      use: { ...devices["Desktop Chrome"] },
    },
  ],
  webServer: {
    command: "pnpm build && pnpm exec vite preview --host 127.0.0.1 --port 4173",
    url: "http://127.0.0.1:4173",
    reuseExistingServer: !process.env.CI,
    timeout: 120_000,
  },
});
```

- [ ] **Step 4: Create deterministic route helpers and visual tests**

Create `tests/visual/building-code-map.visual.spec.ts`:

```ts
import { expect, test, type Page } from "@playwright/test";
import {
  ambiguousResult,
  conflictingResult,
  demoBoundaryFeatures,
  demoLayers,
  demoRefreshStatus,
  resolvedResult,
} from "../../src/test/fixtures";

const transparentTile = Buffer.from(
  "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII=",
  "base64",
);

test.beforeEach(async ({ page }) => {
  await page.clock.setFixedTime(new Date("2026-07-21T12:00:00Z"));
});

async function mockBaseData(page: Page): Promise<void> {
  await page.route("https://tile.openstreetmap.org/**", async (route) => {
    await route.fulfill({
      status: 200,
      contentType: "image/png",
      body: transparentTile,
    });
  });
  await page.route("**/api/layers", async (route) => {
    await route.fulfill({ status: 200, contentType: "application/json", body: JSON.stringify(demoLayers) });
  });
  await page.route("**/api/refresh/status", async (route) => {
    await route.fulfill({
      status: 200,
      contentType: "application/json",
      body: JSON.stringify(demoRefreshStatus),
    });
  });
  await page.route("**/api/boundaries", async (route) => {
    await route.fulfill({
      status: 200,
      contentType: "application/json",
      body: JSON.stringify(demoBoundaryFeatures),
    });
  });
}

async function mockResolution(page: Page, result: unknown): Promise<void> {
  await page.route("**/api/resolve", async (route) => {
    await route.fulfill({
      status: 200,
      contentType: "application/json",
      body: JSON.stringify({
        schema_version: (result as typeof resolvedResult).schemaVersion,
        generated_at: (result as typeof resolvedResult).generatedAt,
        profile_id: (result as typeof resolvedResult).profileId,
        profile_last_verified: (result as typeof resolvedResult).profileLastVerified,
        geography: {
          state_id: (result as typeof resolvedResult).geography.stateId,
          state_fips: (result as typeof resolvedResult).geography.stateFips,
          state_name: (result as typeof resolvedResult).geography.stateName,
          county: (result as typeof resolvedResult).geography.county
            ? {
                layer_family: (result as typeof resolvedResult).geography.county?.layerFamily,
                feature_id: (result as typeof resolvedResult).geography.county?.featureId,
                name: (result as typeof resolvedResult).geography.county?.name,
                source_id: (result as typeof resolvedResult).geography.county?.sourceId,
              }
            : undefined,
          municipality: (result as typeof resolvedResult).geography.municipality
            ? {
                layer_family: (result as typeof resolvedResult).geography.municipality?.layerFamily,
                feature_id: (result as typeof resolvedResult).geography.municipality?.featureId,
                name: (result as typeof resolvedResult).geography.municipality?.name,
                source_id: (result as typeof resolvedResult).geography.municipality?.sourceId,
              }
            : undefined,
          incorporated: (result as typeof resolvedResult).geography.incorporated,
          special_areas: [],
          tribal_areas: [],
          fire_jurisdictions: [],
        },
        code_family: (result as typeof resolvedResult).codeFamily,
        project_type: (result as typeof resolvedResult).projectType,
        applicability_date: (result as typeof resolvedResult).applicabilityDate,
        status: (result as typeof resolvedResult).status,
        policy_basis: (result as typeof resolvedResult).policyBasis
          ? {
              status: (result as typeof resolvedResult).policyBasis?.status,
              required_local_records:
                (result as typeof resolvedResult).policyBasis?.requiredLocalRecords,
              warnings: (result as typeof resolvedResult).policyBasis?.warnings,
              source_ids: (result as typeof resolvedResult).policyBasis?.sourceIds,
              verification: (result as typeof resolvedResult).policyBasis?.verification,
            }
          : undefined,
        authority_candidates: (result as typeof resolvedResult).authorityCandidates.map(
          (candidate) => ({
            kind: candidate.kind,
            authority_id: candidate.authorityId,
            name: candidate.name,
            roles: candidate.roles,
            source_ids: candidate.sourceIds,
            verification: candidate.verification,
          }),
        ),
        authority_path: (result as typeof resolvedResult).authorityPath.map((relationship) => ({
          id: relationship.id,
          from_id: relationship.fromId,
          relationship: relationship.relationship,
          to: relationship.to,
          scope: relationship.scope,
          summary: relationship.summary,
          source_ids: relationship.sourceIds,
          verification: relationship.verification,
        })),
        adoptions: (result as typeof resolvedResult).adoptions.map((adoption) => ({
          id: adoption.id,
          code_family: adoption.codeFamily,
          status: adoption.status,
          state_code_name: adoption.stateCodeName,
          enforcement_model: adoption.enforcementModel,
          dates: adoption.dates,
          source_ids: adoption.sourceIds,
          verification: adoption.verification,
        })),
        applicable_rules: (result as typeof resolvedResult).applicableRules.map((rule) => ({
          id: rule.id,
          kind: rule.kind,
          code_family: rule.codeFamily,
          summary: rule.summary,
          source_ids: rule.sourceIds,
          verification: rule.verification,
        })),
        supporting_claims: (result as typeof resolvedResult).supportingClaims.map((claim) => ({
          id: claim.id,
          subject_id: claim.subjectId,
          field: claim.field,
          status: claim.status,
          value: claim.value,
          conflict_group: claim.conflictGroup,
          source_ids: claim.sourceIds,
          verification: claim.verification,
        })),
        required_local_records: (result as typeof resolvedResult).requiredLocalRecords,
        warnings: (result as typeof resolvedResult).warnings,
        evidence: (result as typeof resolvedResult).evidence.map((source) => ({
          id: source.id,
          title: source.title,
          url: source.url,
          kind: source.kind,
          accessed_at: source.accessedAt,
          last_checked_at: source.lastCheckedAt,
          availability: source.availability,
          caveat: source.caveat,
        })),
      }),
    });
  });
}

async function setTheme(page: Page, theme: "light" | "dark"): Promise<void> {
  await page.addInitScript(
    ({ key, value }) => window.localStorage.setItem(key, value),
    { key: "building-code-map-theme", value: theme },
  );
}

async function resolveDefaultLocation(page: Page): Promise<void> {
  await page.getByRole("button", { name: "Resolve location" }).click();
  await expect(page.getByRole("heading", { name: "City and County of Denver" })).toBeVisible();
}

async function waitForVisibleMap(page: Page): Promise<void> {
  await expect(page.getByText("Boundary map ready")).toBeVisible();
}

for (const theme of ["light", "dark"] as const) {
  test(`desktop ${theme}`, async ({ page }) => {
    await setTheme(page, theme);
    await mockBaseData(page);
    await mockResolution(page, resolvedResult);
    await page.setViewportSize({ width: 1600, height: 1000 });
    await page.goto("/");
    await resolveDefaultLocation(page);
    await waitForVisibleMap(page);
    await expect(page).toHaveScreenshot(`desktop-${theme}.png`, { fullPage: true });
  });

  test(`tablet ${theme}`, async ({ page }) => {
    await setTheme(page, theme);
    await mockBaseData(page);
    await mockResolution(page, resolvedResult);
    await page.setViewportSize({ width: 1024, height: 1100 });
    await page.goto("/");
    await resolveDefaultLocation(page);
    await waitForVisibleMap(page);
    await expect(page).toHaveScreenshot(`tablet-${theme}.png`, { fullPage: true });
  });
}

test("mobile before lookup", async ({ page }) => {
  await setTheme(page, "light");
  await mockBaseData(page);
  await mockResolution(page, resolvedResult);
  await page.setViewportSize({ width: 390, height: 844 });
  await page.goto("/");
  await expect(page.getByRole("button", { name: "Open map explorer" })).toBeVisible();
  await expect(page).toHaveScreenshot("mobile-before-lookup.png", { fullPage: true });
});

for (const [name, result] of [
  ["resolved", resolvedResult],
  ["ambiguous", ambiguousResult],
  ["conflicting", conflictingResult],
] as const) {
  test(`mobile ${name} result`, async ({ page }) => {
    await setTheme(page, "dark");
    await mockBaseData(page);
    await mockResolution(page, result);
    await page.setViewportSize({ width: 390, height: 844 });
    await page.goto("/");
    await resolveDefaultLocation(page);
    await expect(page.getByRole("button", { name: "View jurisdiction on map" })).toBeVisible();
    await expect(page).toHaveScreenshot(`mobile-${name}.png`, { fullPage: true });
  });
}

test("mobile expanded jurisdiction map", async ({ page }) => {
  await setTheme(page, "dark");
  await mockBaseData(page);
  await mockResolution(page, resolvedResult);
  await page.setViewportSize({ width: 390, height: 844 });
  await page.goto("/");
  await resolveDefaultLocation(page);
  await page.getByRole("button", { name: "View jurisdiction on map" }).click();
  await expect(page.getByRole("button", { name: "Close jurisdiction map" })).toBeVisible();
  await waitForVisibleMap(page);
  await expect(page).toHaveScreenshot("mobile-map-expanded.png", { fullPage: true });
});

test("mobile disclosures are keyboard operable with visible focus", async ({ page }) => {
  await setTheme(page, "light");
  await mockBaseData(page);
  await mockResolution(page, resolvedResult);
  await page.setViewportSize({ width: 390, height: 844 });
  await page.goto("/");

  const mapToggle = page.getByRole("button", { name: "Open map explorer" });
  await mapToggle.focus();
  await expect(mapToggle).toBeFocused();
  const outlineStyle = await mapToggle.evaluate((element) => getComputedStyle(element).outlineStyle);
  expect(outlineStyle).not.toBe("none");

  await page.keyboard.press("Enter");
  await expect(page.getByRole("button", { name: "Close map explorer" })).toHaveAttribute(
    "aria-expanded",
    "true",
  );
});

test("unavailable snapshot in both themes", async ({ page }) => {
  await setTheme(page, "light");
  await page.route("https://tile.openstreetmap.org/**", async (route) => {
    await route.fulfill({
      status: 200,
      contentType: "image/png",
      body: transparentTile,
    });
  });
  await page.route("**/api/**", async (route) => {
    await route.fulfill({ status: 503, body: "Unavailable" });
  });
  await page.setViewportSize({ width: 1280, height: 900 });
  await page.goto("/");
  await expect(page.getByText("Cached boundary data could not be loaded.")).toBeVisible();
  await expect(page).toHaveScreenshot("snapshot-unavailable-light.png", { fullPage: true });

  await page.getByRole("button", { name: "Switch to dark theme" }).click();
  await expect(page).toHaveScreenshot("snapshot-unavailable-dark.png", { fullPage: true });
});
```

- [ ] **Step 5: Generate and review the visual baselines**

Run:

```bash
pnpm exec playwright install chromium
pnpm test:visual:update
```

Expected: eleven approved PNG snapshots are created under `tests/visual/__snapshots__/`.

Open every snapshot and verify:

- map remains the desktop visual anchor;
- dark mode is genuinely dark;
- borders, hard shadows, and title bars match the LLC language;
- serif/monospace roles are consistent;
- pink is restrained;
- mobile result precedes map and query;
- map is collapsed by default on mobile;
- ambiguity/conflict are explicit without relying on color;
- OSM attribution remains visible when the map is expanded.

Do not commit snapshots with clipped content, unreadable controls, or unstable timestamps.

- [ ] **Step 6: Add the visual CI job**

Append this job to `.github/workflows/ci.yml`:

```yaml
  frontend-visual:
    name: Frontend visual regression
    runs-on: ubuntu-latest
    timeout-minutes: 20

    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Install pnpm
        uses: pnpm/action-setup@v4
        with:
          run_install: false

      - name: Set up Node.js
        uses: actions/setup-node@v4
        with:
          node-version: 22
          cache: pnpm

      - name: Install dependencies
        run: pnpm install --frozen-lockfile

      - name: Type-check visual tests
        run: pnpm check:visual

      - name: Install Chromium
        run: pnpm exec playwright install --with-deps chromium

      - name: Run visual regression suite
        run: pnpm test:visual

      - name: Upload failure evidence
        if: failure()
        uses: actions/upload-artifact@v4
        with:
          name: building-code-map-visual-failure
          path: |
            .test-results/visual
            playwright-report
          if-no-files-found: ignore
```

- [ ] **Step 7: Document the review workflow**

Create `docs/how-to/review-frontend-visuals.md`:

````md
# Review frontend visual regressions

## Run the checked-in baselines

```bash
pnpm install --frozen-lockfile
pnpm exec playwright install chromium
pnpm test:visual
```

The suite intercepts Building Code Map API requests and OpenStreetMap tiles. It does not require a backend snapshot or live external service.

## Update baselines intentionally

```bash
pnpm test:visual:update
```

Review every changed PNG under `tests/visual/__snapshots__/`. Confirm desktop and tablet light/dark layouts, mobile pre-query and post-query ordering, resolved/ambiguous/conflicting states, expanded mobile map behavior, unavailable-snapshot treatment, readable attribution, and visible focus.

Never accept a baseline change only because the test command produced it. The changed image is the artifact under review.

## Complete verification

```bash
pnpm check
pnpm lint
pnpm format:check
pnpm test
pnpm build
pnpm check:visual
pnpm test:visual
```
````

Add to the README verification section:

````md
For deterministic responsive and theme captures:

```bash
pnpm exec playwright install chromium
pnpm test:visual
```

See [`docs/how-to/review-frontend-visuals.md`](docs/how-to/review-frontend-visuals.md) before updating checked-in baselines.
````

- [ ] **Step 8: Run complete verification**

Run:

```bash
pnpm check
pnpm lint
pnpm format:check
pnpm test
pnpm build
pnpm check:visual
pnpm test:visual
cd backend
go test ./...
go vet ./...
```

Expected: every command exits with status 0.

- [ ] **Step 9: Commit visual verification**

```bash
git add package.json pnpm-lock.yaml eslint.config.js tsconfig.visual.json playwright.visual.config.ts tests/visual docs/how-to/review-frontend-visuals.md .github/workflows/ci.yml README.md
git commit -m "test: add responsive visual regression coverage"
```

---

## Final Review Gate

- [ ] Confirm the implementation branch contains only frontend visual-system work, tests, documentation, and CI required by this plan.
- [ ] Compare the implementation against every acceptance criterion in `docs/superpowers/specs/2026-07-21-building-code-map-visual-system-design.md`.
- [ ] Review desktop light/dark, tablet light/dark, and every mobile state snapshot manually.
- [ ] Verify the exact pull-request head with the full command set from Task 9.
- [ ] Confirm no backend API or regulatory-data contract changed.
- [ ] Confirm OpenStreetMap attribution is visible and readable.
- [ ] Confirm no raw database, local cache, environment secret, or generated runtime snapshot is committed.
- [ ] Request independent review before merge.
