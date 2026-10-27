import assert from "node:assert/strict";
import { readFileSync } from "node:fs";
import { test } from "node:test";
import { runInNewContext } from "node:vm";

const script = readFileSync(new URL("../cmd/bcm/web/app.js", import.meta.url), "utf8");

function element() {
  const listeners = new Map();
  return {
    hidden: false,
    disabled: false,
    value: "",
    placeholder: "",
    textContent: "",
    firstChild: { textContent: "" },
    addEventListener(name, handler) {
      listeners.set(name, handler);
    },
    fire(name) {
      assert.ok(listeners.has(name), `missing ${name} handler`);
      return listeners.get(name)();
    },
  };
}

function startLivePage(readiness) {
  const selectors = [
    "#lookup", "#submit", "#result-panel", "#result-content", "#status",
    "#retry-service", ".mode", "#lookup-title", "#address", "#project-id",
    "#as-of", "#occupancy",
    'label[for="address"] small', "footer span",
  ];
  const nodes = new Map(selectors.map(selector => [selector, element()]));
  const calls = [];
  const document = {
    querySelector(selector) {
      if (selector === 'meta[name="bcm-static-demo"]') return null;
      if (selector === 'meta[name="bcm-live-api"]') {
        return { content: "https://fixture.example" };
      }
      assert.ok(nodes.has(selector), `unexpected selector: ${selector}`);
      return nodes.get(selector);
    },
  };
  const fetch = async (url, options) => {
    calls.push({ url, options });
    assert.equal(url, "https://fixture.example/v1/readiness");
    assert.equal(options.credentials, "omit");
    assert.equal(options.cache, "no-store");
    const step = readiness.shift();
    assert.ok(step, "unexpected extra network request");
    return { ok: step.ok, json: async () => ({ status: step.status }) };
  };
  runInNewContext(script, { document, fetch, AbortController });
  return { nodes, calls };
}

async function settle() {
  // Flush the asynchronous readiness response and status update.
  await new Promise(resolve => setImmediate(resolve));
}

test("live readiness can recover after two failures without project field loss", async () => {
  const page = startLivePage([
    { ok: false },
    { ok: true, status: "not_ready" },
    { ok: true, status: "ready" },
  ]);
  await settle();
  const get = selector => page.nodes.get(selector);
  assert.equal(get("#submit").disabled, true);
  assert.equal(get("#retry-service").hidden, false);
  assert.match(get("#status").textContent, /Retry without re-entering/);

  // The user types their project facts once. A readiness retry must not clear them.
  get("#project-id").value = "ENGINEERING-042";
  get("#address").value = "123 Sample Street";
  get("#as-of").value = "2026-10-27";
  get("#occupancy").value = "business";
  get("#retry-service").fire("click");
  await settle();
  assert.equal(get("#retry-service").hidden, false);
  assert.equal(get("#submit").disabled, true);

  get("#retry-service").fire("click");
  await settle();
  assert.equal(get("#retry-service").hidden, true);
  assert.equal(get("#retry-service").disabled, false);
  assert.equal(get("#submit").disabled, false);
  assert.equal(get("#status").textContent, "");
  assert.equal(get("#project-id").value, "ENGINEERING-042");
  assert.equal(get("#address").value, "123 Sample Street");
  assert.equal(get("#as-of").value, "2026-10-27");
  assert.equal(get("#occupancy").value, "business");
  assert.equal(page.calls.length, 3);
});
