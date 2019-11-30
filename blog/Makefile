

proto: api/blog.pb.go

api/blog.pb.go:
	protoc -I./proto --go_out=plugins=grpc:./api ./proto/blog.proto

image:
	docker build -t blogsvc .
