# 笔记

## Ubuntu 现实问题调整 linux command for desktop display settings

```
gsettings set org.gnome.mutter experimental-features "['x11-randr-fractional-scaling']"
gsettings set org.gnome.mutter experimental-features "['scale-monitor-framebuffer']"
gsettings reset org.gnome.mutter experimental-features
```
## nvidia 驱动安装

参考: https://linuxconfig.org/how-to-install-the-nvidia-drivers-on-ubuntu-22-04
方法一：
```
sudo ubuntu-drivers autoinstall
```
方法二: 手工安装
```
#确定显卡型号并到官网下载驱动
ubuntu-drivers devices

https://www.nvidia.cn/Download/index.aspx?lang=cn

#安装
sudo apt install build-essential libglvnd-dev pkg-config
#
```

验证：
```
lshw -c display | grep driver
```

## 高分辨率屏幕下火狐显示问题 firefox about:config settings

```
 layout.css.devPixelsPerPx=1.3
```

## Ubuntu 开机启动慢问题处理

输入  `systemd-analyze blame ` 定位出启动耗时长的服务，然后禁用

```
     3min 8.471s fstrim.service
         32.099s apt-daily.service
         25.246s plymouth-quit-wait.service
         12.949s snapd.service
         10.433s NetworkManager.service
         10.418s dev-sdb6.device
         10.044s NetworkManager-wait-online.service
          8.340s networkd-dispatcher.service
          8.130s systemd-journal-flush.service
          7.842s plymouth-read-write.service
          6.809s apache2.service
```

禁用 ssd 硬盘相关服务和每日更新软件的服务

```bash
sudo systemctl disable fstrim.service
sudo systemctl disable fstrim.timer
sudo systemctl disable apt-daily.service
sudo systemctl disable apt-daily.timer
sudo systemctl disable NetworkManager-wait-online.service
sudo systemctl disable apache2.service
sudo systemctl mask plymouth-quit-wait.service

```


systemctl list-unit-files | grep fstrim.timer
systemctl list-unit-files | grep fstrim.service




### locate 命令

locate 命令可以用来快速查找文件，比 find 命令快很多，原理是通过检索数据库 /var/lib/mlocate/mlocate.db，
缺点是可能刚创建的文件不能立马检索到，默认是每天更新一次数据库，可以通过 /usr/bin/updatedb 命令来手动更新
locate 命令用到的 mlocate.db 数据库可能不断膨胀， 可以在 /etc/updatedb.conf 配置文件配置过滤


### dns 查看

view /etc/resolv.conf
systemd-resolve --status


### 禁用 ipv6

临时禁用

```
sudo sysctl -w net.ipv6.conf.all.disable_ipv6=1
sudo sysctl -w net.ipv6.conf.default.disable_ipv6=1
sudo sysctl -w net.ipv6.conf.lo.disable_ipv6=1
```

永久禁用
编辑 /et/sysctl.conf

```
net.ipv6.conf.all.disable_ipv6=1
net.ipv6.conf.default.disable_ipv6=1
net.ipv6.conf.lo.disable_ipv6=1
```

使配置立马生效：`sysctl -p`


### 禁用 ubuntu 开机广告

使用 `systemd-analyze blame` 命令查看开机启动时间发现，有个 motd-news 服务耗时 30 多秒，禁用得了。

要禁用广告，编辑文件：
```
sudo vim /etc/default/motd-news
```

将 ENABLED=1 改为 0 
```
ENABLED=0
```


## eclipse 手动安装

先下载 eclipse 压缩包，解压到 /usr/local/eclipse 目录，然后创建快捷方式

```
sudo tar -xzvf eclipse-cpp-2019-12-R-linux-gtk-x86_64.tar.gz -C /usr/local/
sudo mv /usr/local/eclipse /usr/local/eclipse-cpp-2019
sudo ln -s /usr/local/eclipse-cpp-2019/ /usr/local/eclipse

```

创建快捷方式

sudo vim /usr/share/applications/eclipse.desktop

```ini
[Desktop Entry]
Encoding=UTF-8
Name=Eclipse Platfrom
Comment=Eclipse IDE
Exec=/usr/local/eclipse/eclipse
Icon=/usr/local/eclipse/icon.xpm
Terminal=false
StartupNotify=true
Type=Application
Categories=Application;Development;
Keywords=eclipse;

```
