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
	TransactionStatusOk         Status = "ok"
	TransactionStatusFail       Status = "fail"
	TransactionStatusInProgress Status = " in progress"
)

type Service struct {
	nextUserId        int
	nextOrganizationId        int
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
	s.nextUserId++
	s.nextWalletId++
	account := &db.PersonalAccount{
		UserId:s.nextUserId   ,
		UserName: user.UserName,
		Email:    user.Email,
	}
	//check for creating new function for adding wallet!!!????
	wallet := &db.Wallet{
		WalletId: s.nextWalletId,
		UserId:   user.UserId,
		Balance: 0,
	}
	s.personalAccounts = append(s.personalAccounts, account)
	s.wallets = append(s.wallets, wallet)
	return nil
}

func (s *Service) organizationRegister(organization *db.OrganizationAccount) error {
	b := s.checkCustomerExist(organization.Email)
	if b == true {
		return fmt.Errorf("%s", OrganIsExist)
	}
	s.nextOrganizationId++
	s.nextWalletId++
	account := &db.OrganizationAccount{
		OrganizationId:   s.nextOrganizationId,
		OrganizationName: organization.OrganizationName,
		Email:            organization.Email,
	}
	//check for creating new function for adding wallet!!!????
	wallet := &db.Wallet{
		WalletId:s.nextWalletId ,
		UserId:   organization.OrganizationId,
		Balance: 0,
	}
	s.organizationAccounts = append(s.organizationAccounts, account)
	s.wallets = append(s.wallets, wallet)
	return nil
}

func (s *Service) checkCustomerExist(email string) bool {
	ret := s.Called(email)
	err := ret.Error(0)
	if err != nil {
		return false
	}
	return true
}

func deposit()  {
}

func pay() {
}
