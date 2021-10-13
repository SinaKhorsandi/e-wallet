package data

//check
// Wallet : main table for user wallet
type Wallet struct {
	WalletId       int
	UserId         int
	Balance        Money
}

//func GetUserIDByWalletID(email string) UserEmail{
//
//}
//func GetWalletIDByUserID(userID int) UserEmail{
//}
//func GetBalanceByUserID(userID int) UserEmail{
//}
//func GetBalanceByWalletID(userID int) UserEmail{
//}
