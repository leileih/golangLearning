# golangLearning #

已完成：

1. 4 个协程
2. 解析 html 获取链接
3. 最大深度3 （抓取结束不会自己结束程序）

进行中：

1. 接受 gzip 数据。
    - 已添加request.Header.Add("Accept-Encoding", "gzip")
    - 返回数据要修改
2. 存入文件, 离线缓存
3. 符合 go vet

未完成

1. 离线缓存
2. sigterm中断程序
3. go test