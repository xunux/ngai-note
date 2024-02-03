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
        <maven.compile.fork>true</maven.compile.fork>
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

2. maven 多线程加速

以 2 倍核心线程打包
```
mvn -T 2C package
```

2 个线程执行
```
mvn -T 2 package
```

3. `-DskipTests=true` 与 `-Dmaven.test.skip=true` 区别

`-Dmaven.test.skip=true` 表示跳过测试用例的编译，且跳过测试用例的执行
`-DskipTests=true` 不跳过测试用例的编译阶段，但会跳过测试用例的执行

4. maven 下载加速

为了加速 maven 下载速度，除了优化网络环境，配上国内或内网的镜像地址外，还可以配置同时下载个数，默认情况下能同时下载 5 个

```
export MAVEN_OPTS=-Dmaven.artifact.threads=5
mvn -Dmaven.artifact.threads=50
```


5. 重新下载某个项目的所有依赖

```
mvn dependency:purge-local-repository
```

删除并重新下载与某个依赖x有关的依赖包
```
mvn clean dependency:purge-local-repository -DresolutionFuzziness=org.id.of.x
```

## 参考文档

1. [Maven Settings Reference](https://maven.apache.org/settings.html)
2. [Maven Wagon HTTP](https://maven.apache.org/wagon/wagon-providers/wagon-http/)
3. [Guide to Remote repository access through authenticated HTTPS](https://maven.apache.org/guides/mini/guide-repository-ssl.html)
4. [Configuring Apache Maven](https://maven.apache.org/configure.html)
5. [Maven skip tests](https://stackoverflow.com/questions/24727536/maven-skip-tests)
6. [Force re-download of release dependency using Maven](https://stackoverflow.com/questions/7959499/force-re-download-of-release-dependency-using-maven)
7. [Configuring Parallel Artifact Resolution](https://maven.apache.org/guides/mini/guide-configuring-maven.html#configuring-parallel-artifact-resolution)
8. []()