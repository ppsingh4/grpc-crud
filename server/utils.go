package server

import (
	"context"
	userpb "grpc-crud/proto"
)

type UserServiceServer struct {
}

type UserItem struct {
	ID   string
	Name string
	Dob  string
}

var allUsers = make(map[string]UserItem)

func (s *UserServiceServer) CreateUser(ctx context.Context, req *userpb.CreateUserReq) (*userpb.CreateUserRes, error) {
	user := req.GetUser()
	data := UserItem{
		ID:   user.GetId(),
		Name: user.GetName(),
		Dob:  user.GetDob(),
	}
	allUsers[data.ID] = data
	return &userpb.CreateUserRes{User: user}, nil
}

func (s *UserServiceServer) ReadUser(ctx context.Context, req *userpb.ReadUserReq) (*userpb.ReadUserRes, error) {
	id := req.GetId()
	result, ok := allUsers[id]
	if !ok {
		return &userpb.ReadUserRes{}, nil
	}
	response := &userpb.ReadUserRes{
		User: &userpb.User{
			Id:   result.ID,
			Name: result.Name,
			Dob:  result.Dob,
		},
	}
	return response, nil
}
func (s *UserServiceServer) DeleteUser(ctx context.Context, req *userpb.DeleteUserReq) (*userpb.DeleteUserRes, error) {
	id := req.GetId()
	delete(allUsers, id)
	return &userpb.DeleteUserRes{
		Success: true,
	}, nil
}

func (s *UserServiceServer) UpdateUser(ctx context.Context, req *userpb.UpdateUserReq) (*userpb.UpdateUserRes, error) {
	user := req.GetUser()
	data := UserItem{
		ID:   user.GetId(),
		Name: user.GetName(),
		Dob:  user.GetDob(),
	}
	allUsers[user.GetId()] = data
	response := &userpb.UpdateUserRes{
		User: &userpb.User{
			Id:   data.ID,
			Name: data.Name,
			Dob:  data.Dob,
		},
	}
	return response, nil
}

func (s *UserServiceServer) ListUsers(req *userpb.ListUserRequest, stream userpb.UserService_ListUsersServer) error {
	for _, value := range allUsers {
		stream.Send(&userpb.ListUserResponse{
			User: &userpb.User{
				Id:   value.ID,
				Name: value.Name,
				Dob:  value.Dob,
			},
		})
	}

	return nil
}
