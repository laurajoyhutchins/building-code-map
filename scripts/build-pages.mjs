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
    !original.includes('src="./app.js"')) {
  throw Error("Public HTML must use relative static paths for Pages subpath hosting.");
}
const html = original
  .replace("<head>", '<head>\n  <meta name="bcm-static-demo" content="fictional-contract-example">')
  ;
fs.rmSync(target, { recursive: true, force: true });
fs.mkdirSync(target, { recursive: true });
fs.writeFileSync(path.join(target, "index.html"), html);
for (const name of ["app.js", "styles.css"]) {
  fs.copyFileSync(path.join(source, name), path.join(target, name));
}
fs.writeFileSync(path.join(target, "demo-result.json"), JSON.stringify(example, null, 2) + "\n");
console.log("Built public Pages artifact: index.html, app.js, styles.css, demo-result.json");
