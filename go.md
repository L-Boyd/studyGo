# 优势

+ 部署简单
  + 可直接编译成机器码
  + 不依赖其他库
  + 直接运行即可部署
+ 静态类型语言
  + 编译的时候能检查出大多数问题
+ 语言层面的并发
  + 天生基因支持并发
  + 充分利用多核
+ 强大的标准库
  + runtime系统调度机制
  + 高效的GC垃圾回收
  + 丰富的标准库
+ 简单易学
  + 只有25个关键字
  + 支持内嵌C语法
  + 面向对象特征（封装、继承、多态）
  + 跨平台



# 开始前

要先定义package



# fmt包

+ `fmt.`

  + Print、Println

    + ("hello, world")
    + ("a", "b", "c")：Print多个字符串间没空格，Println有空格

  + Printf

    + ("hello, world")

    + ```go
      var a = "aaa"
      var b = "bbb"
      fmt.Printf("a = %v b = %v", a)
      
      /*
      %v：输出值
      %d：输出十进制数
      %b：输出二进制数
      %T：输出数据类型
      */
      ```



# 注释

```go
// 单行注释

/*
	多行注释
*/
```





# 数据类型

## 基本数据类型

### 整型

+ ```
  uint8：0-255（2^8-1) == byte
  uint16：0-65535
  uint32
  uint64
  int8：-128-127
  int16
  int32 == rune
  int64
  int 会在32位或64平台下分别取对应位数的长度
  ```

### 浮点型

+ ```
  float32
  float64
  complex64：32位实数和虚数
  complex128
  ```

+ ```go
  fmt.Println("%v, %f, %.2f")	// 原样输出，保留小数点后六位，保留小数点后两位
  
  var v = 3.14e2	// 表示3.14 * 10^2
  var v = 3.14e-2	// 表示3.14 * 10^(-2)
  ```

+ 精度丢失解决：github.com/shopspring/decimal包

  + ```go
    num1 := 3.1
    num2 := 4.2
    d1 := decimal.newFromFloat(num1)
    d2 := decimal.newFromFloat(num2)
    d3 := d1.Add(d2)
    d4 := d1.Sub(d2)	
    // Mul、Div
    ```

### 布尔型

+ `bool`
+ 默认值为false
+ 不允许整形强转为布尔型
+ 布尔型无法参与数值运算，也无法与其他类型进行转换

### 字符串

+ `string`

+ 默认值为空

+ 转义符

  + `\r`：回车符（返回行首）
  + `\n`：换行符（跳到下一行的同列位置）
  + `\t`：制表符
  + `\'`：单引号
  + `\"`：双引号
  + `\\`：反斜杠

+ 多行字符串

  + 用双引号只能定义单行字符串，多行要用反引号

  + ```go
    str := `line1
    line2
    line3`
    ```

+ 常用操作

  + 求长度：len(str)，一个汉字三个字节，一个英文字母一个字节。string底层是指针实现的，用unsave.SizeOf()会一直得到16
  
  + 拼接字符串：+或`str := fmt.Sprintf("%V %v", str1, str2)`
  
  + 分割字符换：字符串拆分为切片
  
    + ```go
      var str = "123-456-789"
      arr := strings.Split(str, "-")
      ```
  
  + 切片转换为字符串
  
    + ```go
      var str = "123-456-789"
      arr := strings.Split(str, "-")
      str2 := strings.Join(arr, "*")
      ```
  
  + 判断一个字符串是否包含另一个字符串
  
    + ```go
      str1 := "this is str"
      str2 := "this"
      flag := strings.Contains(str) 
      ```
  
  + 前缀/后缀判断
  
    + ```go
      str1 := "this is str"
      str2 := "this"
      flag1 := strings.HasPrefix(str1, str2)
      flag2 := strings.HasSuffix(str1, str2)
      ```
  
  + 找子字符串出现的位置，查不到返回-1，找到返回下标位置
  
    + ```go
      str1 := "this is str"
      str2 := "is"
      num1 := strings.Index(str1, str2)
      num2 := strings.LastIndex(str1, str2)
      ```
  
  + 修改字符串
  
    + ```go
      s1 := "big"
      byteStr := []byte(s1)
      byteStr[0] = 'p'
      fmt.Println(string(byteStr))
      
      s2 := "你好"
      runeStr := []rune(s2)
      runeStr[0] = 'p'
      fmt.Println(string(runeStr))
      ```
  
    + 
  

### byte、rune（字符）

+ byte：uint8类型，代表ASCII码的一个字符

  + ```go
    var a = 'a'	// 单引号定义字符，用%v输出为int类型（ASCII码值），用%c输出原样
    ```

+ rune：int32，代表一个UTF-8字符

  + ```go
    s := "你好golang"
    for i := 0; i < len(s); i++ {
        fmt.Printf("%v(%c)", s[i], s[i])	// 会出问题，因为汉字占三个字节
    }
    
    for _, v := range s {
        fmt.Printf("%v(%c)", v, v)
    }
    ```

  + 





## 复合数据类型

+ 数组
+ 切片
+ 结构体
+ 函数
+ map
+ 通道channel
+ 接口

```go
func main() {
    a := 1
    fmt.Printf("a的类型为：%T\n". a)	// a的类型为int
    unsafe.Sizeof(a)	// 返回a的内存占用字节数
}
```



## 强转类型

可能出现精度丢失，建议低位转高位

```go
// 整型间转换
var a1 int8 = 8
var a2 int16 = 10
fmt.Printf(int16(a1) + a2)

// 浮点型间转换
var b1 float32 = 8
var b2 float64 = 10
fmt.Printf(float64(b1) + b2)

// 整型和浮点型转换
var c1 float32 = 8
var c2 int = 10
fmt.Printf(c1 + float32(c2))

// 其他类型转换成string
var i int = 20
var f float64 = 12.456
var t bool = true
var b byte = 'a'
// 第一种，Sprintf
str1 := fmt.Sprintf("%d", i)
str2 := fmt.Sprintf("%f", f)
str3 := fmt.Sprintf("%t", t)
str4 := fmt.Sprintf("%c", b)
// 第二种，strconv包
str1 := strconv.FormatInt(int64(i), 10)	// int64的数，进制
str2 := strconv.FormatFloat(float64(f), 10)	// 数值，格式化方式('f','b','e','E','g','G')，保留小数点个数-1表示不对小数点格式化，格式化类型（32或64）
str3 := strconv.FormatBool(t)
str4 := strconv.FormatUint(uint64(b), 10)

// string转int
str := "123456"
num, _ := strconv.ParseInt(str, 10, 64)	// 值，进制，位数

// string转float
num2, _ := strconv.ParseFloat(str, 64) // 值， 位数
```



## 数字字面量语法

%v：原样输出

%d：十进制输出

%b：二进制输出

%o：八进制输出

%x：十六进制输出

%f：输出浮点类型值（默认保留小数点后六位）

%T：输出变量类型

```go
v := 0b1001
v2 := 0o377
v3 := 0xac
```





# 运算符

```
逻辑运算符
&&：与，若前面为false，则不会执行后面的
||：或
!：非

位运算符
&：按位与
|：按位或
^：按位异或
<<：左移
>>：右移

&变量：返回变量存储地址
*变量：这个变量是指针变量
var v1 = 4
var ptr *int
ptr = &v1
fmt.Println(*ptr)	// 4
fmt.Println(ptr)	// 0xc00008c058
v1 = 10
fmt.Println(*ptr)	// 10

算术运算符
+：加
-：减
*：乘
/：除，如果运算的数都是整数，相除后会去掉小数部分 
%：取余
++和--：自增、自减，只能单独使用，而且只能放在后面

关系运算符
==
!=
>
<
>=
<=

赋值运算符
=：直接赋值
+=：相加后赋值
-=
*=
/=
%=
```



# 函数

```go
func function_name([parameter list]) [return_types] {
    body
}

func sum(a int, b int) int {
    return a + b
}

// 如果参数都一样，可以只在最后一个变量写类型
func sum(a, b int) int {
    return a + b
}

func sum(a, b int) {	// 没有返回值可以省略返回值类型
    fmt.Println(a+b)
}

func swap(a, b int) (int, int) {	// 可以返回多个值
    return b, a
}

func sum(a int, b int,
         transform func(int) int) int {	// 函数可以作为一个函数的参数
    return transform(a) + transform(b)
}
func main() {
    square := func(x int) int {
        return x * x
    }
    fmt.Println(sum(1, 2, square))
}
```



# 变量

+ 定义位置

  + 函数内定义的变量是局部变量

  + 函数外定义的变量是全局变量

+ 全局变量和局部变量名可以相同，函数内的局部变量会被优先考虑

```go
var d = 2	// 函数外定义变量则必须要用var
func main() {
    a := 1	// 在函数内定义变量可以直接用 :=，并在定义同时赋值
    var b = 2
    var c int // 如果定义时不赋值，则需要指定类型
    fmp.Println(a)
    
    var (
    	username string
        age int
        sex string
    )
    
    var name, age = getUserinfo()
    // _为匿名变量，不会分配内存空间，不存在重复声明
    var name2, _ = getUserinfo()
}

func getUserinfo() (string, int) {
    return "zhangsan", 20
}
```

## 命名规则

+ 由数字、字母、下划线组成，首个字符不能为数字
+ 驼峰命名：age、maxAge（表示私有）、MaxAge（表示公有）、特有名词全大写（如DNS）



# 常量

```go
// 常量定义时必须赋值
const pi = 3.1415926

// 一次声明多个常量
const (
	A = "AAA"
    B = "BBB"
    C	// 没赋值则与上一行相同
)
```

## iota

iota是go的常量计数器，只能在常量的表达式中使用

iota在const关键字出现时被重置为0，const中每新增一行常量声明，将使iota计数一次（iota可理解为const语句块中的行索引）

```go
const a = iota // a=0
const (
	b = iota	// b=0
    c	// c=1
    d	// d=2
    _
    e	// e=4
    f = 100
    g = iota // g=6
)

const (
	n1, n2 = iota+1, iota+2	// 1 2
    n3, n4	// 2 3
    n5, n6	// 3 4
)
```





# 分支结构

## if

```go
func main() {
    a := 10
    if a > 10 {	// 条件不用加括号
        fmt.Println("a大于10")
    } else if a > 5 {
        fmt.Println("a大于5")
    } else {
        fmt.Println("a不大于5")
    }
    
    if b:= 3; a > 10 {	// if里可以初始化变量，只能在if-else if-else语句块内使用，用分号隔开条件
        fmt.Println("a大于10")
    }
}
```

## switch...case

```go
a := 10
switch a {
	case 1:	// case 只会执行一个，所以可以省略break
    fmt.Println("a==1")
    case 5:
    fmt.Println("a==5")
    case 10:
    fmt.Println("a==10")
    default:
    fmt.Println(a)
}
```

## 循环（只有for）

``` go
for i := 1; i < 5; i++ {
	fmt.Println(i)
}

// 模拟while
i := 1
for i < 5 {
    fmt.Println(i)
    i++
}
```









# 数组

```go
var arrayName [arraySize]dataType

a := [5]int{1, 2, 3, 4, 5}
```



# 切片/动态数组

```go
var sliceName []type
var sliceName []type = make([]type, length)
var sliceName []type = make([]type, length, capacity)

len(sliceName)	// 长度
cap(sliceName)	// 容量

a := make([]int, 0)
a = append(a, 1, ,2 ,3 ,4)

a := []int{1, 2, 3, 4}	// 用字面值创建

a[0] = 5	// 修改值
```



# Map

```go
var m = map[key类型]val类型
m := make(map[string]int)

m := map[string]int {
    "a" : 1,
    "b" : 2
}
fmt.Println(m)	// map[a:1 b:2]


```



# 结构体、指针

```go
type struct_variable_type struct {
   member definition
   member definition
   ...
   member definition
}

type Point struct {
    x int
    y int
}
func main() {
    p := Point{1, 2}
    fmt.Println(p)	// {1, 2}
    p.x = 3	// 修改值
    
    q := p
    p.x = 3	// 默认按值复制
    fmt.Println(p)	// {3, 2}
    fmt.Println(q)	// {1, 2}
    
    q := &p
    p.x = 3
    fmt.Println(p)	// {3, 2}
    fmt.Println(*q)	// {3, 2}
    
    q.y = 10	// 可以省略*
    fmt.Println(*q)	// {3, 10}
}
```

## 结构的方法

```go
type Point struct {
    x int
    y int
}

func (p Point) SetPoint(x, y int) {	// 方法名推荐首字母大写
    p.x = x
    p.y = y
}

func main() {
    p := Point{1, 2}
    p.SetPoint(3, 4)
    fmt.Println(p) // {1 2}。没用指针，值传递，方法在内部又复制了一份结构，不影响原结构
}

// 大部分情况方法要加*
func (p *Point) SetPoint(x, y int) {	// 方法名推荐首字母大写
    p.x = x
    p.y = y
}
func main() {
    p := Point{1, 2}
    p.SetPoint(3, 4)
    fmt.Println(p) // {3 4}
}
```



# 接口

```go
type interface_name interface {
    Print() // 规定必须要实现的方法，无需func关键字
}

type Shape interface {
    Print() // 规定必须要实现的方法，无需func关键字
}

// 结构要实现接口，定义同名的方法即可
type Rectangle struct{}
type Triangle struct{}

func (r Rectangle) Print() {
    fmt.Println("矩形")
}
func (r Triangle) Print() {
    fmt.Println("三角形")
}

func main() {
    var s Shape
    s = Rectangle{}
    s.Print()
    s = Triangle{}
    s.Print()
    
    printShape(Rectangle{})
}

func printShape(s Shape) {
    s.Print()
}
```



# 错误处理

```go
n, err := fmt.Println("dd")
if err != nil {
    // 执行正常的代码
} else {
    // 执行异常代码
}
```



# 并发

## Goroutines

```go
func func1() {
    time.Sleep(500 * time.Millisecond)
    fmt.Println("函数1")
}
func func2() {
    fmt.Println("函数2")
}

func main() {
    go func2()	// 开启一个线程执行func2
    func1()
}
```

## Channels——线程间通信

```go
func main() {
    ch := make(chan string) // chan关键字 管道里的数据类型
    go func2(ch)
    res1 := <- ch // 接收管道数据（发送、接收管道数据会阻塞线程）
    
    go func1(ch)
    res2 := <- ch
}

func func1(ch chan string) {
    ch <- "函数1"	// 给管道发送数据
}
func func2(ch chan string) {
    ch <- "函数2"
}
```



# 命令

```bash
go mod init 项目名 # 初始化项目
go mod tidy # 导入需要依赖，清除无用依赖

go get gorm.io/gorm # 下载第三方库到本地，并导入go.mod，更新go.sum校验文件
```

