# golangLearning #

已完成：

1. 4 个协程
2. 解析 html 获取链接
3. 最大深度3
4. 结果转换为 protobuf 格式

进行中：

1. 接受 gzip 数据。
    - 解决400 Bad Request
2. sigterm中断程序
    - 遇到的问题： 只有 ^C 接收到了信号，SIGTERM 信号的 kill -15 pid 指令没有发生变化
3. 已经完成的代码符合 go vet
4. 结果存储为 protobuf 格式的 csv 文件
    - 写入部分要修改，too many open files in system

未完成

1. warn 提示
2. 网络读取错误 3 次重试
3. go test
4. 接收到信号保存当前抓取的进度
5. 下次重开的时候继续
6. 读取数据是错误的重新开始即可，但是要有个 warn 提示 ？