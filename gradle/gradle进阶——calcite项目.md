# gradle 进阶——calcite项目


## calcite 项目打开

下载 calcite 源码 `git clone git@github.com:apache/calcite.git`


打开 IDEA，配置 gradle，选择 gradle 本地安装地址，或者修改 `gradle/gradle-wrapper.properties` 

```
# distributionUrl=https\://services.gradle.org/distributions/gradle-7.6.1-bin.zip
distributionUrl=https\://mirrors.cloud.tencent.com/gradle/gradle-7.6.1-bin.zip
```

## 问题

1. 单元测试用例运行报错

将 Calcite 项目的单元测试用例设为 IntelliJ IDEA 执行

步骤如下：
1. File > Setting (Ctrl+Alt+S)
2. Build, Execution, Deployment > Build Tools > gradle
3. Run Tests using: Intellij IDEA


参考 ["No tests found for given includes" when running Gradle tests in IntelliJ IDEA](https://stackoverflow.com/questions/60228404/no-tests-found-for-given-includes-when-running-gradle-tests-in-intellij-idea)




## 参考文档

1. 