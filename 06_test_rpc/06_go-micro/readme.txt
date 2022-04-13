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



=================================防坑指南==============================

首先区分，Micro 3.0 和 go micro v3：
1、Micro 3.0 是开发go micro 架构的公司的一个云原生开发平台，付费，帮你维护微服务项目，只需要专注业务开发
2、go micro 是一个微服务架构

安装依赖
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go

//注意 不要安装下面这个，下面是Micro3.0 云平台的安装工具
//go get github.com/micro/micro/v3/cmd/protoc-gen-micro

//安装asim/go-micro   这个是go micro 3.0 框架
go get github.com/asim/go-micro/cmd/protoc-gen-micro/v3



安装micro v3：
//需要用到Micro 3.0 环境的micro 工具，可以快速构建项目，但是不使用这个库，用下面的
go get github.com/micro/micro/v3

// 下载go micro 3.0 库，下面库没有上面micro 工具
go get github.com/asim/go-micro/v3

说明：asim/go-micro 比 micro/micro更加完整，多了很多我们开发所用的插件还有方法
例如 asim/go-micro 支持consul，而micro/micro 因为是个云平台，不需要你去考虑注册问题，所以没有


#安装成功检查
##检查GOPATH/bin 目录下是否有3个工具
micro protoc-gen-go  protoc-gen-micro



创建服务：
// 因为go get github.com/micro/micro/v3，所有可以执行micro命令
micro new helloworld
cd helloworld
make proto


如果使用 go get github.com/micro/micro/v3/cmd/protoc-gen-micro
使用    go get github.com/asim/go-micro/cmd/protoc-gen-micro/v3（使用这种）



修改main.go：
//"github.com/micro/micro/v3/service"
 //"github.com/micro/micro/v3/service/logger"
 service "github.com/asim/go-micro/v3"
 "github.com/asim/go-micro/v3/logger"


 // Create service
     // micro/micro 创建服务
  //srv := service.New(
  //  service.Name("helloworld"),
  //  service.Version("latest"),
  //)
     // asim/go-micro 创建服务
  srv := service.NewService(
          service.Name("helloworld"),
          service.Version("latest"),
      )
