default: build

help:
	@echo 'Usage:'
	@echo '    make build           Build docker image.'
	@echo '    make up              Start the service.'
	@echo '    make down 			Stop the service.'
	@echo '    make descriptor      Print base64 encoded proto descriptor'
	@echo

build: api/blog.pb.go
	docker build -t blogsvc .

up:
	kubectl apply -f deployment

down:
	kubectl delete -f deployment

descriptor: tmp/blog.pb
	cat tmp/blog.pb | openssl base64 -A

api/blog.pb.go:
	protoc -I./proto --go_out=plugins=grpc:./api ./proto/blog.proto

tmp:
	mkdir -p tmp

tmp/blog.pb: tmp
	protoc -I./proto --include_imports --include_source_info --descriptor_set_out=tmp/blog.pb blog.proto

clean:
	rm -rf tmp
	docker image rm blogsvc
