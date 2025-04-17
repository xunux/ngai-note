# maven 配置

### 配置 maven 仓库位置

```xml
  <!--<localRepository>${user.home}/.m2/repository</localRepository>-->
  <localRepository>D:/develop/repo/.m2/repository</localRepository>
  <interactiveMode>true</interactiveMode>
  <offline>false</offline>
```

### 插件配置
```xml
  <pluginGroups>
    <pluginGroup>org.eclipse.jetty</pluginGroup>
  </pluginGroups>
```

配上 jetty 插件后可以直接执行 `mvn jetty:run` 代替 `org.eclipse.jetty:jetty-maven-plugin:run`

### 配置 maven 镜像源地址

```xml
  <mirrors>
    <mirror>
      <id>aliyunmaven</id>
      <mirrorOf>central,jcenter</mirrorOf>
      <name>阿里云公共仓库</name>
      <url>https://maven.aliyun.com/repository/public</url>
    </mirror>
    <mirror>
        <id>nexus-tencentyun</id>
        <mirrorOf>!internal-repo,*</mirrorOf>
        <name>Nexus tencentyun</name>
        <url>http://mirrors.cloud.tencent.com/nexus/repository/maven-public/</url>
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

### 全局 maven 仓库配置

```xml
    <profile>
      <id>internal-repo</id>
      <repositories>
        <repository>
          <id>nexus-releases</id>
          <name>nexus-releases</name>
          <url>http://my-compony.com/repository/maven-releases/</url>
          <releases>
            <enabled>true</enabled>
          </releases>
          <snapshots>
            <enabled>false</enabled>
          </snapshots>
        </repository>
        <repository>
          <id>nexus-snapshots</id>
          <name>nexus-snapshots</name>
          <url>http://my.repository.com/repository/maven-snapshots/</url>
          <releases>
            <enabled>false</enabled>
          </releases>
          <snapshots>
            <enabled>true</enabled>
          </snapshots>
        </repository>
      </repositories>
    </profile>
    <profile>
      <id>global-repo</id>
      <repositories>
        <repository>
          <id>google</id>
          <url>https://maven.google.com</url>
        </repository>
        <repository>
          <id>central</id>
          <url>https://maven.google.com</url>
        </repository>
      </repositories>
    </profile>
    <profile>
      <id>download-sources-javadoc</id>
      <properties>
        <downloadSources>true</downloadSources>
        <downloadJavadocs>true</downloadJavadocs>
      </properties>
    </profile>
  </profiles>

  <activeProfiles>
    <activeProfile>internal-repo</activeProfile>
    <activeProfile>global-repo</activeProfile>
    <activeProfile>download-sources-javadoc</activeProfile>
  </activeProfiles>
</settings>
```

## 参考文档

1. [Maven Settings Reference](https://maven.apache.org/settings.html)
2. [Configuring Apache Maven](https://maven.apache.org/configure.html)