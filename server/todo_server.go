package server

import (
	"context"
	"fmt"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"example/backend/db"
	"example/backend/ent"
	pb "example/backend/grpc"
)

func Register(s *grpc.Server) {
	pb.RegisterTodoServer(s, &server{})
}

type server struct {
	pb.UnimplementedTodoServer
}

func (s *server) GetList(ctx context.Context, in *pb.TodoGetListReuqest) (*pb.TodoGetListResponse, error) {
	entities, err := db.Client.Todo.Query().Limit(int(*in.Limit)).Offset(int(*in.Offset)).All(context.Background())
	if err != nil {
		return nil, err
	}
	count, err := db.Client.Todo.Query().Count(context.Background())
	if err != nil {
		return nil, err
	}

	items := []*pb.TodoResponse{}
	for _, e := range entities {
		items = append(items, &pb.TodoResponse{
			Id:       strconv.Itoa(e.ID),
			Title:    e.Title,
			Category: e.Category,
			Content:  &e.Content,
		})
	}
	return &pb.TodoGetListResponse{
		Total: uint32(count),
		Items: items,
	}, nil
}

func (s *server) Create(ctx context.Context, in *pb.TodoCreateRequest) (*pb.TodoResponse, error) {
	entity, err := db.Client.Todo.Create().
		SetTitle(in.Title).
		SetCategory(in.Category).
		SetContent(*in.Content).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return &pb.TodoResponse{
		Id:       strconv.Itoa(entity.ID),
		Title:    entity.Title,
		Category: entity.Category,
		Content:  &entity.Content,
	}, nil
}

func (s *server) Get(ctx context.Context, in *pb.TodoGetRequest) (*pb.TodoResponse, error) {
	id, _ := strconv.Atoi(in.Id)
	entity, err := db.Client.Todo.Get(context.Background(), id)
	if ent.IsNotFound(err) {
		return nil, status.Error(codes.NotFound, "Todo item was not found.")
	} else if err != nil {
		return nil, err
	}

	return &pb.TodoResponse{
		Id:       strconv.Itoa(entity.ID),
		Title:    entity.Title,
		Category: entity.Category,
		Content:  &entity.Content}, nil
}

func (s *server) Update(ctx context.Context, in *pb.TodoUpdateeRequest) (*pb.TodoResponse, error) {
	id, _ := strconv.Atoi(in.Id)
	entity, err := db.Client.Todo.UpdateOneID(id).
		SetTitle(*in.Title).
		SetCategory(*in.Category).
		SetContent(*in.Content).
		Save(context.Background())
	if ent.IsNotFound(err) {
		return nil, status.Error(codes.NotFound, "Todo item was not found.")
	} else if err != nil {
		return nil, err
	}

	return &pb.TodoResponse{
		Id:       strconv.Itoa(entity.ID),
		Title:    entity.Title,
		Category: entity.Category,
		Content:  &entity.Content,
	}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.TodoGetRequest) (*emptypb.Empty, error) {
	id, _ := strconv.Atoi(in.Id)
	if err := db.Client.Todo.DeleteOneID(id).Exec(context.Background()); err != nil {
		if ent.IsNotFound(err) {
			fmt.Println("Not Found")
			return nil, status.Error(codes.NotFound, "Todo item was not found.")
		} else {
			return nil, err
		}
	}
	return &emptypb.Empty{}, nil
}
