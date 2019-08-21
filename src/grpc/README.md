protoc --go_out=plugins=grpc:. route_guide.proto
运行这个命令可以在当前目录中生成下面的文件：
route_guide.pb.go
这些包括：
所有用于填充，序列化和获取我们请求和响应消息类型的 protocol buffer 代码
一个为客户端调用定义在RouteGuide服务的方法的接口类型（或者 存根 ）
一个为服务器使用定义在RouteGuide服务的方法去实现的接口类型（或者 存根 ）

让 RouteGuide 服务工作有两个部分：
实现我们服务定义的生成的服务接口：做我们的服务的实际的“工作”。
运行一个 gRPC 服务器，监听来自客户端的请求并返回服务的响应。