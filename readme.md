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
