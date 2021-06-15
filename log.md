# Memo

1. 消息格式约定(`go struct`格式)

   ````go
   type Message struct {
     sender string
     channel string
     kind int64
     payload []byte
     timestamp time.Timestamp
   }
   
   const (
     text = iota
     picture
     file
     createChannel
   )
   ````

   `payload`字段根据不同的信息类型具有不同语义。

   `message`统一存储到`vuex`中。

   在前端代码中，`message`还需要一个`hash`字段，使用hash函数计算得出，用于`v-for`指令的`key`值。

2. `channel`定义

   ```go
   type Channel struct {
     name string
     subscribers map[string]*Client
   }
   ```

   

   
