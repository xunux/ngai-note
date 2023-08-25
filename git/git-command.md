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


参考地址：
1. 
