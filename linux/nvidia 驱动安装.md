# Nvidia 驱动安装

方法1：最快速方法

```
sudo ubuntu-drivers devices | grep recommended
```

显示结果
```
driver   : nvidia-driver-535 - distro non-free recommended
```

然后执行安装命令 `sudo apt-get install nvidia-driver-535`


方法二：通过系统自带 “软件和更新” 软件更新
