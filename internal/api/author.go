package api

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/tkcrm/mx-example/internal/models"
	"github.com/tkcrm/mx-example/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type authorServer struct {
	name string
	data sync.Map

	pb.UnimplementedAuthorServiceServer
}

func NewAuthorServer() *authorServer {
	return &authorServer{name: "author-server"}
}

// Name of the service
func (s *authorServer) Name() string { return s.name }

// Register service on grpc.Server
func (s *authorServer) Register(srv *grpc.Server) {
	pb.RegisterAuthorServiceServer(srv, s)
}

func (s *authorServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	if req.Id == "" {
		return nil, fmt.Errorf("empty author id")
	}

	item, ok := s.data.Load(req.Id)
	if !ok {
		return nil, fmt.Errorf("author does not exists with id: %s", req.Id)
	}

	author, ok := item.(models.Author)
	if !ok {
		return nil, fmt.Errorf("item does not implemet author")
	}

	return &pb.GetResponse{
		Item: &pb.Author{
			Id:        author.ID,
			Name:      author.Name,
			CreatedAt: timestamppb.New(author.CraetedAt),
		},
	}, nil
}

func (s *authorServer) Find(ctx context.Context, req *emptypb.Empty) (*pb.FindResponse, error) {
	authors := make([]*pb.Author, 0)
	s.data.Range(func(key, value any) bool {
		author, ok := value.(models.Author)
		if !ok {
			return true
		}

		authors = append(authors, &pb.Author{
			Id:        author.ID,
			Name:      author.Name,
			CreatedAt: timestamppb.New(author.CraetedAt),
		})

		return true
	})

	return &pb.FindResponse{
		Items: authors,
	}, nil
}

func (s *authorServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	if req.Name == "" {
		return nil, fmt.Errorf("empty author name")
	}

	author := models.Author{
		ID:        uuid.New().String(),
		Name:      req.Name,
		CraetedAt: time.Now(),
	}

	s.data.Store(author.ID, author)

	return &pb.CreateResponse{
		Item: &pb.Author{
			Id:        author.ID,
			Name:      author.Name,
			CreatedAt: timestamppb.New(author.CraetedAt),
		},
	}, nil
}

func (s *authorServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*emptypb.Empty, error) {
	if req.Id == "" {
		return nil, fmt.Errorf("empty author id")
	}

	if _, ok := s.data.Load(req.Id); !ok {
		return nil, fmt.Errorf("author does not exists with id: %s", req.Id)
	}

	s.data.Delete(req.Id)

	return &emptypb.Empty{}, nil
}
