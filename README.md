# Expense Manager CRUD using gRPC and Go

This a simple CRUD application using gRPC and Go, where you can create, read, update and delete expenses.


# Prerequisites

- [Go](https://golang.org/dl/)
- [Protocol Buffers](https://developers.google.com/protocol-buffers/docs/downloads)

# Setting up a gRPC-Go project
1. Create a new directory for your project and cd into it

```bash
mkdir grpc-go-expense-manager-crud
cd grpc-go-expense-manager-crud
mkdir client server proto
```

2. Installing the gRPC Go plugin

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

export PATH="$PATH:$(go env GOPATH)/bin"
```

3. Initialize a Go module

```bash
go mod init github.com/your_username/grpc-go-expense-manager-crud

go mod tidy
```

4. Create the proto file with the required services and messages in the proto directory

5. Generate .pb.go files from the proto file

```bash
protoc --go_out=. --go_opt=paths=source_relative \ --go-grpc_out=. --go-grpc_opt=paths=source_relative \ proto/expense.proto
```

6. Create the server and client directories and create the main.go files with necessary controllers and services


# Running the application

1. Install the dependencies

```bash
go mod tidy
```

2. Run the server

```bash
go run server/main.go
```

3. Run the client

```bash
go run client/main.go
```

# References
- [What is gRPC](https://grpc.io/docs/what-is-grpc/introduction/)
- [gRPC Go Quickstart](https://grpc.io/docs/languages/go/quickstart/)
