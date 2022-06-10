package main

import (
	"fmt"
	userpb "grpc-crud/proto"
	"io"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	//change address below if you are running locally with docker containers(container IP)
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := userpb.NewUserServiceClient(conn)
	createUser("1", "User1", "041095", c)
	createUser("2", "User2", "041095", c)
	getAllUsers(c)
	getUserById("2", c)
	updateUser("2", "Updated User2", "041095", c)
	getUserById("2", c)
	deleteUser("2", c)
	getUserById("2", c)
	log.Println("END OF PROGRAM")

}

func createUser(id string, name string, dob string, c userpb.UserServiceClient) {
	response, err := c.CreateUser(context.Background(), &userpb.CreateUserReq{User: &userpb.User{Id: id, Name: name, Dob: dob}})
	if err != nil {
		log.Fatalf("Error when calling Server: %s", err)
	}
	log.Printf("Response from server for add user: %s", response.User)
}

func getAllUsers(c userpb.UserServiceClient) {
	log.Printf("Get All User details")
	req := &userpb.ListUserRequest{}
	stream, err := c.ListUsers(context.Background(), req)
	if err != nil {
		log.Fatalf("Error when calling  List server: %s", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error while fetching stream")
			break
		}

		fmt.Println("RESULT from stream is", res.GetUser())
	}
}

func getUserById(id string, c userpb.UserServiceClient) {
	log.Printf("Get User details by Id")
	responseRead, err := c.ReadUser(context.Background(), &userpb.ReadUserReq{Id: "2"})
	if err != nil {
		log.Fatalf("Error when calling Server: %s", err)
	}
	log.Printf("Response from server for read is: %s", responseRead.User)
}
func updateUser(id string, name string, dob string, c userpb.UserServiceClient) {
	log.Printf("Update User details")
	responseUpdate, err := c.UpdateUser(context.Background(), &userpb.UpdateUserReq{User: &userpb.User{Id: id, Name: name, Dob: dob}})
	if err != nil {
		log.Fatalf("Error when calling Server: %s", err)
	}
	log.Printf("Response from server for update is: %s", responseUpdate.User)
}

func deleteUser(id string, c userpb.UserServiceClient) {
	log.Printf("Delete User details")
	responseDelete, err := c.DeleteUser(context.Background(), &userpb.DeleteUserReq{Id: "2"})
	if err != nil {
		log.Fatalf("Error when calling Server: %s", err)
	}
	log.Printf("Response from server for delete is: %s", responseDelete.Success)
}
