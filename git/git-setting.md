# 设置 ssh

```
ssh-keygen 

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
git config --global user.name "lois"
git confgi --global user.email "loissiau@gmail.com"
```

参考地址：
1. [如何用 SSH 密钥在一台机器上管理多个 GitHub 账户](https://www.freecodecamp.org/chinese/news/manage-multiple-github-accounts-the-ssh-way/)
2. 
