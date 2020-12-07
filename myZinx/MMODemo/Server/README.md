# Breeze 服务器端

## MsgId 约定

| 消息id  | 消息功能  | 服务器端函数 | 数据结构 | 
| ----   | ----     |  ---- | ---- |
| 1      |   服务器端告知客户端玩家的ID         |   syncPid  | |
| 2      |   服务器端告知客户端玩家的pos         |   syncPos  | |
| 201    |   客户端发送移动请求和位置信息         |   moveRouter.Handle | |
| 3      |   服务器端广播位置信息，接收201后调用   |   SyncOtherPos | |
