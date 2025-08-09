# Apache Druid 笔记


## Druid 入门

## Druid 系统配置

```
-Xms130g

```

## Druid 特性

## Druid 问题记录

1. Druid 因列缓存设置过大导致在高基维聚合查询情况下促发 historical 节点 full gc的问题
解决方案： 1. 调整 historical 节点参数： druid.processing.columnCache.sizeBytes=10485760，之前是设为了 1G，导致一次查询申请了上千次缓存空间
	2. 高基维查询预防，设置查询超时时间，避免超大查询耗尽集群资源

## 