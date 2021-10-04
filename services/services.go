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
	AccountNotExist				Status = "Account does not exist"
	TransactionStatusOk         Status = "ok"
	TransactionStatusFail       Status = "fail"
	TransactionStatusInProgress Status = " in progress"
)

type Service struct {
	nextAccountId        int

	//Have nextAccountId and nextOrganizationId both???
	//nextOrganizationId        int
	nextWalletId         int
	personalAccounts     []*db.PersonalAccount
	organizationAccounts []*db.OrganizationAccount
	wallets              []*db.Wallet
	payments             []*db.Transaction
	mock.Mock
}

func (s *Service) userRegister(user *db.PersonalAccount) error {
	b := s.checkCustomerExist(user.Email)
	if b == true {
		return fmt.Errorf("%s", UserIsExist)
	}
	s.nextAccountId++
	s.nextWalletId++
	account := &db.PersonalAccount{
		UserId:s.nextAccountId   ,
		UserName: user.UserName,
		Email:    user.Email,
	}
	// create new function for adding wallet!!!????
	wallet := &db.Wallet{
		WalletId: s.nextWalletId,
		UserId:   user.UserId,
		Balance: 0,
	}
	s.personalAccounts = append(s.personalAccounts, account)
	s.wallets = append(s.wallets, wallet)
	return nil
}

func (s *Service) organizationRegister()  {

}

func (s *Service) checkCustomerExist(email string) bool {
	ret := s.Called(email)
	err := ret.Error(0)
	if err != nil {
		return false
	}
	return true
}

func (s *Service) deposit(accountId int,amount db.Money) error  {
	var wallet *db.Wallet
	if amount<=0 {
		return fmt.Errorf("%s\n", NegativeAmount)
	}
	_,err:=s.checkPersonalAccountById(accountId)
	if err != nil {
		fmt.Errorf("%s\n",err)
	}

	//Check
	if wallet.UserId==accountId {
		wallet.Balance+=amount
		return nil
	}
	return nil
	}
func (s *Service) checkPersonalAccountById(accountId int) (*db.PersonalAccount,error)  {
	for _,acc:=range s.personalAccounts{
		if acc.UserId == accountId{
			return acc,nil
		}
	}
	return nil,fmt.Errorf("%s\n",AccountNotExist)
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
