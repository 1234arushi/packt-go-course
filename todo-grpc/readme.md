Implementing to-do application using gRPC


step 1 ->go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31
step 2 -> go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.3

a. developers define the structure of their data using interface definition language(idl)in a plain text .proto file.This schema uses message types to specify the fields and their data types

b. a special compiler ->protoc,reads the .proto file and automatically generates source code in the target programming language.

c. The generated code provides classes and methods to build in-memory data structures,which can then be efficiently converted(serialized)into a compact binary format for storage and transmission.

d. The receiving system uses its own generated code to parse(deserialize)the binary data back into original structured format.


todo.proto -> this defines message types,service methods,request/response structures[it defines the API]
 1. syntax = "proto3"; ->version3
 2. protoc-buff -> to serialize data
 3. option go_package="github.com/packt-go-course/todo-grpc/proto" -> this tells Go-code generator,that when generating Go Files,place them under import path.
 4. service TodoService{
    rpc CreateTodo(NewTodo) returns todo{}
    rpc DeleteTodo(TodoId) returns (Empty){}
    rpc ModifyTod(Todo)returns (Todo){}
    rpc ListTodos(Empty)returns(stream Todo){}

} -> rpc[remote call procedure] : call a function that runs on another machine/server
-> rpc CreateTodo(NewTodo) returns(Todo){} : client sends NewTodo,server processes it and returns Todo
->service body defines the methods that will be called by client and implemented by server
5.  protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/todo.proto(to compile it) but must do->brew install protobuf(to install protoc : compiler)

after compilation of .proto file is done,we get two files: todo_grpc.pb.go and todo.pb.go

1. todo.pb.go -> contains all requests/response structs(data models)
2. todo_grpc.pb.gp -> contains service interface,client code(rpc service logic)
3. these are just generated code layer which handles networking & serialization
4. client calls createTodo()->grpc clinet stub(serialize request,sends binary to http/2)->grpc server receives(deserializes request and calls your function)
