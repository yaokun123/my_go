rpc是自带的包：net/rpc

RPC相关函数：

1、func RegisterName(name string, rcvr interface{}) error
参数1：服务名
参数2：对应rpc对象（就是一个类/struct），该对象绑定的所有方法要满足如下条件
    1、方法必须是导出的  --  包外可见，首字母大写
    2、方法必须有两个参数，都是导出类型、内建类型
    3、方法的第二个参数必须是指针（传出参数）
    4、方法只有一个error类型的返回值
举例：
    type World struct{}
    func (this *World) HelloWorld(name string, resp *string) error{}
    rpc.RegisterName("hello", new(World))

2、func ServeConn(conn io.ReadWriteCloser)   绑定rpc服务
conn就是accept的连接

3、调用远程函数
func (client *Client) Call(serviceMethod string, args interface{}, reply interface) error
serviceMethod:服务名.方法名
args:传入参数
reply：传出参数