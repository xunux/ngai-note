# git 命令

1. 跟踪所有远程分支

在`bash`下执行如下命令 track 所有分支
```
git branch -r | grep -v '\->' | while read remote; do git branch --track "${remote#origin/}" "$remote"; done
```

然后获取更新
```
git fetch --all
```

最后本地代码与远程保持同步
```
git pull --all
```

合起来的命令如下
```
git branch -r | grep -v '\->' | while read remote; do git branch --track "${remote#origin/}" "$remote"; done
git fetch --all
git pull --all

```


git 修改 commit author

修改最近 5 個 commit 的 author 信息
```
git reabase -i HEAD~5
```

然後將要修改 author 的 commit 記錄由 pick 狀態改爲 edit，然後保存退出。

接着重複執行如下命令：

```
git commit --amend --author="user name <uer email>" --no-edit
git rebase --continue
```

如果要修改所有的 commit，包括 第一個 commit 則使用以下命令將第一個 commit 設爲 edit 狀態

```
git rebase -i --root
```



参考地址：
1. 
