# gradle 配置


## gradle 本地安装

国外下载安装包地址很慢，换成 腾讯镜像下载地址 https://mirrors.cloud.tencent.com/gradle/

下载 [gradle-7.6.1-bin.zip](https://mirrors.cloud.tencent.com/gradle/gradle-7.6.1-bin.zip) ,，解压到 `D:\develop\app\gradle-7.6.1`

配置环境变量 

```
GRADLE_HOME=D:\develop\app\gradle-7.6.1
GRADLE_USER_HOME=D:\develop\app\gradle-7.6.1\init.d
PATH=%PATH%;%GRADLE_HOME%\bin
```

指定 M2_HOME 环境变量来保证能使用 MAVEN 本地缓存

```
M2_HOME=D:\develop\app\apache-maven-3.9.2
```


配置全局镜像地址，创建 `%GRADLE_HOME%/init.d/init.gradle`

```
allprojects {

    buildscript {
        repositories {
            mavenLocal()
            maven { url 'https://mirrors.cloud.tencent.com/gradle/'}
            maven { url 'https://maven.aliyun.com/repository/public/' }
            maven { url 'https://maven.aliyun.com/repository/spring/'}
            maven { url 'https://maven.aliyun.com/repository/google/'}
            maven { url 'https://maven.aliyun.com/repository/gradle-plugin/'}
            maven { url 'https://maven.aliyun.com/repository/spring-plugin/'}
            maven { url 'https://maven.aliyun.com/repository/grails-core/'}
            maven { url 'https://maven.aliyun.com/repository/apache-snapshots/'}
            mavenCentral()
        }
    }
    
    repositories {
        mavenLocal()
        maven { url 'https://mirrors.cloud.tencent.com/gradle/'}
        maven { url 'https://maven.aliyun.com/repository/public/' }
        maven { url 'https://maven.aliyun.com/repository/spring/'}
        maven { url 'https://maven.aliyun.com/repository/google/'}
        maven { url 'https://maven.aliyun.com/repository/gradle-plugin/'}
        maven { url 'https://maven.aliyun.com/repository/spring-plugin/'}
        maven { url 'https://maven.aliyun.com/repository/grails-core/'}
        maven { url 'https://maven.aliyun.com/repository/apache-snapshots/'}
        mavenCentral()
    }
}
```




## 参考文档

1. [Initialization Scripts](https://docs.gradle.org/current/userguide/init_scripts.html)
2. [gradle 配置指南](https://developer.aliyun.com/mvn/guide)
3. [gradle中mavenLocal()失效问题](https://developer.aliyun.com/article/789070)
