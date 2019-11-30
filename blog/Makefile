

image:
	docker build -t blogsvc .

proto: api/blog.pb.go

descriptor: build/blog.pb
	cat build/blog.pb | openssl base64 -A

api/blog.pb.go:
	protoc -I./proto --go_out=plugins=grpc:./api ./proto/blog.proto

build:
	mkdir -p build

build/blog.pb: build
	protoc -I./proto --include_imports --include_source_info --descriptor_set_out=build/blog.pb blog.proto
