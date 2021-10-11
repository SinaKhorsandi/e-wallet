package db

//check
// Wallet : main table for user wallet
type Wallet struct {
	WalletId       int
	UserId         int
	OrganizationId int
	Balance        Money
}
