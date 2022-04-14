下载 https://github.com/istio/istio/releases/tag/1.9.9

安装  `./istioctl manifest apply --set profile=demo`

因为是本地测试，修改`istio-ingressgateway`服务为`NodePort`类型

创建一个学习用的命名空间 `kubectl create ns myistio`

### 字段注入istio

显示命名空间是否开启自动注入 `kubectl get namespace -L istio-injection`

给命名空间开启自动注入: `kubectl label namespace myistio istio-injection=enabled`

### 部署kiali

`kubectl apply -f samples/addons/kiali.yaml`

创建访问代理 `kubectl apply -f kiali.yaml`

访问`http://kiali.jtthink.com:32515/`

### 部署service

部署: `kubectl apply -f prodv1.yaml && kubectl apply -f reviewv1.yaml`

服务: `kubectl apply -f prod-svc.yaml`

安装网关和url路径虚拟服务代理 `kubectl apply -f prod-gateway-vs.yaml`

测试访问是否成功`http://prod.jtthink.com:32515/p/12` 或 `http://prod.jtthink.com:32515/r/12`

### prod v2 版本

部署 `kubectl apply -f prodv2.yaml`

连续访问多次 `http://prod.jtthink.com:32515/p/12` 返回不同版本

部署限定版本`rule`: `kubectl apply -f prodv2-rule.yaml`

请求 `http://prod.jtthink.com:32515/v2/p/12` 永远返回`v2`

### prod 轮序

`kubectl apply -f prod-rule-round.yaml`

请求 `http://prod.jtthink.com:32515/simple/12`

由于默认是轮序 所以结果与 `http://prod.jtthink.com:32515/p/12` 是一样的

### 一致性 hash

用于一个用户请求只到一个固定的节点

请求 `http://prod.jtthink.com:32515/hash/12` `header` 加入 `myname` 相同的值的请求会访问同一个节点

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

生成token内容 `cd cert && go run main.go `，粘贴到 `yamls/jtw-test.yaml` 内:

运行 `kubectl apply -f yamls/jtw-test.yaml`

此时访问`http://prod.jtthink.com:32515/p/12` 会没有权限

加入请求`token`，访问 `jtw.io`，选择`RS256`加密方式,`VERIFY SIGNATURE`粘贴`cert/mypub.pem`和`cert/myrsa.pem`，`payload`
内新增`"iss":"user@jtthink.com"`
，获取生成`token`用于请求

`curl -H "Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6InpoYW5nc2FuIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoidXNlckBqdHRoaW5rLmNvbSJ9.T1CLmZQMm1c9LIvLxxVODdGR6rKthVFB67wlArc667O91w-cKRXNGQTSmFrLkhnkF5CDMIj3cNwX4OeVAaIIMEiLF2VNYx-YTfwdg3mPrsBI9JlVIjmCTd6TkqNK_6yDtg2HNp_hQKazFn_2wVzmfPJnsMqxTnwFtg_vz7EFwsMKIrjLOPFK6NY7SKCTtVsFOQfZypmsI5hcpVXRmSh7i01DCPAmxfYzOaOdz3qMS63W6UWHuMfDmJxfP-ehqcb2Fkwq76rbSYOVEVq0_U_O7JokGv3DeHDxiM5yMBErgz-5TujBlpovqw_OaIytsWiDwzEErIo0cPnSr9tlZL_VFg" http://prod.jtthink.com:32515/p/12`

成功访问

### 新增跨域

部署一个允许跨域的路径 `kubectl apply -f yamls/prod-cross-vs.yaml`

打开 `html/index.html` 访问地址 `http://prod.jtthink.com:32515/cross-p/12`

如果`token`错误无法返回错误信息，因为在未加入跨域头时，`JWT`就验证报错打回来了，需要新增一个就算错误也要加入跨域头功能 `kubectl apply -f yamls/prod-cross-vs.yaml`




 
