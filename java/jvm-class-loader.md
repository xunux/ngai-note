# JVM 类加载器

类加载器（Class Loader）是Java虚拟机的重要组成部分，负责在运行时查找和加载类的字节码文件。Java中的类加载器采用了双亲委派模型，这种模型能够保证Java类型体系的完整性和安全性。

## 类加载器的分类

Java中的类加载器主要分为以下几种：

### 1. 启动类加载器（Bootstrap ClassLoader）

- **职责**：加载Java核心类库，如`java.lang.*`、`java.util.*`等
- **特点**：
  - 由C++实现，是JVM的一部分，不是Java类
  - 在Java代码中获取启动类加载器的引用时会返回null
  - 只加载存放在`<JAVA_HOME>/lib`目录或被`-Xbootclasspath`参数指定路径中的类库
- **加载范围**：`rt.jar`、`resources.jar`、`charsets.jar`等核心类库

### 2. 扩展类加载器（Extension ClassLoader）

- **职责**：加载Java的扩展类库
- **特点**：
  - 由Java实现，是`sun.misc.Launcher$ExtClassLoader`类的实例
  - 父加载器为启动类加载器
  - JDK 9之后被平台类加载器（Platform ClassLoader）取代
- **加载范围**：`<JAVA_HOME>/lib/ext`目录或被`java.ext.dirs`系统变量指定路径中的类库

### 3. 应用程序类加载器（Application ClassLoader）

- **职责**：加载应用程序的类
- **特点**：
  - 由Java实现，是`sun.misc.Launcher$AppClassLoader`类的实例
  - 父加载器为扩展类加载器
  - 也被称为系统类加载器（System ClassLoader）
  - 通过`ClassLoader.getSystemClassLoader()`方法可以获取到它
- **加载范围**：classpath指定的目录中的类库

### 4. 自定义类加载器（User-Defined ClassLoader）

- **职责**：加载用户自定义的类
- **特点**：
  - 通过继承`java.lang.ClassLoader`类实现
  - 父加载器通常为应用程序类加载器
  - 可以实现特定的类加载逻辑，如从网络或数据库加载类

## 双亲委派模型

### 工作原理

双亲委派模型（Parents Delegation Model）的工作过程是：

1. 当一个类加载器收到类加载请求时，它首先不会自己尝试加载这个类，而是把这个请求委派给父类加载器去完成。
2. 每一个层次的类加载器都是如此，因此所有的加载请求最终都应该传送到最顶层的启动类加载器中。
3. 只有当父类加载器反馈自己无法完成这个加载请求（它的搜索范围中没有找到所需的类）时，子类加载器才会尝试自己去加载。

```
┌─────────────────────────────────────────────────────┐
│                                                     │
│                 启动类加载器                        │
│             Bootstrap ClassLoader                   │
│                      ▲                             │
│                      │                             │
│                      │                             │
│                 扩展类加载器                        │
│             Extension ClassLoader                   │
│                      ▲                             │
│                      │                             │
│                      │                             │
│               应用程序类加载器                      │
│            Application ClassLoader                  │
│                      ▲                             │
│                      │                             │
│                      │                             │
│                自定义类加载器                       │
│            User-Defined ClassLoader                 │
│                                                     │
└─────────────────────────────────────────────────────┘
```

### 双亲委派模型的好处

1. **确保Java核心类库的类型安全**：例如，用户自定义的`java.lang.Object`类永远不会被加载，因为该类已经被启动类加载器加载过了。

2. **避免类的重复加载**：当父类加载器已经加载了某个类时，子类加载器就不需要再次加载。

3. **保证Java类型体系的统一性**：保证了不会出现同一个类被不同的类加载器加载的情况，避免了类型转换异常。

### 双亲委派模型的实现

双亲委派模型的实现集中在`java.lang.ClassLoader`的`loadClass()`方法中：

```java
protected Class<?> loadClass(String name, boolean resolve) throws ClassNotFoundException {
    synchronized (getClassLoadingLock(name)) {
        // 首先检查该类是否已经被加载过
        Class<?> c = findLoadedClass(name);
        if (c == null) {
            try {
                // 如果有父类加载器，则委托父类加载器加载
                if (parent != null) {
                    c = parent.loadClass(name, false);
                } else {
                    // 如果没有父类加载器，则委托启动类加载器加载
                    c = findBootstrapClassOrNull(name);
                }
            } catch (ClassNotFoundException e) {
                // 如果父类加载器无法加载，则调用自身的findClass方法加载
            }

            if (c == null) {
                // 父类加载器无法加载时，调用自己的findClass方法进行加载
                c = findClass(name);
            }
        }
        if (resolve) {
            resolveClass(c);
        }
        return c;
    }
}
```

## 破坏双亲委派模型

在一些特殊情况下，双亲委派模型会被破坏：

### 1. JDK 1.2之前的兼容性

JDK 1.2之前，自定义类加载器都必须重写`loadClass()`方法，而在JDK 1.2之后，只需要重写`findClass()`方法。为了兼容已有代码，JDK保留了这种破坏双亲委派模型的方式。

### 2. SPI（Service Provider Interface）机制

Java SPI机制需要由启动类加载器加载的`ServiceLoader`类去加载用户提供的SPI实现类，这就必须要由子类加载器去请求父类加载器，破坏了双亲委派模型的委托关系。

### 3. OSGi（Open Service Gateway Initiative）

OSGi实现了一个模块化的、动态的Java系统服务平台。它允许模块（Bundle）自己控制类的可见性，一个Bundle可以动态地向外部发布它的包和服务，也可以动态地导入它所依赖的包和服务。这种灵活的类加载机制打破了传统的双亲委派模型。

### 4. Tomcat类加载器

Tomcat服务器为了实现不同Web应用之间的类隔离，使用了自定义的类加载器结构，部分打破了双亲委派模型：

- Common类加载器：加载Tomcat和Web应用都可见的类
- Catalina类加载器：加载只有Tomcat可见的类
- Shared类加载器：加载所有Web应用都可见但Tomcat不可见的类
- WebApp类加载器：每个Web应用都有一个独立的类加载器，加载仅对当前Web应用可见的类

## 自定义类加载器

### 为什么需要自定义类加载器

1. **隔离加载类**：在一些应用中，需要隔离不同模块的类加载，避免类冲突。
2. **修改类加载方式**：从非标准来源加载类，如网络、数据库等。
3. **防止源码泄露**：通过自定义的加密解密机制保护字节码。
4. **动态创建和卸载类**：在OSGi等框架中，需要动态地加载和卸载模块。

### 如何实现自定义类加载器

实现自定义类加载器通常只需要重写`findClass()`方法：

```java
public class MyClassLoader extends ClassLoader {
    private String classPath;

    public MyClassLoader(String classPath) {
        this.classPath = classPath;
    }

    @Override
    protected Class<?> findClass(String name) throws ClassNotFoundException {
        try {
            byte[] classData = loadClassData(name);
            if (classData == null) {
                throw new ClassNotFoundException();
            }
            return defineClass(name, classData, 0, classData.length);
        } catch (IOException e) {
            throw new ClassNotFoundException("Could not load class: " + name, e);
        }
    }

    private byte[] loadClassData(String name) throws IOException {
        String fileName = classPath + File.separatorChar
                + name.replace('.', File.separatorChar) + ".class";
        try (FileInputStream fis = new FileInputStream(fileName);
             ByteArrayOutputStream baos = new ByteArrayOutputStream()) {
            int bufferSize = 1024;
            byte[] buffer = new byte[bufferSize];
            int bytesRead;
            while ((bytesRead = fis.read(buffer)) != -1) {
                baos.write(buffer, 0, bytesRead);
            }
            return baos.toByteArray();
        }
    }
}
```

### 使用自定义类加载器的示例

```java
public class ClassLoaderTest {
    public static void main(String[] args) throws Exception {
        // 创建自定义类加载器实例
        MyClassLoader myClassLoader = new MyClassLoader("D:\\classes");
        
        // 加载指定类
        Class<?> clazz = myClassLoader.loadClass("com.example.Test");
        
        // 创建实例并调用方法
        Object obj = clazz.newInstance();
        Method method = clazz.getMethod("hello");
        method.invoke(obj);
        
        // 输出类的加载器
        System.out.println(clazz.getClassLoader());
    }
}
```

## 类加载器的常见问题

### 1. 类加载器的命名空间

每个类加载器都有自己的命名空间，命名空间由该加载器及所有父加载器所加载的类组成。在同一个命名空间中，不会出现类的完整名字（包括包名）相同的两个类；在不同的命名空间中，有可能会出现类的完整名字（包括包名）相同的两个类。

### 2. 类的唯一性

对于任意一个类，都需要由加载它的类加载器和这个类本身一同确立其在Java虚拟机中的唯一性。也就是说，比较两个类是否相等，只有在这两个类是由同一个类加载器加载的前提下才有意义。

### 3. 类加载器的内存泄漏

类加载器可能导致内存泄漏，主要有以下几种情况：
- 类加载器的生命周期过长，导致其加载的所有类都无法被卸载
- 类加载器之间的引用关系不当，导致类加载器无法被垃圾回收
- 使用线程上下文类加载器时没有及时清理

### 4. 线程上下文类加载器

线程上下文类加载器（Thread Context ClassLoader）是一种破坏双亲委派模型的方式，它可以通过`Thread.setContextClassLoader()`和`Thread.getContextClassLoader()`方法来设置和获取。

线程上下文类加载器主要用于SPI机制，例如JDBC、JCE、JNDI、JAXB等，这些SPI的接口由启动类加载器加载，但实现类由应用程序类加载器加载。