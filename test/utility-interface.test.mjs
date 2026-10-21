import test from "node:test";
import assert from "node:assert/strict";
import {
  validProjectUtilityContext, renderProjectUtilityContext, installProjectUtilityContext,
} from "../cmd/bcm/web/utility-context.js";

const goodRequest = {project_id:"SITE-1", address:"100 Main St, Denver, CO"};
function response() {
  return {
    project_id:goodRequest.project_id, status:"geography_observed_utility_candidates_unconfirmed",
    geography:{state:"CO", county:"Denver County", county_source_id:"boundary-1", address_precision:"address_point",
      address_source_id:"address-1", provider_bundle_digest:"sha256:reviewed",
      boundary_snapshot_sha256:"b".repeat(64), geocoder_snapshot_sha256:"g".repeat(64)},
    utility_context:{status:"reported_county_candidates_not_site_verified", state:"CO", county:"Denver",
      release_id:"eia861-2025-fixed", reporting_year:2025, source_sha256:"a".repeat(64),
      sqlite_sha256:"c".repeat(64), view_sha256:"d".repeat(64),
      disclaimer:"No site-level serving utility established.",
      candidates:[{eia_id:15466, reported_name:"Example & Utility",source_sheet:"Counties_States",source_row:23,
        activities:[{code:"Distribution",value:"Y"},{code:"Other",value:"N"}]}]},
  };
}
class FakeElement {
  constructor() {this.children=[];this.handlers={};this.hidden=true;this.value="";this.className="";this.textContent="";}
  append(x){this.children.push(x);}
  replaceChildren(){this.children=[];}
  addEventListener(name, callback){this.handlers[name]??=[];this.handlers[name].push(callback);}
  async fire(name){await Promise.all((this.handlers[name]||[]).map(callback=>callback({preventDefault(){}})));}
  text(){return this.textContent + this.children.map(c=>c.text()).join(" ");}
}
function fakeDoc(live=true) {
  const els=new Map();
  for (const key of ["#lookup","#utility-panel","#utility-content","#utility-status","#project-id","#address"])els.set(key,new FakeElement());
  els.get("#project-id").value=goodRequest.project_id;
  els.get("#address").value=goodRequest.address;
  return {
    els, createElement:()=>new FakeElement(),
    querySelector: key => key==='meta[name="bcm-live-api"]' ? (live?{content:"https://bcm-api.example"}:null)
      :key==='meta[name="bcm-static-demo"]' ? (!live?{}:null) : els.get(key),
  };
}
test("accepts pinned candidates, not arbitrary or incomplete geographic facts",()=>{
  const baseline=response();assert.ok(validProjectUtilityContext(baseline,goodRequest));
  for (const mutate of [v=>v.project_id="OTHER",v=>v.geography.county_source_id="",
    v=>v.geography.address_precision="interpolated",v=>v.utility_context.state="WY",
    v=>v.utility_context.view_sha256="missing",v=>v.status="serving_utility_verified",
    v=>v.utility_context.candidates[0].eia_id=0]){
    const v=response();mutate(v);assert.ok(!validProjectUtilityContext(v,goodRequest));
  }
});
test("renders candidate status and provenance using text nodes only",()=>{
  const doc=fakeDoc(),panel=doc.els.get("#utility-content");
  renderProjectUtilityContext(doc,panel,response());
  const text=panel.text();
  assert.match(text,/Example & Utility/);assert.match(text,/Reported county associations \(1\)/);
  assert.match(text,/not a confirmed provider/);assert.match(text,/View SHA-256 d{64}/);
  assert.match(text,/Distribution/);assert.doesNotMatch(text,/Other,/);
});
test("synthetic Pages mode never registers a real utility fetch",async()=>{
  const doc=fakeDoc(false);let fetched=false;
  installProjectUtilityContext(doc,async()=>{fetched=true;throw Error("unexpected")});
  await doc.els.get("#lookup").fire("submit");assert.equal(fetched,false);
  assert.equal(doc.els.get("#utility-panel").hidden,true);
});
test("live lookup independently returns candidates and invalidates on new project facts",async()=>{
  const doc=fakeDoc();let fetchOptions,resolve;
  installProjectUtilityContext(doc,(_url,options)=>{fetchOptions=options;return new Promise(r=>resolve=r)});
  const running=doc.els.get("#lookup").fire("submit");
  assert.equal(JSON.parse(fetchOptions.body).address,goodRequest.address);
  assert.equal(fetchOptions.credentials,"omit");assert.equal(doc.els.get("#utility-panel").hidden,false);
  await doc.els.get("#lookup").fire("input");
  assert.equal(fetchOptions.signal.aborted,true);assert.equal(doc.els.get("#utility-panel").hidden,true);
  resolve({ok:true,json:async()=>response()});await running;
  assert.equal(doc.els.get("#utility-content").text(),"");
});
test("live uncertain county status keeps code-basis independent and shows no candidates",async()=>{
  const doc=fakeDoc();
  installProjectUtilityContext(doc,async()=>({ok:false,status:422}));
  await doc.els.get("#lookup").fire("submit");
  assert.match(doc.els.get("#utility-status").textContent,/could not be independently resolved/);
  assert.equal(doc.els.get("#utility-content").text(),"");
});
