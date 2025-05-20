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

对应配置在 ~/.gitconfig 文件中

```
[user]
	name = xx
	email = xx
```

## 内部和外部项目中使用不同的 Git 用户名和邮箱
为了在内部和外部项目中使用不同的 Git 用户名和邮箱，可以通过 **条件配置** 实现自动化切换。以下是具体步骤：

---

### 方法 1：使用 Git 的 `includeIf` 条件配置（推荐）
#### 步骤 1：创建专用目录
- **内部项目**：统一存放在 `~/work/` 目录下。
- **外部项目**：统一存放在 `~/personal/` 目录下。

#### 步骤 2：创建配置文件
1. **内部项目配置**：  
   创建文件 `~/.gitconfig-work`，内容如下：
   ```ini
   [user]
       name = 内部用户名
       email = 内部邮箱@公司.com
   [http]  
        sslVerify = false
   ```

2. **外部项目配置**：  
   创建文件 `~/.gitconfig-personal`，内容如下：
   ```ini
   [user]
       name = 外部用户名
       email = 个人邮箱@example.com
   ```

#### 步骤 3：修改全局配置
编辑 `~/.gitconfig`，添加条件包含规则：
```ini
[includeIf "gitdir:~/work/**"]  # 所有位于 ~/work/ 下的仓库
    path = .gitconfig-work

[includeIf "gitdir:~/personal/**"]  # 所有位于 ~/personal/ 下的仓库
    path = .gitconfig-personal
```

---

### 方法 2：手动为每个仓库设置（适合临时项目）
进入仓库根目录，执行：
```bash
# 内部项目
git config user.name "内部用户名"
git config user.email "内部邮箱@公司.com"

# 外部项目
git config user.name "外部用户名"
git config user.email "个人邮箱@example.com"
```

---

### 验证配置
在仓库中运行以下命令，确认配置生效：
```bash
git config user.name  # 查看当前用户名
git config user.email # 查看当前邮箱
```

---

### 注意事项
1. **路径匹配**：确保 `gitdir:` 后的路径与实际仓库路径匹配（Windows 用户需注意路径格式，如 `gitdir:C:/Users/username/work/`）。
2. **Git 版本**：`includeIf` 需要 Git 2.13 或更高版本，可通过 `git --version` 检查。
3. **配置文件权限**：确保创建的 `.gitconfig-work` 和 `.gitconfig-personal` 文件有读取权限。

通过以上步骤，Git 会根据仓库路径自动切换用户配置，无需手动干预。

## git 代理设置

**建议使用tun全局代理模式，省得配置代理**

vim ~/.ssh/config
```
Host github.com
   HostName ssh.github.com
   Port 443
   User git
   IdentityFile ~/.ssh/github_id_rsa
   # windows
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
```

git gui 启动报错

vim ~/.gitconfig
```
[gui]

spellingdictionary = none
	
```


参考地址：
1. [如何用 SSH 密钥在一台机器上管理多个 GitHub 账户](https://www.freecodecamp.org/chinese/news/manage-multiple-github-accounts-the-ssh-way/)
2. [ssh_exchange_identification: Connection closed by remote host under Git bash [closed]](https://stackoverflow.com/a/60994276)
