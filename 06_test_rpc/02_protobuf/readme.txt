protobuf时protocol Buffers的简介，它是Google公司开发的一种描述语言，是一种轻便高效的结构化数据存储格式，
可以用于结构化数据串行化，或者说序列化。它很适合做数据存储或RPC数据交换格式。
可以用于通信协议、数据存储等领域的语言无关、平台无关、可扩展的序列化结构数据格式。

优势：
1、序列化后体积相比于json和XML很小，适合网络传输
2、支持跨平台多语言
3、消息格式升级和兼容性还不错
4、序列化反序列化速度很快，快于json的处理速度


劣势：
1、应用不够广泛（相比json/xml）
2、二进制格式导致可读性差
3、缺乏自描述


protobuf的安装：
brew install protobuf
安装之后可以使用protoc -h检验


protobuf文件的编译：
c:  protoc --cpp_out=./ *.proto
go: protoc --go_out=./ *.proto  ->  这个命令太老了，有新版命令


如果在go语言中使用还要安装go语言的插件protoc-gen-go
github.com/golang/protobuf/protoc-gen-go        旧版本
google.golang.org/protobuf/cmd/protoc-gen-go    新版本


grpc:
google.golang.org/grpc/cmd/protoc-gen-go-grpc
