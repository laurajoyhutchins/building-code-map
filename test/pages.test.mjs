import test from "node:test";
import assert from "node:assert/strict";
import fs from "node:fs";
import { execFileSync } from "node:child_process";
import path from "node:path";

test("Pages artifact contains only public UI and the fictional contract fixture", () => {
  execFileSync(process.execPath, ["scripts/build-pages.mjs"]);
  const target = path.resolve("dist/pages");
  assert.deepEqual(fs.readdirSync(target).sort(), ["app.js", "demo-result.json", "index.html", "styles.css"]);
  const html = fs.readFileSync(path.join(target, "index.html"), "utf8");
  assert.match(html, /bcm-static-demo/);
  assert.match(html, /href="\.\/styles\.css"/);
  assert.match(html, /src="\.\/app\.js"/);
  const example = JSON.parse(fs.readFileSync(path.join(target, "demo-result.json"), "utf8"));
  assert.equal(example.schema_version, "0.3");
  assert.equal(example.jurisdiction_id, "DEMO-XX");
  assert.equal(example.verdict, "conditional");
  assert.ok(example.unresolved.length > 0);
});
