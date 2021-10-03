package db

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
	Balance int64
}

// Transaction : table for all transaction in wallet
type Transaction struct {
	TransactionId int
	WalletId int
	Amount int64
	Date int32
}

// Analytic : table for some data that help to create visual analytics
type Analytic struct {
	//Status
	//Result
}

//const (
//	status = 1
//)
