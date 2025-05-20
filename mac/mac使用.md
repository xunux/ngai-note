

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

快速引入包
`control + option + o`
