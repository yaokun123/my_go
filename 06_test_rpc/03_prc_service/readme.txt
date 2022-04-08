添加rpc服务:
service 服务名{
    rpc 函数名(参数 消息体) returns (返回值,消息体)
}


示例：
message People{
    string name = 1;
}
message Student{
    int32 age = 2;
}
service hello{
    rpc HelloWorld(People) return (Student);
}


知识点：
默认，protobuf，编译期间，不编译服务，要想使用编译，需要使用grpc
使用编译指令为：
    protoc --go_out=plugins=grpc:./ *.proto