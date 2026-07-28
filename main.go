package main

import (
	"bank-backend/config"
	"bank-backend/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/api/transactions", handlers.GetTransactions)
	http.HandleFunc("/api/loans", handlers.GetLoans)
	http.HandleFunc("/api/loans/update", handlers.UpdateLoanStatus)

	fs := http.FileServer(http.Dir("D:/Projects/bank-project/bank-project-frontend"))
	http.Handle("/", fs)

	fmt.Println("Server is running on port 8080...")
	fmt.Println("--> Open http://localhost:8080 in your browser to see the UI! <--")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
