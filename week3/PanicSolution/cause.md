```
if len(consumeMSG) >= ConsumeNum {
          //进行异步消费
          go func() {
             m := consumeMSG[:ConsumeNum]
             fn(m)
          }()
          // 更新上次消费时间
          lastConsumeTime = time.Now()
          // 清除插入的数据
          consumeMSG = consumeMSG[ConsumeNum:]
```

--
**异步消费处理和更新消费时间，清除插入数据不同步，消费在新 goroutine，更新，插入在主 goroutine**
另一情况同上
