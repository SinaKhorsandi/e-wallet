package db

type Money int64

// PersonalAccount : table for personal customer
type PersonalAccount struct {
	UserId int
	UserName string
	Email string
}
// OrganizationAccount : table for organization customer
type OrganizationAccount struct {
}
// Usertype : table for both different type of customers
type Usertype struct {
}

// Wallet : main table for user wallet
type Wallet struct {
	WalletId int
	UserId int
	Balance Money
}

// Transaction : table for all transaction in wallet
type Transaction struct {
	TransactionId int
	WalletId int
	Amount Money
	Date int32
	Status string
}

// Analytic : table for some data that help to create visual analytics
type Analytic struct {
	//Status
	//Result
}

//const (
//	status = 1
//)
