# EasyIO

自己包装的一个方便读写的模块

## 下载

```go
import "github.com/cn-maul/easyio"
```

## 使用

### 打开文件

**func Open()**

`func Open(path string, mode string) *os.File`

 - path 为要打开的文件路径
 - mode 为模式选择
   - r 只读
   - c 创建文件并打开
   - wa 只写，追加写
   - wc 只写，覆盖写


### 读文件

**func Read()**

`func Read(file io.Reader) []byte`

 - 从 io.Reader 读取文件，返回 []byte 类型的内容


**func ReadFile()**

`func ReadFile(path string) []byte`

 - 从 path 读取文件，返回 []byte 类型的内容
 - path 为要读取文件的路径


### 写文件

**func WriteBytes()**

`func WriteBytes(writer io.Writer, content []byte)`

 - 将 []byte 类型的 content 内容写入到 writer


**func WriteString()**

`func WriteString(writer io.Writer, content string)`

 - 将 string 类型的 content 内容写入到 writer
