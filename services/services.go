package wallet

import (
	"fmt"
	"github.com/sinakhorsandi/e-wallet/data"
)

type Status string

const (
	UserIsExist                 Status = "user is exist"
	OrganIsExist                Status = "organization is exist"
	NegativeAmount              Status = "Amount must be positive"
	AccountNotExist             Status = "Account does not exist"
	TransactionStatusOk         Status = "ok"
	TransactionStatusFail       Status = "fail"
	TransactionStatusInProgress Status = " in progress"
)

type Service struct {
	nextAccountId int
	nextWalletId  int
	nextRecordId  int
	account map[string]int
	wallet  map[int]data.Money
}



func (s *Service) userRegister(userName, email string) error {
	b := s.checkCustomerExist(email)
	if b == true {
		return fmt.Errorf("%s", UserIsExist)
	}
	s.nextAccountId++
	s.nextRecordId++
	EmailId := &data.UserEmail{
		RecordId: s.nextRecordId,
		UserId:   s.nextAccountId,
		Email:    email,
	}
	s.nextRecordId++
	NameID := &data.UserName{
		RecordId: s.nextRecordId,
		UserId:   s.nextAccountId,
		UserName: userName,
	}
	m := make(map[string]int)
	m[EmailId.Email]=EmailId.UserId
	m[NameID.UserName]=NameID.UserId
	s.account=m

	wall:=s.newWallet(EmailId.UserId)
	w:= make(map[int]data.Money)
	w[wall.WalletId]=wall.Balance
	s.wallet=w
	return nil
}
func (s *Service) checkCustomerExist(email string) bool {
	_,ok:=s.account[email]
	if !ok {
		return false
	}
	return true
}
//Error handling !!!!
func (s *Service) newWallet(accountId int) *data.Wallet {
	s.nextWalletId++
	wallet := &data.Wallet{
		WalletId: s.nextWalletId,
		UserId:   accountId,
		Balance:  0,
	}
	return wallet
}
//check function output
func (s *Service) findIDByEmail(email string) (int, error) {
	v ,ok:= s.account[email]
	if !ok {
		return 0,fmt.Errorf("%s\n",AccountNotExist)
	}
	return v, nil
}
func (s *Service) findIDByName(userName string) (int, error) {
	v ,ok:= s.account[userName]
	if !ok {
		return 0,fmt.Errorf("%s\n",AccountNotExist)
	}
	return v, nil
}

//check output
func (s *Service) getBalance(userID int) (data.Money,error) {
balance,ok:=s.wallet[userID]
	if !ok {

		//check
		return 0,fmt.Errorf("%s\n",AccountNotExist)
	}
	return balance, nil
}

//func showBalance()  {
//}



////Think how can change accountId to another thing
//func (s *Service) deposit(accountId int, amount data.Money) (*data.Wallet, error) {
//	if amount <= 0 {
//		return nil, fmt.Errorf("%s\n", NegativeAmount)
//	}
//	wal, err := s.checkPersonalWallet(accountId)
//	if err != nil {
//		fmt.Printf("%s\n", err)
//	}
//	wal.Balance += amount
//	return wal, nil
//}
//func (s *Service) checkPersonalAccountById(accountId int) (*data.PersonalAccount, error) {
//	for _, acc := range s.personalAccounts {
//		if acc.UserId == accountId {
//			return acc, nil
//		}
//	}
//	return nil, fmt.Errorf("%s\n", AccountNotExist)
//}
//func (s *Service) checkPersonalWallet(accountId int) (*data.Wallet, error) {
//	for _, wal := range s.wallets {
//		if wal.UserId == accountId {
//			return wal, nil
//		}
//	}
//	return nil, fmt.Errorf("%s\n", AccountNotExist)
//}
//
////func (s *Service) checkOrganizationAccountById(accountId int) bool  {
////	ret := s.Called(accountId)
////	err := ret.Error(0)
////	if err != nil {
////		return false
////	}
////	return true
////}
//
//
////Pay : pay function decrease amount of money from an accountId to another accountId
//func pay() {
//}
