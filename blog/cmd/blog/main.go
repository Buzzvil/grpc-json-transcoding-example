package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	// "github.com/Buzzvil/appaccountsvc/internal/app/accountsrv"
	// pb "github.com/Buzzvil/buzzapis/go/appaccount"
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

	var db *sql.DB
	db, err = sql.Open("sqlite3", "file:db/blog.db")

	s := newGrpcServer()
	// Register grpc server
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
