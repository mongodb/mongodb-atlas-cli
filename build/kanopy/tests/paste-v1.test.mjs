import assert from "node:assert/strict";
import { createHash, webcrypto } from "node:crypto";
import { readFile } from "node:fs/promises";
import { test } from "node:test";
import vm from "node:vm";

const html = await readFile(new URL("../oauth-callback/index.html", import.meta.url), "utf8");
const nginx = await readFile(new URL("../nginx.conf", import.meta.url), "utf8");
const vectors = JSON.parse(await readFile(new URL("../testdata/paste-v1.json", import.meta.url), "utf8"));

function inlineContent(tag) {
  const opening = `<${tag}>`;
  const start = html.indexOf(opening) + opening.length;
  const end = html.indexOf(`</${tag}>`, start);
  assert.notEqual(start, opening.length - 1, `missing inline ${tag}`);
  assert.notEqual(end, -1, `unterminated inline ${tag}`);
  return html.slice(start, end);
}

const inlineScript = inlineContent("script");

function cspHash(content) {
  return `sha256-${createHash("sha256").update(content, "utf8").digest("base64")}`;
}

async function renderCallbackPage(search) {
  const ids = [
    "callback-page",
    "status",
    "payload-section",
    "payload",
    "payload-copy-status",
    "cli-acceptance-hint",
    "copy-payload",
    "callback-section",
    "callback-url",
    "callback-copy-status",
    "copy-callback-url",
  ];
  const elements = new Map(ids.map((id) => [id, {
    dataset: {},
    hidden: true,
    value: "",
    textContent: "",
    listeners: {},
    addEventListener(event, callback) {
      this.listeners[event] = callback;
    },
    focus() {},
    select() {},
  }]));
  const location = {
    search,
    hash: "#oauth-fragment",
    origin: "https://callback.example.test",
    pathname: "/atlas-cli/oauth-callback.html",
  };

  vm.runInNewContext(inlineScript, {
    URL,
    URLSearchParams,
    TextEncoder,
    crypto: webcrypto,
    btoa: (value) => Buffer.from(value, "binary").toString("base64"),
    navigator: { clipboard: { writeText: async () => {} } },
    location,
    history: {
      replaceState(_state, _title, pathname) {
        location.pathname = pathname;
        location.search = "";
        location.hash = "";
      },
    },
    document: {
      getElementById(id) {
        assert.ok(elements.has(id), `missing page element: ${id}`);
        return elements.get(id);
      },
    },
  });

  const root = elements.get("callback-page");
  for (let attempt = 0; attempt < 100 && root.dataset.ready !== "true"; attempt++) {
    await new Promise((resolve) => setTimeout(resolve, 10));
  }
  assert.equal(root.dataset.ready, "true", "callback page did not finish processing");
  return { elements, location };
}

test("inline CSP hashes match the page assets and NGINX headers", () => {
  for (const tag of ["script", "style"]) {
    const hash = cspHash(inlineContent(tag));
    assert.ok(html.includes(hash), `HTML CSP is missing the ${tag} hash`);
    assert.ok(nginx.includes(hash), `NGINX CSP is missing the ${tag} hash`);
  }
  assert.match(nginx, /Cache-Control "no-store"/);
  assert.match(nginx, /Referrer-Policy "no-referrer"/);
  assert.match(html, /name="referrer" content="no-referrer"/);
  assert.match(html, /name="robots" content="noindex, nofollow"/);
  assert.match(html, /<noscript>[\s\S]*address bar[\s\S]*<\/noscript>/);
  assert.ok(nginx.includes("$uri"), "NGINX access log must record only the query-free URI");
  assert.ok(!nginx.includes("$request_uri") && !nginx.includes("$args"), "NGINX access log must not include query values");
  assert.doesNotMatch(html, /<script\b[^>]*\bsrc=/i);
  assert.doesNotMatch(inlineScript, /\b(?:fetch|XMLHttpRequest|sendBeacon)\s*\(/);
});

test("page encoder matches the shared Go decoder vector and clears the callback URL", async () => {
  assert.equal(vectors.format, "atlas-cli-paste");
  assert.equal(vectors.version, "1");
  assert.ok(vectors.vectors.some((vector) => vector.issuer === ""), "vectors must cover callbacks without an issuer");
  for (const vector of vectors.vectors) {
    const query = new URLSearchParams({ code: vector.code, state: vector.state });
    if (vector.issuer !== "") query.set("iss", vector.issuer);
    const page = await renderCallbackPage(`?${query}`);

    assert.equal(page.elements.get("payload").value, vector.paste);
    assert.equal(page.elements.get("payload-section").hidden, false);
    assert.equal(page.elements.get("callback-section").hidden, true);
    assert.equal(page.elements.get("callback-page").dataset.queryCleared, "true");

    if (vector.issuer === "") {
      const blankIssuerQuery = new URLSearchParams(query);
      blankIssuerQuery.set("iss", "");
      const blankIssuerPage = await renderCallbackPage(`?${blankIssuerQuery}`);

      assert.equal(blankIssuerPage.elements.get("payload").value, vector.paste);
      assert.equal(blankIssuerPage.elements.get("callback-page").dataset.resultKind, "success");
    }
  }
});

test("copying a response reveals the close hint", async () => {
  const query = new URLSearchParams({ code: "synthetic", state: "synthetic", iss: "https://authorize.mongodb.com" });
  const page = await renderCallbackPage(`?${query}`);
  const hint = page.elements.get("cli-acceptance-hint");

  assert.equal(hint.hidden, true);
  await page.elements.get("copy-payload").listeners.click();
  assert.equal(page.elements.get("payload-copy-status").textContent, "Copied.");
  assert.equal(hint.hidden, false);
});

test("malformed code and state parameters do not produce a paste", async () => {
  const issuer = "https%3A%2F%2Fauthorize.mongodb.com";
  const invalidResponses = [
    ["missing code", `?state=state-123&iss=${issuer}`],
    ["missing state", `?code=code-123&iss=${issuer}`],
    ["empty code", `?code=&state=state-123&iss=${issuer}`],
    ["empty state", `?code=code-123&state=&iss=${issuer}`],
    ["duplicate code", `?code=first&code=second&state=state-123&iss=${issuer}`],
    ["duplicate state", `?code=code-123&state=first&state=second&iss=${issuer}`],
  ];

  for (const [name, query] of invalidResponses) {
    const page = await renderCallbackPage(query);

    assert.equal(page.elements.get("callback-page").dataset.resultKind, "invalid", name);
    assert.equal(page.elements.get("payload-section").hidden, true, name);
    assert.equal(page.elements.get("status").textContent, "Authorization could not be completed. Return to the Atlas CLI and try again.", name);
  }
});

test("OAuth errors show a generic failure and retain a full-URL fallback", async () => {
  const page = await renderCallbackPage("?error=access_denied&error_description=do-not-display&state=state-123");
  const status = page.elements.get("status").textContent;

  assert.equal(status, "Authorization could not be completed. Return to the Atlas CLI and try again.");
  assert.equal(page.elements.get("callback-section").hidden, false);
  assert.match(page.elements.get("callback-url").value, /error=access_denied/);
  assert.doesNotMatch(status, /access_denied|do-not-display/);
  assert.equal(page.elements.get("callback-page").dataset.queryCleared, "true");
});
