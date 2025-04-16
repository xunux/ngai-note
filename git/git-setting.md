# 设置 ssh

```
ssh-keygen -t rsa -f github_id_rsa -C "xx@domain.com"

```

vim ~/.ssh/config
```
Host github.com
   HostName github.com
   User git
   IdentityFile ~/.ssh/github_id_rsa
```

验证
```
ssh -T git@github.com
```

## git 全局配置

```
git config --global user.name "xx"
git config --global user.email "xx@domain.com"
```

## git 代理设置

vim ~/.ssh/config
```
Host github.com
   HostName ssh.github.com
   Port 443
   User git
   IdentityFile ~/.ssh/github_id_rsa
   ProxyCommand connect -S 127.0.0.1:1080 %h %p
   # mac
   #ProxyCommand nc -X 5 -x 127.0.0.1:1080 %h %p
   
```

其中 127.0.0.1:1080 是 sock5 代理，若是 http 代理，则改为 `ProxyCommand connect -H 127.0.0.1:1081 %h %p`，其中 `127.0.0.1:1081` 是 http 代理


设置 git clone http 代理

```
git config --global http.https://github.com.proxy socks5://127.0.0.1:10808
```
或
vim ~/.gitconfig
```
[http "https://github.com"]
	proxy = socks5://127.0.0.1:10808


git gui 启动报错

vim ~/.gitconfig
```
[gui]

spellingdictionary = none
	
```

```

参考地址：
1. [如何用 SSH 密钥在一台机器上管理多个 GitHub 账户](https://www.freecodecamp.org/chinese/news/manage-multiple-github-accounts-the-ssh-way/)
2. [ssh_exchange_identification: Connection closed by remote host under Git bash [closed]](https://stackoverflow.com/a/60994276)
