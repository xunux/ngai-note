# GIT 高级用法

## 1. 忽略已追踪文件的更新
```
git update-index --assume-unchanged <file>
```
## 2. 取消忽略文件的更新
```
git update-index --no-assume-unchanged <file>
```
## 3. 查看忽略文件的更新
```
git ls-files -v
```
