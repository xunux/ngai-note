# Maven 依赖管理

Maven 默认使用"最近依赖原则"解决版本冲突，即使用离当前项目最近的依赖版本。

## 依赖配置
Maven 依赖在 pom.xml 文件中通过 <dependencies> 标签进行配置

### 依赖范围
通过 <scope> 标签控制依赖的可见性和传递性：

常用的依赖范围：

- compile ：默认值，编译、测试、运行时都有效
- provided ：编译和测试有效，运行时由容器提供
- runtime ：测试和运行时有效，编译时无效
- test ：仅测试时有效
- system ：类似 provided，但需指定系统路径
- import ：用于依赖管理，导入 dependencyManagement

## 依赖传递
Maven 会自动传递依赖，即如果项目 A 依赖项目 B，项目 B 依赖项目 C，那么项目 A 就会自动依赖项目 C。

### 依赖版本管理
使用 <dependencyManagement> 集中管理依赖版本：

```xml
<dependencyManagement>
    <dependencies>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-dependencies</artifactId>
            <version>2.7.3</version>
            <type>pom</type>
            <scope>import</scope>
        </dependency>
    </dependencies>
</dependencyManagement>
 ```

子模块引用时可以省略版本号：

```xml
<dependencies>
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-web</artifactId>
        <!-- 无需指定版本，从 dependencyManagement 继承 -->
    </dependency>
</dependencies>
 ```

### 依赖冲突
Maven 默认使用"最近依赖原则"解决版本冲突，可通过以下方式手动处理：

1. 使用 <exclusions> 排除冲突依赖
2. 在 <dependencies> 中显式声明希望使用的版本
3. 使用 mvn dependency:tree 分析依赖树

### 强制更新依赖
```bash
mvn clean install -U
 ```

### 清理本地仓库缓存
```bash
mvn dependency:purge-local-repository
 ```

或手动删除特定依赖：

```bash
rm -rf ~/.m2/repository/com/example/my-dependency
 ```
```

通过合理配置 Maven 依赖，可以有效管理项目的第三方库，提高开发效率和项目质量。
