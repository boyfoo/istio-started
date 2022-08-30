### 安装

下载 https://github.com/istio/istio/releases/tag/1.9.9

安装  `./istioctl manifest apply --set profile=demo`

因为是本地测试，修改`istio-ingressgateway`服务为`NodePort`
类型 `kubectl patch service istio-ingressgateway -n istio-system -p '{"spec":{"type":"NodePort"}}'`

获取地址`kubectl get po -l istio=ingressgateway -n istio-system -o 'jsonpath={.items[0].status.hostIP}'`
和端口`kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].nodePort}'`

创建一个学习用的命名空间 `kubectl create ns myistio`

### 字段注入istio

显示命名空间是否开启自动注入 `kubectl get namespace -L istio-injection`

给命名空间开启自动注入: `kubectl label namespace myistio istio-injection=enabled`

### 部署kiali

`kubectl apply -f samples/addons/kiali.yaml`

创建访问代理 `kubectl apply -f kiali.yaml`

访问`http://kiali.jtthink.com:32515/`

安装`prometheus`，使用示例中提供的 `kb apply -f samples/addons/prometheus.yaml`

### 部署service

部署: `kubectl apply -f prodv1.yaml && kubectl apply -f reviewv1.yaml`

服务: `kubectl apply -f prod-svc.yaml`

安装网关和url路径虚拟服务代理 `kubectl apply -f prod-vs.yaml`，将路径`/p`重写成`/prods`，并代理访问

测试访问是否成功`http://prod.jtthink.com:32515/p/12` 或 `http://prod.jtthink.com:32515/r/12`

### prod v2 版本

部署 `kubectl apply -f prodv2.yaml`

连续访问多次 `http://prod.jtthink.com:32515/p/12` 返回不同版本

原因是`prod-svc.yaml`创建`svc`选择容器时，并没有指定版本标签`version`，所以会轮训所有容器，造成多版本交叉，这种情况就算不使用`istio`也是多版本

部署限定版本`rule`: `kubectl apply -f prodv2-rule.yaml`

请求 `http://prod.jtthink.com:32515/v2/p/12` 永远返回`v2`

### 访问服务prod随机选择pod

`kubectl apply -f prod-rule-round.yaml`

请求 `http://prod.jtthink.com:32515/simple/12`

请求后，会出现连续多次访问出同一个版本的情况，这个在之前默认轮训下是不会发生的

另外 `http://prod.jtthink.com:32515/p/12` 也会变成随机，原因文件内有备注解释

### 清理资源

为了防止后续有冲突，清理下暂时不在当实例的资源

`kubectl delete -f prodv2-rule.yaml`

`kubectl delete -f prod-rule-round.yaml`

### 访问服务一致性hash选择pod

用于一个用户请求只到一个固定的节点

`kubectl apply -f prod-rule-hash.yaml`

请求 `http://prod.jtthink.com:32515/hash/12` `header` 加入 `myname` 相同的值的请求会访问同一个节点

实例体验后删除 `kubectl delete -f prod-rule-hash.yaml`

### 故障注入

延迟`kubectl apply -f prod-vs-err-delay.yaml`

访问`http://prod.jtthink.com:32515/delay/12` 延迟3s相应

http500报错 `kubectl apply -f prod-vs-err-code.yaml`

访问`http://prod.jtthink.com:32515/errcode/12` 报错

删除体验

`kb delete -f prod-vs-err-delay.yaml`

`kb delete -f prod-vs-err-code.yaml`

之前都是给`prod`注入故障，改为给`review`注入故障 `kb apply -f prod-vs-err-delay-review.yaml`

因为`prod`会访问`review`，请求`http://prod.jtthink.com:32515/p/12`也会延迟效果

删除`kb delete -f prod-vs-err-delay-review.yaml`

### JWT

给`ingressgateway` 网关加入`jwt` 认证

生成秘钥

```
$ cd cert
# 生成私钥
$ openssl genrsa -out myrsa.pem 2048
# 生成公钥
$ openssl rsa -in myrsa.pem -pubout -out mypub.pem
```

生成token内容 `cd cert && go run main.go `，粘贴到 `yamls/jwt-test.yaml` 内:

运行 `kubectl apply -f yamls/jwt-test.yaml`

此时访问`http://prod.jtthink.com:32515/p/12` 会没有权限

加入请求`token`，访问 `jwt.io`，选择`RS256`加密方式,`VERIFY SIGNATURE`粘贴`cert/mypub.pem`和`cert/myrsa.pem`，`payload`
内新增`"iss":"user@jtthink.com"`
，获取生成`token`用于请求

在 `payload` 内加入`exp: 123123123`, 时间戳会自动判断是否过期，目前演示都没加

`curl -H "Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6InpoYW5nc2FuIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoidXNlckBqdHRoaW5rLmNvbSJ9.T1CLmZQMm1c9LIvLxxVODdGR6rKthVFB67wlArc667O91w-cKRXNGQTSmFrLkhnkF5CDMIj3cNwX4OeVAaIIMEiLF2VNYx-YTfwdg3mPrsBI9JlVIjmCTd6TkqNK_6yDtg2HNp_hQKazFn_2wVzmfPJnsMqxTnwFtg_vz7EFwsMKIrjLOPFK6NY7SKCTtVsFOQfZypmsI5hcpVXRmSh7i01DCPAmxfYzOaOdz3qMS63W6UWHuMfDmJxfP-ehqcb2Fkwq76rbSYOVEVq0_U_O7JokGv3DeHDxiM5yMBErgz-5TujBlpovqw_OaIytsWiDwzEErIo0cPnSr9tlZL_VFg" http://prod.jtthink.com:32515/p/12`

成功访问

### 新增跨域

部署一个允许跨域的路径 `kubectl apply -f yamls/prod-cross-vs.yaml`

打开 `html/index.html` 访问地址 `http://prod.jtthink.com:32515/cross-p/12`

如果`token`错误则无法返回错误信息，因为在未加入跨域头时，`JWT`就验证报错打回来了，这个时候还没有完成跨域，需要新增一个就算错误也要加入跨域头功能 `kubectl apply -f yamls/jwt-cross.yaml`
，原理就是将新增跨域的策略移植到`token`验证策略之前

删除旧路由，避免后续演示影响`kubectl delete -f prod-vs.yaml,prod-rule-hash.yaml,prod-rule-round.yaml,prod-cross-vs.yaml`

### 重新部署

部署基础资源 `kb apply -f prod-vs.yaml,jwt-test.yaml,jwt-cross.yaml`

### GRPC

**开发阶段**安装基础工具

`porotoc`工具下载`https://github.com/protocolbuffers/protobuf/releases/tag/v3.12.3`

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

上面的如果执行`proto.bat`报错，可能是`protoc-gen-go-grpc`版本问题，可以尝试此链接
https://github.com/grpc/grpc-go/releases/tag/cmd%2Fprotoc-gen-go-grpc%2Fv1.1.0
下载旧版本试试

体验的直接部署： `cd yamls/grpc && kb apply -f .`

设置一个`host`域名代理到服务器`grpc.jtthink.com`

运行`go run test.go`，会提示权限不足（因为后续新建了新的网关，文件内容都变成了选择了新网关的名称，所以这部分应该不会生效），因为之前设置的`jwt`把所有请求都拦截了，需要新增一个新的`istio-ingressgateway`专门来处理`grpc`请求

使用工具打印出自带网关的组件部分的`yaml`: `./istioctl profile dump --config-path components demo`

将内容保存至`sys/demo.yaml`

为防止有全局影响先删除验证`kb delete -f jwt-test.yaml,jwt-cross.yaml`等都部署完在新增

使用工具按照这个文件`./istioctl install -f sys/demo.yaml`

按照之前安装的情况修改为`nodeport`和端口 (k3s可以不改，因为会自动代理)

本地运行 `go run test.go`

### http证书

将`aliyun`下载好的证书绑定入：

`kb create -n istio-system secret tls istio-ingressgateway-certs --key 8384742_istio.k3s.wiki.key --cert 8384742_istio.k3s.wiki.pem`

执行`kb apply -f yamls/prod-vs-https.yaml`

访问`https://istio.k3s.wiki/p/123`


### Filter

删除其他网关配置

部署`kubectl apply -f prod-vs.yaml`

新增响应头`kb apply -f envoy/testfilter.yaml`

请求 `http://prod.jtthink.com:32515/p/12` 会响应`header`内新增`myname`字段

最新增一个根据`myname`字段，自动添加前缀的字段`mynewname`

`kb apply -f envoy/testfilter-pre.yaml` 必须在`testfilter.yaml`完全部署完才能执行

查看新增的Filter内容`kb exec -n istio-system -it istio-ingressgateway-668fb685db-dg9tv -- curl http://localhost:15000/config_dump\?resource\=dynamic_listeners`

### LUA Filter 日志

进入容器内 `kb exec -n istio-system -it istio-ingressgateway-668fb685db-dg9tv -- sh`

设置日志等级为`info`: `curl -X POST  http://localhost:15000/logging?level=info`

设置`request header`头: `kb apply -f testfilter-adduserid.yaml`

设置打印日志：`kb apply -f testfilter-adduserid-log.yaml`

请求几次，查看日志`kb logs -n istio-system istio-ingressgateway-668fb685db-dg9tv`

### LUA Filter 拦截请求

部署：`kb apply -f testfilter-checkappid.yaml`

此时请求会报错`400`，需要在header头里加`appid`值
