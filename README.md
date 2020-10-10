## Insant Message backend(IM)

这是一个**前后端分离** IM （即时通讯）程序**DEMO**版本的**后端部分**应用程序，具备基本的消息收发，用户上、下线通知功能，易扩展。

可以配合前端使用，移步 [前端地址](https://github.com/sh1luo/im_front)

## How to use

首先你需要知道这个软件的一些内容

- 发送 websocket 连接时携带 `username` 将当前用户加入到用户连接池，前端可自由定义自己的逻辑。

- 将消息分为 系统消息 和 用户消息，使用序列化消息体的 `type` 字段，后端只负责消息处理、分发。

- 消息类型的可扩展性强，无限制扩展例如 图片，视频，红包等类型消息。

在了解了这些特性之后，你需要修改端口号消息，然后在随便一个目录下执行：

```
git clone https://github.com/sh1luo/im_backend.git
cd im_backend
go run main.go
```

## Contact

- [x] QQ：690898835
- [x] email：silo1999@163.com