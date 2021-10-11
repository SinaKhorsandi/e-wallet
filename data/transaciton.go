package data

type Transaction struct {
	TransactionId int
	WalletId      int
	Amount        Money
	//Balance 	  Money
	Date          int32
	Status        string
}
