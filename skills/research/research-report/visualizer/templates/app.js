(() => {
'use strict';
const payload = JSON.parse(document.getElementById('run-data').textContent);
const reportMode = document.documentElement.dataset.report === 'true';
const report = reportMode ? payload : null;
const run = report ? report.run : payload;
const research = report?.research ?? null;
const article = report?.article ?? null;
const $ = id => document.getElementById(id);
const arr = value => Array.isArray(value) ? value : [];
const obj = value => value !== null && typeof value === 'object' && !Array.isArray(value);
const str = value => typeof value === 'string' ? value : '';
const zh = /^zh(?:[-_]|$)/i.test(str(article?.language));
const L = zh ? {
 reportKicker:'研究报告', important:'重要限制', readArticle:'阅读正文', sourceNotes:'来源与引文', researchDetails:'研究过程与边界',
 keyLimits:'阅读前请注意', evidenceMap:'证据关系图', evidenceMapIntro:'只显示正文引用的来源、观点和结论。连线表示记录的支持或反驳关系，不代表证据强弱。',
 mapUnavailable:'关系图暂不可用；下方文字关系仍可查看。', mapEmpty:'正文没有可展示的来源—观点关系。', sourceNotesIntro:'编号对应正文中的引文。展开来源可查看实际检索的段落及其用途。',
 evidence:'已核实证据', candidate_lead:'未核实线索', context_source:'背景资料', selected:'已采用', rejected:'未采用', selected_unverified:'台账标为已采用，但来源未核实，不能作为证据',
 sourceType:'资料类型', duplicate:'重复来源', superseded:'已被替代', openSource:'打开原始来源', inspectedPassage:'实际检查的段落', published:'发布于', retrieved:'检索于', upstream:'上游/转引来源', evidenceBalance:'来源集中度提醒',
 usedFor:'用于支持或讨论的观点', supports:'支持', contradicts:'反驳', unverified:'未核实', failed:'获取失败', verified:'已核对原文',
 researchBoundary:'研究范围与停止原因', coverage:'问题覆盖情况', remainingGaps:'尚存缺口与限制', processAndSources:'搜索与来源评估',
 candidateLeads:'未作为证据的线索', contextSources:'背景资料', rounds:'研究轮次', rawData:'完整研究记录',
 noSources:'没有可追溯的来源记录。', noClaims:'正文没有链接到已记录的观点。', noGaps:'没有记录其他研究缺口。', noItems:'无记录。',
 stop_in_progress:'研究仍在进行', stop_saturated:'达到本次范围内的证据饱和', stop_budget_exhausted:'研究预算已用尽', stop_retrieval_blocked:'检索受阻', stop_user_stopped:'按用户要求停止',
 stopNote_saturated:'这表示近期检索的信息增益较低，不代表已经穷尽互联网。',
 coverage_covered:'已有支持', coverage_partial:'部分覆盖', coverage_uncovered:'尚未覆盖',
 round:'轮次', intent:'检索意图', angle:'检索角度', gain:'信息增益', complete:'完成', partial:'部分完成', round_failed:'失败',
 gain_none:'无新增', gain_minor:'少量补充', gain_material:'实质变化', queryTime:'查询记录时间', sourceTime:'来源获取时间',
 noArticle:'此文件是旧版证据台账，没有附带面向读者的文章。', offlineNote:'本报告为离线单文件，不会自动联网；只有读者主动打开来源链接时才访问外部网站。',
 sourceNumber:'来源', source:'来源', claim:'观点', conclusion:'结论', complete:'已完成', partial:'部分完成', failed:'失败', relation_supports:'支持', relation_contradicts:'反驳', relation_derives_from:'归纳自',
 completeReport:'已生成研究报告', reportContents:'目录', processSummary:'范围：', sourcesCount:'个来源', roundsCount:'轮检索',
 rationale:'停止说明', role_evidence:'证据', role_synthesis:'综合判断', role_context:'背景或限制',
 citation:'查看来源', relationship:'关系', externalLink:'在新窗口查看来源',
} : {
 reportKicker:'Research report', important:'Important limits', readArticle:'Read the article', sourceNotes:'Sources and citations', researchDetails:'Research details',
 keyLimits:'Important limits to keep in mind', evidenceMap:'Evidence map', evidenceMapIntro:'Only sources, claims, and conclusions cited by the article appear here. Lines show recorded support or contradiction, not strength.',
 mapUnavailable:'The graph is unavailable; the text relationships below remain accessible.', mapEmpty:'The article has no recorded source-to-claim relationships to show.', sourceNotesIntro:'Numbers match the citations in the article. Expand a source to inspect the retrieved passage and how it was used.',
 evidence:'Verified evidence', candidate_lead:'Candidate lead', context_source:'Context source', selected:'Selected', rejected:'Not selected', selected_unverified:'Ledger marks selected, but retrieval is unverified; not evidence',
 sourceType:'Source type', duplicate:'Duplicate source', superseded:'Superseded', openSource:'Open original source', inspectedPassage:'Inspected passage', published:'Published', retrieved:'Retrieved', upstream:'Upstream or cited origin', evidenceBalance:'Evidence concentration check',
 usedFor:'Claims using this source', supports:'Supports', contradicts:'Contradicts', unverified:'Unverified', failed:'Retrieval failed', verified:'Passage checked',
 researchBoundary:'Research scope and stopping point', coverage:'Question coverage', remainingGaps:'Remaining gaps and limitations', processAndSources:'Search and source assessment',
 candidateLeads:'Leads not used as evidence', contextSources:'Context sources', rounds:'Research rounds', rawData:'Complete research record',
 noSources:'No traceable source records.', noClaims:'The article is not linked to recorded claims.', noGaps:'No additional research gaps were recorded.', noItems:'None recorded.',
 stop_in_progress:'Research is still in progress', stop_saturated:'Evidence saturation reached for this scope', stop_budget_exhausted:'Research budget exhausted', stop_retrieval_blocked:'Retrieval was blocked', stop_user_stopped:'Stopped at the user’s request',
 stopNote_saturated:'This means recent searches added little material information; it does not mean the internet has been exhausted.',
 coverage_covered:'Covered', coverage_partial:'Partly covered', coverage_uncovered:'Not covered',
 round:'Round', intent:'Search intent', angle:'Search angle', gain:'Information gain', complete:'Complete', partial:'Partial', round_failed:'Failed',
 gain_none:'None', gain_minor:'Minor', gain_material:'Material change', queryTime:'Query recorded', sourceTime:'Source retrieved',
 noArticle:'This is a legacy evidence ledger without a reader-facing article.', offlineNote:'This is a self-contained offline report. It makes no automatic network requests; external sites open only when a reader selects a source link.',
 sourceNumber:'Source', source:'Source', claim:'Claim', conclusion:'Conclusion', relation_supports:'supports', relation_contradicts:'contradicts', relation_derives_from:'informs',
 completeReport:'Research report', reportContents:'Contents', processSummary:'Scope: ', sourcesCount:'sources', roundsCount:'research rounds',
 rationale:'Stopping rationale', role_evidence:'Evidence', role_synthesis:'Synthesis', role_context:'Context or limitation',
 citation:'Open source details', relationship:'Relationship', externalLink:'Open source in a new tab',
};
const t = (key, value) => value && L[key + '_' + value] ? L[key + '_' + value] : (L[key] ?? key);
function el(tag, value, cls) { const node = document.createElement(tag); if (value !== undefined && value !== null) node.textContent = String(value); if (cls) node.className = cls; return node; }
function setText(id, value) { const node = $(id); if (node) node.textContent = value; }
function safeURL(value) { try { const url = new URL(value); return ['http:', 'https:'].includes(url.protocol) ? url.href : ''; } catch (_) { return ''; } }
function sourceRole(source) {
 const explicit = research?.source_roles?.[source.id];
 if (explicit) return explicit;
 const refs = arr(run.claims).flatMap(claim => arr(claim.evidence).filter(e => e.source_id === source.id));
 if (source.status === 'selected' && refs.some(e => e.verification === 'verified') && source.verification === 'verified') return 'evidence';
 if (source.status === 'selected' && source.verification === 'verified' && refs.length === 0) return 'context_source';
 return 'candidate_lead';
}
const sources = arr(run.sources);
const sourceByID = new Map(sources.map(source => [String(source.id), source]));
const claims = arr(run.claims);
const claimByID = new Map(claims.map(claim => [String(claim.id), claim]));
const conclusions = arr(run.conclusions);
const conclusionByID = new Map(conclusions.map(item => [String(item.id), item]));
const articleClaimIDs = new Set();
const articleConclusionIDs = new Set();
for (const section of arr(article?.sections)) for (const block of arr(section.blocks)) {
 for (const id of arr(block.claim_ids)) articleClaimIDs.add(String(id));
 for (const id of arr(block.conclusion_ids)) articleConclusionIDs.add(String(id));
}
for (const id of articleConclusionIDs) for (const claimID of arr(conclusionByID.get(id)?.claim_ids)) articleClaimIDs.add(String(claimID));
const citedSourceIDs = new Set();
for (const id of articleClaimIDs) for (const evidence of arr(claimByID.get(id)?.evidence)) citedSourceIDs.add(String(evidence.source_id));
const citedSources = [...citedSourceIDs].map(id => sourceByID.get(id)).filter(Boolean);
const sourceNumbers = new Map(citedSources.map((source, index) => [String(source.id), index + 1]));

function applyLocale() {
 document.documentElement.lang = str(article?.language) || 'en';
 document.querySelectorAll('[data-i18n]').forEach(node => { node.textContent = t(node.dataset.i18n); });
 $('toc').setAttribute('aria-label', t('reportContents'));
 $('graph').setAttribute('aria-label', t('evidenceMap'));
 $('graph-list').setAttribute('aria-label', t('evidenceMap'));
}
function renderTitle() {
 const title = str(article?.title) || str(run.metadata?.title) || str(run.metadata?.topic) || str(run.contract?.question) || t('completeReport');
 setText('report-title', title);
 document.title = title;
 if (article) {
  setText('article-lead', article.lead);
  $('article-panel').hidden = false;
 } else if (!report) {
  setText('article-lead', t('noArticle'));
  $('reader-limits').hidden = false;
  setText('reader-limit-content', t('noArticle'));
 }
}
function evidenceBalanceMessage() {
 const question = str(run.contract?.question);
 if (!/(compare|comparison|versus|\bvs\.?\b|对比|比较|差异)/i.test(question)) return '';
 const verified = citedSources.filter(source => sourceRole(source) === 'evidence');
 if (!verified.length) return '';
 const counts = new Map();
 for (const source of verified) {
  try { const host = new URL(source.url).hostname.replace(/^www\./i, ''); if (host) counts.set(host, (counts.get(host) || 0) + 1); } catch (_) {}
 }
 if (!counts.size) return '';
 const [host, count] = [...counts.entries()].sort((a,b) => b[1] - a[1])[0];
 if (count / verified.length < .7) return '';
 return zh
  ? `已核实引文主要来自 ${host}。这是来源集中提示，不代表存在偏见；正文应说明其他一方或独立来源是否缺席。`
  : `Most verified citations come from ${host}. This flags source concentration, not bias; the article should say whether the other side or independent sources are missing.`;
}
function renderLimitations() {
 if (!research) return;
 const stop = t('stop', research.stop.reason);
 const gaps = arr(research.gaps);
 const balance = evidenceBalanceMessage();
 const needsCallout = research.stop.reason !== 'saturated' || gaps.length > 0 || Boolean(balance);
 $('reader-limits').hidden = !needsCallout;
 if (!needsCallout) return;
 const box = $('reader-limit-content'); box.replaceChildren();
 box.append(el('p', stop));
 if (research.stop.reason === 'saturated') box.append(el('p', t('stopNote_saturated'), 'muted'));
 if (balance) box.append(el('p', `${t('evidenceBalance')}: ${balance}`));
 if (gaps.length) {
  const list = el('ul', undefined, 'callout-list');
  for (const gap of gaps.slice(0, 4)) list.append(el('li', gap));
  if (gaps.length > 4) list.append(el('li', `+${gaps.length - 4}`));
  box.append(list);
 }
}
function renderArticle() {
 if (!article) return;
 const body = $('article-body'); body.replaceChildren();
 for (const section of arr(article.sections)) {
  const sectionNode = el('section', undefined, 'article-section');
  sectionNode.id = 'section-' + section.id;
  sectionNode.append(el('h2', section.heading));
  for (const block of arr(section.blocks)) {
   const content = el('div', undefined, 'article-block');
   content.dataset.role = block.role;
   const role = t('role', block.role);
   if (role) content.append(el('span', role, 'role'));
   if (block.kind === 'paragraph') content.append(el('p', block.text));
   else if (block.kind === 'list') { const list = el('ul'); for (const item of arr(block.items)) list.append(el('li', item)); content.append(list); }
   else if (block.kind === 'table') {
    const table = el('table'), head = el('thead'), header = el('tr'), rows = el('tbody');
    for (const item of arr(block.headers)) header.append(el('th', item)); head.append(header);
    for (const row of arr(block.rows)) { const tr = el('tr'); for (const cell of arr(row)) tr.append(el('td', cell)); rows.append(tr); }
    table.append(head, rows); const wrap = el('div', undefined, 'scroll'); wrap.append(table); content.append(wrap);
   }
   const cited = new Set();
   for (const claimID of arr(block.claim_ids)) for (const evidence of arr(claimByID.get(String(claimID))?.evidence)) cited.add(String(evidence.source_id));
   for (const conclusionID of arr(block.conclusion_ids)) for (const claimID of arr(conclusionByID.get(String(conclusionID))?.claim_ids)) for (const evidence of arr(claimByID.get(String(claimID))?.evidence)) cited.add(String(evidence.source_id));
   const numbers = [...cited].filter(id => sourceNumbers.has(id));
   if (numbers.length) {
    const citeList = el('span', undefined, 'cite-list');
    for (const id of numbers) {
     const number = sourceNumbers.get(id), link = el('a', `［${number}］`, 'cite');
     link.href = `#source-${number}`; link.title = `${t('citation')} ${number}`;
     link.addEventListener('click', event => { event.preventDefault(); openSource(id); history.replaceState(null, '', link.href); });
     citeList.append(link);
    }
    content.append(citeList);
   }
   sectionNode.append(content);
  }
  body.append(sectionNode);
 }
}
function sourceCard(source, number, compact = false) {
 const role = sourceRole(source), card = el('article', undefined, 'source-card');
 if (number) card.id = 'source-' + number;
 const head = el('div');
 if (number) head.append(el('span', `［${number}］ `, 'badge'));
 head.append(el('span', source.title || source.url || t('sourceNumber'), 'source-title'));
 head.append(el('span', t(role), `badge${role === 'candidate_lead' ? ' candidate' : ''}`));
 card.append(head);
 const meta = el('p', undefined, 'source-meta');
 if (source.type) meta.append(document.createTextNode(`${t('sourceType')}: ${source.type} · `));
 const selection = role === 'candidate_lead' && source.status === 'selected' ? t('selected_unverified') : t(source.status, source.status);
 meta.append(document.createTextNode(`${t(source.verification, source.verification)} · ${selection}`));
 card.append(meta);
 const href = safeURL(str(source.url));
 if (href) { const a = el('a', t('openSource')); a.href = href; a.target = '_blank'; a.rel = 'noopener noreferrer'; card.append(a); }
 const details = el('details', undefined, 'source-detail');
 details.append(el('summary', compact ? t('inspectedPassage') : t('usedFor')));
 if (source.content) { details.append(el('p', t('inspectedPassage'), 'source-meta')); details.append(el('blockquote', source.content, 'quote')); }
 if (source.published_at) details.append(el('p', `${t('published')}: ${source.published_at}`, 'source-meta'));
 if (source.retrieved_at) details.append(el('p', `${t('retrieved')}: ${source.retrieved_at}`, 'source-meta'));
 const upstream = arr(source.upstream_ids).map(id => sourceByID.get(String(id))).filter(Boolean);
 if (upstream.length) {
  details.append(el('p', t('upstream'), 'source-meta'));
  for (const origin of upstream) {
   const originLink = el('a', origin.title || origin.url); const originURL = safeURL(str(origin.url));
   if (originURL) { originLink.href = originURL; originLink.target = '_blank'; originLink.rel = 'noopener noreferrer'; }
   details.append(originLink, document.createElement('br'));
  }
 }
 const relevantClaims = claims.filter(claim => arr(claim.evidence).some(evidence => evidence.source_id === source.id));
 for (const claim of relevantClaims) {
  const block = el('div', undefined, 'record'); block.append(el('strong', claim.text));
  for (const evidence of arr(claim.evidence).filter(item => item.source_id === source.id)) {
   block.append(el('p', `${t(evidence.relation, evidence.relation)} · ${t(evidence.verification, evidence.verification)}`, 'source-meta'));
   block.append(el('blockquote', evidence.quote, 'quote'));
   if (evidence.locator) block.append(el('p', evidence.locator, 'source-meta'));
  }
  details.append(block);
 }
 if (!source.content && !relevantClaims.length) details.append(el('p', t('noItems'), 'empty-state'));
 card.append(details);
 return card;
}
function renderSources() {
 const list = $('source-list'); list.replaceChildren();
 if (!citedSources.length) list.append(el('p', t('noSources'), 'empty-state'));
 citedSources.forEach((source, index) => list.append(sourceCard(source, index + 1)));
 for (const role of ['candidate_lead', 'context_source']) {
  const target = $(role === 'candidate_lead' ? 'candidate-list' : 'context-list');
  target.replaceChildren();
  const items = sources.filter(source => sourceRole(source) === role && !citedSourceIDs.has(String(source.id)));
  if (!items.length) target.append(el('p', t('noItems'), 'empty-state'));
  for (const source of items) target.append(sourceCard(source, null, true));
 }
}
function renderCoverageAndAudit() {
 if (!research) return;
 const reason = t('stop', research.stop.reason);
 const rationale = str(research.stop.rationale);
 setText('stop-summary', `${reason}. ${rationale}${research.stop.reason === 'saturated' ? ' ' + t('stopNote_saturated') : ''}`);
 const coverage = $('coverage-summary'); coverage.replaceChildren();
 const rows = arr(research.coverage);
 if (!rows.length) coverage.append(el('p', t('noItems'), 'empty-state'));
 for (const row of rows) {
  const question = arr(run.contract?.questions).find(item => item.id === row.question_id)?.text || row.question_id;
  const section = el('section', undefined, 'record');
  section.append(el('strong', question), el('span', ' · ' + t('coverage', row.status), 'badge'));
  if (arr(row.gaps).length) { const list = el('ul', undefined, 'gap-list'); for (const gap of arr(row.gaps)) list.append(el('li', gap)); section.append(list); }
  coverage.append(section);
 }
 const gaps = $('research-gaps'); gaps.replaceChildren();
 if (!arr(research.gaps).length) gaps.append(el('p', t('noGaps'), 'empty-state'));
 else { const list = el('ul', undefined, 'gap-list'); for (const gap of arr(research.gaps)) list.append(el('li', gap)); gaps.append(list); }
 const process = $('process-summary'); process.replaceChildren();
 process.append(el('p', `${L.processSummary}${research.audience} · ${research.rounds.length} ${L.roundsCount} · ${sources.length} ${L.sourcesCount}`, 'muted'));
 const roundList = $('round-list'); roundList.replaceChildren();
 if (!arr(research.rounds).length) roundList.append(el('p', t('noItems'), 'empty-state'));
 for (const round of arr(research.rounds)) {
  const details = el('details', undefined, 'record');
  details.append(el('summary', `${t('round')} ${round.id}: ${round.search_angle} · ${t('gain', round.material_gain)} · ${t(round.status === 'failed' ? 'round_failed' : round.status)}`));
  const queries = arr(round.query_ids).map(id => arr(run.queries).find(query => query.id === id)).filter(Boolean);
  for (const query of queries) details.append(el('p', `${query.text} · ${L.queryTime}: ${query.at}`, 'source-meta'));
  for (const id of arr(round.assessed_source_ids)) {
   const source = sourceByID.get(String(id));
   if (source) details.append(el('p', `${source.title || source.url} · ${L.sourceTime}: ${source.retrieved_at}`, 'source-meta'));
  }
  if (round.materiality_reason) details.append(el('p', round.materiality_reason));
  if (arr(round.gaps).length) { const list = el('ul', undefined, 'gap-list'); for (const gap of round.gaps) list.append(el('li', gap)); details.append(list); }
  roundList.append(details);
 }
}

function graphModel() {
 const nodes = [], edges = [], nodeIDs = new Set(), edgeIDs = new Set();
 const addNode = (id, type, label, sourceID = '') => {
  if (!id || nodeIDs.has(id)) return;
  nodeIDs.add(id); nodes.push({ data: { id, type, label: label || t(type), sourceID } });
 };
 for (const source of citedSources) addNode(`source:${source.id}`, 'source', source.title || source.url, String(source.id));
 for (const id of articleClaimIDs) { const claim = claimByID.get(id); if (claim) addNode(`claim:${id}`, 'claim', claim.text); }
 for (const id of articleConclusionIDs) { const item = conclusionByID.get(id); if (item) addNode(`conclusion:${id}`, 'conclusion', item.text); }
 const addEdge = (from, to, type) => {
  const id = `${from}|${to}|${type}`;
  if (edgeIDs.has(id) || !nodeIDs.has(from) || !nodeIDs.has(to)) return;
  edgeIDs.add(id); edges.push({ data: { id: `edge-${edges.length}`, source: from, target: to, type, label: t('relation', type) } });
 };
 for (const id of articleClaimIDs) {
  const claim = claimByID.get(id); if (!claim) continue;
  for (const evidence of arr(claim.evidence)) addEdge(`source:${evidence.source_id}`, `claim:${id}`, evidence.relation);
 }
 for (const conclusionID of articleConclusionIDs) {
  const conclusion = conclusionByID.get(conclusionID); if (!conclusion) continue;
  for (const claimID of arr(conclusion.claim_ids)) if (articleClaimIDs.has(String(claimID))) addEdge(`claim:${claimID}`, `conclusion:${conclusionID}`, 'derives_from');
 }
 return { nodes, edges };
}
function renderGraph() {
 const model = graphModel();
 const details = $('evidence-map');
 if (!report || model.nodes.length === 0) { details.hidden = true; return; }
 const list = $('graph-list'); list.replaceChildren();
 const nodeLabel = new Map(model.nodes.map(node => [node.data.id, node.data.label]));
 for (const edge of model.edges) {
  const source = nodeLabel.get(edge.data.source), target = nodeLabel.get(edge.data.target);
  const button = el('button', `${source} — ${edge.data.label} → ${target}`, 'graph-item');
  button.type = 'button'; button.addEventListener('click', () => {
   const id = edge.data.source.startsWith('source:') ? edge.data.source : edge.data.target.startsWith('source:') ? edge.data.target : '';
   const sourceID = model.nodes.find(node => node.data.id === id)?.data.sourceID;
   if (sourceID) openSource(sourceID);
  });
  list.append(button);
 }
 for (const node of model.nodes) {
  const button = el('button', `${t(node.data.type)}: ${node.data.label}`, 'graph-item');
  button.type = 'button';
  if (node.data.sourceID) button.addEventListener('click', () => openSource(node.data.sourceID));
  else button.addEventListener('click', () => {
   const related = model.edges.find(edge => edge.data.source === node.data.id || edge.data.target === node.data.id);
   const sourceNode = related && [related.data.source, related.data.target].map(id => model.nodes.find(item => item.data.id === id)).find(item => item?.data.sourceID);
   if (sourceNode) openSource(sourceNode.data.sourceID);
  });
  list.append(button);
 }
 let cy;
 try {
  cy = cytoscape({ container: $('graph'), elements: model.nodes.concat(model.edges), style: [
   { selector: 'node', style: { 'background-color': '#347f76', 'label': 'data(label)', 'color': '#1f3032', 'font-size': 12, 'text-wrap': 'wrap', 'text-max-width': 180, 'text-valign': 'bottom', 'text-margin-y': 7, 'width': 34, 'height': 34 } },
   { selector: 'node[type="claim"]', style: { 'background-color': '#d3a64a', 'shape': 'round-rectangle', 'width': 'label', 'height': 'label', 'padding': 8 } },
   { selector: 'node[type="conclusion"]', style: { 'background-color': '#77927b', 'shape': 'hexagon', 'width': 'label', 'height': 'label', 'padding': 8 } },
   { selector: 'edge', style: { 'width': 2, 'line-color': '#7b8a86', 'target-arrow-color': '#7b8a86', 'target-arrow-shape': 'triangle', 'curve-style': 'bezier', 'label': 'data(label)', 'font-size': 10, 'text-background-color': '#fff', 'text-background-opacity': 1, 'text-background-padding': 3 } },
   { selector: 'edge[type="contradicts"]', style: { 'line-color': '#a6402e', 'target-arrow-color': '#a6402e', 'line-style': 'dashed' } },
  ], layout: { name: 'breadthfirst', directed: true, padding: 28, spacingFactor: 1.25 }, wheelSensitivity: .2 });
  cy.on('tap', 'node', event => { const data = event.target.data(); if (data.sourceID) openSource(data.sourceID); });
 } catch (error) { setText('graph-message', t('mapUnavailable')); }
 const fit = () => { if (cy) { cy.resize(); cy.layout({ name: 'breadthfirst', directed: true, padding: 28, spacingFactor: 1.25 }).run(); cy.fit(undefined, 28); } };
 details.addEventListener('toggle', () => { if (details.open) requestAnimationFrame(fit); });
}
function openSource(id) {
 const number = sourceNumbers.get(String(id));
 const card = number && $(`source-${number}`);
 if (card) {
  const disclosure = card.querySelector('details'); if (disclosure) disclosure.open = true;
  card.scrollIntoView({ behavior: 'smooth', block: 'start' });
  card.setAttribute('tabindex', '-1'); card.focus({ preventScroll: true });
 }
}
function renderLegacyAudit() {
 if (!report) {
  const candidate = $('candidate-list'); candidate.replaceChildren();
  for (const source of sources.filter(item => sourceRole(item) !== 'context_source')) candidate.append(sourceCard(source, null, true));
  const context = $('context-list'); context.replaceChildren();
  for (const source of sources.filter(item => sourceRole(item) === 'context_source')) context.append(sourceCard(source, null, true));
  $('evidence-map').hidden = true;
 }
}
function renderRaw() { $('raw').textContent = JSON.stringify(payload, null, 2); }

applyLocale();
renderTitle();
renderLimitations();
renderArticle();
renderSources();
if (research) renderCoverageAndAudit();
renderLegacyAudit();
renderGraph();
renderRaw();
})();
