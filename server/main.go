package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	pb "github.com/thegeorgenikhil/grpc-go-expense-manager-crud/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

// will be using this slice as our DB
var expenses []*pb.Expense

type expenseServer struct {
	pb.UnimplementedExpenseServiceServer
}

func main() {
	seedDatabase()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
	// create a new gRPC server
	grpcServer := grpc.NewServer()
	// register the expense service
	pb.RegisterExpenseServiceServer(grpcServer, &expenseServer{})
	log.Printf("Server started at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}

// seedDatabase will seed the database with some dummy data
func seedDatabase() {
	expenses = append(expenses, &pb.Expense{
		Id:        "10001",
		Title:     "Movie Ticket",
		Amount:    100,
		Timestamp: time.Now().Unix(),
	})
	expenses = append(expenses, &pb.Expense{
		Id:        "10002",
		Title:     "Lunch",
		Amount:    200,
		Timestamp: time.Now().Unix(),
	})
}


func (s *expenseServer) GetExpenses(ctx context.Context, v *pb.NoParam) (*pb.Expenses, error) {
	res := &pb.Expenses{Expenses: expenses}
	return res, nil
}

func (s *expenseServer) GetExpense(ctx context.Context, v *pb.ExpenseId) (*pb.Expense, error) {
	res := &pb.Expense{}
	for _, expense := range expenses {
		if expense.GetId() == v.GetId() {
			res = expense
			break
		}
	}
	return res, nil
}

func (s *expenseServer) AddExpense(ctx context.Context, v *pb.ExpenseInfo) (*pb.Expense, error) {
	newExpense := &pb.Expense{}

	newExpense.Id = strconv.Itoa(rand.Intn(10000000))
	newExpense.Timestamp = time.Now().Unix()
	newExpense.Title = v.GetTitle()
	newExpense.Amount = v.GetAmount()
	expenses = append(expenses, newExpense)

	return newExpense, nil
}

func (s *expenseServer) UpdateExpense(ctx context.Context, v *pb.UpdateExpenseReq) (*pb.ExpenseStatusResponse, error) {
	updatedExpense := &pb.Expense{}
	res := &pb.ExpenseStatusResponse{}
	for index, expense := range expenses {
		if expense.GetId() == v.GetId() {
			expenses = append(expenses[:index], expenses[index+1:]...)
			updatedExpense = expense
			if v.GetTitle() != "" {
				updatedExpense.Title = v.GetTitle()
			}
			if v.GetAmount() != 0 {
				updatedExpense.Amount = v.GetAmount()
			}
			res.Status = 1
			res.Id = updatedExpense.GetId()
			break
		}
	}
	expenses = append(expenses, updatedExpense)
	return res, nil
}

func (s *expenseServer) DeleteExpense(ctx context.Context, v *pb.ExpenseId) (*pb.ExpenseStatusResponse, error) {
	res := &pb.ExpenseStatusResponse{}
	for index, expense := range expenses {
		if expense.GetId() == v.GetId() {
			expenses = append(expenses[:index], expenses[index+1:]...)
			res.Status = 1
			res.Id = expense.GetId()
		}
	}
	return res, nil
}
