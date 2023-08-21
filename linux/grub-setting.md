# grub 设置

编辑 grub 配置文件 `vim /etc/default/grub`

```
#记住上次选择的开机项, 与 GRUB_DEFAULT 设为 saved 配合使用
GRUB_SAVEDEFAULT=true
#开机启动默认项，默认为 0
```

然后执行 `update-grub` 命令，最后重启
