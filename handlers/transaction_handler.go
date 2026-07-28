package handlers

import (
	"bank-backend/config"
	"bank-backend/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := `
		SELECT t.transaction_id, c.name, t.transaction_type, t.amount, c.current_balance
		FROM dbo.Transactions t
		JOIN dbo.Customers c ON t.customer_id = c.customer_id
		ORDER BY t.transaction_id DESC
	`
	rows, err := config.DB.Query(query)
	if err != nil {
		fmt.Println("SQL Error:", err.Error())
		http.Error(w, "Database query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var txn models.Transaction
		if err := rows.Scan(&txn.TransactionID, &txn.CustomerName, &txn.Type, &txn.Amount, &txn.Balance); err != nil {
			fmt.Println("Row Scan Error:", err.Error())
			http.Error(w, "Error reading data", http.StatusInternalServerError)
			return
		}
		transactions = append(transactions, txn)
	}

	if transactions == nil {
		transactions = []models.Transaction{}
	}

	json.NewEncoder(w).Encode(transactions)
}
