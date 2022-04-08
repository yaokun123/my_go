服务发现的种类
consul: 常应用于go-micro中
mdns:   go-micro中默认自带的服务发现
etcd:   k8s内嵌的服务发现
zookeeper:java中常用


consul的安装：brew install consul


consul常用的命令：
consul agent
    -dev 以开发者模式启动，不需要额外的配置文件，使用的都是默认的配置
    -bind 指定consul所在机器的IP地址，默认0.0.0.0
    -client=127.0.0.1 表明哪些机器是可以访问consul
    -config-dir=foo 所有主动注册服务的描述信息
    -data-dir=path  存储所有注册过来的srv机器的详细信息
    -node=hostname  服务发现的名字
    -rejoin         consul启动的时候，加入到的consul集群
    -server         以服务方式开启consul，允许其他consul连接开启的consul上（形成集群）
    -bootstrap-expect分布式集群的个数
    -ui             可以使用web页面来查看服务发现的详情
    -http-port=8500 consul自带的一个web访问的端口

示例
consul agent -server -bootstrap-expect 1 -data-dir /tmp/consul -node=n1
-bind=127.0.0.1 -ui -config-dir=/etc/consul.d/ -client 0.0.0.0


basic.json内容： 默认端口分别是:
{
    "ports": {
        "server": 8300
        "serf_lan": 8301,
        "serf_wan": 8302,
        "rpc": 8400,
        "http": 8500 ,
        "dns": 8600,
    }
