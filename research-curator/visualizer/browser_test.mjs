// Real-browser regression at the published HTML boundary. Node >=22; no npm packages.
// CHROME=<chromium executable> node visualizer/browser_test.mjs <CLI executable> <report.json>
import {spawn, spawnSync} from 'node:child_process';
import {mkdtemp, mkdir, readFile, writeFile, rename} from 'node:fs/promises';
import {tmpdir} from 'node:os';
import path from 'node:path';
import {pathToFileURL} from 'node:url';
import assert from 'node:assert/strict';
const [cli,input]=process.argv.slice(2);
if(!cli||!input||!process.env.CHROME) throw Error('Set CHROME and pass CLI executable and report JSON');
const tempRoot=path.join(tmpdir(),'pi-agent');await mkdir(tempRoot,{recursive:true});
const work=await mkdtemp(path.join(tempRoot,'curator-browser-'));
const fixture=JSON.parse(await readFile(input,'utf8'));
const original=fixture.run.sources[0];
const firstClaim=fixture.run.claims[0];
fixture.article.sections[0].blocks.push({id:'browser-claim',kind:'paragraph',role:'evidence',text:'Browser regression citation',claim_ids:[firstClaim.id],conclusion_ids:[]});
fixture.article.sections[0].blocks.push({id:'hostile',kind:'paragraph',role:'context',text:'</script><img src=https://example.org/x onerror=alert(1)>',claim_ids:[],conclusion_ids:[]});
const fixturePath=path.join(work,'input.json');await writeFile(fixturePath,JSON.stringify(fixture));
const output=path.join(work,'topic');const pub=spawnSync(cli,['publish','-in',fixturePath,'-out',output],{encoding:'utf8'});assert.equal(pub.status,0,pub.stderr);
const moved=path.join(work,'moved');await rename(output,moved);
const child=spawn(process.env.CHROME,['--headless=new','--disable-gpu','--disable-background-networking','--no-first-run','--remote-debugging-port=0','--user-data-dir='+path.join(work,'profile'),'about:blank'],{stdio:'ignore'});
let ws;const pending=new Map();let nextID=0;const external=[],errors=[];
const delay=ms=>new Promise(resolve=>setTimeout(resolve,ms));
try {
 let port;
 for(let i=0;i<100;i++){try {port=Number((await readFile(path.join(work,'profile','DevToolsActivePort'),'utf8')).split('\n')[0]);break;}catch{await delay(100)}}
 assert.ok(port,'Chrome did not expose CDP');
 const targets=await(await fetch(`http://127.0.0.1:${port}/json`)).json();
 ws=new WebSocket(targets.find(t=>t.type==='page').webSocketDebuggerUrl);
 await new Promise((resolve,reject)=>{ws.addEventListener('open',resolve,{once:true});ws.addEventListener('error',reject,{once:true})});
 ws.addEventListener('message',({data})=>{const m=JSON.parse(data);if(m.method==='Network.requestWillBeSent'&&/^https?:/.test(m.params.request.url))external.push(m.params.request.url);if(m.method==='Runtime.exceptionThrown')errors.push(m.params.exceptionDetails);if(m.id){const p=pending.get(m.id);pending.delete(m.id);if(p)m.error?p.reject(Error(m.error.message)):p.resolve(m.result)}});
 const send=(method,params={})=>new Promise((resolve,reject)=>{const id=++nextID;pending.set(id,{resolve,reject});ws.send(JSON.stringify({id,method,params}))});
 const evaluate=async expression=>{const r=await send('Runtime.evaluate',{expression,returnByValue:true,awaitPromise:true});if(r.exceptionDetails)throw Error(JSON.stringify(r.exceptionDetails));return r.result.value};
 await send('Runtime.enable');await send('Page.enable');await send('Network.enable');
 const load=async url=>{await send('Page.navigate',{url});for(let i=0;i<100;i++){if(await evaluate(`location.href===${JSON.stringify(url)}&&document.readyState==='complete'&&document.querySelectorAll('.citation').length>0`))return;await delay(100)}throw Error('report did not render')};
 await load(pathToFileURL(path.join(moved,'index.html')).href);
 assert.equal(await evaluate(`document.querySelectorAll('img').length`),0,'hostile text created markup');
 await evaluate(`document.querySelector('#type').value='source';document.querySelector('#type').dispatchEvent(new Event('change'));document.querySelector('[data-claim-id=${JSON.stringify(firstClaim.id)}]').click()`);
 const selected=await evaluate(`document.getElementById('graph')._cyreg.cy.$(':selected').map(n=>n.id())`);
 assert.ok(selected.includes(firstClaim.id)&&selected.includes(original.id),'citation under active filter did not select claim and source: '+JSON.stringify(selected));
 await evaluate(`[...document.querySelectorAll('#graph-list button')].find(b=>b.textContent.startsWith('source:')).click()`);
 const detail=await evaluate(`document.querySelector('#detail-content').textContent`);
 assert.ok(detail.includes(firstClaim.evidence[0].locator),'source detail omits evidence locator');
 assert.ok(detail.includes('Used in the article'),'source backlink missing');
 await evaluate(`document.querySelector('.article-mention').click()`);
 assert.match(await evaluate(`document.activeElement.id`),/^article-block-/);
 await evaluate(`[...document.querySelectorAll('#graph-list button')].find(b=>b.textContent.startsWith('goal:')).click()`);
 assert.ok((await evaluate(`document.querySelector('#detail-content').textContent`)).includes('No article reference'),'uncited node has no explicit empty state');
 await send('Emulation.setDeviceMetricsOverride',{width:390,height:844,deviceScaleFactor:1,mobile:true});await delay(400);
 assert.ok(await evaluate(`document.documentElement.scrollWidth<=document.documentElement.clientWidth`),'mobile overflow');
 await send('Emulation.setEmulatedMedia',{media:'print'});
 assert.notEqual(await evaluate(`getComputedStyle(document.querySelector('.citation')).display`),'none','printed citation hidden');
 await send('Emulation.setEmulatedMedia',{media:''});
 const html=await readFile(path.join(moved,'index.html'),'utf8');
 const fallback=path.join(work,'fallback.html');await writeFile(fallback,html.replace(/<script>\/\* CYTOSCAPE_VENDOR \*\/[\s\S]*?<\/script>/,'<script>/* simulated missing graph asset */</script>'));
 await load(pathToFileURL(fallback).href);
 await evaluate(`document.querySelector('.citation').click()`);
 assert.ok((await evaluate(`document.querySelector('#detail-content').textContent`)).includes(firstClaim.evidence[0].quote),'fallback citation lost evidence');
 assert.ok((await evaluate(`document.querySelector('#warnings').textContent`)).includes('Graph unavailable'));
 const legacyInput=path.join(work,'legacy.json'),legacyPage=path.join(work,'legacy.html');
 await writeFile(legacyInput,JSON.stringify({...fixture.run,run:{unexpected:'legacy extension'}}));
 // CLI intentionally validates v1; use a tiny Go caller to exercise the renderer's
 // documented permissive legacy seam without changing the production validator.
 const caller=path.join(work,'legacy.go');await writeFile(caller,'package main\nimport("os";"researchcurator/visualizer")\nfunc main(){b,e:=os.ReadFile(os.Args[1]);if e!=nil{panic(e)};h,e:=visualizer.Render(b);if e!=nil{panic(e)};if e=os.WriteFile(os.Args[2],h,0600);e!=nil{panic(e)}}');
 const render=spawnSync('go',['run',caller,legacyInput,legacyPage],{encoding:'utf8'});assert.equal(render.status,0,render.stderr);
 await send('Page.navigate',{url:pathToFileURL(legacyPage).href});
 for(let i=0;i<100;i++){if(await evaluate(`document.querySelector('#graph-list')?.children.length>0`))break;await delay(100)}
 assert.ok(await evaluate(`document.querySelector('#graph-list').children.length>0`),'legacy unknown run field broke graph');
 assert.deepEqual(external,[],'network requests (captured from before navigation)');assert.deepEqual(errors,[],'page exceptions');
 console.log(JSON.stringify({result:'PASS',work,checks:['legacy unknown-field compatibility','moved file','hostile text','filtered citation selects sources','source locator','backlink','uncited node','mobile','print citation','graph fallback','zero network requests']}));
} finally {ws?.close();child.kill();}
