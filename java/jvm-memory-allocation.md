# JVM 内存分配

## 对象的内存分配

Java程序运行过程中，创建对象是最常见的内存分配场景。在JVM中，对象主要分配在堆上，具体的分配策略和机制如下：

### 对象创建的内存分配方式

#### 1. 指针碰撞（Bump the Pointer）

- 适用于内存规整的情况（如使用Serial、ParNew等带有压缩整理过程的收集器）
- 已使用内存和空闲内存之间有一个指针作为分界点
- 分配内存时只需要把指针向空闲空间方向移动对象大小的距离

```
已分配区域 | 指针 | 空闲区域
```

#### 2. 空闲列表（Free List）

- 适用于内存不规整的情况（如使用CMS这种基于标记-清除算法的收集器）
- 虚拟机维护一个列表，记录哪些内存块是可用的
- 分配内存时从列表中找到一块足够大的空间分配给对象，并更新列表

### 内存分配的并发问题

创建对象在虚拟机中是非常频繁的行为，可能会存在多个线程同时创建对象的情况，这就需要考虑线程安全问题。主要有两种解决方案：

#### 1. 同步处理（CAS+失败重试）

- 采用CAS（Compare And Swap）配上失败重试的方式保证更新操作的原子性
- 当多个线程尝试更新同一内存区域时，只有一个线程能成功，其他线程需要重试

#### 2. 本地线程分配缓冲（TLAB）

- Thread Local Allocation Buffer，为每个线程预先分配一小块内存
- 线程在TLAB上分配对象时不需要同步，只有TLAB用完并分配新的TLAB时才需要同步
- 可以通过`-XX:+/-UseTLAB`参数来设置是否使用TLAB

## 堆内存分配策略

### 1. 对象优先在Eden区分配

- 大多数情况下，对象在新生代Eden区分配
- 当Eden区没有足够空间进行分配时，虚拟机将发起一次Minor GC

### 2. 大对象直接进入老年代

- 大对象是指需要大量连续内存空间的对象，如长字符串或数组
- 通过`-XX:PretenureSizeThreshold`参数设置大对象的阈值
- 目的是避免在Eden区和Survivor区之间发生大量的内存复制

### 3. 长期存活的对象进入老年代

- 虚拟机给每个对象定义了一个对象年龄（Age）计数器
- 对象在Eden出生并经过第一次Minor GC后仍然存活，将被移动到Survivor区，年龄设为1
- 对象在Survivor区每熬过一次Minor GC，年龄就增加1
- 当年龄增加到一定程度（默认为15，可通过`-XX:MaxTenuringThreshold`设置），就会晋升到老年代

### 4. 动态对象年龄判定

- 如果在Survivor空间中相同年龄所有对象大小的总和大于Survivor空间的一半，年龄大于或等于该年龄的对象可以直接进入老年代
- 这种动态机制使得对象的晋升更加灵活

### 5. 空间分配担保

- 在发生Minor GC之前，虚拟机会先检查老年代最大可用的连续空间是否大于新生代所有对象总空间
- 如果成立，那么Minor GC可以确保是安全的
- 如果不成立，则查看`-XX:HandlePromotionFailure`参数是否允许担保失败
  - 如果允许，会继续检查老年代最大可用的连续空间是否大于历次晋升到老年代对象的平均大小
  - 如果大于，将尝试进行一次Minor GC，但这次Minor GC是有风险的
  - 如果小于，或者`-XX:HandlePromotionFailure`设置不允许冒险，那么改为进行一次Full GC

## 内存分配示例

### 示例1：新对象分配

```java
public class AllocationExample {
    private static final int _1MB = 1024 * 1024;
    
    public static void testAllocation() {
        byte[] allocation1, allocation2, allocation3, allocation4;
        allocation1 = new byte[2 * _1MB];
        allocation2 = new byte[2 * _1MB];
        allocation3 = new byte[2 * _1MB];
        allocation4 = new byte[4 * _1MB]; // 出现一次Minor GC
    }
    
    public static void main(String[] args) {
        testAllocation();
    }
}
```

运行参数：`-Xms20M -Xmx20M -Xmn10M -XX:+PrintGCDetails -XX:SurvivorRatio=8`

这个例子中，堆内存大小为20MB，其中新生代10MB，Eden:Survivor=8:1，即Eden区8MB，两个Survivor区各1MB。当分配allocation4时，Eden区已经被占用6MB，不足以分配4MB大小的对象，会触发一次Minor GC。

### 示例2：大对象直接进入老年代

```java
public class BigObjectExample {
    private static final int _1MB = 1024 * 1024;
    
    public static void testPretenureSizeThreshold() {
        byte[] allocation = new byte[4 * _1MB]; // 直接分配在老年代
    }
    
    public static void main(String[] args) {
        testPretenureSizeThreshold();
    }
}
```

运行参数：`-Xms20M -Xmx20M -Xmn10M -XX:+PrintGCDetails -XX:SurvivorRatio=8 -XX:PretenureSizeThreshold=3145728`

这个例子中，我们设置了`-XX:PretenureSizeThreshold=3145728`（3MB），所以超过3MB的对象会直接在老年代分配。

### 示例3：长期存活的对象进入老年代

```java
public class TenuringThresholdExample {
    private static final int _1MB = 1024 * 1024;
    
    public static void testTenuringThreshold() {
        byte[] allocation1, allocation2, allocation3;
        allocation1 = new byte[_1MB / 4]; // 256KB
        
        // 什么时候进入老年代取决于XX:MaxTenuringThreshold设置
        allocation2 = new byte[4 * _1MB];
        allocation3 = new byte[4 * _1MB];
        allocation3 = null;
        allocation3 = new byte[4 * _1MB];
    }
    
    public static void main(String[] args) {
        testTenuringThreshold();
    }
}
```

运行参数：`-Xms20M -Xmx20M -Xmn10M -XX:+PrintGCDetails -XX:SurvivorRatio=8 -XX:MaxTenuringThreshold=1 -XX:+PrintTenuringDistribution`

这个例子中，allocation1对象会在第一次GC后进入Survivor区，并且年龄为1。由于我们设置了`-XX:MaxTenuringThreshold=1`，所以当第二次GC发生时，allocation1会被晋升到老年代。

## 内存分配调优建议

1. **合理设置堆内存大小**
   - 避免内存不足导致频繁GC
   - 避免内存过大导致GC时间过长
   - 可以通过`-Xms`和`-Xmx`参数设置

2. **优化新生代和老年代的比例**
   - 新生代太小会导致对象过早进入老年代
   - 新生代太大会减少老年代空间
   - 可以通过`-XX:NewRatio`参数设置

3. **调整对象晋升阈值**
   - 根据对象的生命周期特点调整`-XX:MaxTenuringThreshold`
   - 短生命周期对象应该在新生代被回收，避免进入老年代

4. **使用TLAB提高分配效率**
   - 对于多线程应用，开启TLAB可以减少线程同步开销
   - 可以通过`-XX:+UseTLAB`参数开启

5. **避免创建大量临时对象**
   - 减少不必要的对象创建
   - 考虑使用对象池或缓存机制
   - 使用StringBuilder替代String的+操作

6. **合理设置大对象阈值**
   - 根据应用特点设置`-XX:PretenureSizeThreshold`
   - 避免大对象在新生代和老年代之间复制

7. **监控内存分配情况**
   - 使用`-XX:+PrintGCDetails`、`-XX:+PrintTenuringDistribution`等参数查看GC和对象晋升情况
   - 使用工具如JVisualVM、MAT分析内存使用情况