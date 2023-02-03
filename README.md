Golang 快速语法
===

# 特性

- 函数式编程语言
- 静态类型
- 语法类似C语言（没有分号），结构类似Oberon-2
- 编译成机器代码（没有 JVM）
- 没有类，但是 struct 可以带方法
- 接口类型 interface
- 没有继承，只允许类型嵌套/组合
- 函数是一等公民
- 函数可以返回多个值
- 有闭包
- 有指针，但是没有指针运算
- 内置并发原语（go/chan）

# 内置类型

- bool
- string
- int,int8,int16,int32,int64
- uint,uint8,uint16,uint32,uint64
- byte
- rune
- float32,float64
- complex64,complex128

# 包

- 每个 Go 源文件的最开头（不包括注释）添加包声明
- 可执行文件在 main 包里面
- 惯例：包名 == 包导入路径的最后一个名称（导入路径 math/rand => 包名 rand
- 大写字母开头的标识符：表示可导出（exported），对其他包可见
- 小写字母开头的标识符：表示私有（private），对其他包不可见

# 数组、切片

- 数组是值类型，切片是引用类型
