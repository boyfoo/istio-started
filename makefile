prods-build:
	docker run --rm -it -v $(pwd):/app -w /app -e CGO_ENABLED=0 -e GO111MODULE=on  golang:1.14-alpine3.12 go build -o build/prods pord/main.go


docker-prods-build:
	docker build -t prods:1.0 -f ./Dockerfile-prods .

docker-prods-push:
	docker tag prods:1.0 registry.cn-hangzhou.aliyuncs.com/boyfoo/prod:1.0 && docker push registry.cn-hangzhou.aliyuncs.com/boyfoo/prod:1.0

docker-review-build:
	docker build -t review:1.0 -f ./Dockerfile-review .

docker-review-push:
	docker tag review:1.0 registry.cn-hangzhou.aliyuncs.com/boyfoo/review:1.0 && docker push registry.cn-hangzhou.aliyuncs.com/boyfoo/review:1.0