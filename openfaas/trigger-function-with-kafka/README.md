# 使用 KAFKA 消息触发 Function

OpenFaaS支持在Kafka和Function建立订阅关系，除了通过网关使用HTTP访问Function外，你还可以使用发送Kafka消息触发Function

## 在 K8S 上部署 KAFKA

下载仓库代码

```bash
$ git clone https://github.com/openfaas-incubator/kafka-connector && cd kafka-connector/yaml/kubernetes
```

部署kafka

```bash
$ kubectl apply -f kafka-broker-dep.yml,kafka-broker-svc.yml
```

部署zookeeper

```bash
$ kubectl apply -f zookeeper-dep.yaml,zookeeper-svc.yaml
```

## 部署 KAFKA CONNECTOR

Kafka Connector 将Kafka主题连接到OpenFaaS功能。部署kafka-connector 并将其指向您的代理后，您可以通过在函数的 stack.yml 文件中添加简单的注释来将函数连接到主题。

```bash
$ kubectl apply -f connector-dep.yml
```

发送一个消息验证一下

获取 kafka 容器名称

```bash
$ BROKER=$(kubectl get pods -n openfaas -l component=kafka-broker -o name|cut -d'/' -f2)
```

启动命令行 kafka-console-producer，并发送一个消息

```bash
$ kubectl exec -n openfaas -t -i $BROKER -- /opt/kafka_2.12-0.11.0.1/bin/kafka-console-producer.sh --broker-list kafka:9092 --topic faas-request
>hello world
```

启动命令行 kafka-console-consumer，查看上一步发送的消息

```bash
$ kubectl exec -n openfaas -t -i $BROKER -- /opt/kafka_2.12-0.11.0.1/bin/kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic faas-request --from-beginning
```

## Function 绑定 Topic

修改Function绑定Topic faas-request [function-golang-http-helloworld](https://github.com/coolbeevip/serverless-labs/blob/master/openfaas/function-golang-http-helloworld)

修改  [stack.yml](https://github.com/coolbeevip/serverless-labs/blob/master/openfaas/function-golang-http-helloworld/stack.yml) 增加 topic annotation

```yaml
version: 1.0
provider:
  name: faas
  gateway: http://10.22.1.191:31112
functions:
  golang-http-helloworld:
    lang: Dockerfile
    skip_build: true
    image: coolbeevip/openfaas-function-golang-http-helloworld:latest
    annotations:
      topic: faas-request
```

重新部署

```bash
$ faas-cli deploy stack.yml
```

## 测试

发送 {"text":"Hello Topic"}

```bash
$ BROKER=$(kubectl get pods -n openfaas -l component=kafka-broker -o name|cut -d'/' -f2)
$ kubectl exec -n openfaas -t -i $BROKER -- /opt/kafka_2.12-0.11.0.1/bin/kafka-console-producer.sh --broker-list kafka:9092 --topic faas-request

>{"text":"Hello Topic"}
```

查看 kafka connector 日志

```bash
$ CONNECTOR=$(kubectl get pods -n openfaas -o name|grep kafka-connector|cut -d'/' -f2) & kubectl logs -n openfaas -f --tail 100 $CONNECTOR
```

可以看到消息已经收到并触发了golang-http-helloworld

```bash
2019/12/03 06:21:01 Syncing topic map
[#4] Received on [faas-request,0]: '{"text":"Hello Topic"}'
2019/12/03 06:21:04 Invoke function: golang-http-helloworld
2019/12/03 06:21:04 connector-sdk got result: [202] faas-request => golang-http-helloworld (0) bytes
[202] faas-request => golang-http-helloworld

2019/12/03 06:21:04 Syncing topic map
2019/12/03 06:21:07 Syncing topic map
```

