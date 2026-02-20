package db

const (
	Users TableName = iota
	Products
	Sales
	SalesSummary
	Purchases
	PurchaseSummary
	Expenses
	ExpenseSummary
	ExpenseByCategory
)

type TableName int

func (t TableName) String() string {
	switch t {
	case Users:
		return "Users"
	case Products:
		return "Products"
	case Sales:
		return "Sales"
	case SalesSummary:
		return "SalesSummary"
	case Purchases:
		return "Purchases"
	case PurchaseSummary:
		return "PurchaseSummary"
	case Expenses:
		return "Expenses"
	case ExpenseSummary:
		return "ExpenseSummary"
	case ExpenseByCategory:
		return "ExpenseByCategory"
	default:
		return "Invalid Table name"
	}
}
