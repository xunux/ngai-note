# Python 笔记

## 目录
- [Python 环境管理](#python环境管理)
  - [环境管理工具对比](#环境管理工具对比)
  - [使用 pyenv 进行多版本安装和管理](#使用-pyenv-进行多版本安装和管理)
    - [安装指定版本 python](#安装指定版本-python)
    - [切换 python 版本](#切换-python-版本)
    - [虚拟环境最佳实践](#虚拟环境最佳实践)
- [依赖管理](#依赖管理)
  - [典型工作流](#典型工作流)
  - [依赖管理工具对比](#依赖管理工具对比)
- [镜像配置技巧](#镜像配置技巧)
  - [国内镜像配置](#国内镜像配置)

## Python环境管理

### 环境管理工具对比

| 工具                | 适用场景                      | 推荐指数 |
|-------------------|---------------------------|--------|
| pyenv ★★★★★      | 多版本并行管理               | 必备工具 |
| pyenv-virtualenv ★★★★☆ | 虚拟环境管理（推荐组合）      | 生产推荐 |
| venv ★★★☆☆        | 轻量级虚拟环境（Python≥3.3） | 内置方案 |
| virtualenv ★★☆☆☆        | Python 3.3 以下版本 | 老项目使用 |

组合方案：
1. 使用pyenv管理基础Python版本
2. 使用pyenv-virtualenv创建项目专属虚拟环境
3. 使用venv作为轻量级替代方案

### 使用 pyenv 进行多版本安装和管理

#### 安装指定版本 python

pyenv 文档：[官方文档](https://github.com/pyenv/pyenv)

```bash
# 安装pyenv版本管理工具
curl https://pyenv.run | bash
# 查找Python版本
pyenv install --list | grep 3.10
# 使用指定Python版本
pyenv install 3.10.16
# 使用指定Python版本（自动读取.python-version）
pyenv install
```
如果想要安装的 python 版本 **3.13.3** 不在 `pyenv install --list` 列表中，您可以尝试以下方法：
1. 更新 pyenv ：
```bash
pyenv update
 ```
2. 手动安装 ：即使版本不在列表中，您也可以尝试直接指定版本号进行安装：
```bash
pyenv install 3.13.3
 ```

#### 切换 python 版本
```bash
# 查看当前安装的Python版本
pyenv versions
# 切换Python版本
pyenv global 3.13.3
```

#### 虚拟环境最佳实践

```bash
# 创建环境
pyenv virtualenv 3.13.3 my_project_env

# 激活环境
pyenv activate my_project_env

# 退出环境
pyenv deactivate

# 删除环境
pyenv uninstall my_project_env
```

🔔 内置方案（venv）:
```bash
python -m venv .venv  # 创建
source .venv/bin/activate  # 激活
```

1. 使用 pyenv-virtualenv 创建虚拟环境

pyenv-virtualenv 是一个 pyenv 插件，它提供了一个命令行工具，用于创建和管理虚拟环境。
文档：：https://github.com/pyenv/pyenv-virtualenv

`pyenv-virtualenv` uses `python -m venv` if it is available and the `virtualenv` command is not available.

```bash
pyenv version # 查看当前Python版本
pyenv virtualenv 3.13.3 new_venv # 创建虚拟环境
 ```

pyenv-virtualenv 常用命令：
```bash
pyenv virtualenvs # 列出所有虚拟环境
pyenv activate new_venv # 激活虚拟环境
pyenv deactivate # 退出虚拟环境
pyenv uninstall new_venv # 删除虚拟环境
 ```


2. 使用 venv 创建虚拟环境

```bash
pyenv global 3.13.3 # 切换到Python 3.13.3
python -m venv new_venv # 创建新的虚拟环境
source new_venv/bin/activate # 激活虚拟环境
pip install -r requirements.txt # 安装依赖
 ```

3. 使用 virtualenv 创建虚拟环境

virtualenv 是一个第三方的虚拟环境管理工具（支持Python 2/3）：
- 文档：[virtualenv官方文档](https://virtualenv.pypa.io/)
- 典型场景：需要兼容旧版Python（<3.3）的项目

```bash
# 安装virtualenv
pip install virtualenv
# 创建新的虚拟环境
virtualenv new_venv
# 激活新环境
source new_venv/bin/activate
 ```

## 依赖管理

### 典型工作流
1. 在激活的虚拟环境中操作
2. 生成依赖文件：
```bash
pip freeze > requirements.txt  # 生成精确版本依赖
# 或使用 pip-compile 生成分层依赖（需安装pip-tools）
```
3. 安装项目依赖：
```bash
pip install -r requirements.txt  # 安装全部依赖
pip install package==1.0.0       # 安装指定版本
```

### 依赖管理工具对比
| 方式                | 优点                      | 缺点                 |
|---------------------|--------------------------|----------------------|
| requirements.txt    | 简单直观                  | 无版本冲突解决机制    |
| pipenv              | 自动管理依赖锁            | 需要额外安装          |
| poetry              | 全生命周期管理            | 学习曲线较高          |
```bash
pip install -r requirements.txt
```

## 镜像配置技巧

<details>
<summary>📌 国内镜像配置（点击展开）</summary>

```bash
# 临时使用
pip install -i https://pypi.tuna.tsinghua.edu.cn/simple package

# 永久配置
mkdir -p ~/pip && cat > ~/pip/pip.ini << EOF
[global]
index-url = https://pypi.tuna.tsinghua.edu.cn/simple
[install]
trusted-host = pypi.tuna.tsinghua.edu.cn
EOF
```

支持镜像源：
- 清华：`https://pypi.tuna.tsinghua.edu.cn/simple`
- 阿里云：`https://mirrors.aliyun.com/pypi/simple/`
</details>
