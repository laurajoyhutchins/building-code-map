const form = document.querySelector("#lookup");
const submit = document.querySelector("#submit");
const panel = document.querySelector("#result-panel");
const content = document.querySelector("#result-content");
const status = document.querySelector("#status");
let pending = null;
const staticDemo = document.querySelector('meta[name="bcm-static-demo"]') !== null;
if (staticDemo) {
  form.hidden = true;
  document.querySelector("#lookup-title").textContent = "Sample project";
  document.querySelector(".mode").textContent = "Sample data";
  const note = document.createElement("p");
  note.textContent = "Example: 100 Demo Plaza, DEMO-XX. Live lookup is not available.";
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
  submit.firstChild.textContent = "Run demo ";
}
form.addEventListener("input", invalidate);
form.addEventListener("change", invalidate);

function render(result) {
  content.replaceChildren();
  const heading = element("div", "verdict");
  addText(heading, "strong", result.verdict.replaceAll("_", " ").toUpperCase());
  addText(heading, "span", result.jurisdiction.primary_jurisdiction_name);
  content.append(heading);
  const codeHeading = element("h3", "", "Codes");
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
  const section = element("h3", "", "Open items");
  content.append(section);
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
  const footer = element("p", "provenance");
  footer.textContent = "As of " + result.as_of;
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
  status.textContent = "Checking sample project…";
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
    status.textContent = "Demo complete.";
  } catch (error) {
    if (!controller.signal.aborted) status.textContent = "Not verified: " + error.message;
  } finally {
    if (pending === controller) {
      pending = null;
      submit.disabled = false;
      submit.firstChild.textContent = "Run demo ";
    }
  }
});

if (staticDemo) {
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
}
