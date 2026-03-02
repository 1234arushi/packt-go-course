package main

//client is another go program/service or even the same machine that calls your gRPC server

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/packt-go-course/todo-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	ADDRESS = "localhost:50051"
)

type TodoTask struct {
	Name        string
	Description string
	Done        bool
}

func main() { //this is client
	var id string

	conn, err := grpc.NewClient(ADDRESS, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := proto.NewTodoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	todos := []TodoTask{
		{Name: "Code Review", Description: "Review new feauture code", Done: false},
		{Name: "Make Youtube video", Description: "Create new youtube video", Done: false},
		{Name: "Gym", Description: "go to the gym", Done: false},
		{Name: "Grocery", Description: "go to the grocery", Done: false},
		{Name: "Meeting", Description: "Meet about blockers in project", Done: false},
	}
	for _, todo := range todos {
		res, err := c.CreateTodo(ctx, &proto.NewTodo{Name: todo.Name, Description: todo.Description, Done: todo.Done})
		if err != nil {
			log.Fatalf("Could not create task : %v", err)
		}
		log.Printf(`
		ID: %s
		Name: %s
		Description: %s
		Done: %v

		`, res.GetId(), res.GetName(), res.GetDescription(), res.GetDone())
		id = res.GetId()
	}
	modifiedTodo := &proto.Todo{
		Id:          id,
		Name:        "Updated Name",
		Description: "Updated description",
		Done:        true,
	}
	modifiedTodoRes, err := c.ModifyTodo(ctx, modifiedTodo)
	if err != nil {
		log.Printf("Todo with ID: %v", id)
		log.Fatalf("Could not modify todo %v", err)
	}
	log.Printf("Modeified Todo: %v", modifiedTodoRes)
	deletedRes, err := c.DeleteTodo(ctx, &proto.TodoId{Id: id})
	if err != nil {
		log.Fatalf("Could not delete todo: %v", err)
	}
	log.Printf("Deleted todo: %v", deletedRes)
	listStream, err := c.ListTodos(ctx, &proto.Empty{})
	if err != nil {
		log.Fatalf("Could not list todos: %v", err)
	}
	for {
		todo, err := listStream.Recv()
		if err == io.EOF {
			log.Printf("end of list")
			return
		}
		if err != nil {
			log.Fatalf("Could not receive todo from stream : %v", err)
		}
		log.Printf("Listed todo: %v", todo)
	}
}
