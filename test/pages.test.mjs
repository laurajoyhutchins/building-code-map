import test from "node:test";
import assert from "node:assert/strict";
import fs from "node:fs";
import { execFileSync } from "node:child_process";
import path from "node:path";

test("Pages artifact contains only public UI and the fictional contract fixture", () => {
  execFileSync(process.execPath, ["scripts/build-pages.mjs"]);
  const target = path.resolve("dist/pages");
  assert.deepEqual(fs.readdirSync(target).sort(), ["app.js", "demo-result.json", "index.html", "styles.css", "utility-context.js"]);
  const html = fs.readFileSync(path.join(target, "index.html"), "utf8");
  assert.match(html, /bcm-static-demo/);
  assert.match(html, /<h1>Building Code Map<\/h1>/);
  assert.doesNotMatch(html, /A clearer basis for|every building|ENGINE · CANONICAL/);
  assert.match(html, /href="\.\/styles\.css"/);
  assert.match(html, /src="\.\/app\.js"/);
  assert.match(html, /src="\.\/utility-context\.js"/);
  const example = JSON.parse(fs.readFileSync(path.join(target, "demo-result.json"), "utf8"));
  assert.equal(example.schema_version, "0.3");
  assert.equal(example.jurisdiction_id, "DEMO-XX");
  assert.equal(example.verdict, "conditional");
  assert.ok(example.unresolved.length > 0);
});


test("Live Pages build uses HTTPS API origin and contains no synthetic result", () => {
  execFileSync(process.execPath, ["scripts/build-pages.mjs"], {
    env: { ...process.env, BCM_PUBLIC_API_BASE_URL: "https://bcm-api.example.org" },
  });
  const target = path.resolve("dist/pages");
  assert.deepEqual(fs.readdirSync(target).sort(), ["app.js", "index.html", "styles.css", "utility-context.js"]);
  const html = fs.readFileSync(path.join(target, "index.html"), "utf8");
  assert.match(html, /bcm-live-api/);
  assert.match(html, /https:\/\/bcm-api\.example\.org/);
  assert.doesNotMatch(html, /bcm-static-demo/);
});

test("Live Pages build rejects insecure, credential-bearing or path-based origins", () => {
  const invalid = [
    "http://bcm-api.example.org",
    "https://user:password@bcm-api.example.org",
    "https://bcm-api.example.org/v1",
    "https://laurajoyhutchins.github.io",
  ];
  for (const endpoint of invalid) {
    assert.throws(() => execFileSync(process.execPath, ["scripts/build-pages.mjs"], {
      env: { ...process.env, BCM_PUBLIC_API_BASE_URL: endpoint },
      stdio: "pipe",
    }), endpoint);
  }
});
