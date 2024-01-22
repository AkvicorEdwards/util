# Util

```go
import "github.com/AkvicorEdwards/util"
```

- [x] CRC32
- [x] MD5 (File, reader, string, bytes)
- [x] SHA1 (File, reader, string, bytes)
- [x] SHA256 (File, reader, string, bytes)
- [x] SHA512 (File, reader, string, bytes)
- [x] Base64
- [x] bytes to int (uint16, uint32, uint64)
- [x] int to bytes (uint16, uint32, uint64)
- [x] Bytes Combine
- [x] Bit Set
- [x] JSON
- [x] File Stat
- [x] Path Split
- [x] Random String
- [x] IP
- [x] Time Calculate
- [x] Tcp Port Checker
- [x] Get Client IP & HTTP GET/POST/PUT Request & HTTP Response & Redirect

# CRC32 & MD5 & SHA1 & SHA256 & SHA512

各类型提供的函数方法**名称类似**，只需要将函数名中的`CRC32`替换为其他类型名即可使用

以下以CRC32为例子

- 计算结果为`*CRC32Result`类型，可通过以下方法获取不同类型结果
- `.Value() uint32` CRC32独有
- `.Array() [4]byte` 各类型长度不一致
- `.Slice() []byte`
- `.Upper() string`
- `.Lower() string`
- `.Error() error` 获取计算过程中产生的错误

## NewCRC32()

```go
// 示例
fmt.Println(NewCRC32().FromString("Akvicor").Lower())
```

- `.FromReader()` 从Reader中读取数据并计算
- `.FromFile()` 从文件中读取数据并计算
- `.FromReaderChunk()` 从Reader中分块读取数据并计算
- `.FromFileChunk()` 从文件中分块读取数据并计算
- `.FromBytes()` 计算byte数组
- `.FromString()` 计算字符串

## NewCRC32Pip()

```go
// 示例
cp := NewCRC32Pip()
io.WriteString(cp, "Akvicor")
fmt.Println(cp.Result().Value())
```

- `.Write()` Writer实现
- `.WriteBytes()` 写入`[]byte`
- `.WriteString()` 写入`string`
- `.Result()` 获取计算结果

# Base64

- 计算结果为`*Base64Result`类型，可通过以下方法获取不同类型结果
- `.Bytes() []byte`
- `.String() string`
- `.Error() error` 获取计算过程中产生的错误

## NewBase64()

- `.EncodeBytes()` 编码[]byte
- `.EncodeString()` 编码string
- `.DecodeBytes()` 解码[]byte
- `.DecodeString()` 解码string

# `[]byte <-> uint`

支持16/32/64位无符号整型，以下仅列出16位相关函数

- `.UInt16ToBytesSlice()` `uint16` 转 `[]byte`
- `.BytesSliceToUInt16()` `[]byte` 转 `uint16`
- `.UInt16ToBytesArray()` `uint16` 转 `[...]byte`
- `.BytesArrayToUInt16()` `[...]byte` 转 `uint16`

# Bytes Combine

合并多个`[]byte`

```go
b := BytesCombine([]byte{0x11, 0x22}, []byte{0x33, 0x44, 0x55}, []byte{0x66})
```

# Bit Set

设置整型变量二进制位，支持`byte`,`uint8`,`int8`,`uint16`,`int16`,`uint32`,`int32`,`uint64`,`int64`

```go
b8 := uint8(0)
BitSet(&b8, 0, true)
```

# JSON

```go
type TestStruct struct {
Name string `json:"name"`
Age  int    `json:"age"`
}
test1 := TestStruct{
Name: "Akvicor",
Age:  17,
}

fmt.Println(NewJSON(&test1, false).String())
fmt.Println(NewJSONResult([]byte(`{"name":"Akvicor","age":17}`)).Map())
```

- `NewJSONResult(data []byte) *JSON` 传入json格式数据并储存，自动判断是json对象还是json数组
- `NewJSON(v any, isArray bool) *JSON` 将传入变量转为json并储存，通过isArray指名传入的是否是数组
- 计算结果为`*JSON`类型，可通过以下方法获取不同类型结果
- `.Bytes() []byte`
- `.String() string`
- `.Map(idx ...int) map[string]any` 返回json对象，如果变量中保存的是json数组，返回idx位置的对象，默认为0
- `.MapArray() []map[string]any` 返回json数组，如果变量中保存的是json对象，自动创建一个数组并将变量作为数组第一个元素
- `.Error() error` 获取计算过程中产生的错误

# File Stat

判断文件类型（不存在，是文件夹，是文件

```go
FileStat(filename string) int
```

通过以下常量和返回值确定文件类型

```go
FileStatNotExist = 0
FileStatIsDir    = 1
FileStatIsFile   = 2
```

# Path Split

三种方法分割路径字符串

- `SplitPathSkip(p string, skip int) (head, tail string)` 跳过skip个`/`后，在下一个`/`处分割，且`/`保留在tail中
- `SplitPath(p string) (head, tail string)` 在第一个`/`处分割，且`/`保留在tail中
- `SplitPathRepeat(p string, repeat int) (head, tail string)` 在第一个`/`处分割，且`/`保留在tail中，将tail作为参数再次分割，重复repeat次

# Random String

返回随机字符串

字符

- `RandomLower` 小写字母
- `RandomUpper` 大写字母
- `RandomDigit` 数字
- `RandomSpecial` 特殊字符
- `RandomAlpha` 字母
- `RandomAll` 各类字符拼接成的单个字符串
- `RandomSlice` 各类型字符组成的字符串数组

方法

- `RandomString(length int, str ...string)` 随机生成，默认通过`RandomAll`生成，也可通过传入str来自定义
- `RandomStringAtLeastOnce(length int, str ...string)` 每种类型至少包含一个，每个str元素为一个类型，默认通过`RandomSlice`生成，也可通过传入str来自定义
- `RandomStringWithTimestamp(length int, unix ...int64)` 长度至少为8，返回包含时间戳的随机字符串
- `ParseRandomStringWithTimestamp(str string) (int64, string)` 解析包含时间戳的随机字符串

# IP

- `IsIPAddr` 判断是否是ip地址
- `IsLocalIPAddr` 判断ip地址是否是本地ip
- `IsLocalIP` 判断ip是否是本地ip
- `IPAddrToUint32` ip地址转为uint32
- `Uint32ToIPAddr` uint32转为ip地址
- `IPToUint32` ip转为uint32
- `Uint32ToIP` uint32转为ip

# Time Calculate

- `TimeNowFormat` 返回当前时间的格式化字符串
- `TimeNowToBase36` 返回当前时间的36进制字符串
- `TimeUnixToFormat` unix转为格式化字符串
- `TimeUnixToBase36` unix转为36进制字符串
- `TimeBase36ToUnix` 36进制字符串转为unix
- `YearBetweenTwoDate` 计算两个日期年份相减，不计算月等等
- `YearBetweenTwoTime` 计算两个日期间隔的年份，计算到秒
- `MonthBetweenTwoDate` 计算两个日期相差月份，不计算日等等
- `MonthBetweenTwoTime` 计算两个日期相差月份，计算到秒
- `DayBetweenTwoDate` 计算两个日期相差天数，不计算小时等等
- `DayBetweenTwoTime` 计算两个日期相差天数，计算到秒
- `HourBetweenTwoTime` 计算两个日期相差小时
- `MinuteBetweenTwoTime` 计算两个日期相差分钟
- `SecondBetweenTwoTime` 计算两个日期相差秒

# Tcp Port Checker

检测端口是否开放

- `TcpPortIsOpen(ip, port string)`
- `TcpPortIsOpenByAddr(ipPort string)`

# Get Client IP & HTTP GET/POST/PUT Request & HTTP Response & Redirect

## 获取IP

- `RemoteIP` 通过RemoteAddr获取ip
- `GetClientIP` 获取客户端ip,包含内网地址
- `GetClientPublicIP` 获取客户端ip

## 页面跳转

- `RespRedirect` 重定向
- `LastPage` 上一个页面
- `Reload` 刷新页面

## API响应

- `NewHTTPResp` 创建响应数据结构体变量
- `ParseHTTPResp` 解析json字符串为响应数据结构体变量
- `WriteHTTPRespAPIOk` 返回 ok
- `WriteHTTPRespAPIFailed` 返回 failed
- `WriteHTTPRespAPIInvalidKey` 返回 非法的key
- `WriteHTTPRespAPIInvalidInput` 返回 非法的输入
- `WriteHTTPRespAPIProcessingFailed` 返回 处理失败

## GET & POST

Content Type

- `HTTPContentTypeUrlencoded`
- `HTTPContentTypeJson`

方法

- `HttpGet(u string, args any)` u为url，args为可序列化变量
- `HttpPost(contentType string, u string, args any)`
- `HttpPostGet(contentType string, u string, argsGET, argsPOST any)` 

