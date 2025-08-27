
## Doris Broker load 导数

```sql

# 查看load任务
show load;

# 取消load任务
cancel load from <db_name> where LABEL = "<load_label>";

# 删除load任务记录，可以重复使用该 label，便于程序重复执行
clean label <load_label> from <db_name>;

```

### 从 S3 导数

```sql
-- -- 清除导入任务 label，供重复使用该 label
-- CLEAN LABEL <lable_name> from <db_name>;

-- 从 S3 导入 parquet 文件数据

LOAD LABEL <db_name>.load_s3_parquet_xxx
(
    DATA INFILE("s3://<path>/*.parq")
    INTO TABLE <table_name>
    FORMAT AS "parquet"
    (
        field_1,
        field_2,
        field_xx
    )
    set (
        field_1=`field_1`,
        
    )
)
WITH S3
(
    "provider" = "S3",
    "s3.endpoint" = "xx.amazonaws.com",
    "s3.region" = "xxx",
    "s3.access_key" = "xxx",
    "s3.secret_key" = "xxxx"
)
PROPERTIES
(
    "max_filter_ratio" = "0.05",
    "timeout" = "3600"
);
```