package models

type Transaction struct {
	TransactionID int     `json:"transactionId"`
	CustomerName  string  `json:"customerName"`
	Type          string  `json:"type"`
	Amount        float64 `json:"amount"`
	Balance       float64 `json:"balance"`
}

type Loan struct {
	LoanID          int     `json:"loanId"`
	CustomerName    string  `json:"customerName"`
	CollateralType  string  `json:"collateralType"`
	CollateralWorth float64 `json:"collateralWorth"`
	LoanRequested   float64 `json:"loanAmount"`
	Status          string  `json:"status"`
}

type LoanUpdateRequest struct {
	LoanID     int     `json:"loanId"`
	LoanAmount float64 `json:"loanAmount"`
	Status     string  `json:"status"`
}
