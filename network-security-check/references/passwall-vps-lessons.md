# PassWall VPS Lessons

This reference records the safe procedure learned from the AWS Tokyo + PassWall + Claude troubleshooting session.

## Official facts used

- sing-box VLESS inbound requires `users.uuid`; `flow` supports `xtls-rprx-vision`; TLS is configured separately. Source: https://sing-box.sagernet.org/configuration/inbound/vless/
- sing-box Reality uses server-side `private_key`, client-side `public_key`, and `short_id`; direction matters. Source: https://sing-box.sagernet.org/configuration/shared/tls/
- AWS EC2 security groups act as virtual firewalls; inbound rules control traffic reaching the instance. Source: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-security-groups.html
- dnsmasq forwards DNS to recursive upstream servers and caches responses; negative/stale client behavior matters after DNS restart. Source: https://thekelleys.org.uk/dnsmasq/docs/dnsmasq-man.html
- Cloudflare bot scoring uses request features, browser/session signals, heuristics, and ML; curl challenges do not always equal a bad route. Source: https://developers.cloudflare.com/bots/concepts/bot-score/

## Correct AWS VPS validation order

Before installing a proxy service, test the VPS IP from the server itself:

```sh
curl -4 https://api.ipify.org
curl -4 -I https://api.anthropic.com
curl -4 -I https://claude.ai
```

Interpretation:

```text
api.anthropic.com 404/405: API edge reachable.
claude.ai 403 + cf-mitigated: challenge: curl/browser-fingerprint challenge, not automatically unusable.
Real Chrome over SSH SOCKS can login/reply for 5 minutes: IP is likely usable for Claude Web/Desktop.
```

Test real browser without changing router:

```powershell
ssh -i "$env:USERPROFILE\.ssh\<your-key>.pem" -N -D 127.0.0.1:10808 ubuntu@<AWS_VPS_IP>
& "C:\Program Files\Google\Chrome\Application\chrome.exe" --user-data-dir="$env:TEMP\chrome-aws-test" --proxy-server="socks5://127.0.0.1:10808" https://claude.ai
```

## Correct sing-box Reality shape

Server side:

```json
{
  "type": "vless",
  "listen": "0.0.0.0",
  "listen_port": 443,
  "users": [{ "uuid": "...", "flow": "xtls-rprx-vision" }],
  "tls": {
    "enabled": true,
    "server_name": "www.microsoft.com",
    "reality": {
      "enabled": true,
      "handshake": { "server": "www.microsoft.com", "server_port": 443 },
      "private_key": "SERVER_PRIVATE_KEY",
      "short_id": ["SHORT_ID"]
    }
  }
}
```

Client side:

```json
{
  "type": "vless",
  "server": "<AWS_VPS_IP>",
  "server_port": 443,
  "uuid": "...",
  "flow": "xtls-rprx-vision",
  "network": "tcp",
  "tls": {
    "enabled": true,
    "server_name": "www.microsoft.com",
    "utls": { "enabled": true, "fingerprint": "chrome" },
    "reality": {
      "enabled": true,
      "public_key": "SERVER_PUBLIC_KEY",
      "short_id": "SHORT_ID"
    }
  }
}
```

## Mandatory Direct rule for self-hosted node

When a router uses a proxy node, the IP address of that proxy server must be direct. Otherwise the router may connect to the new node through the old active node, causing nested proxying and Reality failures.

For this user's AWS VPS, ensure Direct IP list includes:

```text
<AWS_VPS_IP>/32
```

Symptom when missing:

```text
router temporary sing-box client: reality verification failed
AWS service log: REALITY: processed invalid connection
source IP in AWS log may be old proxy exit, e.g. 185.220.239.50
```

Minimal validation after adding Direct IP:

```sh
# Use a temporary local sing-box client, do not change PassWall tcp_node.
curl -4 -sS --socks5-hostname 127.0.0.1:20808 --connect-timeout 10 --max-time 20 https://api.ipify.org
curl -4 -sS -I --socks5-hostname 127.0.0.1:20808 --connect-timeout 10 --max-time 20 https://api.anthropic.com
```

Success:

```text
api.ipify.org -> <AWS_VPS_IP>
api.anthropic.com -> HTTP 404/405
```

## DNS failure runbook

If foreign domains fail after PassWall changes:

1. Check generated chinadns-ng config:

```sh
cat /tmp/etc/passwall/acl/default/chinadns_ng.conf
```

Bad pattern:

```text
trust-dns 127.0.0.1#15353
```

and no listener on 15353:

```sh
netstat -lntup | grep 15353
```

Fix for this user's stable setup:

```sh
uci set passwall.@global[0].dns_mode='tcp'
uci commit passwall
/etc/init.d/passwall stop
/etc/init.d/passwall start
```

Expected generated config:

```text
trust-dns tcp://1.1.1.1#53
group-upstream tcp://1.1.1.1#53
```

2. If duplicate `dnsmasq_default` exists after repeated restart, do clean stop/start:

```sh
/etc/init.d/passwall stop
ps w
netstat -lntup
/etc/init.d/passwall start
```

3. Flush Windows negative DNS cache:

```powershell
ipconfig /flushdns
```

## Known bad change from the incident

Do not repeat this pattern:

```sh
uci set passwall.myshunt.AIGC='LIdDZJG3'
uci set passwall.@global[0].tcp_node='myshunt'
# plus DNS mode accidentally xray
```

Why it broke:

- It changed the global traffic path instead of only testing Claude.
- The DNS mode generated a trust DNS upstream on a non-listening local port.
- Windows then cached DNS failures, so `curl` kept failing until `ipconfig /flushdns`.

## Safer Claude routing plan

Do not implement automatically. Ask first.

1. Keep global TCP node as the known-good node.
2. Keep AWS node imported and validated.
3. Keep AWS VPS IP in Direct list.
4. Use PassWall UI only after identifying the exact feature in this build: access control, shunt UI, or domain rule binding. Do not infer from another PassWall version.
5. Make one change and validate with Claude domains only.
