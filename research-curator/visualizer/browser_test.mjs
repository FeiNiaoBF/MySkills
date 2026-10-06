// Reader-first/offline browser check at the published HTML boundary. Node >=22; no npm packages.
// CHROME=<chromium executable> node visualizer/browser_test.mjs <CLI executable> <report.json>
import {spawn, spawnSync} from 'node:child_process';
import {mkdtemp, mkdir, readFile, writeFile, rename} from 'node:fs/promises';
import {tmpdir} from 'node:os';
import path from 'node:path';
import {pathToFileURL} from 'node:url';
import assert from 'node:assert/strict';
const [cli, input] = process.argv.slice(2);
if (!cli || !input || !process.env.CHROME) throw Error('Set CHROME and pass CLI executable and report JSON');
const root = path.join(tmpdir(), 'pi-agent'); await mkdir(root, {recursive: true});
const work = await mkdtemp(path.join(root, 'curator-reader-'));
const fixture = JSON.parse(await readFile(input, 'utf8'));
fixture.article.language = 'zh-CN';
const section = fixture.article.sections[0];
const claim = fixture.run.claims[0];
const firstArticleClaim = fixture.run.claims.find(item => item.id === section.blocks.find(block => block.claim_ids?.length)?.claim_ids[0]);
if (!firstArticleClaim?.evidence?.length) throw Error('first article evidence block has no source evidence');
const firstArticleSourceID = firstArticleClaim.evidence[0].source_id;
const articleBlocks = fixture.article.sections.flatMap(articleSection => articleSection.blocks);
const citedClaimIDs = new Set(articleBlocks.flatMap(block => block.claim_ids ?? []));
for (const conclusionID of articleBlocks.flatMap(block => block.conclusion_ids ?? [])) {
 const conclusion = fixture.run.conclusions.find(item => item.id === conclusionID);
 for (const claimID of conclusion?.claim_ids ?? []) citedClaimIDs.add(claimID);
}
const citedSourceIDs = new Set();
for (const claimID of citedClaimIDs) {
 const claim = fixture.run.claims.find(item => item.id === claimID);
 for (const evidence of claim?.evidence ?? []) citedSourceIDs.add(evidence.source_id);
}
const citedSources = [...citedSourceIDs].map(id => fixture.run.sources.find(source => source.id === id)).filter(Boolean);
const firstArticleSourceNumber = citedSources.findIndex(source => source.id === firstArticleSourceID) + 1;
if (!firstArticleSourceNumber) throw Error('first article citation source is absent from cited sources');
const firstArticleSourceSelector = `#source-${firstArticleSourceNumber}`;
const firstArticleEvidence = firstArticleClaim.evidence.find(item => item.source_id === firstArticleSourceID);
section.blocks.push({id:'browser-citation',kind:'paragraph',role:'evidence',text:'Browser citation check.',claim_ids:[claim.id],conclusion_ids:[]});
section.blocks.push({id:'hostile',kind:'paragraph',role:'context',text:'</script><img src=https://example.org/x onerror=alert(1)>',claim_ids:[],conclusion_ids:[]});
const fixturePath = path.join(work, 'report.json'); await writeFile(fixturePath, JSON.stringify(fixture));
const output = path.join(work, 'topic');
const pub = spawnSync(cli, ['publish','-in',fixturePath,'-out',output], {encoding:'utf8'});
assert.equal(pub.status, 0, pub.stderr);
const moved = path.join(work, 'moved'); await rename(output, moved);
const child = spawn(process.env.CHROME, ['--headless=new','--disable-gpu','--disable-background-networking','--no-first-run','--remote-debugging-port=0','--user-data-dir='+path.join(work,'profile'),'about:blank'], {stdio:'ignore'});
let ws; const pending = new Map(); let nextID = 0; const external = [], errors = [];
const delay = ms => new Promise(resolve => setTimeout(resolve, ms));
try {
 let port;
 for (let i=0;i<100;i++) { try { port=Number((await readFile(path.join(work,'profile','DevToolsActivePort'),'utf8')).split('\n')[0]); break; } catch { await delay(100); } }
 assert.ok(port, 'Chrome did not expose CDP');
 const targets = await (await fetch(`http://127.0.0.1:${port}/json`)).json();
 ws = new WebSocket(targets.find(target=>target.type==='page').webSocketDebuggerUrl);
 await new Promise((resolve,reject)=>{ws.addEventListener('open',resolve,{once:true});ws.addEventListener('error',reject,{once:true});});
 ws.addEventListener('message',({data})=>{const message=JSON.parse(data);if(message.method==='Network.requestWillBeSent'&&/^https?:/.test(message.params.request.url))external.push(message.params.request.url);if(message.method==='Runtime.exceptionThrown')errors.push(message.params.exceptionDetails);if(message.id){const item=pending.get(message.id);pending.delete(message.id);if(item)message.error?item.reject(Error(message.error.message)):item.resolve(message.result);}});
 const send=(method,params={})=>new Promise((resolve,reject)=>{const id=++nextID;pending.set(id,{resolve,reject});ws.send(JSON.stringify({id,method,params}));});
 const evaluate=async expression=>{const result=await send('Runtime.evaluate',{expression,returnByValue:true,awaitPromise:true});if(result.exceptionDetails)throw Error(JSON.stringify(result.exceptionDetails));return result.result.value;};
 await send('Runtime.enable'); await send('Page.enable'); await send('Network.enable');
 const load=async url=>{await send('Page.navigate',{url});for(let i=0;i<100;i++){if(await evaluate(`location.href===${JSON.stringify(url)}&&document.readyState==='complete'&&document.querySelectorAll('.cite').length>0`))return;await delay(100);}throw Error('report did not render');};
 await load(pathToFileURL(path.join(moved,'index.html')).href);
 assert.equal(await evaluate('document.documentElement.lang'), 'zh-CN', 'UI did not follow article.language');
 assert.deepEqual(await evaluate(`['article-panel','evidence-register','evidence-map','audit-panel'].map(id=>Array.from(document.querySelector('main').children).findIndex(node=>node.id===id))`), [0,1,2,3], 'reader-first section order changed');
 assert.equal(await evaluate(`document.getElementById('evidence-map').open`), false, 'evidence map should start collapsed');
 assert.equal(await evaluate(`document.querySelectorAll('img').length`), 0, 'hostile text created markup');
 await evaluate(`document.querySelector('.cite').click()`);
 assert.equal(await evaluate(`document.querySelector(${JSON.stringify(firstArticleSourceSelector + ' details')}).open`), true, 'citation did not open its source passage');
 const sourceDetail = await evaluate(`document.querySelector(${JSON.stringify(firstArticleSourceSelector)}).textContent`);
 assert.ok(sourceDetail.includes(firstArticleEvidence.locator), 'source locator missing');
 assert.ok(sourceDetail.includes(firstArticleEvidence.quote), 'exact source passage missing');
 await evaluate(`document.getElementById('evidence-map').open=true`); await delay(200);
 const graphText = await evaluate(`document.querySelector('#graph-list').textContent`);
 assert.ok(graphText.includes(fixture.run.sources[0].title) && graphText.includes(fixture.run.claims[0].text), 'map omits meaningful source/claim labels');
 const graphLabels = await evaluate(`Array.from(document.querySelectorAll('#graph-list button')).map(button=>button.textContent.trim())`);
 assert.ok(!graphLabels.some(label=>label===fixture.run.sources[0].id||label===claim.id), 'internal IDs shown as graph labels');
 await send('Emulation.setDeviceMetricsOverride',{width:390,height:844,deviceScaleFactor:1,mobile:true}); await delay(250);
 assert.ok(await evaluate('document.documentElement.scrollWidth<=document.documentElement.clientWidth'),'mobile overflow');
 await send('Emulation.setEmulatedMedia',{media:'print'});
 assert.notEqual(await evaluate(`getComputedStyle(document.querySelector('.cite')).display`),'none','print citation hidden');
 await send('Emulation.setEmulatedMedia',{media:''});
 const html = await readFile(path.join(moved,'index.html'),'utf8');
 const missingGraph = path.join(work,'missing-graph.html');
 await writeFile(missingGraph, html.replace(/<script>\/\* CYTOSCAPE_VENDOR \*\/[\s\S]*?<\/script>/,'<script>/* simulated missing graph */</script>'));
 await send('Page.navigate',{url:pathToFileURL(missingGraph).href}); await delay(300);
 assert.ok(await evaluate(`document.getElementById('evidence-register').textContent.includes(${JSON.stringify(fixture.run.sources[0].title)})`),'article sources disappeared when graph was unavailable');
 assert.deepEqual(external, [], 'report made runtime network requests');
 assert.deepEqual(errors, [], 'uncaught page exceptions');
 console.log(JSON.stringify({result:'PASS',work,checks:['Chinese UI follows article.language','article and evidence precede collapsed graph/audit','citation opens exact source passage','meaningful graph labels hide internal IDs','hostile text inert','mobile layout','print citations','source access without Cytoscape','zero network requests']}));
} finally { ws?.close(); child.kill(); }
