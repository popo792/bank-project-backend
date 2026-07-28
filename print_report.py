import os
import pyodbc
import pandas as pd
from datetime import datetime

reports_folder = os.path.join(os.path.dirname(os.path.abspath(__file__)), "reports")
if not os.path.exists(reports_folder):
    os.makedirs(reports_folder)

connection_string = (
    "Driver={ODBC Driver 17 for SQL Server};"
    "Server=localhost,1433;"
    "Database=bank_db;"
    "UID=sa;"
    "PWD=sabfy1995;"
)
database_connection = pyodbc.connect(connection_string)

transactions_query = """
SELECT t.transaction_id, c.name, t.transaction_type, t.amount, c.current_balance
FROM dbo.Transactions t
JOIN dbo.Customers c ON t.customer_id = c.customer_id
"""

loans_query = """
SELECT l.loan_id, c.name, l.collateral_type, l.collateral_worth, l.loan_requested, l.loan_passed, l.status
FROM dbo.Loans l
JOIN dbo.Customers c ON l.customer_id = c.customer_id
"""

transactions_dataframe = pd.read_sql(transactions_query, database_connection)
loans_dataframe = pd.read_sql(loans_query, database_connection)

current_time = datetime.now().strftime('%Y%m%d_%H%M%S')
excel_path = os.path.join(reports_folder, f"Bank_Report_{current_time}.xlsx")

with pd.ExcelWriter(excel_path, engine='openpyxl') as writer:
    transactions_dataframe.to_excel(writer, index=False, sheet_name="Transactions")
    loans_dataframe.to_excel(writer, index=False, sheet_name="Loans")

print(f"READY:{excel_path}")