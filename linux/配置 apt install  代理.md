# 配置 apt install 代理

基于已经安装好 Clash Verge，并启动情况下

sudo vim /etc/apt/apt.conf.d/proxy.conf 
```
Acquire::http::Proxy "http://127.0.0.1:7897/";
```

参考：

