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
	account       map[string]int
	balance       map[int]data.Money
	wallet        map[int]int
}

var (
	acc = make(map[string]int)
	bal = make(map[int]data.Money)
	wal = make(map[int]int)
)

func (s *Service) userRegister(userName, email string) error {
	b := s.checkCustomerExist(email)
	if b == true {
		return fmt.Errorf("%s", UserIsExist)
	}
	s.nextAccountId++
	s.nextRecordId++
	EmailID := &data.UserEmail{
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
	acc[EmailID.Email] = EmailID.UserId
	acc[NameID.UserName] = NameID.UserId
	s.account = acc

	wallet := s.newWallet(EmailID.UserId)
	bal[wallet.WalletId] = wallet.Balance
	wal[wallet.UserId] = wallet.WalletId
	s.wallet = wal
	s.balance = bal
	return nil
}
func (s *Service) checkCustomerExist(email string) bool {
	_, ok := s.account[email]
	if !ok {
		return false
	}
	return true
}

//Error handling !!!!
func (s *Service) newWallet(userID int) *data.Wallet {
	s.nextWalletId++
	wallet := &data.Wallet{
		WalletId: s.nextWalletId,
		UserId:   userID,
		Balance:  0,
	}
	return wallet
}

//check function output
func (s *Service) findIDByEmail(email string) (int, error) {
	userID, ok := s.account[email]
	if !ok {
		return 0, fmt.Errorf("%s\n", AccountNotExist)
	}
	return userID, nil
}
func (s *Service) findIDByName(userName string) (int, error) {
	userID, ok := s.account[userName]
	if !ok {
		return 0, fmt.Errorf("%s\n", AccountNotExist)
	}
	return userID, nil
}

//check output
func (s *Service) getBalance(userID int) (data.Money, error) {
	balance, ok := s.balance[userID]
	if !ok {

		//check
		return 0, fmt.Errorf("%s\n", AccountNotExist)
	}
	return balance, nil
}

//func showBalance()  {
//}

//Think how can change accountId to another thing
func (s *Service) deposit(walletID int, amount data.Money) (*data.Wallet, error) {
	if amount <= data.Money(0) {
		return nil, fmt.Errorf("%s\n", NegativeAmount)
	}
	wallet, err := s.getWallet(walletID)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	wallet.Balance += amount
	return wallet, nil
}
func (s *Service) getWallet(walletID int) (*data.Wallet, error) {
	userId := s.wallet[walletID]
	b, ok := s.balance[walletID]
	if !ok {
		return nil, fmt.Errorf("%s\n", AccountNotExist)
	}
	wall := &data.Wallet{
		WalletId: walletID,
		UserId:   userId,
		Balance:  b,
	}
	return wall, nil
}

func (s *Service) findWalletID(userID int) (int, error) {
	walletID, ok := s.wallet[userID]
	if !ok {
		return 0, fmt.Errorf("%s\n", AccountNotExist)
	}
	return walletID, nil
}

func deleteAccount() {

}
//len map
func (s *Service) countAccount() int{
	return len(s.account)
}