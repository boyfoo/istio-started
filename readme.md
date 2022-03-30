下载 https://github.com/istio/istio/releases/tag/1.9.9

安装  `./istioctl manifest apply --set profile=demo`

创建一个学习用的命名空间 `kubectl create ns myistio`

下载脚手架 `go get -u github.com/shenyisyn/goft-gin@v0.5.0`

显示命名空间是否开启自动注入 `kubectl get namespace -L istio-injection`

给命名空间开启自动注入: `kubectl label namespace myistio istio-injection=enabled`

部署: `kubectl apply -f app.yaml`

因为是本地测试，修改`istio-ingressgateway`服务为`NodePort`类型