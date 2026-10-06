(() => {
'use strict';
const payload = JSON.parse(document.getElementById('run-data').textContent);
const obj = value => value !== null && typeof value === 'object' && !Array.isArray(value);
const report = document.documentElement.dataset.report === 'true' ? payload : null;
const run = report ? report.run : payload;
const article = report?.article ?? null;
const research = report?.research ?? null;
const $ = id => document.getElementById(id);
const arr = value => Array.isArray(value) ? value : [];
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
const groups = ['sources','claims','conclusions','queries','events','decisions','rejected_sources'];
const records = new Map();
for(const group of groups) for(const item of arr(run[group])) if(item?.id != null) {
 const key = String(item.id); const entries=records.get(key)||[]; entries.push({group,item}); records.set(key,entries);
}
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
 let engineID='__edge_'+i; while(ids.has(engineID)) engineID+='_' ; ids.add(engineID);
 edges.push({...e,source:String(e.source),target:String(e.target),type:text(e.type??'related'),engineID});
}
function references(node) {
 const id = String(node?.reference_id ?? node?.id ?? '');
 const all = records.get(id)||[];
 const group = {source:'sources',claim:'claims',conclusion:'conclusions'}[node?.type];
 return group ? all.filter(r=>r.group===group) : all;
}
function nodeStatus(node) { return text(node.status ?? references(node)[0]?.item?.status ?? ''); }
function articleTextForNode(id) {
 const parts=[];
 for(const section of arr(article?.sections)) for(const block of arr(section.blocks)) {
  if(arr(block.claim_ids).includes(id) || arr(block.conclusion_ids).includes(id)) parts.push(block.text, ...arr(block.items), ...arr(block.rows).flat());
 }
 return parts.join(' ').toLowerCase();
}
function nodeSearch(node) { return pretty([node,...references(node).map(r=>r.item)]).toLowerCase()+' '+articleTextForNode(node.id); }
function articleMentions(id) {
 const related=new Set([String(id)]), node=nodes.find(item=>item.id===String(id));
 if(node?.type==='source') for(const claim of arr(run.claims)) if(arr(claim.evidence).some(evidence=>String(evidence.source_id)===String(id))) related.add(String(claim.id));
 if(node?.type==='conclusion') { const conclusion=arr(run.conclusions).find(item=>String(item.id)===String(id)); for(const claimID of arr(conclusion?.claim_ids)) related.add(String(claimID)); }
 const matches=[];
 for(const section of arr(article?.sections)) for(const block of arr(section.blocks)) {
  if([...arr(block.claim_ids),...arr(block.conclusion_ids)].some(reference=>related.has(String(reference)))) matches.push({section,block});
 }
 return matches;
}
function addReadableRecord(parent, group, item) {
 parent.append(el('h3',label(item)));
 if(group === 'sources') {
  parent.append(el('p',item.type ? 'Source type: '+item.type : 'Source','muted'));
  if(item.url) links(parent,item.url);
  if(item.content) { parent.append(el('p','Inspected passage','muted'),el('blockquote',item.content)); }
  if(item.published_at) parent.append(el('p','Published: '+item.published_at));
  if(item.retrieved_at) parent.append(el('p','Retrieved: '+item.retrieved_at));
  if(item.verification) parent.append(el('p','Retrieval check: '+item.verification));
  if(item.reason) parent.append(el('p','Selection note: '+item.reason));
  for(const claim of arr(run.claims)) for(const evidence of arr(claim.evidence)) if(evidence.source_id===item.id) {
   const box=el('section',undefined,'record');
   box.append(el('h3',label(claim)),el('p',evidence.relation==='contradicts'?'Contradictory evidence':'Supporting evidence'),el('blockquote',evidence.quote));
   box.append(el('p','Location: '+text(evidence.locator)),el('p','Retrieval check: '+text(evidence.verification)+'. This is a recorded assessment, not guaranteed truth.','muted'));
   parent.append(box);
  }
 } else if(group === 'claims') {
  if(item.text) parent.append(el('p',item.text));
  for(const evidence of arr(item.evidence)) {
   const source=arr(run.sources).find(s=>s.id===evidence.source_id);
   const box=el('section',undefined,'record');
   box.append(el('p',evidence.relation==='contradicts'?'Contradictory evidence':'Supporting evidence','muted'));
   box.append(el('p',evidence.quote));
   if(evidence.locator) box.append(el('p','Location: '+evidence.locator));
   if(evidence.verification) box.append(el('p','Retrieval check: '+evidence.verification));
   if(source) { const b=el('button','Open source: '+label(source)); b.type='button'; b.addEventListener('click',()=>focusEvidence(source.id)); box.append(b); }
   parent.append(box);
  }
 } else if(group === 'conclusions') {
  if(item.text) parent.append(el('p',item.text));
  parent.append(el('p','Assessment: '+text(item.status??'not recorded'),'muted'));
 } else {
  jsonBlock(parent,item);
 }
 const details=el('details'); details.append(el('summary','Full record')); jsonBlock(details,item); parent.append(details);
}
function focusArticleBlock(match) {
 const node=document.getElementById('article-block-'+match.block.id);
 if(node) { node.scrollIntoView({behavior:'instant',block:'center'}); node.focus({preventScroll:true}); }
}
function showDetail(value, kind, nodeID) {
 const parent=$('detail-content'); empty(parent); parent.append(el('h3',kind+': '+label(value)));
 if(nodeID) {
  const n=nodes.find(n=>n.id===nodeID);
  for(const record of references(n)) addReadableRecord(parent,record.group,record.item);
  parent.append(el('h3','Recorded relationships'));
  const related=edges.filter(e=>e.source===nodeID||e.target===nodeID);
  if(!related.length) parent.append(el('p','No explicit relationships recorded.','muted'));
  for(const edge of related) { const b=el('button',edge.source+' → '+edge.type+' → '+edge.target); b.type='button'; b.addEventListener('click',()=>showDetail(edge,'Relationship')); parent.append(b); }
  const mentions=articleMentions(nodeID);
  if(mentions.length) {
   parent.append(el('h3','Used in the article'));
   for(const match of mentions) { const b=el('button',match.section.heading); b.type='button'; b.className='article-mention'; b.addEventListener('click',()=>focusArticleBlock(match)); parent.append(b); }
  } else if(article) { parent.append(el('p','No article reference for this record.','muted')); }
 } else {
  jsonBlock(parent,value); links(parent,value);
 }
 $('detail').scrollIntoView({behavior:'instant',block:'nearest'});
}
function recordList(parent, items, kind) {
 if(!items.length) { parent.append(el('p','No '+kind+' recorded.','muted')); return; }
 for(const item of items) {
  const detail=el('details',undefined,'record'); detail.append(el('summary',label(item)));
  const b=el('button','Inspect '+kind); b.type='button'; b.addEventListener('click',()=>{showDetail(item,kind); $('detail').scrollIntoView({behavior:'instant',block:'start'});});
  detail.append(b); jsonBlock(detail,item); links(detail,item); parent.append(detail);
 }
}
function renderArticle() {
 if(!article) return;
 $('article-panel').hidden=false;
 $('article-lead').textContent=article.lead;
 document.documentElement.lang=article.language||'en';
 const body=$('article-body'); empty(body);
 for(const section of arr(article.sections)) {
  const sectionNode=el('section',undefined,'article-section'); sectionNode.append(el('h2',section.heading));
  for(const block of arr(section.blocks)) {
   const content=el('div',undefined,'article-block'); content.id='article-block-'+block.id; content.tabIndex=-1; content.dataset.role=block.role;
   content.append(el('span',block.role,'role'));
   if(block.kind==='paragraph') content.append(el('p',block.text));
   else if(block.kind==='list') { const list=el('ul'); for(const item of arr(block.items)) list.append(el('li',item)); content.append(list); }
   else if(block.kind==='table') {
    const table=el('table'), head=el('thead'), header=el('tr'), rows=el('tbody');
    for(const item of arr(block.headers)) header.append(el('th',item)); head.append(header);
    for(const row of arr(block.rows)) { const tr=el('tr'); for(const cell of arr(row)) tr.append(el('td',cell)); rows.append(tr); }
    table.append(head,rows); const scroll=el('div',undefined,'scroll'); scroll.append(table); content.append(scroll);
   }
   const citationIDs=[...new Set([...arr(block.claim_ids),...arr(block.conclusion_ids)])];
   if(citationIDs.length) {
    const citations=el('span',undefined,'citations');
    for(const id of citationIDs) { const button=el('button',(arr(block.claim_ids).includes(id)?'Claim ':'Conclusion ')+id,'citation'); button.type='button'; button.dataset.claimId=id; button.addEventListener('click',()=>focusEvidence(id)); citations.append(button); }
    content.append(citations);
   }
   sectionNode.append(content);
  }
  body.append(sectionNode);
 }
}
const reportTitle = article?.title ?? run.metadata?.title ?? run.metadata?.topic ?? run.contract?.question ?? 'Research report';
$('report-title').textContent = reportTitle;
document.title = reportTitle;
if(report) $('report-subtitle').textContent=article.output_form+'. Research stopped because: '+research.stop.reason.replaceAll('_',' ')+'.';
$('metadata').textContent=pretty({version:run.version ?? run.schema_version,id:run.id,metadata:run.metadata});
$('raw').textContent=pretty(payload);
const contract=run.contract??{},contractView=$('contract');empty(contractView);
const questions=arr(contract.questions).map(question=>text(question?.text??question)).filter(Boolean);
const mainQuestion=text(contract.question??questions[0]??'No research question recorded.');
contractView.append(el('p',mainQuestion,'lead'));
const subquestions=questions.filter(question=>question!==mainQuestion);
if(subquestions.length>0){contractView.append(el('h3','Core questions'));const list=el('ul');for(const question of subquestions)list.append(el('li',question));contractView.append(list);}
for(const [key,title] of [['types','Source types'],['excludes','Excluded'],['freshness','Freshness'],['preferences','Source preferences']]) {
 const value=Array.isArray(contract[key])?contract[key].join('; '):text(contract[key]??'');
 if(value) contractView.append(el('p',title+': '+value,'muted'));
}
const contractDetails=el('details');contractDetails.append(el('summary','Full research contract'));jsonBlock(contractDetails,contract);contractView.append(contractDetails);
for(const group of groups) $('counts').append(el('span',group.replaceAll('_',' ')+': '+arr(run[group]).length,'badge'));
$('counts').append(el('span','nodes: '+rawNodes.length,'badge'),el('span','relationships: '+rawEdges.length,'badge'));
$('process').append(el('h3','Pipeline stages'));
recordList($('process'),arr(run.stages),'stages');
$('process').append(el('h3','Recorded candidate changes'));
const countEvents=arr(run.events).filter(e=>Number.isInteger(e.before_count)&&Number.isInteger(e.after_count));
if(!countEvents.length) $('process').append(el('p','No structured before/after counts recorded.','muted'));
for(const event of countEvents) $('process').append(el('p',event.stage+' · '+event.count_scope+': '+event.before_count+' → '+event.after_count));
for(const group of ['queries','events']) { $('process').append(el('h3',group)); recordList($('process'),arr(run[group]),group); }
recordList($('conclusions'),arr(run.conclusions),'conclusions');
const selected=arr(run.sources).filter(s=>s?.selected===true || ['selected','accepted','final'].includes(s?.status));
$('conclusions').append(el('h3','Explicitly selected materials'));
recordList($('conclusions'),selected,'selected sources');
for(const group of ['sources','claims']) { $('evidence').append(el('h3',group)); recordList($('evidence'),arr(run[group]),group); }
for(const group of ['decisions','rejected_sources']) { $('decisions').append(el('h3',group.replaceAll('_',' '))); recordList($('decisions'),arr(run[group]),group); }
function renderCoverage(coverage) {
 const parent=$('coverage'); empty(parent);
 parent.append(el('p','Verification records an assessment, not guaranteed truth. Original-evidence origin groups do not establish independent publishers or statistical corroboration.','muted'));
 if(coverage == null) { parent.append(el('p','No coverage recorded.','muted')); return; }
 const rows=Array.isArray(coverage)?coverage:obj(coverage)?Object.entries(coverage).map(([dimension,value])=>({dimension,value})):[{value:coverage}];
 const table=el('table'), head=el('thead'), body=el('tbody'), tr=el('tr');
 const keys=[...new Set(rows.flatMap(r=>obj(r)?Object.keys(r):['value']))];
 for(const key of keys) tr.append(el('th',key)); head.append(tr); table.append(head,body);
 for(const row of rows) { const tr=el('tr'); for(const key of keys) { const v=obj(row)?row[key]:row; tr.append(el('td',typeof v==='object'?pretty(v):v??'')); } body.append(tr); }
 const scroll=el('div',undefined,'scroll'); scroll.append(table); parent.append(scroll);
}
renderCoverage(research?.coverage ?? run.coverage);
function renderResearch() {
 if(!research) return;
 $('saturation-panel').hidden=false;
 $('stop-summary').textContent='Research stopped: '+research.stop.reason.replaceAll('_',' ')+'. '+research.stop.rationale;
 const metrics=$('research-data'); empty(metrics);
 metrics.append(el('p','Audience: '+research.audience));
 metrics.append(el('p','Purpose: '+research.purpose));
 if(research.origin_types_rationale) metrics.append(el('p','Source types: '+research.origin_types_rationale,'muted'));
 metrics.append(el('p',research.rounds.length+' rounds; low-gain window '+research.low_gain_window+'; budget '+research.max_rounds,'badge'));
 if(research.resource_budget) { const b=research.resource_budget; metrics.append(el('p','Resource budget: '+b.kind+' — '+b.used+' / '+b.limit+' '+b.unit)); }
 const summary=el('div',undefined,'scroll'), table=el('table'), head=el('thead'), header=el('tr'), body=el('tbody');
 for(const key of ['Round','Search angle','Intent','Coverage','Open gaps','New claims','New origins','New contradictions','New questions','Redundant','Assessed','Redundant ratio','Gain']) header.append(el('th',key));
 head.append(header);
 for(const round of arr(research.rounds)) {
  const assessed=arr(round.assessed_source_ids).length, redundant=arr(round.redundant_source_ids).length;
  const ratio=assessed?Math.round(100*redundant/assessed)+'%':'not measured';
  const snapshot=arr(round.coverage_snapshot), covered=snapshot.filter(item=>item.status==='covered').length;
  const values=[round.id,round.search_angle,round.search_intent,covered+'/'+snapshot.length,arr(round.gaps).length,arr(round.new_claim_ids).length,arr(round.new_origin_ids).length,arr(round.new_contradiction_refs).length,arr(round.new_question_ids).length,redundant,assessed,ratio,round.material_gain];
  const tr=el('tr'); for(const value of values) tr.append(el('td',value)); body.append(tr);
 }
 table.append(head,body); summary.append(table); metrics.append(summary);
 const roundDetails=el('div'); recordList(roundDetails,arr(research.rounds),'round details'); metrics.append(roundDetails);
 const gapParent=$('research-gaps'); empty(gapParent);
 if(!arr(research.gaps).length) gapParent.append(el('p','No additional research gaps recorded.','muted'));
 else { const list=el('ul'); for(const gap of research.gaps) list.append(el('li',gap)); gapParent.append(list); }
}
renderArticle();
renderResearch();
for(const [id,values] of [['type',nodes.map(n=>n.type)],['status',nodes.map(nodeStatus)],['relation',edges.map(e=>e.type)]]) {
 for(const value of [...new Set(values)].filter(Boolean).sort()) { const option=el('option',value); option.value=value; $(id).append(option); }
}
let cy;
try {
 cy=cytoscape({container:$('graph'),elements:[],style:[
  {selector:'node',style:{'background-color':'#477e79','label':'data(label)','color':'#202c31','font-size':12,'text-wrap':'wrap','text-max-width':130,'text-valign':'bottom','text-margin-y':6,'width':30,'height':30}},
  {selector:'node[type="claim"]',style:{'background-color':'#d0a444','color':'#202c31','shape':'round-rectangle'}},
  {selector:'node[type="conclusion"]',style:{'background-color':'#55796d','shape':'diamond'}},
  {selector:'edge',style:{'width':2,'line-color':'#748580','target-arrow-color':'#748580','target-arrow-shape':'triangle','curve-style':'bezier','label':'data(type)','color':'#202c31','font-size':10,'text-background-color':'#f7f8f5','text-background-opacity':1,'text-background-padding':3}},
  {selector:'edge[type="contradicts"],edge[type="conflicts"]',style:{'line-color':'#a6402e','target-arrow-color':'#a6402e','line-style':'dashed'}},
  {selector:':selected',style:{'border-width':3,'border-color':'#263e45','line-color':'#263e45','target-arrow-color':'#263e45'}}
 ],layout:{name:'preset'},wheelSensitivity:0.2});
 cy.on('tap','node',event=>{const n=nodes.find(n=>n.id===event.target.id());showDetail(n,'Node',n.id);});
 cy.on('tap','edge',event=>showDetail(edges.find(e=>e.engineID===event.target.id()),'Relationship'));
} catch(error) { problems.push('Graph unavailable: '+error.message+'. Article citations and source records remain accessible.'); }
function focusEvidence(id) {
 const node=nodes.find(n=>n.id===id);
 if(node) {
  showDetail(node,'Node',node.id);
  if(cy) {
   for(const filter of ['search','type','status','relation']) $(filter).value='';
   update();
   const selected=cy.getElementById(node.id);
   let neighborhood=selected.closedNeighborhood();
   if(node.type==='conclusion') neighborhood=neighborhood.union(neighborhood.nodes('[type="claim"]').closedNeighborhood());
   cy.elements().unselect(); neighborhood.select(); cy.fit(neighborhood,50);
  }
 }
 $('graph-panel').scrollIntoView({behavior:'instant',block:'start'});
}
function update() {
 const query=$('search').value.trim().toLowerCase(), type=$('type').value, status=$('status').value, relation=$('relation').value;
 const visible=nodes.filter(n=>(!query||nodeSearch(n).includes(query))&&(!type||n.type===type)&&(!status||nodeStatus(n)===status));
 const visibleIDs=new Set(visible.map(n=>n.id));
 const visibleEdges=edges.filter(e=>visibleIDs.has(e.source)&&visibleIDs.has(e.target)&&(!relation||e.type===relation));
 $('graph-summary').textContent=visible.length+' / '+nodes.length+' valid nodes · '+visibleEdges.length+' / '+edges.length+' valid relationships';
 $('warnings').textContent=problems.length?problems.join(' '):!nodes.length?'No explicit graph nodes recorded. Browse the report records below.':'';
 if(cy) { cy.elements().remove(); cy.add(visible.map(n=>({data:{id:n.id,type:n.type,label:label(n)}})).concat(visibleEdges.map(e=>({data:{id:e.engineID,source:e.source,target:e.target,type:e.type}})))); cy.layout({name:'breadthfirst',directed:true,animate:false,padding:30,spacingFactor:1.25}).run(); }
 empty($('graph-list'));
 for(const n of visible) { const b=el('button',n.type+': '+label(n)); b.type='button'; b.addEventListener('click',()=>{showDetail(n,'Node',n.id);if(cy){cy.elements().unselect();cy.getElementById(n.id).select();}}); $('graph-list').append(b); }
 for(const e of visibleEdges) { const b=el('button',e.source+' → '+e.type+' → '+e.target);b.type='button';b.addEventListener('click',()=>showDetail(e,'Relationship'));$('graph-list').append(b); }
}
for(const id of ['search','type','status','relation']) $(id).addEventListener(id==='search'?'input':'change',update);
$('reset').addEventListener('click',()=>{for(const id of ['search','type','status','relation']) $(id).value='';update();});
$('fit').addEventListener('click',()=>cy?.fit(undefined,30));
if(cy && typeof ResizeObserver!=='undefined') new ResizeObserver(()=>cy.resize()).observe($('graph'));
window.addEventListener('resize',()=>cy?.resize());
update();
})();
