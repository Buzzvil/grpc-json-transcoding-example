package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/Buzzvil/grpc-json-transcoding-example/blog/api"
	"github.com/Buzzvil/grpc-json-transcoding-example/blog/pkg/blogsrv"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Starting the server..")

	s := newGrpcServer()
	blogServer := blogsrv.New()
	pb.RegisterBlogServiceServer(s, blogServer)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func newGrpcServer() *grpc.Server {
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())
	opts := []grpc_logrus.Option{}
	grpc_logrus.ReplaceGrpcLogger(logrusEntry)
	return grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.UnaryServerInterceptor(logrusEntry, opts...),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.StreamServerInterceptor(logrusEntry, opts...),
		),
	)
}
