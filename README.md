# Go Util

```go
import "github.com/AkvicorEdwards/util"
```

- [x] CRC32
- [x] MD5 (File, reader, string, bytes)
- [x] SHA1 (File, reader, string, bytes)
- [x] SHA256 (File, reader, string, bytes)
- [x] SHA512 (File, reader, string, bytes)
- [x] AES (CBC)
- [x] Base64
- [x] bytes to int (uint16, uint32, uint64)
- [x] int to bytes (uint16, uint32, uint64)
- [x] Bytes Combine
- [x] Bit Set
- [x] JSON
- [x] Slice Remove Duplicates
- [x] File Stat
- [x] Dir List
- [x] Mkdir
- [x] Dir Size
- [x] Size Unit, Decimal System and Binary System
- [x] Path Split
- [x] Random String, Key
- [x] IP
- [x] Time Calculate
- [x] Tcp Port Checker
- [x] Port
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

# AES

- 计算结果为`*AESResult`类型，可通过以下方法获取不同类型结果
- `.Bytes() []byte`
- `.Base64() *Base64Result`
- `.String() string`
- `.Upper() string`
- `.Lower() string`
- `.Encrypted() bool`
- `.Decrypt(key []byte, iv ...[]byte) *AESResult`
- `.Encrypt(key []byte, iv ...[]byte) *AESResult`
- `.Error() error` 获取计算过程中产生的错误

## NewAESResult()

- `.EncryptCBC(origData, key []byte, iv ...[]byte) *AESResult`
- `.DecryptCBC(encrypted, key []byte, iv ...[]byte) *AESResult`

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

# Slice Remove Duplicates

Slice去除重复元素

- `NewRemoveDuplicates() *RemoveDuplicates`
- `.String(s []string) []string`
- `.Byte(s []byte) []byte`
- `.Int8(s []int8) []int8`
- `.Int16(s []int16) []int16`
- `.Int(s []int) []int`
- `.Int32(s []int32) []int32`
- `.Int64(s []int64) []int64`
- `.UInt8(s []uint8) []uint8`
- `.UInt16(s []uint16) []uint16`
- `.UInt(s []uint) []uint`
- `.UInt32(s []uint32) []uint32`
- `.UInt64(s []uint64) []uint64`
- `.Float32(s []float32) []float32`
- `.Float64(s []float64) []float64`

# File Stat

判断文件类型（不存在，是文件夹，是文件

```go
FileStat(filename string) *FileStatModel
```

通过以下函数确定文件类型

```go
IsFile()
NotFile()
IsDir()
NotDir()
IsExist()
NotExist()
IsDenied()
NotDenied()
IsError()
NotError()
```

# Dir List

获取文件列表

```go
DirList(p string) *DirListModel

type DirListModel struct {
    Files []DirListUnitModel
    Dirs  []DirListUnitModel
    Error error
}

type DirListUnitModel struct {
    Name    string
    Size    int64
    Mode    fs.FileMode
    ModTime time.Time
}
```

# Mkdir

创建文件夹

```go
MkdirP(p string, perm ...os.FileMode) error
```

# Mkdir

获取文件夹大小

```go
DirSize(p string) (int64, error)
```

# Size Unit, Decimal System and Binary System

将大小单位`B`转换为`K,M,G,T,P,E`，并支持返回格式化后的字符串

- 以`KiB,MiB,GiB,TiB,PiB,EiB`为单位（1024进制
- 以`KB,MB,GB,TB,PB,EB`为单位（1000进制
- 以`K,M,G,T,P,E`为单位(转换为`iB`或`B`时，按照1024进制处理
- 将B单位的值格式化到各个单位（如`1025B`返回`1KiB 1B`，返回`string`或`[]*SizeFormatModel`
- 截取最大的单位的大小（如1M2B返回1M

```go
Size(1*SizeB + 2*SizeKiB).Format(",", true)
Size(1*SizeB + 2*SizeKB).Format(",", true)
```

# Path Split

三种方法分割路径字符串

- `SplitPathSkip(p string, skip int) (head, tail string)` 跳过skip个`/`后，在下一个`/`处分割，且`/`保留在tail中
- `SplitPath(p string) (head, tail string)` 在第一个`/`处分割，且`/`保留在tail中
- `SplitPathRepeat(p string, repeat int) (head, tail string)` 在第一个`/`处分割，且`/`保留在tail中，将tail作为参数再次分割，重复repeat次

# Random String, Key

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
- `RandomKey(l int) *KeyResult` 生成一个长度为l的key的byte数组, 返回包含key的结构体
- `KeyResultFromHexString(hx string) *KeyResult` 解析hex字符串，转换为KeyResult

KeyResult下的方法

- `.Bytes` 返回key的byte数组
- `.Upper` 返回hex后的大写字符串，长度为Bytes的两倍
- `.Lower` 返回hex后的小写字符串，长度为Bytes的两倍
- `.Error`

# IP

- `IsIPAddr` 判断是否是ip地址
- `IsLocalIPAddr` 判断ip地址是否是本地ip
- `IsLocalIP` 判断ip是否是本地ip
- `IPAddrToUint32` ip地址转为uint32
- `Uint32ToIPAddr` uint32转为ip地址
- `IPToUint32` ip转为uint32
- `Uint32ToIP` uint32转为ip
- `GetLocalIp` 获取本地IP

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

# Port

获取可用端口

- `GetAvailablePort() (int, error)`

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
- `HTTPContentTypeUrlencodedUTF8`
- `HTTPContentTypeJson`
- `HTTPContentTypeJsonUTF8`
- `HTTPContentTypeXml`
- `HTTPContentTypeXmlUTF8`
- `HTTPContentTypePlain`
- `HTTPContentTypePlainUTF8`
- `HTTPContentTypeHtml`
- `HTTPContentTypeHtmlUTF8`
- `HTTPContentTypeFormData`
- `HTTPContentTypeFormDataUTF8`

方法

- `HttpGet(u string, args any)` u为url，args为可序列化变量
- `HttpPost(contentType string, u string, args any)`
- `HttpPostGet(contentType string, u string, argsGET, argsPOST any)` 

