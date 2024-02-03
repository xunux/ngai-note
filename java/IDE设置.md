# IDE 配置


## eclipse 设置

打开 eclipse.exe 文件同级目录下的 **eclipse.ini** 文件，添加如下额外设置参数
```ini
-vm
D:/develop/app/jdk-16.0.1/bin
-Duser.name=Lois Siau
-javaagent:D:\developer\app\eclipse-jee-2021-03\plugins\org.projectlombok.agent_1.18.22/lombok.jar
```


## idea 设置

打开 idea64.exe 文件同级目录下的 **idea64.exe.vmoptions** 文件设置额外参数
```
-Duser.name=Lois Siau
```

<p>

### IDEA 插件

1. IdeaVim
   Ctrl+C、Ctrl+V、Ctrl+F 组合键与VIM功能会有冲突，在 Settings->Editors->Vim 设置即可
2. IdeaVimExtension
   `set keep-english-in-normala-and-restore-in-insert`加入到`~/.ideavimrc` 来开启自动切换中英文输入法
3. IdeaVim-EasyMotion
4. AceJump
5. Key PromoterX
6. Lombok
7. Maven Helper
8. Rainbow Brackets
9.  Spring Boot Assistant
10. Alibaba Java Coding Guideline


