下载 https://github.com/istio/istio/releases/tag/1.9.9

安装  `./istioctl manifest apply --set profile=demo`

创建一个学习用的命名空间 `kubectl create ns myistio`

下载脚手架 `go get -u github.com/shenyisyn/goft-gin@v0.5.0`

显示命名空间是否开启自动注入 `kubectl get namespace -L istio-injection`

给命名空间开启自动注入: `kubectl label namespace myistio istio-injection=enabled`

部署: `kubectl apply -f app.yaml`

因为是本地测试，修改`istio-ingressgateway`服务为`NodePort`类型

安装代理 `kubectl apply -f ingress.yaml`

测试访问是否成功`http://review.jtthink.com:32515/review/12`

### 安装kiali

`kubectl apply -f samples/addons/kiali.yaml`

创建访问代理 `kubectl apply -f kiali.yaml`

访问`http://kiali.jtthink.com:32515/`