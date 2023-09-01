# maven 配置

配置 maven 仓库位置

```xml
  <!--<localRepository>${user.home}/.m2/repository</localRepository>-->
  <localRepository>D:/develop/repo/.m2/repository</localRepository>
  <interactiveMode>true</interactiveMode>
  <offline>false</offline>
```

插件配置
```xml
  <pluginGroups>
    <pluginGroup>org.eclipse.jetty</pluginGroup>
  </pluginGroups>
```

配上 jetty 插件后可以直接执行 `mvn jetty:run` 代替 `org.eclipse.jetty:jetty-maven-plugin:run`

maven 代理配置
```xml
  <proxies>
    <!-- proxy
     | Specification for one proxy, to be used in connecting to the network.
     |
    <proxy>
      <id>optional</id>
      <active>true</active>
      <protocol>http</protocol>
      <username>proxyuser</username>
      <password>proxypass</password>
      <host>proxy.host.net</host>
      <port>80</port>
      <nonProxyHosts>local.net|some.host.com</nonProxyHosts>
    </proxy>
    -->
    <proxy>
      <id>optional</id>
      <active>true</active>
      <protocol>http</protocol>
      <username></username>
      <password></password>
      <host>127.0.0.1</host>
      <port>1081</port>
      <nonProxyHosts>local.net|some.host.com</nonProxyHosts>
    </proxy>
  </proxies>
```

配置 maven 镜像源地址

```xml
  <mirrors>
    <!-- mirror
     | Specifies a repository mirror site to use instead of a given repository. The repository that
     | this mirror serves has an ID that matches the mirrorOf element of this mirror. IDs are used
     | for inheritance and direct lookup purposes, and must be unique across the set of mirrors.
     |
    <mirror>
      <id>mirrorId</id>
      <mirrorOf>repositoryId</mirrorOf>
      <name>Human Readable Name for this Mirror.</name>
      <url>http://my.repository.com/repo/path</url>
    </mirror>
     -->
    <mirror>
      <id>huaweicloud</id>
      <mirrorOf>central,jcenter,jboss,grails-core,grails-plugins,goolge,spring-release,spring-plugins</mirrorOf>
      <url>https://repo.huaweicloud.com/repository/maven/</url>
    </mirror>
    <mirror>
      <id>aliyunmaven</id>
      <mirrorOf>central,jcenter</mirrorOf>
      <name>阿里云公共仓库</name>
      <url>https://maven.aliyun.com/repository/public</url>
    </mirror>
    <mirror>
      <id>aliyun-spring</id>
      <mirrorOf>spring</mirrorOf>
      <name>阿里云代理仓库</name>
      <url>https://maven.aliyun.com/repository/spring</url>
    </mirror>
    <mirror>
      <id>aliyun-spring-plugin</id>
      <mirrorOf>spring-plugin</mirrorOf>
      <name>阿里云代理仓库</name>
      <url>https://maven.aliyun.com/repository/spring-plugin</url>
    </mirror>
    <mirror>
      <id>aliyun-google</id>
      <mirrorOf>google</mirrorOf>
      <name>阿里云代理谷歌仓库</name>
      <url>https://maven.aliyun.com/repository/google</url>
    </mirror>
    <mirror>
      <id>aliyun-apache-snapshots</id>
      <mirrorOf>apache.snapshots</mirrorOf>
      <name>阿里云代理仓库</name>
      <url>https://maven.aliyun.com/repository/apache-snapshots</url>
    </mirror>
    <!--<mirror>
      <id>maven-default-http-blocker</id>
      <mirrorOf>external:http:*</mirrorOf>
      <name>Pseudo repository to mirror external repositories initially using HTTP.</name>
      <url>http://0.0.0.0/</url>
      <blocked>true</blocked>
    </mirror>-->
  </mirrors>
```

maven 设置代理后报 ssl 证书认证错误，方法有 2 种

方法一：忽略 ssl 检查

在 home 目录建立 .mavenrc 文件，然后在里面配置 MAVEN_OPTS 环境变量

```shell
MAVEN_OPTS="-Dmaven.wagon.http.ssl.insecure=true -Dmaven.wagon.http.ssl.allowall=true -Dmaven.wagon.http.ssl.ignore.validity.dates=true"
```

或在执行 mvn 命令时带上 `-Dmaven.wagon.http.ssl.insecure=true -Dmaven.wagon.http.ssl.allowall=true -Dmaven.wagon.http.ssl.ignore.validity.dates=true`
 参数

甚至可以在项目的根目录建立一个 `.mvn/maven.config` 文件，内容为：

```
-Dmaven.wagon.http.ssl.insecure=true
-Dmaven.wagon.http.ssl.allowall=true
-Dmaven.wagon.http.ssl.ignore.validity.dates=true
```

方法二：https 请求经过 http 代理后相当于是中间人攻击，需要下载 http 代理的 https 根证书并安装到 java 里才行，操作如下

```
keytool -keystore ${JAVA_HOME}/jre/lib/security/cacerts -list
keytool -keystore ${JAVA_HOME}/jre/lib/security/cacerts -importcert -alias YOUR_CA_ALIAS -file xxx-ca.cert
```

## 参考文档

1. [Maven Settings Reference](https://maven.apache.org/settings.html)
2. [Maven Wagon HTTP](https://maven.apache.org/wagon/wagon-providers/wagon-http/)
3. [Guide to Remote repository access through authenticated HTTPS](https://maven.apache.org/guides/mini/guide-repository-ssl.html)
4. [Configuring Apache Maven](https://maven.apache.org/configure.html)