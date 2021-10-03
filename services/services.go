package wallet

import (
	"github.com/sinakhorsandi/testify/mock"
)

type Status string
const (
	UserIsExist 				Status = "user is exist"
	NegativeAmount				Status = "Amount must be positive"
	TransactionStatusOk         Status = "ok"
	TransactionStatusFail       Status = "fail"
	TransactionStatusInProgress Status = " in progress"
)

type Service struct {
	//accounts []*db.PersonalAccount
	//wallets      []*db.Wallet
	//payments      []*db.Transaction
	mock.Mock
}


//type Call struct {
//
//}

//func (s *Service) userRegister(user *db.PersonalAccount) error {
//	b:=s.checkCustomerExist(user.Email)
//	if b==true {
//		return fmt.Errorf("%s",UserIsExist)
//	}
//	account :=&db.PersonalAccount{
//		UserId:user.UserId ,
//		UserName: user.UserName,
//		Email: user.Email,
//	}
//
//}

func organizationRegister()  {
}


func (s *Service) checkCustomerExist(email string) bool {
	ret := s.Called(email)
	err := ret.Error(0)
	if err != nil {
		return false
	}
	return true
}

//func (s *Service) deposit(WalletId int, amount db.Money) error {
//	if amount <= 0 {
//		return fmt.Errorf("%s",NegativeAmount)
//	}
//
//	var wallet *db.Wallet
//	for _, wal := range s.wallets {
//		if wal.WalletId == WalletId {
//			wallet = wal
//			break
//		}
//	}
//
//	if wallet == nil {
//		return fmt.Errorf("error")
//	}
//
//	wallet.Balance += amount
//	return nil
//}
//
//func pay()  {
//
//}


