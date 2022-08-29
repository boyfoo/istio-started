prods-build:
	docker run --rm -it -v $(pwd):/app -w /app -e CGO_ENABLED=0 -e GO111MODULE=on golang:1.15-alpine3.12 go build -o build/prods prod/main.go

docker-prods-build:
	docker build -t prods:1.1 -f ./Dockerfile-prods .

docker-prods-push:
	docker tag prods:1.1 registry.cn-hangzhou.aliyuncs.com/boyfoo/prod:1.1 && docker push registry.cn-hangzhou.aliyuncs.com/boyfoo/prod:1.1




review-build:
	docker run --rm -it -v $(pwd):/app -w /app -e CGO_ENABLED=0 -e GO111MODULE=on golang:1.15-alpine3.12 go build -o build/review review/main.go

docker-review-build:
	docker build -t review:1.1 -f ./Dockerfile-review .

docker-review-push:
	docker tag review:1.1 registry.cn-hangzhou.aliyuncs.com/boyfoo/review:1.1 && docker push registry.cn-hangzhou.aliyuncs.com/boyfoo/review:1.1




prodsv2-build:
	docker run --rm -it -v $(pwd):/app -w /app -e CGO_ENABLED=0 -e GO111MODULE=on  golang:1.15-alpine3.12 go build -o build/prodsv2 prodv2/main.go
docker-prodsv2-build:
	docker build -t prods:2.1 -f ./Dockerfile-prodsv2 .

docker-prodsv2-push:
	docker tag prods:2.1 registry.cn-hangzhou.aliyuncs.com/boyfoo/prod:2.1 && docker push registry.cn-hangzhou.aliyuncs.com/boyfoo/prod:2.1



gprods-build:
	docker run --rm -it -v $(pwd):/app -w /app -e CGO_ENABLED=0 -e GO111MODULE=on -e GOPROXY=https://goproxy.cn,direct golang:1.15-alpine3.12 go build -o build/gprdos gProdService.go

docker-gprods-build:
	docker build -t gprods:1.0 -f ./Dockerfile-gprods .
docker-gprods-push:
	docker tag gprods:1.0 registry.cn-hangzhou.aliyuncs.com/boyfoo/gprods:1.0 && docker push registry.cn-hangzhou.aliyuncs.com/boyfoo/gprods:1.0