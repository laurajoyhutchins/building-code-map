const form = document.querySelector("#lookup");
const submit = document.querySelector("#submit");
const panel = document.querySelector("#result-panel");
const content = document.querySelector("#result-content");
const status = document.querySelector("#status");
let pending = null;
const staticDemo = document.querySelector('meta[name="bcm-static-demo"]') !== null;
if (staticDemo) {
  form.hidden = true;
  document.querySelector("#lookup-title").textContent = "Fictional project example";
  document.querySelector(".mode").textContent = "Static showcase · synthetic data only";
  const note = document.createElement("p");
  note.textContent = "This page displays a prewritten DEMO-XX contract example. It does not perform a live regulatory lookup. Run the public Go reference app locally to try the synthetic verifier.";
  form.parentNode.insertBefore(note, form);
}


function element(tag, className, value) {
  const node = document.createElement(tag);
  if (className) node.className = className;
  if (value !== undefined) node.textContent = value;
  return node;
}
function addText(parent, tag, text, className = "") {
  parent.append(element(tag, className, String(text)));
}
function invalidate() {
  if (pending) pending.abort();
  pending = null;
  panel.hidden = true;
  content.replaceChildren();
  status.textContent = "";
  submit.disabled = false;
  submit.firstChild.textContent = "Verify demo project ";
}
form.addEventListener("input", invalidate);
form.addEventListener("change", invalidate);

function render(result) {
  content.replaceChildren();
  const heading = element("div", "verdict");
  addText(heading, "strong", result.verdict.replaceAll("_", " ").toUpperCase());
  addText(heading, "span", result.jurisdiction.primary_jurisdiction_name);
  content.append(heading);
  const codeHeading = element("h3", "", "Code set");
  content.append(codeHeading);
  const list = element("ul", "records");
  for (const code of result.code_set || []) {
    const item = element("li");
    addText(item, "strong", code.family + (code.edition ? " · " + code.edition : ""));
    addText(item, "span", code.status.replaceAll("_", " "));
    if (code.basis) addText(item, "p", code.basis);
    list.append(item);
  }
  content.append(list);
  const unresolved = result.unresolved || [];
  const section = element("h3", "", "Still to verify");
  content.append(section);
  if (!unresolved.length) {
    addText(content, "p", "No unresolved requirements in this synthetic result.");
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
  const footer = element("p", "provenance");
  footer.textContent = "Schema " + result.schema_version + " · " + result.as_of
    + " · bundle " + (result.provenance?.bundle_id || "unavailable");
  content.append(footer);
  panel.hidden = false;
}

form.addEventListener("submit", async event => {
  event.preventDefault();
  invalidate();
  const controller = new AbortController();
  pending = controller;
  submit.disabled = true;
  submit.firstChild.textContent = "Verifying… ";
  status.textContent = "Checking the fictional project against the public engine.";
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
    const response = await fetch("/v1/project-code-basis", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(request),
      signal: controller.signal,
    });
    const payload = await response.json();
    if (controller.signal.aborted) return;
    if (!response.ok) throw Error(payload.error?.message || "Verification unavailable");
    render(payload);
    status.textContent = "Synthetic verification complete. This is not a real code determination.";
  } catch (error) {
    if (!controller.signal.aborted) status.textContent = "Not verified: " + error.message;
  } finally {
    if (pending === controller) {
      pending = null;
      submit.disabled = false;
      submit.firstChild.textContent = "Verify demo project ";
    }
  }
});

if (staticDemo) {
  status.textContent = "Loading the fictional example.";
  fetch("./demo-result.json", { cache: "no-store" })
    .then(response => {
      if (!response.ok) throw Error("Sample data could not be loaded.");
      return response.json();
    })
    .then(example => {
      if (example.jurisdiction_id !== "DEMO-XX" || example.verdict !== "conditional") {
        throw Error("The synthetic example did not pass publication checks.");
      }
      render(example);
      status.textContent = "Prewritten fictional example only. No live verification performed.";
    })
    .catch(error => { status.textContent = "Example unavailable: " + error.message; });
}
