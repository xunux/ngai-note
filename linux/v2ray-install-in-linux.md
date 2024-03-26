# Linux Ubuntu 上科学上网

直奔主题， v2rayA 安装了，但依赖的 v2ray 内核安装不起来，后面发现 clash meta 的继任者 clash-verge 的后继者 clash-verge-rev 好用，
clash-verge-rev 在 linux windows mac 都可以用， android 下则继续用 clash meta

下载地址：https://clash-verge-rev.github.io/install.html，选择 AppImage或 deb 格式的下载

备用地址：

软件下载地址

AMD平台(X86架构) 
https://app.chongjin.icu/Linux/Clash_Verge_REV/clash-verge_amd64.AppImage 
https://app.chongjin.icu/Linux/Clash_Verge_REV/clash-verge_amd64.deb

ARM 
https://app.chongjin.icu/Linux/Clash_Verge_REV/clash-verge_arm64.deb



第1步 复制订阅地址， 然后打开安装的软件
第2步点击 配置 将 订阅地址 粘贴到 输入框 点击 导入
第3步 设置 关闭 IPV6 设置 勾选 系统代理

## AppImage 文件添加快捷方式

```
sudo add-apt-repository ppa:appimagelauncher-team/stable
sudo apt update
sudo apt install appimagelauncher
```

参考：
[如何在 Linux 中使用 AppImage [完整指南]](https://cn.linux-console.net/?p=19002)
