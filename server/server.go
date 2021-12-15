package server

// TODO: uncomment below once you've generated your rpc files and you're stubbing methods
import (
	"context"
	"fmt"
	groceryProto "grocery-list/rpc/grocery"
)

// TODO: uncomment below once you're ready to fully write out methods
// import (
//   config "grocery-list/config"
//   "github.com/twitchtv/twirp"
// )

type Server struct{}

// TODO: write the expected functions for our API (CreateGrocery, UpdateGrocery, etc...)
// use the GroceryService interface generated in service.twirp.go as a guide

func (*Server) CreateGrocery(ctx context.Context, req *groceryProto.CreateGroceryRequest) (*groceryProto.CreateGroceryResponse, error) {
	fmt.Println("In CreateGrocery")
	return &groceryProto.CreateGroceryResponse{}, nil
}

func (*Server) GetGrocery(ctx context.Context, req *groceryProto.GetGroceryRequest) (*groceryProto.GetGroceryResponse, error) {
	fmt.Println("In GetGrocery")
	return &groceryProto.GetGroceryResponse{}, nil
}

func (*Server) UpdateGrocery(ctx context.Context, req *groceryProto.UpdateGroceryRequest) (*groceryProto.UpdateGroceryResponse, error) {
	fmt.Println("In UpdateGrocery")
	return &groceryProto.UpdateGroceryResponse{}, nil
}

func (*Server) DeleteGrocery(ctx context.Context, req *groceryProto.DeleteGroceryRequest) (*groceryProto.DeleteGroceryResponse, error) {
	fmt.Println("In DeleteGrocery")
	return &groceryProto.DeleteGroceryResponse{}, nil
}

func (*Server) ListGrocery(ctx context.Context, req *groceryProto.ListGroceryRequest) (*groceryProto.ListGroceryResponse, error) {
	fmt.Println("In ListGrocery")
	return &groceryProto.ListGroceryResponse{}, nil
}
