# maven 配置

配置 maven 仓库位置

```xml
  <localRepository>D:/develop/repo/.m2/repository</localRepository>
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
      <id>aliyunmaven</id>
      <mirrorOf>central,jcenter</mirrorOf>
      <name>阿里云公共仓库</name>
      <url>https://maven.aliyun.com/repository/public</url>
    </mirror>
    <mirror>
      <id>aliyun-apache-snapshots</id>
      <mirrorOf>snapshots</mirrorOf>
      <name>阿里云代理仓库</name>
      <url>https://maven.aliyun.com/repository/apache-snapshots</url>
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
    <!--<mirror>
      <id>maven-default-http-blocker</id>
      <mirrorOf>external:http:*</mirrorOf>
      <name>Pseudo repository to mirror external repositories initially using HTTP.</name>
      <url>http://0.0.0.0/</url>
      <blocked>true</blocked>
    </mirror>-->
  </mirrors>
```
