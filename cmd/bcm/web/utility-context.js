// Separate, non-regulatory discovery. No address-to-county guess or serving-utility claim.
export const projectUtilityContextPath = "/v1/project-utility-context";

const sha256 = /^[a-f0-9]{64}$/;

export function validProjectUtilityContext(value, request) {
  const geo = value?.geography;
  const context = value?.utility_context;
  return value?.project_id === request.project_id &&
    value?.status === "geography_observed_utility_candidates_unconfirmed" &&
    geo && typeof geo.county === "string" && geo.county.length > 0 &&
    typeof geo.county_source_id === "string" && geo.county_source_id.length > 0 &&
    geo.address_precision === "address_point" &&
    typeof geo.address_source_id === "string" && geo.address_source_id.length > 0 &&
    typeof geo.provider_bundle_digest === "string" && geo.provider_bundle_digest.length > 0 &&
    typeof geo.boundary_snapshot_sha256 === "string" && geo.boundary_snapshot_sha256.length > 0 &&
    typeof geo.geocoder_snapshot_sha256 === "string" && geo.geocoder_snapshot_sha256.length > 0 &&
    context?.status === "reported_county_candidates_not_site_verified" &&
    context?.state === geo.state && typeof context?.county === "string" &&
    /^[A-Z]{2}$/.test(context.state) &&
    context.county.length > 0 &&
    typeof context.release_id === "string" && context.release_id.startsWith("eia861-") &&
    sha256.test(context.source_sha256) && sha256.test(context.sqlite_sha256) &&
    sha256.test(context.view_sha256) &&
    typeof context.disclaimer === "string" && context.disclaimer.length > 0 &&
    Array.isArray(context.candidates) && context.candidates.every(candidate =>
      Number.isSafeInteger(candidate.eia_id) && candidate.eia_id > 0 &&
      typeof candidate.reported_name === "string" && candidate.reported_name.trim().length > 0 &&
      typeof candidate.source_sheet === "string" &&
      Number.isSafeInteger(candidate.source_row) && candidate.source_row > 0 &&
      Array.isArray(candidate.activities));
}

function addText(doc, parent, tag, value, className = "") {
  const node = doc.createElement(tag);
  if (className) node.className = className;
  node.textContent = String(value);
  parent.append(node);
  return node;
}

export function renderProjectUtilityContext(doc, container, value) {
  const { geography, utility_context: context } = value;
  container.replaceChildren();
  addText(doc, container, "p", `${geography.county}, ${geography.state} · Source-bound geographic observation`, "utility-geography");
  addText(doc, container, "p", "Candidate electric utilities, not a confirmed provider at this address. This section does not affect the building-code determination.", "utility-caveat");
  const candidates = context.candidates;
  addText(doc, container, "h3", `Reported county associations (${candidates.length})`);
  if (candidates.length === 0) {
    addText(doc, container, "p", "No reported county candidates in this release. This does not mean the address lacks electrical service.");
  } else {
    const listing = doc.createElement("ul");
    listing.className = "records utility-candidates";
    for (const candidate of candidates.slice(0, 100)) {
      const item = doc.createElement("li");
      addText(doc, item, "strong", candidate.reported_name);
      addText(doc, item, "span", `EIA respondent ${candidate.eia_id} · ${candidate.source_sheet} row ${candidate.source_row}`);
      const reported = candidate.activities.filter(a => a.value === "Y").map(a => a.code);
      addText(doc, item, "p", reported.length ? `Reported activities: ${reported.join(", ")}` : "No affirmative activity marker available; activity is unknown, not absent.");
      listing.append(item);
    }
    container.append(listing);
    if (candidates.length > 100) {
      addText(doc, container, "p", `Showing the first 100 of ${candidates.length} reported candidates. Confirm the actual provider independently.`);
    }
  }
  addText(doc, container, "p", context.disclaimer, "utility-caveat");
  addText(doc, container, "p", `EIA ${context.reporting_year} · ${context.release_id} · View SHA-256 ${context.view_sha256} · Provider bundle ${geography.provider_bundle_digest}`, "provenance");
}

export function installProjectUtilityContext(doc = globalThis.document, getFetch = globalThis.fetch) {
  const api = doc.querySelector('meta[name="bcm-live-api"]')?.content || "";
  // The Pages demo may never call a real geographic endpoint or show utility claims.
  if (!api || doc.querySelector('meta[name="bcm-static-demo"]')) return;
  const form = doc.querySelector("#lookup");
  const panel = doc.querySelector("#utility-panel");
  const content = doc.querySelector("#utility-content");
  const status = doc.querySelector("#utility-status");
  if (!form || !panel || !content || !status) throw Error("missing utility context presentation controls");
  let pending = null;
  function invalidate() {
    if (pending) pending.abort();
    pending = null;
    panel.hidden = true;
    content.replaceChildren();
    status.textContent = "";
  }
  form.addEventListener("input", invalidate);
  form.addEventListener("change", invalidate);
  form.addEventListener("submit", async () => {
    invalidate();
    const request = {
      project_id: doc.querySelector("#project-id").value.trim(),
      address: doc.querySelector("#address").value.trim(),
    };
    const controller = new AbortController();
    pending = controller;
    panel.hidden = false;
    status.textContent = "Checking reported utility candidates using source-bound geographic evidence…";
    try {
      const response = await getFetch(api + projectUtilityContextPath, {
        method: "POST", headers: { "Content-Type": "application/json" },
        body: JSON.stringify(request), signal: controller.signal,
        credentials: "omit", cache: "no-store",
      });
      if (controller.signal.aborted || pending !== controller) return;
      if (!response.ok) {
        // A configured code-basis API may not yet have utility-context data.
        throw Error(response.status === 422 ? "The location could not be independently resolved to a unique county."
          : "Utility context is not available for this project yet.");
      }
      const result = await response.json();
      if (controller.signal.aborted || pending !== controller) return;
      if (!validProjectUtilityContext(result, request)) throw Error("Utility context source identity is incomplete or inconsistent.");
      renderProjectUtilityContext(doc, content, result);
      status.textContent = "Reported EIA candidates only. Independently confirm the utility serving the address.";
    } catch (error) {
      if (!controller.signal.aborted && pending === controller) {
        content.replaceChildren();
        status.textContent = error.message || "Utility context unavailable.";
      }
    } finally {
      if (pending === controller) pending = null;
    }
  });
}

if (typeof document !== "undefined") installProjectUtilityContext();
