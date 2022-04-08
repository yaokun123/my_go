go-micro    // 核心库
micro       // 微服务运行环境，命令，创建微服务空项目
go-plugins  // 插件
protoc-gen-micro    // 生成micro相关代码


安装go-micro: go get -u -v github.com/micro/go-micro
安装工具集：    go get -u -v giehub.com/micro/micro

安装protobuf插件
go get -u github.com/golang/protobuf/{proto,proto-gen-go}
go get -u github.com/micrro/protoc-gen-micro


micro new
--namespace 命名空间 = 包名
--type  要创建的微服务类型
    srv 微服务
    web 基于微服务的web项目
    api api项目
示例：
micro new --type srv bj38


micro有默认的服务发现mdns,可以更换为consul