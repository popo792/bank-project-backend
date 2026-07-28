package handlers

import (
	"bank-backend/config"
	"bank-backend/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetLoans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := `
		SELECT l.loan_id, c.name, l.collateral_type, l.collateral_worth, l.loan_requested, l.status
		FROM dbo.Loans l
		JOIN dbo.Customers c ON l.customer_id = c.customer_id
	`
	rows, err := config.DB.Query(query)
	if err != nil {
		fmt.Println("SQL Error:", err.Error())
		http.Error(w, "Database query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var loans []models.Loan
	for rows.Next() {
		var loan models.Loan
		if err := rows.Scan(&loan.LoanID, &loan.CustomerName, &loan.CollateralType, &loan.CollateralWorth, &loan.LoanRequested, &loan.Status); err != nil {
			fmt.Println("Row Scan Error:", err.Error())
			return
		}
		loans = append(loans, loan)
	}

	if loans == nil {
		loans = []models.Loan{}
	}

	json.NewEncoder(w).Encode(loans)
}

func UpdateLoanStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req models.LoanUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	query := `UPDATE dbo.Loans SET loan_passed = ?, status = ? WHERE loan_id = ?`
	_, err := config.DB.Exec(query, req.LoanAmount, req.Status, req.LoanID)
	if err != nil {
		fmt.Println("SQL Update Error:", err.Error())
		http.Error(w, "Failed to update loan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Loan updated successfully"})
}
