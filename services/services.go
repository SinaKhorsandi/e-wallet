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
	nextWalletId         int
	nextRecordId	int

	//hAccount *hashTable.HashTable
	//hWallet *hashTable.HashTable
	EmailIdAccount     []*data.UserEmail
	NameIdAccount     []*data.UserName
}


func (s *Service) userRegister(userName, email string) error {
	//b := s.checkCustomerExist(email)
	//if b == true {
	//	return fmt.Errorf("%s", UserIsExist)
	//}
	s.nextAccountId++
	s.nextRecordId++
	EmailId := &data.UserEmail{
		RecordId: s.nextRecordId,
		UserId:   s.nextAccountId,
		Email:    email,
	}
	s.nextRecordId++
	NameID:= &data.UserName{
		RecordId: s.nextRecordId,
		UserId:   s.nextAccountId,
		UserName:  userName,
	}
	s.EmailIdAccount=append(s.EmailIdAccount,EmailId)
	s.NameIdAccount=append(s.NameIdAccount,NameID)
	//s.emailIDMap[EmailId.Email]=EmailId.UserId
	//s.emailIDMap[NameID.UserName]=NameID.UserId

	//s.hAccount.Set(email,EmailId.UserId)
	//s.hAccount.Set(userName,NameID.UserId)

	//wallet:=s.newWallet(account.UserId)
	//s.wallets = append(s.wallets, wallet)
	return nil
}
//check function output
func (s *Service) FindEmailByID(userID int) (*data.UserEmail, error) {
	for _, user := range s.EmailIdAccount {
		if user.UserId == userID {
			return user, nil
		}
	}
	return nil, fmt.Errorf("%s\n",AccountNotExist)
}
func (s *Service) FindNameByID(userID int) (*data.UserName, error) {
	for _, user := range s.NameIdAccount {
		if user.UserId == userID {
			return user, nil
		}
	}
	return nil, fmt.Errorf("%s\n",AccountNotExist)
}
//Error handling !!!!
func (s *Service) newWallet(accountId int) *data.Wallet {
	s.nextWalletId++
	wallet := &data.Wallet{
		WalletId: s.nextWalletId,
		UserId:   accountId,
		Balance:  0,
	}
	//wh:=hashTable.NewHash()
	//wh.Set()
	return wallet
}

func (s *Service) checkCustomerExist(email string) bool {
	//v:=s.hAccount.Get(email)
	//if v != nil {
	//	return true
	//}
	return false
}
//
//func (s *Service) organizationRegister() {
//}
//
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
