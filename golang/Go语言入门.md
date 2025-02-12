# go 语言入门



## Go 数据类型

数据类型有 bool int int8 int16 int32 int64 float float64 uint8 uint16 uint32 uint64 string byte rune 等


## Go 控制流程

if/else 不需要在 if 中使用括号() 将条件包裹，if 条件可以定义局部变量

switch/case 默认是执行完语句就 break 的，跟 C语言不一样，如果要让 case 内的代码执行完后再执行下一个 case 可以加 fallthrough，
 另外 case 支持多个表达式

for 循环同样不需要括号()，for{} 代表死循环

defer panic recover 处理异常

### 数组&切片

数组声明和定义
```golang
var arr [3]int
var strs = [...]string{"ni", "wo", "ta"}
var strs2 = []string{"ni", "wo", "ta"}
```

数组默认会初始化，与 C 语言不同

切片是动态的，数组大小是固定的，好比是 java array 与 list 的区别， 切片 slice 用 append 插入数据，无法删除数据，但可以通过再次切片形势删除数据

### map

### map 定義及初始化
```golang
// 定义map
var m map[string]int
// 初始化 map
m = make(map[string]int)

// 定义并初始化 map
m2 := map[string]int{}

// 定义泛型 map
m3 = map[string]any{}
```

### map 操作

map 不是线程安全的，需要使用线程安全的 map 得使用 sync.map
map 赋值

```golang
m["user"] = "lois.蕭"
```

刪除元素

```golang
delete(m, "user")
```
   
遍历 map

```go
for k, v := range m {
    fmt.Println("key: ", k, "value: ", v)
}
```

### 结构体 struct

struct 声明
```golang

type User struct{
    Id int `json:"id"`
    Age int `json:"age"`
    Name string `json:"name"`
}

var user1 User
user2 := User{1, 18, "xiaoli"}
```



## 入门学习材料

1. https://learn.microsoft.com/zh-cn/training/paths/go-first-steps/
2. 