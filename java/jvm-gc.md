# JVM 垃圾回收

## 垃圾回收概述

垃圾回收（Garbage Collection，GC）是Java虚拟机自动管理内存的一种机制，它能自动识别和清除不再使用的对象，释放内存空间。垃圾回收的主要目标是：

1. 释放不再使用的对象占用的内存
2. 减少内存碎片，提高内存利用率
3. 尽可能减少应用暂停时间（Stop-The-World）

## 判断对象是否存活的算法

### 1. 引用计数法

- **基本原理**：为对象添加一个引用计数器，每当有一个地方引用它时，计数器值加1；当引用失效时，计数器值减1；任何时刻计数器为0的对象就是不可能再被使用的。
- **优点**：实现简单，判定效率高
- **缺点**：无法解决对象之间相互循环引用的问题
- **应用**：Python、PHP等语言的垃圾回收机制采用引用计数法，但Java虚拟机没有选用此算法

### 2. 可达性分析算法

- **基本原理**：通过一系列称为「GC Roots」的对象作为起始点，从这些节点开始向下搜索，搜索所走过的路径称为引用链。当一个对象到GC Roots没有任何引用链相连时，则证明此对象是不可用的。
- **GC Roots包括**：
  - 虚拟机栈（栈帧中的本地变量表）中引用的对象
  - 方法区中类静态属性引用的对象
  - 方法区中常量引用的对象
  - 本地方法栈中JNI（Native方法）引用的对象
  - Java虚拟机内部的引用（如基本数据类型对应的Class对象、常驻的异常对象等）
  - 所有被同步锁（synchronized关键字）持有的对象
  - 反映Java虚拟机内部情况的JMXBean、JVMTI中注册的回调、本地代码缓存等

## Java中的引用类型

Java中提供了四种引用类型，用于灵活控制对象的生命周期：

### 1. 强引用（Strong Reference）

- 最常见的引用类型，如`Object obj = new Object()`
- 只要强引用存在，垃圾收集器永远不会回收被引用的对象

### 2. 软引用（Soft Reference）

- 用于描述一些还有用但非必需的对象
- 在系统将要发生内存溢出异常之前，会把这些对象列入回收范围进行第二次回收
- 通过`SoftReference`类实现

```java
SoftReference<Object> softRef = new SoftReference<>(new Object());
```

### 3. 弱引用（Weak Reference）

- 比软引用更弱，被弱引用关联的对象只能生存到下一次垃圾收集发生之前
- 当垃圾收集器工作时，无论当前内存是否足够，都会回收掉只被弱引用关联的对象
- 通过`WeakReference`类实现

```java
WeakReference<Object> weakRef = new WeakReference<>(new Object());
```

### 4. 虚引用（Phantom Reference）

- 最弱的引用关系，一个对象是否有虚引用的存在，完全不会对其生存时间构成影响
- 为一个对象设置虚引用关联的唯一目的是能在这个对象被收集器回收时收到一个系统通知
- 通过`PhantomReference`类实现，必须和引用队列（ReferenceQueue）联合使用

```java
ReferenceQueue<Object> queue = new ReferenceQueue<>();
PhantomReference<Object> phantomRef = new PhantomReference<>(new Object(), queue);
```

## 垃圾收集算法

### 1. 标记-清除算法（Mark-Sweep）

- **执行过程**：
  1. 标记阶段：标记出所有需要回收的对象
  2. 清除阶段：统一回收所有被标记的对象
- **优点**：基础算法，后续算法都是基于它改进的
- **缺点**：
  - 效率问题：标记和清除两个过程的效率都不高
  - 空间问题：标记清除后会产生大量不连续的内存碎片

### 2. 复制算法（Copying）

- **执行过程**：
  1. 将可用内存按容量划分为大小相等的两块
  2. 每次只使用其中的一块
  3. 当这一块的内存用完了，就将还存活着的对象复制到另外一块上面
  4. 然后再把已使用过的内存空间一次清理掉
- **优点**：实现简单，运行高效，不会产生内存碎片
- **缺点**：将内存缩小为原来的一半，代价太高
- **应用**：新生代垃圾收集，HotSpot虚拟机的新生代使用复制算法，但不是按照1:1的比例划分内存，而是将内存分为一块较大的Eden空间和两块较小的Survivor空间（默认8:1:1）

### 3. 标记-整理算法（Mark-Compact）

- **执行过程**：
  1. 标记阶段：与标记-清除算法一样
  2. 整理阶段：让所有存活的对象都向一端移动，然后直接清理掉端边界以外的内存
- **优点**：不会产生内存碎片
- **缺点**：移动对象需要更新所有引用，效率较低
- **应用**：老年代垃圾收集，因为老年代对象存活率高，使用复制算法效率会降低

### 4. 分代收集算法（Generational Collection）

- **基本思想**：根据对象存活周期的不同将内存划分为几块，不同块采用不同的收集算法
- **实现**：
  - 新生代：对象存活率低，使用复制算法
  - 老年代：对象存活率高，使用标记-清除或标记-整理算法
- **优点**：结合多种算法的优点，提高垃圾收集效率
- **应用**：当前商业虚拟机的垃圾收集都采用分代收集算法

## 垃圾收集器

HotSpot虚拟机提供了多种垃圾收集器，每种收集器都有各自的特点和适用场景：

### 1. Serial收集器

- **特点**：单线程收集器，在垃圾收集时，必须暂停所有的工作线程（Stop The World）
- **应用**：Client模式下的默认新生代收集器
- **适用场景**：单CPU环境下的运行，或者内存较小的嵌入式设备
- **参数**：`-XX:+UseSerialGC`启用Serial + Serial Old的收集器组合

### 2. ParNew收集器

- **特点**：Serial收集器的多线程版本
- **应用**：Server模式下的默认新生代收集器
- **适用场景**：多CPU环境，与CMS收集器配合使用
- **参数**：`-XX:+UseParNewGC`启用ParNew + Serial Old的收集器组合

### 3. Parallel Scavenge收集器

- **特点**：新生代收集器，复制算法，多线程，关注吞吐量（高效率利用CPU）
- **应用**：吞吐量优先的应用
- **适用场景**：后台运算而不需要太多交互的任务
- **参数**：
  - `-XX:+UseParallelGC`启用Parallel Scavenge + Serial Old的收集器组合
  - `-XX:+UseParallelOldGC`启用Parallel Scavenge + Parallel Old的收集器组合
  - `-XX:MaxGCPauseMillis`设置最大垃圾收集停顿时间
  - `-XX:GCTimeRatio`设置吞吐量大小

### 4. Serial Old收集器

- **特点**：Serial收集器的老年代版本，单线程收集器，使用标记-整理算法
- **应用**：Client模式下的默认老年代收集器
- **适用场景**：与Serial收集器搭配使用，或者作为CMS收集器的后备预案

### 5. Parallel Old收集器

- **特点**：Parallel Scavenge收集器的老年代版本，多线程，使用标记-整理算法
- **应用**：吞吐量优先的应用
- **适用场景**：与Parallel Scavenge收集器搭配使用
- **参数**：`-XX:+UseParallelOldGC`启用Parallel Scavenge + Parallel Old的收集器组合

### 6. CMS收集器（Concurrent Mark Sweep）

- **特点**：以获取最短回收停顿时间为目标的收集器，基于标记-清除算法
- **执行过程**：
  1. 初始标记（CMS initial mark）：标记GC Roots能直接关联到的对象，速度很快，需要Stop The World
  2. 并发标记（CMS concurrent mark）：进行GC Roots Tracing，与用户线程并发执行
  3. 重新标记（CMS remark）：修正并发标记期间因用户程序继续运作而导致标记产生变动的那一部分对象的标记记录，需要Stop The World
  4. 并发清除（CMS concurrent sweep）：清除标记为死亡的对象，与用户线程并发执行
- **优点**：并发收集，低停顿
- **缺点**：
  - 对CPU资源非常敏感
  - 无法处理浮动垃圾（Floating Garbage）
  - 产生大量空间碎片
- **应用**：互联网站或B/S系统的服务端上，重视服务的响应速度
- **参数**：
  - `-XX:+UseConcMarkSweepGC`启用CMS收集器
  - `-XX:CMSInitiatingOccupancyFraction`设置CMS收集器在老年代空间被使用多少后触发垃圾收集

### 7. G1收集器（Garbage-First）

- **特点**：面向服务端应用的垃圾收集器，基于标记-整理算法，可以精确控制停顿
- **执行过程**：
  1. 初始标记（Initial Marking）：标记GC Roots能直接关联到的对象，需要Stop The World
  2. 并发标记（Concurrent Marking）：进行GC Roots Tracing，与用户线程并发执行
  3. 最终标记（Final Marking）：修正并发标记期间因用户程序继续运作而导致标记产生变动的那一部分对象的标记记录，需要Stop The World
  4. 筛选回收（Live Data Counting and Evacuation）：对各个Region的回收价值和成本进行排序，根据用户所期望的GC停顿时间来制定回收计划，需要Stop The World
- **优点**：
  - 并行与并发
  - 分代收集
  - 空间整合
  - 可预测的停顿
- **应用**：JDK 9及以后的默认垃圾收集器
- **参数**：
  - `-XX:+UseG1GC`启用G1收集器
  - `-XX:MaxGCPauseMillis`设置期望的最大GC停顿时间

### 8. ZGC收集器（Z Garbage Collector）

- **特点**：可伸缩的低延迟垃圾收集器，目标是在任何堆大小下都能将STW时间控制在10ms以内
- **执行过程**：基于Region内存布局，使用了读屏障、染色指针和内存多重映射等技术来实现并发的标记-整理算法
- **优点**：超低停顿时间，在大内存场景下表现尤为出色
- **应用**：JDK 11引入，JDK 15成为正式特性
- **参数**：`-XX:+UseZGC`启用ZGC收集器

### 9. Shenandoah收集器

- **特点**：与ZGC类似，目标是低延迟垃圾收集器，但实现方式不同
- **执行过程**：通过读写屏障和Brooks转发指针来实现并发的标记-整理算法
- **优点**：低停顿时间，与ZGC相比更加注重通用性
- **应用**：JDK 12引入，JDK 15成为正式特性
- **参数**：`-XX:+UseShenandoahGC`启用Shenandoah收集器

## GC日志分析

### GC日志参数

- `-XX:+PrintGC`：输出GC日志
- `-XX:+PrintGCDetails`：输出详细GC日志
- `-XX:+PrintGCTimeStamps`：输出GC的时间戳（JVM启动到当前的总时长）
- `-XX:+PrintGCDateStamps`：输出GC的时间戳（日期）
- `-XX:+PrintHeapAtGC`：在进行GC的前后打印出堆的信息
- `-Xloggc:filename`：将GC日志输出到文件

### GC日志示例与分析

#### Minor GC日志示例

```
[GC (Allocation Failure) [PSYoungGen: 65536K->10752K(76288K)] 65536K->10760K(251392K), 0.0115712 secs] [Times: user=0.06 sys=0.01, real=0.01 secs]
```

分析：
- `[GC (Allocation Failure)`：表示这是一次Minor GC，原因是分配失败
- `[PSYoungGen: 65536K->10752K(76288K)]`：表示年轻代从65536K减少到10752K，总容量为76288K
- `65536K->10760K(251392K)`：表示整个堆从65536K减少到10760K，总容量为251392K
- `0.0115712 secs`：表示GC耗时
- `[Times: user=0.06 sys=0.01, real=0.01 secs]`：表示用户态耗时0.06秒，内核态耗时0.01秒，实际耗时0.01秒

#### Full GC日志示例

```
[Full GC (Metadata GC Threshold) [PSYoungGen: 10752K->0K(76288K)] [ParOldGen: 8K->10674K(175104K)] 10760K->10674K(251392K), [Metaspace: 20856K->20856K(1069056K)], 0.0631695 secs] [Times: user=0.23 sys=0.00, real=0.06 secs]
```

分析：
- `[Full GC (Metadata GC Threshold)`：表示这是一次Full GC，原因是元空间达到阈值
- `[PSYoungGen: 10752K->0K(76288K)]`：表示年轻代从10752K减少到0K，总容量为76288K
- `[ParOldGen: 8K->10674K(175104K)]`：表示老年代从8K增加到10674K，总容量为175104K
- `10760K->10674K(251392K)`：表示整个堆从10760K减少到10674K，总容量为251392K
- `[Metaspace: 20856K->20856K(1069056K)]`：表示元空间从20856K保持不变，总容量为1069056K
- `0.0631695 secs`：表示GC耗时
- `[Times: user=0.23 sys=0.00, real=0.06 secs]`：表示用户态耗时0.23秒，内核态耗时0.00秒，实际耗时0.06秒

## GC调优策略

### 1. 选择合适的垃圾收集器

- 吞吐量优先：Parallel Scavenge + Parallel Old
- 响应时间优先：ParNew + CMS或G1
- 大内存低延迟：ZGC或Shenandoah

### 2. 调整内存大小

- 堆内存：`-Xms`和`-Xmx`
- 新生代：`-Xmn`或`-XX:NewRatio`
- Eden和Survivor比例：`-XX:SurvivorRatio`

### 3. 调整GC触发时机

- CMS：`-XX:CMSInitiatingOccupancyFraction`
- G1：`-XX:InitiatingHeapOccupancyPercent`

### 4. 调整GC线程数

- `-XX:ParallelGCThreads`：并行GC线程数
- `-XX:ConcGCThreads`：并发GC线程数

### 5. 减少Full GC

- 增大堆内存和元空间大小
- 及时清理不再使用的对象引用
- 避免创建大量临时对象
- 使用软引用和弱引用

### 6. 监控GC情况

- 使用GC日志分析工具：GCViewer、GCEasy等
- 使用监控工具：JVisualVM、JConsole等
- 关注GC频率、GC耗时和内存使用情况

## 实际案例分析

### 案例1：频繁的Minor GC

**现象**：应用程序每隔几秒就会发生一次Minor GC，但每次回收的内存很少。

**原因分析**：
1. 新生代空间太小
2. 应用程序创建了大量临时对象

**解决方案**：
1. 增加新生代空间：`-Xmn`或调整`-XX:NewRatio`
2. 优化代码，减少临时对象的创建
3. 使用对象池或缓存机制

### 案例2：频繁的Full GC

**现象**：应用程序频繁发生Full GC，导致应用响应时间变长。

**原因分析**：
1. 内存泄漏
2. 老年代空间不足
3. 大对象直接进入老年代

**解决方案**：
1. 使用内存分析工具（如MAT）查找内存泄漏
2. 增加老年代空间
3. 调整大对象阈值：`-XX:PretenureSizeThreshold`
4. 使用CMS或G1收集器减少停顿时间

### 案例3：长时间的GC暂停

**现象**：应用程序偶尔会出现长时间的GC暂停，影响用户体验。

**原因分析**：
1. 使用了会导致长时间暂停的垃圾收集器
2. 堆内存过大，导致GC时间延长

**解决方案**：
1. 使用低延迟垃圾收集器：CMS、G1、ZGC或Shenandoah
2. 调整堆内存大小，找到合适的平衡点
3. 调整`-XX:MaxGCPauseMillis`参数（G1收集器）
4. 增加并行GC线程数：`-XX:ParallelGCThreads`