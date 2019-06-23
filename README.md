## 从零写kafka生产者消费者

```
1. 什么是Kafka?
2. 什么是消息中间件?
3. 消息中间件的应用场景?
4. Kafka优势和劣势?
5. Kafka应用场景?
6. kafka名词解释?
    代理(Broker)
    主题(Topic)
    分区(partition)
    生产者(Producer)
    消费群组(Consumer Group)
    消费者(Consumer)
7. 代理(Broker)
9. 主题(Topic)
9. 分区(Partition)
10. 生产者(Producer)
11. 消费者(Consumer)


```

## 1. 什么是Kafka?
> Kafka是消息中间件的一种,它是一个分布式、可分区、可复制的高性能消息系统,它提供普通消息系统的功能队列、发布订阅
> Kafka使用队列模型时,它将处理工作平均分配给消费组(Consumer Group)中的消费者成员
> Kafka使用发布订阅模式时,它可以将消息广播给多个消费组,也允许消息被多个消费者订阅

## 2. 什么是消息中间件?
> 消息中间件是一种异步通信协议,也可以说试一种通信软件,它可以为分布式应用、系统或者实体提供跨平台、异步、松耦合、高可靠性并且安全的通信功能

## 3. 消息中间件的应用场景?
> 1. 应用程序解耦
> 2. 异步处理
> 3. 流量削

```
1. 应用程序解耦
    比如用户注册送红包发送注册短信,传统做法是先用户注册然后送红包最后发送注册短信,
    使用消息队列用户注册成功后发送1个消息立即返回, 同时发送用户注册和送注册红包
    传统如果用户注册送红包发送注册短信分别需要用时50ms,就需要150ms=(50ms 用户注册 + 50ms 送红包 + 50ms发送注册短信)
    消息中间件用户注册成功后发送1个消息立即返回,同时发送用户注册和送注册红包,就需要100ms=(50ms 用户注册 + 50ms(送红包 + 发送注册短信))
    如果不需要发送短信,停用短信消费者即可,需要时启用短信消费者就好
    如果不需要送红包,停用送红包消费者即可,需要的启用送红包消费者就好
    从而提高吞吐量和应用解耦
    
2. 异步处理
    发送短信、发送邮件、注册送红包等

3. 流量削
    主要用于高并发业务场景,缓解短时间内高流量压垮应用比如秒杀和活动等,将用户请求放入消息队列中,如果消息队列成功最大数量,
    直接抛弃用户请求或跳转到错误页面
```


## 4. kafka优势和劣势
```sh
优势:
    1. 可靠性: kafka是分布式、分区、复制和容错来保存可靠性
        kafka分布式的,一个数据多个副本,少数集群宕机,不会丢失数据,也不会导致不可用
    2. 可扩展性: 可以通过简单的增加服务器横向扩展,无需停机
    3. 持久性: kafka消息持久化(零拷贝zero-copy)到磁盘上,副本机制实现数据沉余,保证数据不会丢失
    4. 高性能: kafka对于发布订阅订阅都具有很高吞吐量
        单机写入TPS约在百万条/秒,即使在非常廉价的商用机器上也能做到单机支持每秒100K条消息的传输
        零拷贝技术(sendfile zero copy)
        分布式存储
        点对点消息传递: 分区(partition)可以在多个节点上(例如: 1个节点1个主题3分区吞吐量是300MB, 分成3个节点分别存放个主题3分区即300MB*3=900MB的吞吐量)
        磁盘顺序读顺序写
        批量读批量写
        kafka不保存消息状态,消费者通过记录主题分区中的offset表示之前的消息都被消费,offset之后的都是未消费,offset可以任意移动
        
        
缺点:
    消息乱序: Kafka只能保证单个partition(分区)内部的消息有序,如果单个topic(主题)有多个分区,分区之间无法保证消息有序
    重复消息: Kafka保证每条消息至少送达一次,虽然几率很小,但一条消息可能被送达多次.
    复杂性: Kafka需要Zookeeper的支持,主题(Topic)一般需要人工创建,部署和维护比一般MQ成本更高

```
