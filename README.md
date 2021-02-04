# go-service-example
基于“服务”的服务器框架的示例代码
运行例子时请先安装好redis，保证进程能够连上，redis主要做服务发现用。
例子介绍：
- serviceA-call-serviceB 
	两个service通信，相互rpc通信
- client-call-serviceA
	客户端调用service, 相互rpc通信
- entity-call-entity
	客户端登录，创建实体，实体rpc通信，实体属性同步
- matchservice, battleservice暂未完成，不能使用