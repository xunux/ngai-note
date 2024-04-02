## firefox 安装

方案1：基于 ppa 安装，，参考： https://sysin.cn/blog/ubuntu-remove-snap/#%E5%AE%89%E8%A3%85-DEB-%E6%A0%BC%E5%BC%8F%E7%9A%84-Firefox

安装 DEB 格式的 Firefox

添加 Firefox 官方 PPA（Personal Package Archives）仓库：
```
sudo add-apt-repository ppa:mozillateam/ppa
```
将 Firefox 官方 PPA 仓库中的 firefox 设为高优先级：

```
sudo sh -c "cat > /etc/apt/preferences.d/mozillateam-firefox.pref" << EOL
Package: firefox*
Pin: release o=LP-PPA-mozillateam
Pin-Priority: 501
EOL
```

现在可以安装最新版的 deb 版 Firefox：

```
sudo apt update
apt list firefox  #现在已经不是 Snap firefox
sudo apt install firefox
# 或者 Firefox ESR
sudo apt install firefox-esr
```

方案2：基于 mozilla 安装，参考： https://support.mozilla.org/zh-CN/kb/install-firefox-linux?utm_source=www.mozilla.org&utm_medium=referral&utm_campaign=firefox-download-thanks#w_shi-yong-deb-bao-she-zhi-firefox-yu-yan

使用基于 Debian 发行版的 .deb 包安装 Firefox

要通过 APT 库安装基于 .deb 包，请：

    创建一个保存 APT 库密钥的目录：
```
sudo install -d -m 0755 /etc/apt/keyrings
```
    导入 Mozilla APT 密钥环：
```
wget -q https://packages.mozilla.org/apt/repo-signing-key.gpg -O- | sudo tee /etc/apt/keyrings/packages.mozilla.org.asc > /dev/null
```
    如果没有安装 wget，请通过命令 sudo apt-get install wget 安装。

    密钥指纹应该是 35BAA0B33E9EB396F59CA838C0BA5CE6DC6315A3。你可以用以下命令检查：
```
gpg -n -q --import --import-options import-show /etc/apt/keyrings/packages.mozilla.org.asc | awk '/pub/{getline; gsub(/^ +| +$/,""); if($0 == "35BAA0B33E9EB396F59CA838C0BA5CE6DC6315A3") print "\nThe key fingerprint matches ("$0").\n"; else print "\nVerification failed: the fingerprint ("$0") does not match the expected one.\n"}'
```
    把 Mozilla APT 库添加到源列表中：
```
echo "deb [signed-by=/etc/apt/keyrings/packages.mozilla.org.asc] https://packages.mozilla.org/apt mozilla main" | sudo tee -a /etc/apt/sources.list.d/mozilla.list > /dev/null
```
    配置 APT 优先使用 Mozilla 库中的包：
```
echo '
Package: *
Pin: origin packages.mozilla.org
Pin-Priority: 1000
' | sudo tee /etc/apt/preferences.d/mozilla
```
    更新软件列表并安装 Firefox .deb 包： 
```
sudo apt-get update && sudo apt-get install firefox
```
使用 .deb 包设置 Firefox 语言

需要在不是英语的环境下使用 Firefox 的用户，请使用 .deb 包里带有的 Firefox 语言包。使用以下命令安装其他语言包，注意把其中的 fr 换成目标语言代码：
```
sudo apt-get install firefox-l10n-fr
```
安装 Mozilla APT 库之后，可以使用 sudo apt-get update:
```
apt-cache search firefox-l10n
```
列出所有语言。 

## firefox 配置

1. firefox 放大比例调整

firefox about:config 
```
layout.css.devPixelsPerPx=1.0
```
