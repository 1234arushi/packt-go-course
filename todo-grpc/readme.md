# Implementing To-Do Application using gRPC

---

## 🔹 Setup Steps

### Step 1
Install Go protobuf plugin:

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31

### Step 2
Install gRPC Go plugin:

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

### Step 3
Install protoc compiler (if not installed):

brew install protobuf

---

## 🔹 gRPC + Protocol Buffers Flow

### a. Define Structure using IDL (.proto file)

- Developers define the structure of data using Interface Definition Language (IDL).
- Written in a plain text `.proto` file.
- Uses message types to specify:
  - Fields
  - Data types
- Defines:
  - Service methods
  - Request/Response structure

---

### b. Compile the .proto File

- protoc (Protocol Buffer Compiler) reads the `.proto` file.
- Automatically generates source code in target language (Go here).

Command:

protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       proto/todo.proto

---

### c. Serialization

- Generated code provides:
  - Structs
  - Helper methods
- Converts in-memory data into compact binary format.
- Used for:
  - Storage
  - Network transmission

---

### d. Deserialization

- Receiving system:
  - Uses its own generated code
  - Parses binary data
  - Converts back into structured format

---

## 🔹 todo.proto Explanation

The `.proto` file defines:

- Message types
- Service methods
- Request/Response structures
- Complete API contract

---

### 1. Syntax Version

syntax = "proto3";

- Specifies Protocol Buffer version 3.

---

### 2. Protocol Buffers

- Used for:
  - Serialization
  - Compact binary data transfer

---

### 3. go_package Option

option go_package = "github.com/packt-go-course/todo-grpc/proto";

- Tells Go code generator:
  - Where to place generated files
  - What import path to use

---

### 4. Service Definition

service TodoService {
    rpc CreateTodo(NewTodo) returns (Todo);
    rpc DeleteTodo(TodoId) returns (Empty);
    rpc ModifyTod(Todo) returns (Todo);
    rpc ListTodos(Empty) returns (stream Todo);
}

---

### Meaning of RPC

- rpc = Remote Procedure Call
- Client calls a function that runs on another machine/server.

Example:

rpc CreateTodo(NewTodo) returns (Todo);

- Client sends NewTodo
- Server processes request
- Server returns Todo

---

### Service Body

- Defines methods:
  - Called by client
  - Implemented by server

---

## 🔹 After Compilation

Two files are generated:

1) todo.pb.go
   - Contains:
     - Request structs
     - Response structs
     - Data models

2) todo_grpc.pb.go
   - Contains:
     - Service interface
     - Client stub
     - Server registration logic

Important:

- These are generated code layers.
- They handle:
  - Networking
  - Serialization
  - HTTP/2 communication

---

## 🔹 gRPC Request Flow

1. Client calls CreateTodo()
2. gRPC client stub:
   - Serializes request
   - Sends binary data over HTTP/2
3. gRPC server:
   - Receives request
   - Deserializes data
   - Calls your server function

---

## 🔹 Docker Concepts

docker build -t todo-server .

- Builds an image
- -t → name/tag
- . → use current folder as build context

Image → blueprint  
Container → running instance of that blueprint  

Class → Object  
Image → Container  

---

RUN go build -o server ./server/server.go

- Compiles Go code into a binary file
- Binary is named "server"

---

CMD

- Tells image what to execute when container starts

---

docker run -p 50051:50051 todo-server

- Creates & starts a container

Port mapping format:

host_port : container_port

Meaning:

- If someone hits localhost:50051 on your Mac
- Forward that traffic to port 50051 inside container

---

## 🔹 Why We Containerize

- So it runs everywhere
- No dependency conflicts
- Same environment in all systems
- Production-ready deployment