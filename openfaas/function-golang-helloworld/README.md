# OpenFaaS : Golang Function

## 创建一个 Golang 应用

编译

```bash
$ cd function-golang-helloworld
$ go build cmd/dockerd/main.go 
```

测试

```bash
$ echo 'hello' | ./main
Hi,I'm OpenFaaS. I have received your message 'hello'
```

## 编译 Docker Images

多阶段编译本地应用镜像

```bash
$ docker build -t coolbeevip/openfaas-function-golang-helloworld .
```

启动本地镜像

```bash
$ docker run -p 8080:8080 coolbeevip/openfaas-function-golang-helloworld:latest
2019/11/29 03:26:37 OperationalMode: streaming
2019/11/29 03:26:37 Timeouts: read: 10s, write: 10s hard: 10s.
2019/11/29 03:26:37 Listening on port: 8080
2019/11/29 03:26:37 Writing lock-file to: /tmp/.lock
2019/11/29 03:26:37 Metrics listening on port: 8081
```

测试镜像服务

```bash
$ curl -H "Content-Type:application/json" -X POST -d '{"text": "Hello Docker"}' http://127.0.0.1:8080
{"text":"Hi,I'm OpenFaaS. I have received your message 'Hello Docker'"}
```

## 上传镜像到 DockerHub

在 DockerHub上创建仓库 openfaas-function-golang-helloworld

登录 Docker Hub

```bash
$ docker login
```

上传镜像

```bash
$ docker push coolbeevip/openfaas-function-golang-helloworld
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
Deploying: golang-helloworld.
WARNING! Communication is not secure, please consider using HTTPS. Letsencrypt.org offers free SSL/TLS certificates.

Deployed. 202 Accepted.
URL: http://127.0.0.1:31112/function/golang-helloworld
```

测试

> 在命令行中输入以下命令进行测试

```bash
$ curl -H "Content-Type:application/json" -X POST -d '{"text": "Hello Serverless"}' http://127.0.0.1:31112/function/golang-helloworld
{"text":"Hi,I'm OpenFaaS. I have received your message 'Hello Serverless'"}
```

压力测试

> 部署三节点K8S集群，模拟 100 并发调用 50万次

```bash
$ echo '{"text": "Hello Serverless"}' > data.json
$ ab -n 50000  -c 50 -p 'data.json' -T 'application/json' http://10.22.1.191:31112/function/golang-helloworld
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

Document Path:          /function/golang-helloworld
Document Length:        78 bytes

Concurrency Level:      100
Time taken for tests:   536.831 seconds
Complete requests:      500000
Failed requests:        0
Write errors:           0
Total transferred:      139000000 bytes
Total body sent:        96500000
HTML transferred:       39000000 bytes
Requests per second:    931.39 [#/sec] (mean)
Time per request:       107.366 [ms] (mean)
Time per request:       1.074 [ms] (mean, across all concurrent requests)
Transfer rate:          252.86 [Kbytes/sec] received
                        175.55 kb/s sent
                        428.40 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1  10.9      0    1004
Processing:     2  106  41.6    100    2087
Waiting:        2  106  41.6    100    2087
Total:          3  107  43.1    100    2111

Percentage of the requests served within a certain time (ms)
  50%    100
  66%    103
  75%    104
  80%    106
  90%    112
  95%    120
  98%    184
  99%    403
 100%   2111 (longest request)
```
