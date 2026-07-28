import pyodbc
import pandas as pd
from datetime import datetime

conn_str = (
    "Driver={ODBC Driver 17 for SQL Server};"
    "Server=localhost,1433;"
    "Database=bank_db;"
    "UID=sa;"
    "PWD=your_password;"
)
conn = pyodbc.connect(conn_str)

query_transactions = """
SELECT t.transaction_id, c.name, t.transaction_type, t.amount, c.current_balance
FROM dbo.Transactions t
JOIN dbo.Customers c ON t.customer_id = c.customer_id
"""

query_loans = """
SELECT l.loan_id, c.name, l.collateral_type, l.collateral_worth, l.loan_requested, l.loan_passed, l.status
FROM dbo.Loans l
JOIN dbo.Customers c ON l.customer_id = c.customer_id
"""

try:
    df_transactions = pd.read_sql(query_transactions, conn)
    df_loans = pd.read_sql(query_loans, conn)
except Exception as e:
    print(f"Error reading database: {e}")
    exit()

filename = f"Bank_Report_{datetime.now().strftime('%Y%m%d_%H%M')}.txt"

with open(filename, "w") as file:
    file.write("="*50 + "\n")
    file.write("BANK MANAGEMENT SYSTEM - OFFICIAL REPORT\n")
    file.write(f"Generated on: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}\n")
    file.write("="*50 + "\n\n")

    file.write("--- 1. CUSTOMER TRANSACTIONS (DEBITS/CREDITS) ---\n")
    file.write(df_transactions.to_string(index=False))
    file.write("\n\n")

    file.write("--- 2. LOAN LEDGER (APPROVED/REJECTED/CONSIDERATION) ---\n")
    file.write(df_loans.to_string(index=False))
    file.write("\n\n")
    
    file.write("="*50 + "\n")
    file.write("END OF REPORT\n")

print(f"Success! Printed document saved as '{filename}'")