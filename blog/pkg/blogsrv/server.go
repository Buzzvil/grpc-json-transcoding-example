package blogsrv

import (
	"context"

	pb "github.com/Buzzvil/grpc-json-transcoding-example/blog/api"
)

// Server implements blog server.
type Server struct {
}

// New creates a blog server.
func New() *Server {
	return &Server{}
}

// CreatePost creates a post.
func (s *Server) CreatePost(ctx context.Context, in *pb.CreatePostRequest) (*pb.Post, error) {
	newPost := pb.Post{
		Id:    1,
		Title: in.Post.GetTitle(),
		Body:  in.Post.GetBody(),
		Tags:  in.Post.GetTags(),
	}
	return &newPost, nil
}

// ListPosts returns a list of posts.
func (s *Server) ListPosts(ctx context.Context, in *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	res := pb.ListPostsResponse{
		Posts: []*pb.Post{
			&pb.Post{Id: 1, Title: "Sample Post Title #1", Body: "Sample Post Body #1", Tags: []string{"tag1", "tag2"}},
			&pb.Post{Id: 2, Title: "Sample Post Title #2", Body: "Sample Post Body #2", Tags: []string{"tag1", "tag3"}},
		},
	}
	return &res, nil
}

// UpdatePost updates the post record.
func (s *Server) UpdatePost(ctx context.Context, in *pb.UpdatePostRequest) (*pb.Post, error) {
	return in.Post, nil
}
