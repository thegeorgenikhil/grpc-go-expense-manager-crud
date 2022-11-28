package main

import (
	"context"
	"log"
	"time"

	pb "github.com/thegeorgenikhil/grpc-go-expense-manager-crud/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	// connecting to the gRPC server
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	// new grpc client
	client := pb.NewExpenseServiceClient(conn)

	// CRUD OPERATIONS

	GetAllExpensesFromClient(client)
	// GetExpenseFromClient(client, "10001")
	// AddExpenseFromClient(client, "Rent", 3000)
	// UpdateExpenseFromClient(client, "10001", "Car Rental", 1000)
	// DeleteExpenseFromClient(client, "10002")
}

func GetAllExpensesFromClient(client pb.ExpenseServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.NoParam{}

	res, err := client.GetExpenses(ctx, req)

	if err != nil {
		log.Fatalf("Error when calling GetExpenses: %v", err)
	}

	for _, expense := range res.Expenses {
		log.Printf("Expense: %v", expense)
	}
}

func GetExpenseFromClient(client pb.ExpenseServiceClient, expenseId string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.ExpenseId{Id: expenseId}

	res, err := client.GetExpense(ctx, req)

	if err != nil {
		log.Fatalf("Error when calling GetExpense: %v", err)
	}
	log.Printf("Expense %v", res)
}

func AddExpenseFromClient(client pb.ExpenseServiceClient, title string, amount float32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.ExpenseInfo{Title: title, Amount: float32(amount)}

	res, err := client.AddExpense(ctx, req)

	if err != nil {
		log.Fatalf("Error when calling AddExpense: %v", err)
	}
	log.Printf("Successfully Added Expense %v!", res.Id)
}

func UpdateExpenseFromClient(client pb.ExpenseServiceClient, id string, title string, amount float32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.UpdateExpenseReq{Id: id, Title: title, Amount: float32(amount)}

	res, err := client.UpdateExpense(ctx, req)

	if err != nil {
		log.Fatalf("Error when calling UpdateExpense: %v", err)
	}
	log.Printf("Successfully Updated Expense %v!", res.Id)
}

func DeleteExpenseFromClient(client pb.ExpenseServiceClient, id string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.ExpenseId{Id: id}

	res, err := client.DeleteExpense(ctx, req)

	if err != nil {
		log.Fatalf("Error when calling DeleteExpense: %v", err)
	}
	log.Printf("Successfully Deleted Expense %v!", res.Id)
}
