## gim

这是一个**前后端分离** IM （即时通讯）程序**DEMO**版本的**后端部分**应用程序，具备基本的消息收发，用户上、下线通知功能，可以配合各种前端使用。

我也为它写了一个前端应用的 DEMO，使用 Vue.js 开源框架，可以作为你的参考 :)

## How to use

首先你需要知道这个软件的一些内容

- 发送 websocket 连接时携带 `username` 将当前用户加入到用户连接池，前端可自由定义自己的逻辑。

- 将消息分为 系统消息 和 用户消息，使用序列化消息体的 `type` 字段，后端只负责消息处理、分发。

- 可对消息进行无限制扩展，例如图片，视频，红包等类型消息。

在了解了这些特性之后，你需要修改端口号消息，然后在随便一个目录下执行：

```
git clone https://github.com/sh1luo/gim.git
cd gim
go run main.go
```

## Other

我还为它写了一个简单的压力测试程序，用来测试并发环境下的吞吐量，在`backend/cmd/benchmark` 目录下。

## Contact

- [x] email：silo1999@163.com