(() => {
'use strict';
const run = JSON.parse(document.getElementById('run-data').textContent);
const $ = id => document.getElementById(id);
const arr = value => Array.isArray(value) ? value : [];
const obj = value => value !== null && typeof value === 'object' && !Array.isArray(value);
const text = value => typeof value === 'string' ? value : JSON.stringify(value ?? '');
const label = item => typeof item === 'string' ? item : text(item?.title ?? item?.text ?? item?.label ?? item?.name ?? item?.id ?? 'Untitled record');
const pretty = value => JSON.stringify(value, null, 2);
function el(tag, value, cls) { const node = document.createElement(tag); if(value !== undefined) node.textContent = text(value); if(cls) node.className = cls; return node; }
function empty(parent) { parent.replaceChildren(); }
function jsonBlock(parent, value) { parent.append(el('pre', pretty(value))); }
function links(parent, value) {
 const seen = new Set();
 function walk(v) {
  if(typeof v === 'string' && /^https?:\/\//i.test(v)) {
   try { const u = new URL(v); if(!['http:', 'https:'].includes(u.protocol) || seen.has(u.href)) return; seen.add(u.href); const a = el('a',v); a.href=u.href; a.target='_blank'; a.rel='noopener noreferrer'; parent.append(a,el('br')); } catch(_) {}
  } else if(Array.isArray(v)) v.forEach(walk); else if(obj(v)) Object.values(v).forEach(walk);
 }
 walk(value);
}
const groups = ['sources','claims','conclusions','queries','adjudications','events','decisions','rejected_sources'];
const records = new Map();
for(const group of groups) for(const item of arr(run[group])) if(item?.id != null) {
 const key = String(item.id); const entries=records.get(key)||[]; entries.push({group,item}); records.set(key,entries);
}
function references(node) {
 const id = String(node.reference_id ?? node.id ?? '');
 const all = records.get(id)||[];
 const group = {source:'sources',claim:'claims',conclusion:'conclusions',query:'queries'}[node.type];
 return group ? all.filter(r=>r.group===group) : all;
}
function nodeSearch(node) { return pretty([node,...references(node).map(r=>r.item)]).toLowerCase(); }
function nodeStatus(node) { return text(node.status ?? references(node)[0]?.item?.status ?? ''); }
const rawNodes = arr(run.graph?.nodes ?? run.nodes);
const rawEdges = arr(run.graph?.edges ?? run.edges).map(edge => ({...edge, source: edge.from ?? edge.source, target: edge.to ?? edge.target}));
const nodes=[], ids=new Set(), problems=[];
for(const n of rawNodes) {
 if(!obj(n) || n.id == null || String(n.id)==='' || ids.has(String(n.id))) { problems.push('Skipped a node with missing or duplicate ID.'); continue; }
 const id=String(n.id); ids.add(id); nodes.push({...n,id,type:text(n.type??'unknown').toLowerCase()});
}
const edges=[];
for(const [i,e] of rawEdges.entries()) {
 if(!obj(e) || !ids.has(String(e.source)) || !ids.has(String(e.target))) { problems.push('Skipped an edge with missing endpoints.'); continue; }
 // Engine IDs are independent of external edge IDs, avoiding collisions.
 let engineID='__edge_'+i; while(ids.has(engineID)) engineID+='_' ; ids.add(engineID);
 edges.push({...e,source:String(e.source),target:String(e.target),type:text(e.type??'related'),engineID});
}
function showDetail(value, kind, nodeID) {
 const parent=$('detail-content'); empty(parent); parent.append(el('h3',kind+': '+label(value))); jsonBlock(parent,value); links(parent,value);
 if(nodeID) {
  const n=nodes.find(n=>n.id===nodeID);
  for(const r of references(n)) { parent.append(el('h3',r.group)); jsonBlock(parent,r.item); links(parent,r.item); }
  parent.append(el('h3','Supports, conflicts, and other explicit relationships'));
  const related=edges.filter(e=>e.source===nodeID||e.target===nodeID);
  if(!related.length) parent.append(el('p','No explicit relationships recorded.','muted'));
  for(const e of related) { const b=el('button',e.source+' → '+e.type+' → '+e.target); b.type='button'; b.addEventListener('click',()=>showDetail(e,'Relationship')); parent.append(b); }
 }
}
function recordList(parent, items, kind) {
 if(!items.length) { parent.append(el('p','No '+kind+' recorded.','muted')); return; }
 for(const item of items) {
  const detail=el('details',undefined,'record'); detail.append(el('summary',label(item)));
  const b=el('button','Inspect '+kind); b.type='button'; b.addEventListener('click',()=>{showDetail(item,kind); $('detail').scrollIntoView({behavior:'instant',block:'start'});});
  detail.append(b); jsonBlock(detail,item); links(detail,item); parent.append(detail);
 }
}
const reportTitle = run.metadata?.title ?? run.metadata?.topic ?? run.contract?.question ?? 'Research evidence report';
$('report-title').textContent = reportTitle;
document.title = reportTitle;
$('metadata').textContent=pretty({version:run.version ?? run.schema_version,id:run.id,metadata:run.metadata});
$('raw').textContent=pretty(run);
jsonBlock($('contract'),run.contract??{});
for(const group of groups) $('counts').append(el('span',group.replaceAll('_',' ')+': '+arr(run[group]).length,'badge'));
$('counts').append(el('span','nodes: '+rawNodes.length,'badge'),el('span','edges: '+rawEdges.length,'badge'));
$('process').append(el('h3','Pipeline stages'));
recordList($('process'),arr(run.stages),'stages');
$('process').append(el('h3','Recorded candidate changes'));
const countEvents=arr(run.events).filter(e=>Number.isInteger(e.before_count)&&Number.isInteger(e.after_count));
if(!countEvents.length) $('process').append(el('p','No structured before/after counts recorded.','muted'));
for(const event of countEvents) $('process').append(el('p',event.stage+' · '+event.count_scope+': '+event.before_count+' → '+event.after_count));
for(const group of ['queries','events']) { $('process').append(el('h3',group)); recordList($('process'),arr(run[group]),group); }
$('process').append(el('h3','Query → candidate provenance'));
for(const q of arr(run.queries)) {
 const block=el('details',undefined,'record'); block.append(el('summary',q.text+' · '+(q.provider||'')+'/'+(q.tool||'')));
 block.append(el('p','Retrieval reference: '+text(q.retrieval_reference||''),'muted'));
 for(const id of arr(q.candidate_source_ids)) { const source=arr(run.sources).find(s=>s.id===id); block.append(el('p',id+' → '+text(source?.status||'unknown')+' · '+text(source?.reason||''))); }
 $('process').append(block);
}
recordList($('adjudications'),arr(run.adjudications),'conflict adjudications');
recordList($('conclusions'),arr(run.conclusions),'conclusions');
const selected=arr(run.sources).filter(s=>s?.selected===true || ['selected','accepted','final'].includes(s?.status));
$('conclusions').append(el('h3','Explicitly selected materials'));
recordList($('conclusions'),selected,'selected sources');
for(const group of ['sources','claims']) { $('evidence').append(el('h3',group)); recordList($('evidence'),arr(run[group]),group); }
for(const group of ['decisions','rejected_sources']) { $('decisions').append(el('h3',group.replaceAll('_',' '))); recordList($('decisions'),arr(run[group]),group); }
// Generic coverage remains faithful to the input: no invented coverage scores.
const coverage=run.coverage;
$('coverage').append(el('p','Verification records an assessment, not guaranteed truth. independent_sources counts recorded original-evidence provenance groups; it does not establish independent publishers or statistical corroboration.','muted'));
if(coverage == null) $('coverage').append(el('p','No coverage recorded.','muted'));
else {
 const rows=Array.isArray(coverage)?coverage:obj(coverage)?Object.entries(coverage).map(([dimension,value])=>({dimension,value})):[{value:coverage}];
 const table=el('table'), head=el('thead'), body=el('tbody'), tr=el('tr');
 const keys=[...new Set(rows.flatMap(r=>obj(r)?Object.keys(r):['value']))];
 for(const key of keys) tr.append(el('th',key)); head.append(tr); table.append(head,body);
 for(const row of rows) { const tr=el('tr'); for(const key of keys) { const v=obj(row)?row[key]:row; tr.append(el('td',typeof v==='object'?pretty(v):v??'')); } body.append(tr); }
 const scroll=el('div',undefined,'scroll'); scroll.append(table); $('coverage').append(scroll);
}
for(const [id,values] of [['type',nodes.map(n=>n.type)],['status',nodes.map(nodeStatus)],['relation',edges.map(e=>e.type)]]) {
 for(const value of [...new Set(values)].filter(Boolean).sort()) { const option=el('option',value); option.value=value; $(id).append(option); }
}
let cy;
try {
 cy=cytoscape({container:$('graph'),elements:[],style:[
  {selector:'node',style:{'background-color':'#8ecbff','label':'data(label)','color':'#eef3fa','font-size':12,'text-wrap':'wrap','text-max-width':130,'text-valign':'bottom','text-margin-y':6,'width':30,'height':30}},
  {selector:'node[type="claim"]',style:{'background-color':'#ffdc80','shape':'round-rectangle'}},
  {selector:'node[type="conclusion"]',style:{'background-color':'#91e3b6','shape':'diamond'}},
  {selector:'edge',style:{'width':2,'line-color':'#bac7da','target-arrow-color':'#bac7da','target-arrow-shape':'triangle','curve-style':'bezier','label':'data(type)','color':'#eef3fa','font-size':10,'text-background-color':'#111b29','text-background-opacity':1,'text-background-padding':3}},
  {selector:'edge[type="conflicts"],edge[type="contradicts"]',style:{'line-color':'#ffafc0','target-arrow-color':'#ffafc0','line-style':'dashed'}},
  {selector:':selected',style:{'border-width':3,'border-color':'#ffffff','line-color':'#ffffff','target-arrow-color':'#ffffff'}}
 ],layout:{name:'preset'},wheelSensitivity:0.2});
 cy.on('tap','node',event=>{const n=nodes.find(n=>n.id===event.target.id());showDetail(n,'Node',n.id);});
 cy.on('tap','edge',event=>showDetail(edges.find(e=>e.engineID===event.target.id()),'Relationship'));
} catch(error) { problems.push('Graph unavailable: '+error.message+'. All evidence remains accessible below.'); }
function update() {
 const query=$('search').value.trim().toLowerCase(), type=$('type').value, status=$('status').value, relation=$('relation').value;
 const visible=nodes.filter(n=>(!query||nodeSearch(n).includes(query))&&(!type||n.type===type)&&(!status||nodeStatus(n)===status));
 const visibleIDs=new Set(visible.map(n=>n.id));
 const visibleEdges=edges.filter(e=>visibleIDs.has(e.source)&&visibleIDs.has(e.target)&&(!relation||e.type===relation));
 $('graph-summary').textContent=visible.length+' / '+nodes.length+' valid nodes · '+visibleEdges.length+' / '+edges.length+' valid relationships';
 $('warnings').textContent=problems.length?problems.join(' '):!nodes.length?'No explicit graph nodes recorded. Browse the run records below.':'';
 if(cy) { cy.elements().remove(); cy.add(visible.map(n=>({data:{id:n.id,type:n.type,label:label(n)}})).concat(visibleEdges.map(e=>({data:{id:e.engineID,source:e.source,target:e.target,type:e.type}})))); cy.layout({name:'breadthfirst',directed:true,animate:false,padding:30,spacingFactor:1.25}).run(); }
 empty($('graph-list'));
 for(const n of visible) { const b=el('button',n.type+': '+label(n)); b.type='button'; b.addEventListener('click',()=>{showDetail(n,'Node',n.id);if(cy){cy.elements().unselect();cy.getElementById(n.id).select();}}); $('graph-list').append(b); }
 for(const e of visibleEdges) { const b=el('button',e.source+' → '+e.type+' → '+e.target);b.type='button';b.addEventListener('click',()=>showDetail(e,'Relationship'));$('graph-list').append(b); }
}
for(const id of ['search','type','status','relation']) $(id).addEventListener(id==='search'?'input':'change',update);
$('reset').addEventListener('click',()=>{for(const id of ['search','type','status','relation']) $(id).value='';update();});
$('fit').addEventListener('click',()=>cy?.fit(undefined,30));
if(cy && typeof ResizeObserver!=='undefined') new ResizeObserver(()=>cy.resize()).observe($('graph'));
update();
})();
