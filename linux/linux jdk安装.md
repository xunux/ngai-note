# linux jdk 安装


## 安装方式一：通过 apt 或 yum 命令安装

查看 jdk 列表
```
sudo apt search cache openjdk*
```

安装 jdk 和 源码
```
sudo apt install openjdk-11-jdk
sudo apt install openjdk-11-source
```

若 jdk 存在多版本，则还需要设置 jdk 版本
```
sudo update-alternatives --config java
```

<pre><font color="#26A269"><b>root@my-pc</b></font>:<font color="#12488B"><b>~</b></font>$ sudo update-alternatives --config java
有 2 个候选项可用于替换 java (提供 /usr/bin/java)。

  选择       路径                                          优先级  状态
------------------------------------------------------------
* 0            /usr/lib/jvm/java-11-openjdk-amd64/bin/java      1111      自动模式
  1            /usr/lib/jvm/java-11-openjdk-amd64/bin/java      1111      手动模式
  2            /usr/lib/jvm/java-8-openjdk-amd64/jre/bin/java   1081      手动模式

</pre>


设置 JAVA_HOME 环境变量

通过 update-alternatives 命令找到 jdk 安装路径，然后可以在 `/et/environment` 或 `/etc/profile` 或 `/etc/bash.bashrc` 或 `~/.bashrc` 或 `~/.bash_profile` 设置环境变量

```
vim /etc/environment

export JAVA_HOME=/usr/lib/jvm/java-11-openjdk-amd64
```

卸载 jdk

```
sudo apt remove openjdk-11-jdk
sudo apt remove openjdk-11-source
```

## 安装方式二：下载安装包安装




## 参考

1. [如何在 Ubuntu 20.04 上安装 Java](https://zhuanlan.zhihu.com/p/137114682)
2. 


