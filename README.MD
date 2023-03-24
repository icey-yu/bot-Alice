# bot-Alice
高性能ですから!

开启ChatGPT服务实例：
```bash
ChatGPT-API-server 8848 Alice
```

### 开发注意事项
- 如果要在`Reply`类型的消息的开始部位加入`At`消息，需要加入两次才能生效。 
如：`sendingMessage.Elements = append(sendingMessage.Elements, at, reply, at, msgStr)`或
`sendingMessage.Elements = append(sendingMessage.Elements, reply, at, at, msgStr)`