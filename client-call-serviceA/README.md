### go-service介绍
  Go Service是一个使用go语言开发的分布式游戏服务器、app服务器框架，特点是开发效率高，方便上手。
### go-service的功能
- 可用vscode单步调试的分布式服务端，N变1

  * 一般来说，分布式服务端要启动很多进程，一旦进程多了，单步调试就变得非常困难，导致服务端开发基本上靠打log来查找问题。平常开发游戏逻辑也得开启一大堆进程，不仅启动慢，而且查找问题及其不方便，要在一堆堆日志里面查问题，这感觉非常糟糕，go-service框架使用了service设计，所有服务端内容都拆成了一个个service，启动时根据服务器类型挂载自己所需要的service。
- 随意可拆分功能的分布式服务端，1变N
  * 分布式服务端要开发多种类型的服务器进程，比如Login server，gate server，battle server，chat server friend server等等一大堆各种server，传统开发方式需要预先知道当前的功能要放在哪个服务器上，当功能越来越多的时候，比如聊天功能之前在一个中心服务器上，之后需要拆出来单独做成一个服务器，这时会牵扯到大量迁移代码的工作，烦不胜烦。go-service框架在平常开发的时候根本不太需要关心当前开发的这个功能会放在什么server上，只用一个进程进行开发，功能开发成组件的形式。发布的时候使用一份多进程的配置即可发布成多进程的形式，是不是很方便呢？随便你怎么拆分服务器。只需要修改极少的代码就可以进行拆分。不同的server挂上不同的组件就行了嘛！

- go语言天生跨平台
  * 提供windows ,linux的一键运行脚本
- go mod支持，能够自动下载需要的三方库，和框架的底层库

### 本例子功能
- client是go语言写的client的例子
- client连接serviceA(app只有servcieA一个service)
- client往服务器发送登录消息，登录之后服务器会创建entity,并且通知client创建entity
- client entity往服务器发送hello的rpc消息, 服务器entity回一个hello的rpc消息

### 例子启动
- 详见启动步骤.md