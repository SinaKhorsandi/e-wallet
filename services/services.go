package wallet

import (
	"fmt"
	"github.com/sinakhorsandi/e-wallet/db"
	"github.com/sinakhorsandi/testify/mock"
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

	//Have nextAccountId and nextOrganizationId both???
	//nextOrganizationId        int
	nextWalletId         int
	personalAccounts     []*db.PersonalAccount
	organizationAccounts []*db.OrganizationAccount
	wallets              []*db.Wallet
	payments             []*db.Transaction
	mock.Mock
}

func (s *Service) userRegister(userName, email string) error {
	b := s.checkCustomerExist(email)
	if b == true {
		return fmt.Errorf("%s", UserIsExist)
	}
	s.nextAccountId++
	s.nextWalletId++
	account := &db.PersonalAccount{
		UserId:   s.nextAccountId,
		UserName: userName,
		Email:    email,
	}
	// create new function for adding wallet!!!????
	wallet := &db.Wallet{
		WalletId: s.nextWalletId,
		UserId:   account.UserId,
		Balance:  0,
	}
	s.personalAccounts = append(s.personalAccounts, account)
	s.wallets = append(s.wallets, wallet)
	return nil
}

func (s *Service) organizationRegister() {

}

func (s *Service) checkCustomerExist(email string) bool {
	for _, account := range s.personalAccounts {
		if account.Email == email {
			return true
		}
	}
	return false
	//ret := s.Called(email)
	//err := ret.Error(0)
	//if err != nil {
	//	return false
	//}
	//return true
}

//Think how can change accountId to another thing
func (s *Service) deposit(accountId int, amount db.Money) error {
	if amount <= 0 {
		return fmt.Errorf("%s\n", NegativeAmount)
	}
	wal, err := s.checkPersonalWallet(accountId)
	if err != nil {
		fmt.Errorf("%s\n", err)
	}
	wal.Balance += amount
	return nil
}
func (s *Service) checkPersonalAccountById(accountId int) (*db.PersonalAccount, error) {
	for _, acc := range s.personalAccounts {
		if acc.UserId == accountId {
			return acc, nil
		}
	}
	return nil, fmt.Errorf("%s\n", AccountNotExist)
}
func (s *Service) checkPersonalWallet(accountId int) (*db.Wallet, error) {
	for _, wal := range s.wallets {
		if wal.UserId == accountId {
			return wal, nil
		}
	}
	return nil, fmt.Errorf("%s\n", AccountNotExist)
}

//func (s *Service) checkOrganizationAccountById(accountId int) bool  {
//	ret := s.Called(accountId)
//	err := ret.Error(0)
//	if err != nil {
//		return false
//	}
//	return true
//}

func pay() {
}
