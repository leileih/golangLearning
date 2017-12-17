# golangLearning #

已完成：

1. 4 个协程
2. 解析 html 获取链接
3. 最大深度3

进行中：

1. 接受 gzip 数据。
    - 解决400 Bad Request
2. sigterm中断程序
    - 遇到的问题： 接受信号以后直接中断程序
                s := <-c //阻塞直至有信号传入，阻塞之后的代码不执行, 官网的例子也运行不成功
3. 已经完成的代码符合 go vet

未完成

1. 离线缓存， protobuf 编码，warn 提示
2. go test