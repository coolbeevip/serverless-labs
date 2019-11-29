# OpenFaaS : Quarkus Native Application

## 创建一个 Quarkus 应用

>  Quarkus Application 可以参考 https://quarkus.io/guides/getting-started

启动服务

```bash
$ cd function-quarkus-helloworld
./mvnw compile quarkus:dev
```

测试服务

```bash
$ curl -H "Content-Type:application/json" -X POST -d '{"text": "Hello"}' http://127.0.0.1:9000
{"text":"Hi,I'm OpenFaaS. I have received your message 'Hello'"}
```

## 编译 Native Docker Images

多阶段编译本地应用镜像

```bash
$ docker build -f src/main/docker/Dockerfile.multistage -t coolbeevip/openfaas-function-quarkus-helloworld .
```

启动本地镜像（started in 0.008s 毫秒级启动 🏃‍♀️🏃🏃‍♀️🏃🏃‍♀️🏃）

```bash
$ docker run -p 8080:8080 coolbeevip/openfaas-function-quarkus-helloworld:latest
Forking - /work/application []
2019/11/25 14:22:12 Started logging stderr from function.
2019/11/25 14:22:12 Started logging stdout from function.
2019/11/25 14:22:12 OperationalMode: http
2019/11/25 14:22:12 Timeouts: read: 25s, write: 25s hard: 20s.
2019/11/25 14:22:12 Listening on port: 8080
2019/11/25 14:22:12 Writing lock-file to: /tmp/.lock
2019/11/25 14:22:12 Metrics listening on port: 8081
2019/11/25 14:22:12 stdout: 2019-11-25 14:22:12,592 INFO  [io.quarkus] (main) function-quarkus-helloworld 1.0.0-SNAPSHOT (running on Quarkus 0.25.0) started in 0.008s. Listening on: http://0.0.0.0:9000
2019/11/25 14:22:12 stdout: 2019-11-25 14:22:12,592 INFO  [io.quarkus] (main) Profile prod activated. 
2019/11/25 14:22:12 stdout: 2019-11-25 14:22:12,592 INFO  [io.quarkus] (main) Installed features: [cdi, resteasy]
```

测试镜像服务

```bash
$ curl -H "Content-Type:application/json" -X POST -d '{"text": "Hello Docker"}' http://127.0.0.1:8080
{"text":"Hi,I'm OpenFaaS. I have received your message 'Hello Docker'"}
```

## 上传镜像到 DockerHub

在 DockerHub上创建仓库 openfaas-function-quarkus-helloworld

登录 Docker Hub

```bash
$ docker login
```

上传镜像

```bash
$ docker push coolbeevip/openfaas-function-quarkus-helloworld
```

## 部署到 OpenFaaS

登录 OpenFaaS

```bash
$ export OPENFAAS_URL=http://127.0.0.1:31112
$ faas-cli login --password 78b1d4c29831bbd9040d2ffe6da2c9b9c7845bf2
```

部署

```bash
$ faas-cli deploy stack.yml 
Deploying: quarkus-helloworld.
WARNING! Communication is not secure, please consider using HTTPS. Letsencrypt.org offers free SSL/TLS certificates.

Deployed. 202 Accepted.
URL: http://127.0.0.1:31112/function/quarkus-helloworld
```

查看部署UI

> 在浏览器打开 http://127.0.0.1:31112

可以看到 quarkus-helloworld 已经部署完毕，选择 JSON 后点击 INVOKE 按钮可以测试函数

![image-20191126175756472](assets/image-20191126175756472.png)

打开 K8s 控制台可以看到函数已经部署并拉起

![image-20191126180606517](assets/image-20191126180606517.png)

命令行测试

> 在命令行中输入以下命令进行测试

```bash
$ curl -H "Content-Type:application/json" -X POST -d '{"text": "Hello Serverless"}' http://127.0.0.1:31112/function/quarkus-helloworld
{"text":"Hi,I'm OpenFaaS. I have received your message 'Hello Serverless'"}
```

压力测试

> 部署三节点K8S集群，模拟 100 并发调用 50万次

```bash
$ echo '{"text": "Hello Serverless"}' > data.json
$ ab -n 500000  -c 100 -p 'data.json' -T 'application/json' http://10.22.1.191:31112/function/quarkus-helloworld
This is ApacheBench, Version 2.3 <$Revision: 1430300 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 10.22.1.191 (be patient)
Completed 50000 requests
Completed 100000 requests
Completed 150000 requests
Completed 200000 requests
Completed 250000 requests
Completed 300000 requests
Completed 350000 requests
Completed 400000 requests
Completed 450000 requests
Completed 500000 requests
Finished 500000 requests


Server Software:
Server Hostname:        10.22.1.191
Server Port:            31112

Document Path:          /function/quarkus-helloworld
Document Length:        75 bytes

Concurrency Level:      100
Time taken for tests:   131.603 seconds
Complete requests:      500000
Failed requests:        0
Write errors:           0
Total transferred:      148500000 bytes
Total body sent:        97000000
HTML transferred:       37500000 bytes
Requests per second:    3799.31 [#/sec] (mean)
Time per request:       26.321 [ms] (mean)
Time per request:       0.263 [ms] (mean, across all concurrent requests)
Transfer rate:          1101.95 [Kbytes/sec] received
                        719.79 kb/s sent
                        1821.74 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1  10.0      1    1005
Processing:     1   24 156.0     20   16208
Waiting:        1   24 156.0     19   16208
Total:          2   26 156.3     21   16211

Percentage of the requests served within a certain time (ms)
  50%     21
  66%     24
  75%     27
  80%     29
  90%     34
  95%     39
  98%     47
  99%     55
 100%  16211 (longest request)
```

在K8S控制台可以看到副本自动增加

![image-20191126180934949](assets/image-20191126180934949.png)

观察日志可以看到单个镜像服务启动耗时22毫秒

![image-20191126181124820](assets/image-20191126181124820.png)












