# ubuntu 安装 idea

1. 下载 idea 社区版本，地址：https://www.jetbrains.com/idea/download/?section=linux, 直接下载地址：[ideaIntelliJ IDEA Community Edition](https://www.jetbrains.com/idea/download/download-thanks.html?platform=linux&code=IIC)

2. 解压安装


```bash
sudo mv ideaIC-2023-2.tar.gz /usr/local
cd /usr/local/
tar -xzvf ideaIC-2023-2.tar.gz
ln -s idea-IC-232.8660.185/ ideaICln -s
# 运行 idea
./idea.sh
```

创建桌面快捷方式

```bash
cd ~
touch idea.desktop
chmod a+x idea.desktop
vim idea.desktop
```

内如如下：

```ini
[Desktop Entry]
Encoding=UTF-8
Name=IntelliJ IDEA
Comment=IntelliJ IDEA
Exec=/usr/local/ideaIC/bin/idea.sh
Icon=/usr/local/ideaIC/bin/idea.svg
Terminal=false
Type=Application
Categories=Application;Development;

```

## 常用插件

1. IdeaVim
2. Rainbow Brackets
3. Maven Helper
4. aiXcoder
5. Alibaba Cloud AI Coding Assistant (Cosy)
6. Alibaba Java Coding Guidelines
7. Key Promoter X
8. ignore
9. Lombok
10. markdown
11. Spring boot Assistant

## 其他配置

1. java properties 文件编码设为 utf-8
2. maven 配置
3. jdk 配置
4. spring boot assistant 配置关联文件
