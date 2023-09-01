# Maven 进阶使用

1. maven 下载依赖包源码

```
mvn dependency:sources
```

或者执行 
```
-DdownloadSources=true -DdownloadJavadocs=true
```

或者在 settings.xml 配置

```xml
  <profiles>
    <profile>
      <id>env-default</id>
      <properties>
        <downloadSources>true</downloadSources>
        <downloadJavadocs>true</downloadJavadocs>
      </properties>
    </profile>
  </profiles>

  <activeProfiles>
    <activeProfile>env-default</activeProfile>
  </activeProfiles>
```

其他以 `-D` 配置的参数都可以在 `properties` 内配置

```xml
  <profiles>
    <profile>
      <id>env-default</id>
      <properties>
        <downloadSources>true</downloadSources>
        <downloadJavadocs>true</downloadJavadocs>
        <maven.test.skip>true</maven.test.skip>
        <maven.wagon.http.ssl.insecure>true</maven.wagon.http.ssl.insecure>
        <maven.wagon.http.ssl.allowall>true</maven.wagon.http.ssl.allowall>
        <maven.wagon.http.ssl.ignore.validity.dates>true</maven.wagon.http.ssl.ignore.validity.dates>
      </properties>
    </profile>
  </profiles>

  <activeProfiles>
    <activeProfile>env-default</activeProfile>
  </activeProfiles>
```


## 参考文档

1. [Maven Settings Reference](https://maven.apache.org/settings.html)
2. [Maven Wagon HTTP](https://maven.apache.org/wagon/wagon-providers/wagon-http/)
3. [Guide to Remote repository access through authenticated HTTPS](https://maven.apache.org/guides/mini/guide-repository-ssl.html)
4. [Configuring Apache Maven](https://maven.apache.org/configure.html)