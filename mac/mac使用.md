
## Mac 必备工具安装

1. HomeBrew 安装

安装国内源的 homebrew
```shell
/bin/zsh -c "$(curl -fsSL https://gitee.com/cunkai/HomebrewCN/raw/master/Homebrew.sh)"
```

2. 安装 fish

```shell
brew install fish
fish
fish_add_path /opt/homebrew/bin
echo "/opt/homebrew/bin/fish" | sudo tee -a /etc/shells
chsh -s /opt/homebrew/bin/fish
```

## 清除 dns

```
sudo dscacheutil -flushcache
sudo killall -HUP mDNSResponder
```

## tcpdump

```
sudo tcpdump -i any port 53 -nn
```

## ip工具

```
brew install iproute2mac
```

## idea 快捷键的使用

快速引入包 `control + option + o`
