#### MOA Client使用方式

* 安装：
     安装ZooKeeper
        $Zookeeper/bin/zkServer.sh start

    运行 sh build.sh
    
    运行 ./demo-server
    
    运行 ./demo-client
    
    得到结果
    
    result:error:%!s(<nil>),resp:moa
    
    或使用redis客户端调用
      
    redis-cli -p 13800 get '{"action":"/service/bibi/go-moa","params":{"m":"getName","args":["a"]}}'

    结果：
    
    "{\"ec\":0,\"em\":\"\\u003cnil\\u003e\",\"result\":\"moa\"}"

