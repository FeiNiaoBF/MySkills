# 只读基线命令

从 [network-diagnostics](../SKILL.md) 拆出的按需参考；进入 Baseline 阶段时读。

客户端基线：

```powershell
curl.exe -4 --max-time 15 https://api.ipify.org
curl.exe -6 --max-time 15 https://api64.ipify.org
nslookup baidu.com <ROUTER_IP>
nslookup api.ipify.org <ROUTER_IP>
curl.exe -4 -I --max-time 15 https://api.anthropic.com
```

路由器只读基线：

```sh
uci show passwall.@global[0]
uci show passwall.myshunt
uci show passwall.AIGC
ps w
netstat -lntup
free
swapon -s
```
