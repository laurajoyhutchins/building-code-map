// Publish only an allowlisted static shell and the existing public DEMO-XX example.
import fs from "node:fs";
import path from "node:path";
import { fileURLToPath } from "node:url";

const root = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "..");
const target = path.join(root, "dist", "pages");
const source = path.join(root, "cmd", "bcm", "web");
const fixturePath = path.join(root, "demo", "DEMO-XX", "project-code-basis.json");
const example = JSON.parse(fs.readFileSync(fixturePath, "utf8"));
if (example.schema_version !== "0.3" || example.jurisdiction_id !== "DEMO-XX" ||
    example.verdict !== "conditional" || !example.unresolved?.length ||
    example.provenance?.bundle_id !== "DEMO-BUNDLE") {
  throw Error("Public synthetic contract fixture has changed or is unsafe for display.");
}
const original = fs.readFileSync(path.join(source, "index.html"), "utf8");
if (!original.includes('<head>') || !original.includes('href="./styles.css"') ||
    !original.includes('src="./app.js"') ||
    !original.includes('src="./utility-context.js"')) {
  throw Error("Public HTML must use relative static paths for Pages subpath hosting.");
}
const configured = process.env.BCM_PUBLIC_API_BASE_URL?.trim() || "";
let apiBase = "";
if (configured) {
  let url;
  try { url = new URL(configured); } catch { throw Error("BCM_PUBLIC_API_BASE_URL must be a valid HTTPS origin"); }
  if (url.protocol !== "https:" || url.username || url.password || url.search || url.hash ||
      url.pathname !== "/" || !url.hostname || url.hostname === "laurajoyhutchins.github.io") {
    throw Error("BCM_PUBLIC_API_BASE_URL must be a separate HTTPS API origin with no path or credentials");
  }
  apiBase = url.origin;
}
const tag = apiBase
  ? '<meta name="bcm-live-api" content="' + apiBase + '">'
  : '<meta name="bcm-static-demo" content="fictional-contract-example">';
const html = original.replace("<head>", "<head>\n  " + tag);
fs.rmSync(target, { recursive: true, force: true });
fs.mkdirSync(target, { recursive: true });
fs.writeFileSync(path.join(target, "index.html"), html);
for (const name of ["app.js", "styles.css", "utility-context.js"]) {
  fs.copyFileSync(path.join(source, name), path.join(target, name));
}
if (!apiBase) {
  fs.writeFileSync(path.join(target, "demo-result.json"), JSON.stringify(example, null, 2) + "\n");
}
console.log(apiBase ? "Built live API client." : "Built sample-only Pages site.");
