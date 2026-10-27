const form = document.querySelector("#lookup");
const submit = document.querySelector("#submit");
const panel = document.querySelector("#result-panel");
const content = document.querySelector("#result-content");
const status = document.querySelector("#status");
const retry = document.querySelector("#retry-service");
const staticDemo = document.querySelector('meta[name="bcm-static-demo"]') !== null;
const liveAPI = document.querySelector('meta[name="bcm-live-api"]')?.content || "";
let pending = null;
let liveReady = !liveAPI;

function element(tag, className, value) {
  const node = document.createElement(tag);
  if (className) node.className = className;
  if (value !== undefined) node.textContent = String(value);
  return node;
}

function addText(parent, tag, value, className = "") {
  parent.append(element(tag, className, value));
}

function invalidate() {
  if (pending) pending.abort();
  pending = null;
  panel.hidden = true;
  content.replaceChildren();
  status.textContent = "";
  submit.disabled = !liveReady;
  submit.firstChild.textContent = liveAPI ? "Check code basis " : "Run demo ";
}

form.addEventListener("input", invalidate);
form.addEventListener("change", invalidate);

function render(result) {
  content.replaceChildren();
  const heading = element("div", "verdict");
  addText(heading, "strong", result.verdict.replaceAll("_", " ").toUpperCase());
  addText(heading, "span", result.jurisdiction.primary_jurisdiction_name);
  content.append(heading);

  addText(content, "h3", "Codes");
  const list = element("ul", "records");
  for (const code of result.code_set || []) {
    const item = element("li");
    addText(item, "strong", code.family + (code.edition ? " · " + code.edition : ""));
    addText(item, "span", code.status.replaceAll("_", " "));
    if (code.basis) addText(item, "p", code.basis);
    list.append(item);
  }
  content.append(list);

  addText(content, "h3", "Open items");
  const unresolved = result.unresolved || [];
  if (!unresolved.length) {
    addText(content, "p", "None recorded.");
  } else {
    const items = element("ul", "requirements");
    for (const entry of unresolved) {
      const item = element("li");
      addText(item, "strong", entry.kind.replaceAll("_", " "));
      addText(item, "p", entry.message);
      items.append(item);
    }
    content.append(items);
  }

  if (liveAPI && Array.isArray(result.evidence) && result.evidence.length) {
    addText(content, "h3", "Source references");
    const refs = element("ul", "records");
    for (const evidence of result.evidence) {
      const item = element("li");
      addText(item, "span", evidence.id);
      refs.append(item);
    }
    content.append(refs);
  }

  const footer = element("p", "provenance");
  footer.textContent = "As of " + result.as_of;
  if (liveAPI && result.provenance?.bundle_id) {
    footer.textContent += " · Bundle: " + result.provenance.bundle_id;
  }
  content.append(footer);
  panel.hidden = false;
}

function validLiveResult(value, request) {
  return value && value.schema_version === "0.3" &&
    value.project_id === request.project_id &&
    value.as_of === request.applicability_date &&
    value.inputs?.address === request.address &&
    typeof value.jurisdiction_id === "string" &&
    !value.jurisdiction_id.startsWith("DEMO") &&
    typeof value.jurisdiction?.primary_jurisdiction_name === "string" &&
    Boolean(value.provenance?.bundle_id) &&
    ["verified", "conditional", "not_verified"].includes(value.verdict) &&
    Array.isArray(value.code_set) &&
    Array.isArray(value.unresolved);
}

form.addEventListener("submit", async event => {
  event.preventDefault();
  if (!liveReady) return;
  invalidate();
  const controller = new AbortController();
  pending = controller;
  submit.disabled = true;
  submit.firstChild.textContent = "Checking… ";
  status.textContent = liveAPI ? "Checking location and code records…" : "Checking sample project…";

  const facts = {};
  const occupancy = document.querySelector("#occupancy").value.trim();
  if (occupancy) facts.occupancy = occupancy;
  const request = {
    project_id: document.querySelector("#project-id").value.trim(),
    address: document.querySelector("#address").value.trim(),
    applicability_date: document.querySelector("#as-of").value,
    facts,
  };

  try {
    const response = await fetch((liveAPI || "") + "/v1/project-code-basis", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(request),
      signal: controller.signal,
      credentials: "omit",
      cache: "no-store",
    });
    const payload = await response.json();
    if (controller.signal.aborted || pending !== controller) return;
    if (!response.ok) throw Error(payload.error?.message || "Lookup unavailable");
    if (liveAPI && !validLiveResult(payload, request)) {
      throw Error("The service returned an invalid project result.");
    }
    render(payload);
    status.textContent = liveAPI
      ? "Review open items and confirm applicable codes with the jurisdiction."
      : "Demo complete.";
  } catch (error) {
    if (!controller.signal.aborted) status.textContent = "Not verified: " + error.message;
  } finally {
    if (pending === controller) {
      pending = null;
      submit.disabled = !liveReady;
      submit.firstChild.textContent = liveAPI ? "Check code basis " : "Run demo ";
    }
  }
});

async function checkLiveService() {
  liveReady = false;
  submit.disabled = true;
  retry.hidden = true;
  retry.disabled = true;
  status.textContent = "Checking service availability…";
  try {
    const response = await fetch(liveAPI + "/v1/readiness", {
      credentials: "omit",
      cache: "no-store",
    });
    if (!response.ok) throw Error("Service unavailable");
    const readiness = await response.json();
    if (readiness.status !== "ready") throw Error("Service unavailable");
    liveReady = true;
    submit.disabled = false;
    status.textContent = "";
  } catch {
    status.textContent = "Live lookup is unavailable. Retry without re-entering project details.";
    retry.hidden = false;
  } finally {
    retry.disabled = false;
  }
}

if (staticDemo) {
  form.hidden = true;
  document.querySelector("#lookup-title").textContent = "Sample project";
  document.querySelector(".mode").textContent = "Sample data";
  const note = element("p", "", "Example: 100 Demo Plaza, DEMO-XX. Live lookup is not available.");
  form.parentNode.insertBefore(note, form);
  status.textContent = "Loading example…";
  fetch("./demo-result.json", { cache: "no-store" })
    .then(response => {
      if (!response.ok) throw Error("Sample data could not be loaded.");
      return response.json();
    })
    .then(example => {
      if (example.jurisdiction_id !== "DEMO-XX" || example.verdict !== "conditional") {
        throw Error("Invalid example data.");
      }
      render(example);
      status.textContent = "";
    })
    .catch(error => { status.textContent = "Could not load example: " + error.message; });
} else if (liveAPI) {
  document.querySelector(".mode").textContent = "Live lookup";
  document.querySelector("#lookup-title").textContent = "Project location";
  document.querySelector("#address").value = "";
  document.querySelector("#address").placeholder = "Street address, city, state";
  document.querySelector("#project-id").value = "";
  document.querySelector("#project-id").placeholder = "Project identifier";
  document.querySelector('label[for="address"] small').textContent = "";
  document.querySelector("footer span").textContent = "Confirm the code basis with the authority having jurisdiction.";
  submit.firstChild.textContent = "Check code basis ";
  retry.addEventListener("click", checkLiveService);
  checkLiveService();
}
